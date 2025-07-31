package api

import (
	"io"
	"net/http"

	"smrtmart-go-postgresql/internal/models"
	"smrtmart-go-postgresql/internal/service"

	"github.com/gin-gonic/gin"
)

type PaymentHandler struct {
	service service.PaymentService
}

func NewPaymentHandler(service service.PaymentService) *PaymentHandler {
	return &PaymentHandler{service: service}
}

// CreateCheckoutSession godoc
// @Summary Create Stripe checkout session
// @Description Create a Stripe checkout session for payment processing
// @Tags payments
// @Accept json
// @Produce json
// @Param checkout body CheckoutRequest true "Checkout data"
// @Success 200 {object} models.APIResponse{data=CheckoutResponse}
// @Failure 400 {object} models.APIResponse
// @Failure 500 {object} models.APIResponse
// @Router /orders/checkout [post]
func (h *PaymentHandler) CreateCheckoutSession(c *gin.Context) {
	var req CheckoutRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.APIResponse{
			Success: false,
			Message: "Invalid request data",
			Error: &models.APIError{
				Code:    "INVALID_REQUEST",
				Message: err.Error(),
			},
		})
		return
	}

	// Validate required fields
	if len(req.Items) == 0 {
		c.JSON(http.StatusBadRequest, models.APIResponse{
			Success: false,
			Message: "No items in cart",
			Error: &models.APIError{
				Code:    "EMPTY_CART",
				Message: "Cart must contain at least one item",
			},
		})
		return
	}

	if req.CustomerEmail == "" {
		c.JSON(http.StatusBadRequest, models.APIResponse{
			Success: false,
			Message: "Customer email is required",
			Error: &models.APIError{
				Code:    "MISSING_EMAIL",
				Message: "Customer email is required for checkout",
			},
		})
		return
	}

	// Convert request items to service items
	var items []service.CheckoutItem
	for _, item := range req.Items {
		items = append(items, service.CheckoutItem{
			ProductID:   item.ProductID,
			Name:        item.Name,
			Description: item.Description,
			Price:       item.Price,
			Quantity:    item.Quantity,
			Images:      item.Images,
		})
	}

	// Set default URLs if not provided
	successURL := req.SuccessURL
	if successURL == "" {
		successURL = "http://localhost:3000/checkout/success?session_id={CHECKOUT_SESSION_ID}"
	}
	
	cancelURL := req.CancelURL
	if cancelURL == "" {
		cancelURL = "http://localhost:3000/checkout/cancel"
	}

	// Create checkout session
	session, err := h.service.CreateCheckoutSession(items, req.CustomerEmail, successURL, cancelURL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.APIResponse{
			Success: false,
			Message: "Failed to create checkout session",
			Error: &models.APIError{
				Code:    "CHECKOUT_FAILED",
				Message: err.Error(),
			},
		})
		return
	}

	c.JSON(http.StatusOK, models.APIResponse{
		Success: true,
		Message: "Checkout session created successfully",
		Data: CheckoutResponse{
			SessionID:  session.ID,
			SessionURL: session.URL,
		},
	})
}

// StripeWebhook godoc
// @Summary Handle Stripe webhooks
// @Description Handle Stripe webhook events for payment processing
// @Tags payments
// @Accept json
// @Produce json
// @Success 200 {object} models.APIResponse
// @Failure 400 {object} models.APIResponse
// @Failure 500 {object} models.APIResponse
// @Router /webhooks/stripe [post]
func (h *PaymentHandler) StripeWebhook(c *gin.Context) {
	payload, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.APIResponse{
			Success: false,
			Message: "Failed to read request body",
			Error: &models.APIError{
				Code:    "INVALID_PAYLOAD",
				Message: err.Error(),
			},
		})
		return
	}

	signature := c.GetHeader("Stripe-Signature")
	if signature == "" {
		c.JSON(http.StatusBadRequest, models.APIResponse{
			Success: false,
			Message: "Missing Stripe signature",
			Error: &models.APIError{
				Code:    "MISSING_SIGNATURE",
				Message: "Stripe-Signature header is required",
			},
		})
		return
	}

	if err := h.service.HandleWebhook(payload, signature); err != nil {
		c.JSON(http.StatusBadRequest, models.APIResponse{
			Success: false,
			Message: "Failed to process webhook",
			Error: &models.APIError{
				Code:    "WEBHOOK_FAILED",
				Message: err.Error(),
			},
		})
		return
	}

	c.JSON(http.StatusOK, models.APIResponse{
		Success: true,
		Message: "Webhook processed successfully",
	})
}

// Request/Response types
type CheckoutRequest struct {
	Items         []CheckoutItemRequest `json:"items" binding:"required"`
	CustomerEmail string                `json:"customer_email" binding:"required,email"`
	SuccessURL    string                `json:"success_url,omitempty"`
	CancelURL     string                `json:"cancel_url,omitempty"`
}

type CheckoutItemRequest struct {
	ProductID   string   `json:"product_id" binding:"required"`
	Name        string   `json:"name" binding:"required"`
	Description string   `json:"description"`
	Price       float64  `json:"price" binding:"required,gt=0"`
	Quantity    int      `json:"quantity" binding:"required,gt=0"`
	Images      []string `json:"images"`
}

type CheckoutResponse struct {
	SessionID  string `json:"session_id"`
	SessionURL string `json:"session_url"`
}