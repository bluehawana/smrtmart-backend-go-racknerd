# 🚀 SmartMart Cloud Deployment Guide

## Backend Deployment (Heroku + PostgreSQL)

### Step 1: Deploy Backend to Heroku
1. Clone the repository and navigate to the backend folder
2. Make sure you have the Heroku CLI installed
3. Run the deployment script:
   ```bash
   chmod +x deploy-heroku.sh
   ./deploy-heroku.sh
   ```
4. The script will automatically create a PostgreSQL database

### Step 2: Manual Heroku Setup (Alternative)
If you prefer manual setup:
```bash
# Login to Heroku
heroku login

# Create app
heroku create your-app-name

# Add PostgreSQL addon
heroku addons:create heroku-postgresql:essential-0 --app your-app-name

# Set environment variables
heroku config:set GIN_MODE=release --app your-app-name
heroku config:set DB_SSLMODE=require --app your-app-name

# Deploy
git push heroku main
```

### Step 3: Run Database Migrations
The migrations will run automatically when the app starts. You can also run them manually:
```bash
heroku run ./main migrate --app your-app-name
```

### Step 4: Get Your API URL
- Heroku will provide a URL like: `https://your-app-name.herokuapp.com`
- Your API will be available at: `https://your-app-name.herokuapp.com/api/v1`

### Step 5: Database Management
You can manage your database using:
```bash
# Connect to database
heroku pg:psql --app your-app-name

# View database info
heroku pg:info --app your-app-name

# View database logs
heroku logs --tail --app your-app-name
```

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
NEXT_PUBLIC_API_URL=https://your-heroku-app.herokuapp.com/api/v1
NEXT_PUBLIC_IMAGE_BASE_URL=https://your-heroku-app.herokuapp.com/uploads
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