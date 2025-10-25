#!/bin/bash

set -e  # Exit on any error

# Configuration
VPS_HOST="107.175.235.220"
VPS_USER="harvad"
VPS_PASSWORD="11"
DEPLOY_PATH="/opt/smrtmart"
SERVICE_NAME="smrtmart"
BINARY_NAME="server"
LOCAL_BUILD_DIR="."

# Colors for output
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
RED='\033[0;31m'
NC='\033[0m' # No Color

echo -e "${GREEN}================================================${NC}"
echo -e "${GREEN}    SmrtMart VPS Deployment Script${NC}"
echo -e "${GREEN}================================================${NC}"

# Check if sshpass is installed
if ! command -v sshpass &> /dev/null; then
    echo -e "${RED}Error: sshpass is not installed${NC}"
    echo "Install it with: sudo apt install sshpass"
    exit 1
fi

# Step 1: Build the application locally
echo -e "\n${YELLOW}[1/7] Building application locally...${NC}"
cd "$LOCAL_BUILD_DIR"
GOOS=linux GOARCH=amd64 go build -o "$BINARY_NAME" ./cmd/server
if [ ! -f "$BINARY_NAME" ]; then
    echo -e "${RED}Build failed!${NC}"
    exit 1
fi
echo -e "${GREEN}✓ Build successful${NC}"

# Step 2: Create backup on VPS
echo -e "\n${YELLOW}[2/7] Creating backup of current deployment...${NC}"
BACKUP_TIMESTAMP=$(date +%Y%m%d_%H%M%S)
sshpass -p "$VPS_PASSWORD" ssh -o StrictHostKeyChecking=no "$VPS_USER@$VPS_HOST" \
    "cp $DEPLOY_PATH/$BINARY_NAME $DEPLOY_PATH/${BINARY_NAME}.backup.$BACKUP_TIMESTAMP 2>/dev/null || echo 'No existing binary to backup'"
echo -e "${GREEN}✓ Backup created${NC}"

# Step 3: Stop the service
echo -e "\n${YELLOW}[3/7] Stopping service on VPS...${NC}"
sshpass -p "$VPS_PASSWORD" ssh -o StrictHostKeyChecking=no "$VPS_USER@$VPS_HOST" \
    "sudo systemctl stop $SERVICE_NAME"
echo -e "${GREEN}✓ Service stopped${NC}"

# Step 4: Upload new binary
echo -e "\n${YELLOW}[4/7] Uploading new binary...${NC}"
sshpass -p "$VPS_PASSWORD" scp -o StrictHostKeyChecking=no "$BINARY_NAME" "$VPS_USER@$VPS_HOST:$DEPLOY_PATH/$BINARY_NAME"
echo -e "${GREEN}✓ Binary uploaded${NC}"

# Step 5: Set permissions
echo -e "\n${YELLOW}[5/7] Setting permissions...${NC}"
sshpass -p "$VPS_PASSWORD" ssh -o StrictHostKeyChecking=no "$VPS_USER@$VPS_HOST" \
    "chmod +x $DEPLOY_PATH/$BINARY_NAME"
echo -e "${GREEN}✓ Permissions set${NC}"

# Step 6: Upload static files and templates (if they exist)
echo -e "\n${YELLOW}[6/7] Uploading static assets and templates...${NC}"
if [ -d "static" ]; then
    sshpass -p "$VPS_PASSWORD" scp -r -o StrictHostKeyChecking=no static "$VPS_USER@$VPS_HOST:$DEPLOY_PATH/"
    echo -e "${GREEN}✓ Static files uploaded${NC}"
fi
if [ -d "templates" ]; then
    sshpass -p "$VPS_PASSWORD" scp -r -o StrictHostKeyChecking=no templates "$VPS_USER@$VPS_HOST:$DEPLOY_PATH/"
    echo -e "${GREEN}✓ Templates uploaded${NC}"
fi

# Step 7: Start the service
echo -e "\n${YELLOW}[7/7] Starting service...${NC}"
sshpass -p "$VPS_PASSWORD" ssh -o StrictHostKeyChecking=no "$VPS_USER@$VPS_HOST" \
    "sudo systemctl start $SERVICE_NAME"

# Wait a moment for service to start
sleep 3

# Check service status
echo -e "\n${YELLOW}Checking service status...${NC}"
SERVICE_STATUS=$(sshpass -p "$VPS_PASSWORD" ssh -o StrictHostKeyChecking=no "$VPS_USER@$VPS_HOST" \
    "systemctl is-active $SERVICE_NAME")

if [ "$SERVICE_STATUS" == "active" ]; then
    echo -e "${GREEN}✓ Service started successfully${NC}"

    # Show recent logs
    echo -e "\n${YELLOW}Recent logs:${NC}"
    sshpass -p "$VPS_PASSWORD" ssh -o StrictHostKeyChecking=no "$VPS_USER@$VPS_HOST" \
        "sudo journalctl -u $SERVICE_NAME -n 10 --no-pager"

    echo -e "\n${GREEN}================================================${NC}"
    echo -e "${GREEN}    Deployment Successful! 🎉${NC}"
    echo -e "${GREEN}================================================${NC}"
    echo -e "Your API is now running on: http://$VPS_HOST:8080"
    echo -e "Backup saved as: ${BINARY_NAME}.backup.$BACKUP_TIMESTAMP"
    echo -e "\nUseful commands:"
    echo -e "  View logs: ssh $VPS_USER@$VPS_HOST 'sudo journalctl -u $SERVICE_NAME -f'"
    echo -e "  Restart:   ssh $VPS_USER@$VPS_HOST 'sudo systemctl restart $SERVICE_NAME'"
    echo -e "  Status:    ssh $VPS_USER@$VPS_HOST 'sudo systemctl status $SERVICE_NAME'"
else
    echo -e "${RED}✗ Service failed to start!${NC}"
    echo -e "\n${YELLOW}Error logs:${NC}"
    sshpass -p "$VPS_PASSWORD" ssh -o StrictHostKeyChecking=no "$VPS_USER@$VPS_HOST" \
        "sudo journalctl -u $SERVICE_NAME -n 20 --no-pager"

    echo -e "\n${YELLOW}Rolling back to previous version...${NC}"
    sshpass -p "$VPS_PASSWORD" ssh -o StrictHostKeyChecking=no "$VPS_USER@$VPS_HOST" \
        "cp $DEPLOY_PATH/${BINARY_NAME}.backup.$BACKUP_TIMESTAMP $DEPLOY_PATH/$BINARY_NAME && \
         sudo systemctl start $SERVICE_NAME"

    echo -e "${RED}Deployment failed! Rolled back to previous version.${NC}"
    exit 1
fi

# Clean up local build artifact (optional)
echo -e "\n${YELLOW}Cleaning up local build...${NC}"
rm -f "$BINARY_NAME"
echo -e "${GREEN}✓ Local cleanup complete${NC}"
