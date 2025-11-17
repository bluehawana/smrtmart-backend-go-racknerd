package api

import (
	"io"
	"net/http"

	"smrtmart-go-postgresql/internal/models"
	"smrtmart-go-postgresql/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/stripe/stripe-go/v76"
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

	// Check if this is a full customer info request or simple email-only request
	var session *stripe.CheckoutSession
	var err error
	
	if req.CustomerInfo.FirstName != "" && req.CustomerInfo.LastName != "" && 
	   req.ShippingAddress.AddressLine1 != "" && req.ShippingAddress.City != "" {
		// Full customer info checkout
		session, err = h.service.CreateCheckoutSessionWithFullInfo(
			items, 
			service.CustomerInfo{
				FirstName: req.CustomerInfo.FirstName,
				LastName:  req.CustomerInfo.LastName,
				Phone:     req.CustomerInfo.Phone,
				Email:     req.CustomerInfo.Email,
			},
			service.Address{
				FirstName:    req.ShippingAddress.FirstName,
				LastName:     req.ShippingAddress.LastName,
				Company:      req.ShippingAddress.Company,
				AddressLine1: req.ShippingAddress.AddressLine1,
				AddressLine2: req.ShippingAddress.AddressLine2,
				City:         req.ShippingAddress.City,
				State:        req.ShippingAddress.State,
				PostalCode:   req.ShippingAddress.PostalCode,
				Country:      req.ShippingAddress.Country,
				Phone:        req.ShippingAddress.Phone,
			},
			func() *service.Address {
				if req.BillingAddress == nil {
					return nil
				}
				return &service.Address{
					FirstName:    req.BillingAddress.FirstName,
					LastName:     req.BillingAddress.LastName,
					Company:      req.BillingAddress.Company,
					AddressLine1: req.BillingAddress.AddressLine1,
					AddressLine2: req.BillingAddress.AddressLine2,
					City:         req.BillingAddress.City,
					State:        req.BillingAddress.State,
					PostalCode:   req.BillingAddress.PostalCode,
					Country:      req.BillingAddress.Country,
					Phone:        req.BillingAddress.Phone,
				}
			}(),
			successURL, 
			cancelURL,
		)
	} else {
		// Simple email-only checkout (fallback to original method)
		session, err = h.service.CreateCheckoutSession(items, req.CustomerEmail, successURL, cancelURL)
	}
	
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
		// Log error but return 200 to acknowledge receipt
		c.JSON(http.StatusOK, models.APIResponse{
			Success: false,
			Message: "Failed to read request body",
		})
		return
	}

	signature := c.GetHeader("Stripe-Signature")
	if signature == "" {
		// Log error but return 200 to acknowledge receipt
		c.JSON(http.StatusOK, models.APIResponse{
			Success: false,
			Message: "Missing Stripe signature",
		})
		return
	}

	if err := h.service.HandleWebhook(payload, signature); err != nil {
		// Log error but ALWAYS return 200 to prevent Stripe from retrying
		// Stripe requires 2xx status code to consider webhook delivered
		c.JSON(http.StatusOK, models.APIResponse{
			Success: false,
			Message: "Failed to process webhook",
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
	Items           []CheckoutItemRequest `json:"items" binding:"required"`
	CustomerEmail   string                `json:"customer_email" binding:"required,email"`
	CustomerInfo    CustomerInfo          `json:"customer_info,omitempty"`
	ShippingAddress Address               `json:"shipping_address,omitempty"`
	BillingAddress  *Address              `json:"billing_address,omitempty"` // Optional, if different from shipping
	SuccessURL      string                `json:"success_url,omitempty"`
	CancelURL       string                `json:"cancel_url,omitempty"`
}

type CustomerInfo struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
}

type Address struct {
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Company      string `json:"company,omitempty"`
	AddressLine1 string `json:"address_line1"`
	AddressLine2 string `json:"address_line2,omitempty"`
	City         string `json:"city"`
	State        string `json:"state"`
	PostalCode   string `json:"postal_code"`
	Country      string `json:"country"`
	Phone        string `json:"phone,omitempty"`
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