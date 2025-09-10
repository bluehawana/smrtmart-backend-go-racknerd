#!/bin/bash

echo "🚀 SmrtMart Go Backend Deployment Script"
echo "=========================================="

# Check if we're on dev branch
BRANCH=$(git branch --show-current)
if [ "$BRANCH" != "dev" ]; then
    echo "⚠️  Please switch to dev branch first: git checkout dev"
    exit 1
fi

echo "✅ On dev branch - ready for deployment"

# Add and commit any changes
echo "📦 Committing latest changes..."
git add .
git commit -m "🚀 Deploy: Latest changes for cloud deployment" || echo "No changes to commit"

# Push to GitHub
echo "📤 Pushing to GitHub..."
git push origin dev

echo ""
echo "🎉 Ready for Railway deployment!"
echo ""
echo "Next steps:"
echo "1. Go to https://railway.app"
echo "2. Create new project from GitHub"
echo "3. Select smrtmart-backend repository"
echo "4. Choose 'dev' branch"
echo "5. Add PostgreSQL database"
echo "6. Configure environment variables"
echo ""
echo "Your API will be available at: https://your-app-name.railway.app/api/v1"