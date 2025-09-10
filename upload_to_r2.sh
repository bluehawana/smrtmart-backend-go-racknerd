#!/bin/bash

# SmrtMart - Upload product images to Cloudflare R2
# Configure these variables with your R2 details

R2_BUCKET_NAME="smrtmart"
R2_ENDPOINT="https://2a35af424f8734e497a5d707344d79d5.r2.cloudflarestorage.com"
R2_PUBLIC_DOMAIN="https://2a35af424f8734e497a5d707344d79d5.r2.cloudflarestorage.com/smrtmart"

echo "🚀 Uploading SmrtMart product images to Cloudflare R2..."

# Create products directory in R2
echo "📁 Creating products directory..."

# Upload all product images
cd uploads

echo "📸 Uploading product images..."

# Electronics & Computers
aws s3 cp "dell-xps-15-2023.jpg" s3://$R2_BUCKET_NAME/products/dell-xps-15-2023.jpg --endpoint-url $R2_ENDPOINT
aws s3 cp "macbook.jpg" s3://$R2_BUCKET_NAME/products/macbook.jpg --endpoint-url $R2_ENDPOINT  
aws s3 cp "asus.jpg" s3://$R2_BUCKET_NAME/products/asus.jpg --endpoint-url $R2_ENDPOINT
aws s3 cp "dell.jpg" s3://$R2_BUCKET_NAME/products/dell.jpg --endpoint-url $R2_ENDPOINT
aws s3 cp "xps.jpg" s3://$R2_BUCKET_NAME/products/xps.jpg --endpoint-url $R2_ENDPOINT
aws s3 cp "ultra.jpg" s3://$R2_BUCKET_NAME/products/ultra.jpg --endpoint-url $R2_ENDPOINT

# Audio
aws s3 cp "airpods2.jpg" s3://$R2_BUCKET_NAME/products/airpods2.jpg --endpoint-url $R2_ENDPOINT
aws s3 cp "sony.jpg" s3://$R2_BUCKET_NAME/products/sony.jpg --endpoint-url $R2_ENDPOINT

# Smartphones & Accessories  
aws s3 cp "iphone.jpg" s3://$R2_BUCKET_NAME/products/iphone.jpg --endpoint-url $R2_ENDPOINT
aws s3 cp "iphone16 promaxcase.jpg" s3://$R2_BUCKET_NAME/products/iphone16-promax-case.jpg --endpoint-url $R2_ENDPOINT
aws s3 cp "iphone-magsafe-case.jpg" s3://$R2_BUCKET_NAME/products/iphone-magsafe-case.jpg --endpoint-url $R2_ENDPOINT

# Wearables
aws s3 cp "huaweismartwatch.jpg" s3://$R2_BUCKET_NAME/products/huawei-smartwatch.jpg --endpoint-url $R2_ENDPOINT
aws s3 cp "huawei-gt2-pro.jpg" s3://$R2_BUCKET_NAME/products/huawei-gt2-pro.jpg --endpoint-url $R2_ENDPOINT

# Accessories & Cables
aws s3 cp "macbook m4 charging cable.png" s3://$R2_BUCKET_NAME/products/macbook-m4-charging-cable.png --endpoint-url $R2_ENDPOINT
aws s3 cp "macbookair adaptor and cable.png" s3://$R2_BUCKET_NAME/products/macbookair-adaptor-cable.png --endpoint-url $R2_ENDPOINT
aws s3 cp "8k data cable dell.jpg" s3://$R2_BUCKET_NAME/products/8k-data-cable-dell.jpg --endpoint-url $R2_ENDPOINT
aws s3 cp "dell-thunderbolt-cable.jpg" s3://$R2_BUCKET_NAME/products/dell-thunderbolt-cable.jpg --endpoint-url $R2_ENDPOINT
aws s3 cp "usb c iphone cable.jpg" s3://$R2_BUCKET_NAME/products/usb-c-iphone-cable.jpg --endpoint-url $R2_ENDPOINT
aws s3 cp "magnetic-charging-cable.jpg" s3://$R2_BUCKET_NAME/products/magnetic-charging-cable.jpg --endpoint-url $R2_ENDPOINT
aws s3 cp "magsafe3-cable-blue.jpg" s3://$R2_BUCKET_NAME/products/magsafe3-cable-blue.jpg --endpoint-url $R2_ENDPOINT
aws s3 cp "apple-29w-adapter.jpg" s3://$R2_BUCKET_NAME/products/apple-29w-adapter.jpg --endpoint-url $R2_ENDPOINT

# Cases & Protection
aws s3 cp "macbook-air-case-green.jpg" s3://$R2_BUCKET_NAME/products/macbook-air-case-green.jpg --endpoint-url $R2_ENDPOINT
aws s3 cp "macbookair m3 weaving case.jpg" s3://$R2_BUCKET_NAME/products/macbookair-m3-weaving-case.jpg --endpoint-url $R2_ENDPOINT

# Tracking & Smart Devices
aws s3 cp "mtrackingtag.jpg" s3://$R2_BUCKET_NAME/products/m-tracking-tag.jpg --endpoint-url $R2_ENDPOINT
aws s3 cp "mtag-tracker.jpg" s3://$R2_BUCKET_NAME/products/mtag-tracker.jpg --endpoint-url $R2_ENDPOINT
aws s3 cp "smart tracking card.jpg" s3://$R2_BUCKET_NAME/products/smart-tracking-card.jpg --endpoint-url $R2_ENDPOINT

# AI & Translation
aws s3 cp "ai-translate-pro.jpg" s3://$R2_BUCKET_NAME/products/ai-translate-pro.jpg --endpoint-url $R2_ENDPOINT
aws s3 cp "smart-translator.jpg" s3://$R2_BUCKET_NAME/products/smart-translator.jpg --endpoint-url $R2_ENDPOINT

echo "✅ Upload complete!"
echo "🔗 Images are now available at: $R2_PUBLIC_DOMAIN/products/"

cd ..

echo ""
echo "📝 Next steps:"
echo "1. Update your R2 bucket settings to allow public access for the 'products/' folder"
echo "2. Update the database migration files with the new R2 URLs"
echo "3. Test the image URLs in your browser"