# Stripe Webhook Quick Setup

## ✅ Status: CHECKOUT WORKING, WEBHOOK CONFIGURED

### Webhook Configuration in Stripe Dashboard

**1. Add Webhook Endpoint**
```
URL: https://api.smrtmart.com/api/v1/webhooks/stripe
```

**2. Select These Events:**
```
✅ checkout.session.completed
✅ payment_intent.succeeded
✅ payment_intent.payment_failed
```

**3. Copy Webhook Secret**
```
Format: whsec_...
```

**Note:** Your webhook secret is already configured on VPS. Only update if you generate a new one in Stripe.

**4. If You Generated a New Secret, Update VPS:**
```bash
ssh harvad@107.175.235.220
nano /opt/smrtmart/.env
# Update STRIPE_WEBHOOK_SECRET=whsec_NEW_SECRET
sudo systemctl restart smrtmart
```

---

## 🧪 Quick Test

### Test Checkout
```bash
curl -X POST https://api.smrtmart.com/api/v1/orders/checkout \
  -H "Content-Type: application/json" \
  -d '{
    "items": [{
      "product_id": "test",
      "name": "Test Product",
      "description": "Test",
      "price": 10.00,
      "quantity": 1
    }],
    "customer_email": "test@example.com"
  }'
```

### Monitor Webhooks
```bash
ssh harvad@107.175.235.220 'sudo journalctl -u smrtmart -f | grep webhook'
```

---

## 📋 Verification Checklist

- [ ] Add webhook URL in Stripe Dashboard
- [ ] Select events (checkout.session.completed, etc.)
- [ ] Copy webhook secret
- [ ] Update VPS .env if secret changed
- [ ] Restart service if secret changed
- [ ] Test payment flow end-to-end
- [ ] Verify webhook delivery in Stripe Dashboard

---

## 🔗 Full Documentation

See `STRIPE_WEBHOOK_GUIDE.md` for complete documentation.
