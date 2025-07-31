# 🚀 SmartMart Cloud Deployment Guide

## Backend Deployment (Railway)

### Step 1: Create Railway Account
1. Go to [railway.app](https://railway.app)
2. Sign up with GitHub
3. Connect your GitHub account

### Step 2: Deploy Backend
1. Click "New Project" → "Deploy from GitHub repo"
2. Select `smrtmart-backend` repository
3. Choose `dev` branch for testing
4. Railway will auto-detect Go and deploy

### Step 3: Add Database
1. In Railway dashboard, click "New" → "Database" → "PostgreSQL"
2. Copy the connection details
3. Add environment variables in Railway:
   ```
   DB_HOST=your-railway-postgres-host
   DB_PORT=5432
   DB_USER=postgres
   DB_PASSWORD=your-railway-db-password
   DB_NAME=railway
   DB_SSLMODE=require
   GIN_MODE=release
   PORT=8080
   CORS_ORIGINS=https://smrtmart.com,https://www.smrtmart.com
   JWT_SECRET=your-secure-jwt-secret
   STRIPE_SECRET_KEY=sk_test_your_stripe_key
   ```

### Step 4: Run Migrations
1. In Railway, go to your service
2. Open "Deploy Logs"
3. Migrations will run automatically on startup

### Step 5: Get Your API URL
- Railway will provide a URL like: `https://your-app-name.railway.app`
- Your API will be available at: `https://your-app-name.railway.app/api/v1`

## Frontend Deployment (Vercel)

### Step 1: Create Vercel Account
1. Go to [vercel.com](https://vercel.com)
2. Sign up with GitHub
3. Connect your GitHub account

### Step 2: Deploy Frontend
1. Click "New Project"
2. Import `smrtmart-frontend` repository
3. Choose `dev` branch
4. Vercel will auto-detect Next.js

### Step 3: Add Environment Variables
In Vercel dashboard, add:
```
NEXT_PUBLIC_API_URL=https://your-railway-app.railway.app/api/v1
NEXT_PUBLIC_IMAGE_BASE_URL=https://your-railway-app.railway.app/uploads
NEXT_PUBLIC_STRIPE_PUBLISHABLE_KEY=pk_test_your_stripe_publishable_key
```

### Step 4: Configure Custom Domain
1. In Vercel, go to Project Settings → Domains
2. Add your domain: `smrtmart.com`
3. Configure DNS in Cloudflare:
   - Add CNAME record: `www` → `cname.vercel-dns.com`
   - Add A record: `@` → Vercel IP addresses

## Testing Checklist

### Backend Testing
- [ ] Health check: `https://your-api.railway.app/health`
- [ ] Products API: `https://your-api.railway.app/api/v1/products`
- [ ] Featured products: `https://your-api.railway.app/api/v1/products/featured`
- [ ] Stripe checkout: Test with sample data

### Frontend Testing
- [ ] Homepage loads with products
- [ ] Shopping cart functionality
- [ ] Checkout flow with Stripe
- [ ] Responsive design on mobile
- [ ] All images load correctly

### Domain Testing
- [ ] `https://smrtmart.com` loads correctly
- [ ] `https://www.smrtmart.com` redirects properly
- [ ] SSL certificate is valid
- [ ] All API calls work from custom domain

## Production Deployment

Once dev branch testing is successful:

1. **Merge to main:**
   ```bash
   git checkout main
   git merge dev
   git push origin main
   ```

2. **Update deployment branches:**
   - Railway: Switch to `main` branch
   - Vercel: Switch to `main` branch

3. **Update environment variables:**
   - Use live Stripe keys
   - Update CORS origins
   - Set production database

## Monitoring & Maintenance

- **Railway**: Monitor logs and metrics in dashboard
- **Vercel**: Check deployment logs and analytics
- **Cloudflare**: Monitor DNS and security settings
- **Stripe**: Monitor payments in dashboard

## Support

If you encounter issues:
1. Check deployment logs in Railway/Vercel
2. Verify environment variables
3. Test API endpoints individually
4. Check database connections