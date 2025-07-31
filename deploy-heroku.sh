#!/bin/bash

echo "🚀 SmartMart Go Backend - Heroku Deployment Script"
echo "=================================================="

# Check if logged into Heroku
if ! heroku auth:whoami > /dev/null 2>&1; then
    echo "❌ Please login to Heroku first: heroku login"
    exit 1
fi

# Create unique app name
APP_NAME="smrtmart-go-backend-$(date +%s)"
echo "📱 Creating Heroku app: $APP_NAME"

# Create app
heroku create $APP_NAME

# Add PostgreSQL addon
echo "🐘 Adding PostgreSQL addon..."
heroku addons:create heroku-postgresql:mini --app $APP_NAME

# Set environment variables
echo "⚙️  Setting environment variables..."
heroku config:set GIN_MODE=release --app $APP_NAME
heroku config:set DB_SSLMODE=require --app $APP_NAME

# Add git remote if it doesn't exist
if ! git remote get-url heroku > /dev/null 2>&1; then
    heroku git:remote --app $APP_NAME
fi

# Deploy
echo "🚀 Deploying to Heroku..."
git add .
git commit -m "Deploy to Heroku" || echo "No changes to commit"
git push heroku master

# Get app URL
APP_URL=$(heroku apps:info --app $APP_NAME --json | jq -r '.app.web_url')

echo ""
echo "🎉 Deployment complete!"
echo "🌐 Your API is available at: ${APP_URL}api/v1"
echo "📚 API Documentation: ${APP_URL}swagger/index.html"
echo "🔧 App name: $APP_NAME"
echo ""
echo "Next steps:"
echo "1. Update your frontend environment variables with: ${APP_URL}api/v1"
echo "2. Test your API endpoints"
echo "3. Set up custom domain if needed"