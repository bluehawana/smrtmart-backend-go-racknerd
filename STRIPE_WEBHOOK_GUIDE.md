# Stripe Webhook Configuration Guide

## ✅ Current Status

**Checkout Endpoint:** ✅ WORKING
**Webhook Endpoint:** ✅ CONFIGURED
**Environment:** ✅ PRODUCTION (Live Keys)

### Endpoints

- **Checkout:** `POST https://api.smrtmart.com/api/v1/orders/checkout`
- **Webhook:** `POST https://api.smrtmart.com/api/v1/webhooks/stripe`
- **API Health:** `GET https://api.smrtmart.com/api/v1/health`

### Environment Variables (VPS)

```bash
STRIPE_SECRET_KEY=sk_live_YOUR_SECRET_KEY_HERE
STRIPE_WEBHOOK_SECRET=whsec_YOUR_WEBHOOK_SECRET_HERE
```

**Note:** Your actual keys are configured on the VPS in `/opt/smrtmart/.env`

---

## 🔧 Stripe Dashboard Configuration

### Step 1: Configure Webhook in Stripe Dashboard

1. Go to **Stripe Dashboard** → **Developers** → **Webhooks**
2. Click **Add endpoint**
3. Enter endpoint URL:
   ```
   https://api.smrtmart.com/api/v1/webhooks/stripe
   ```

### Step 2: Select Events to Listen

Enable these events for your webhook:

#### Required Events:
- ✅ `checkout.session.completed` - When checkout is successful
- ✅ `payment_intent.succeeded` - When payment succeeds
- ✅ `payment_intent.payment_failed` - When payment fails

#### Optional but Recommended:
- `checkout.session.expired` - Checkout session expired
- `charge.succeeded` - Charge succeeded
- `charge.failed` - Charge failed
- `charge.refunded` - Charge refunded
- `customer.created` - New customer created
- `invoice.paid` - Invoice paid
- `invoice.payment_failed` - Invoice payment failed

### Step 3: Get Webhook Signing Secret

1. After creating the webhook, Stripe will show you a **Signing secret**
2. It starts with `whsec_`
3. Copy this secret

### Step 4: Update VPS Environment

```bash
# SSH into VPS
ssh harvad@107.175.235.220

# Edit environment file
nano /opt/smrtmart/.env

# Update or add:
STRIPE_WEBHOOK_SECRET=whsec_YOUR_NEW_SECRET_HERE

# Save and exit (Ctrl+X, Y, Enter)

# Restart service
sudo systemctl restart smrtmart

# Verify
sudo systemctl status smrtmart
```

---

## 📝 Webhook Implementation Details

### Current Implementation

The webhook handler (in `internal/service/payment_service.go:229`) currently handles:

1. **`checkout.session.completed`** - Payment successful
   - Logs session ID
   - TODO: Create order in database
   - TODO: Update inventory
   - TODO: Send confirmation email

2. **`payment_intent.succeeded`** - Payment intent succeeded
   - Logs payment intent ID

3. **`payment_intent.payment_failed`** - Payment failed
   - Logs failure

### Webhook Security

- ✅ Validates Stripe signature using `stripe.webhook.ConstructEvent()`
- ✅ Verifies webhook secret from environment
- ✅ Rejects unsigned or invalid requests

---

## 🧪 Testing

### Test Checkout Endpoint

```bash
curl -X POST https://api.smrtmart.com/api/v1/orders/checkout \
  -H "Content-Type: application/json" \
  -d '{
    "items": [{
      "product_id": "prod_123",
      "name": "Test Product",
      "description": "Test product description",
      "price": 29.99,
      "quantity": 1,
      "images": []
    }],
    "customer_email": "customer@example.com",
    "success_url": "https://smrtmart.com/checkout/success?session_id={CHECKOUT_SESSION_ID}",
    "cancel_url": "https://smrtmart.com/checkout/cancel"
  }'
```

**Expected Response:**
```json
{
  "success": true,
  "message": "Checkout session created successfully",
  "data": {
    "session_id": "cs_live_...",
    "session_url": "https://checkout.stripe.com/..."
  }
}
```

### Test Webhook with Stripe CLI

```bash
# Install Stripe CLI
brew install stripe/stripe-cli/stripe  # Mac
# or
snap install stripe  # Linux

# Login
stripe login

# Forward webhooks to local/VPS
stripe listen --forward-to https://api.smrtmart.com/api/v1/webhooks/stripe

# Trigger test events
stripe trigger checkout.session.completed
stripe trigger payment_intent.succeeded
stripe trigger payment_intent.payment_failed
```

### Test from Frontend

```javascript
// Example frontend code
const checkoutData = {
  items: [
    {
      product_id: "prod_123",
      name: "Premium T-Shirt",
      description: "High quality cotton t-shirt",
      price: 29.99,
      quantity: 2,
      images: ["tshirt-main.jpg"]
    }
  ],
  customer_email: "customer@example.com",
  success_url: "https://smrtmart.com/checkout/success?session_id={CHECKOUT_SESSION_ID}",
  cancel_url: "https://smrtmart.com/checkout/cancel"
};

const response = await fetch('https://api.smrtmart.com/api/v1/orders/checkout', {
  method: 'POST',
  headers: {
    'Content-Type': 'application/json'
  },
  body: JSON.stringify(checkoutData)
});

const result = await response.json();

if (result.success) {
  // Redirect to Stripe Checkout
  window.location.href = result.data.session_url;
}
```

---

## 🔍 Monitoring & Debugging

### Check Webhook Logs on VPS

```bash
# View real-time webhook events
ssh harvad@107.175.235.220 'sudo journalctl -u smrtmart -f | grep -i webhook'

# View recent webhook events
ssh harvad@107.175.235.220 'sudo journalctl -u smrtmart -n 100 | grep -i "webhook\|stripe\|payment"'

# Check for errors
ssh harvad@107.175.235.220 'sudo journalctl -u smrtmart -p err -n 50'
```

### Stripe Dashboard Monitoring

1. Go to **Stripe Dashboard** → **Developers** → **Webhooks**
2. Click on your webhook endpoint
3. View **Recent deliveries** tab
4. Check delivery status and response codes

### Common Response Codes

- **200** ✅ - Webhook processed successfully
- **400** ❌ - Invalid signature or payload
- **500** ❌ - Server error processing webhook

---

## 🛠️ Troubleshooting

### Issue: Webhook Returns 400 (Invalid Signature)

**Cause:** Webhook secret mismatch or wrong secret

**Solution:**
```bash
# 1. Get correct secret from Stripe Dashboard
# 2. Update on VPS
ssh harvad@107.175.235.220
nano /opt/smrtmart/.env
# Update STRIPE_WEBHOOK_SECRET
sudo systemctl restart smrtmart
```

### Issue: Webhook Returns 500

**Cause:** Application error

**Solution:**
```bash
# Check logs for specific error
ssh harvad@107.175.235.220 'sudo journalctl -u smrtmart -n 100'

# Check if service is running
ssh harvad@107.175.235.220 'systemctl status smrtmart'
```

### Issue: Webhooks Not Receiving Events

**Checklist:**
- [ ] Webhook URL is correct in Stripe Dashboard
- [ ] Webhook endpoint is HTTPS (required by Stripe)
- [ ] Events are enabled in Stripe Dashboard
- [ ] Service is running on VPS
- [ ] Firewall allows HTTPS traffic
- [ ] DNS is pointing to correct server

---

## 📊 Webhook Events Flow

```
Customer → Stripe Checkout → Payment Successful
                                    ↓
                          Stripe sends webhook to:
                  https://api.smrtmart.com/api/v1/webhooks/stripe
                                    ↓
                        Backend verifies signature
                                    ↓
                         Process event type:
                                    ↓
              ┌─────────────────────┼─────────────────────┐
              ↓                     ↓                     ↓
    checkout.session.    payment_intent.      payment_intent.
       completed            succeeded          payment_failed
              ↓                     ↓                     ↓
        Create Order         Log Success          Log Failure
        Update Inventory                      Notify Customer
        Send Email
```

---

## 🔐 Security Best Practices

### 1. Always Verify Webhook Signatures
✅ Already implemented using `stripe.webhook.ConstructEvent()`

### 2. Use HTTPS
✅ Your API is already using HTTPS

### 3. Keep Webhook Secret Secure
- ✅ Stored in environment variables
- ✅ Not committed to Git
- ✅ File permissions: `chmod 600 /opt/smrtmart/.env`

### 4. Implement Idempotency
**TODO:** Implement idempotency to handle duplicate webhook events

```go
// Example: Store processed event IDs
processedEvents := make(map[string]bool)

func (s *paymentService) HandleWebhook(payload []byte, signature string) error {
    event, err := webhook.ConstructEvent(payload, signature, s.stripeConfig.WebhookSecret)

    // Check if already processed
    if processedEvents[event.ID] {
        log.Printf("Duplicate event: %s", event.ID)
        return nil
    }

    // Process event...

    // Mark as processed
    processedEvents[event.ID] = true
}
```

### 5. Rate Limiting
**TODO:** Consider implementing rate limiting for webhook endpoint

---

## 📋 Next Steps / TODO

### Immediate Tasks

- [ ] **Verify Webhook URL in Stripe Dashboard**
  - Ensure URL is: `https://api.smrtmart.com/api/v1/webhooks/stripe`
  - Ensure correct events are selected

- [ ] **Test End-to-End Flow**
  - Create test checkout session
  - Complete payment
  - Verify webhook is received
  - Check logs on VPS

### Development Tasks

- [ ] Implement order creation on `checkout.session.completed`
- [ ] Add database transaction for order creation
- [ ] Implement inventory update logic
- [ ] Add email notification service
- [ ] Implement webhook event idempotency
- [ ] Add webhook retry logic for failed processing
- [ ] Create admin dashboard for order management
- [ ] Add webhook event logging to database

### Monitoring Tasks

- [ ] Set up webhook monitoring alerts
- [ ] Configure log rotation for webhook logs
- [ ] Add metrics for webhook processing time
- [ ] Set up error alerting (email/SMS)

---

## 📞 Support Resources

- **Stripe Webhooks Documentation:** https://stripe.com/docs/webhooks
- **Stripe Testing:** https://stripe.com/docs/testing
- **Stripe CLI:** https://stripe.com/docs/stripe-cli
- **GitHub Repository:** https://github.com/bluehawana/smrtmart-backend-go-racknerd

---

## ✅ Verification Checklist

### Before Going Live

- [x] Checkout endpoint tested and working
- [x] Webhook endpoint accessible via HTTPS
- [x] Webhook secret configured on VPS
- [x] Service running and healthy
- [ ] Webhook URL configured in Stripe Dashboard
- [ ] Test payment completed successfully
- [ ] Webhook event received and logged
- [ ] Error handling tested
- [ ] Production monitoring in place

---

**Last Updated:** October 25, 2025
**Status:** ✅ Ready for Stripe Dashboard Configuration
