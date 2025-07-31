package service

import (
	"encoding/json"
	"fmt"
	"log"

	"smrtmart-go-postgresql/internal/config"

	"github.com/stripe/stripe-go/v76"
	"github.com/stripe/stripe-go/v76/checkout/session"
	"github.com/stripe/stripe-go/v76/webhook"
)

type PaymentService interface {
	CreateCheckoutSession(items []CheckoutItem, customerEmail string, successURL, cancelURL string) (*stripe.CheckoutSession, error)
	HandleWebhook(payload []byte, signature string) error
}

type CheckoutItem struct {
	ProductID   string  `json:"product_id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Quantity    int     `json:"quantity"`
	Images      []string `json:"images"`
}

type paymentService struct {
	stripeConfig config.StripeConfig
}

func NewPaymentService(stripeConfig config.StripeConfig) PaymentService {
	// Initialize Stripe
	stripe.Key = stripeConfig.SecretKey
	return &paymentService{stripeConfig: stripeConfig}
}

func (s *paymentService) CreateCheckoutSession(items []CheckoutItem, customerEmail string, successURL, cancelURL string) (*stripe.CheckoutSession, error) {
	var lineItems []*stripe.CheckoutSessionLineItemParams

	for _, item := range items {
		// Convert price to cents (Stripe uses cents)
		priceInCents := int64(item.Price * 100)

		lineItem := &stripe.CheckoutSessionLineItemParams{
			PriceData: &stripe.CheckoutSessionLineItemPriceDataParams{
				Currency: stripe.String("usd"),
				ProductData: &stripe.CheckoutSessionLineItemPriceDataProductDataParams{
					Name:        stripe.String(item.Name),
					Description: stripe.String(item.Description),
				},
				UnitAmount: stripe.Int64(priceInCents),
			},
			Quantity: stripe.Int64(int64(item.Quantity)),
		}

		// Add product images if available
		if len(item.Images) > 0 {
			var images []*string
			for _, img := range item.Images {
				// Construct full image URL
				imageURL := fmt.Sprintf("http://localhost:8080/uploads/%s", img)
				images = append(images, stripe.String(imageURL))
			}
			lineItem.PriceData.ProductData.Images = images
		}

		lineItems = append(lineItems, lineItem)
	}

	params := &stripe.CheckoutSessionParams{
		PaymentMethodTypes: stripe.StringSlice([]string{"card"}),
		LineItems:          lineItems,
		Mode:               stripe.String(string(stripe.CheckoutSessionModePayment)),
		SuccessURL:         stripe.String(successURL),
		CancelURL:          stripe.String(cancelURL),
		CustomerEmail:      stripe.String(customerEmail),
		
		// Enable shipping address collection
		ShippingAddressCollection: &stripe.CheckoutSessionShippingAddressCollectionParams{
			AllowedCountries: stripe.StringSlice([]string{"US", "CA", "GB", "DE", "FR", "ES", "IT", "NL", "BE", "AT", "CH", "SE", "NO", "DK", "FI"}),
		},
		
		// Add metadata for order tracking
		Metadata: map[string]string{
			"source": "smrtmart_website",
		},
	}

	sess, err := session.New(params)
	if err != nil {
		return nil, fmt.Errorf("failed to create checkout session: %w", err)
	}

	return sess, nil
}

func (s *paymentService) HandleWebhook(payload []byte, signature string) error {
	event, err := webhook.ConstructEvent(payload, signature, s.stripeConfig.WebhookSecret)
	if err != nil {
		return fmt.Errorf("failed to verify webhook signature: %w", err)
	}

	switch event.Type {
	case "checkout.session.completed":
		var session stripe.CheckoutSession
		if err := json.Unmarshal(event.Data.Raw, &session); err != nil {
			return fmt.Errorf("failed to unmarshal session: %w", err)
		}
		
		log.Printf("Payment successful for session: %s", session.ID)
		// TODO: Create order in database, update inventory, send confirmation email
		
	case "payment_intent.succeeded":
		var paymentIntent stripe.PaymentIntent
		if err := json.Unmarshal(event.Data.Raw, &paymentIntent); err != nil {
			return fmt.Errorf("failed to unmarshal payment intent: %w", err)
		}
		
		log.Printf("Payment intent succeeded: %s", paymentIntent.ID)
		
	case "payment_intent.payment_failed":
		var paymentIntent stripe.PaymentIntent
		if err := json.Unmarshal(event.Data.Raw, &paymentIntent); err != nil {
			return fmt.Errorf("failed to unmarshal payment intent: %w", err)
		}
		
		log.Printf("Payment failed: %s", paymentIntent.ID)
		
	default:
		log.Printf("Unhandled event type: %s", event.Type)
	}

	return nil
}