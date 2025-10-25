#!/bin/bash

set -e

# Quick deployment test using existing binary

VPS_HOST="107.175.235.220"
VPS_USER="harvad"
VPS_PASSWORD="11"
DEPLOY_PATH="/opt/smrtmart"
SERVICE_NAME="smrtmart"

GREEN='\033[0;32m'
YELLOW='\033[1;33m'
RED='\033[0;31m'
NC='\033[0m'

echo -e "${GREEN}Testing deployment process...${NC}"

# Check if binary exists
if [ ! -f "server" ]; then
    echo -e "${RED}Error: server binary not found!${NC}"
    echo "Please build it first: make build"
    exit 1
fi

echo -e "\n${YELLOW}[1/4] Testing SSH connection...${NC}"
if sshpass -p "$VPS_PASSWORD" ssh -o StrictHostKeyChecking=no "$VPS_USER@$VPS_HOST" "echo 'Connected successfully'"; then
    echo -e "${GREEN}✓ SSH connection works${NC}"
else
    echo -e "${RED}✗ SSH connection failed${NC}"
    exit 1
fi

echo -e "\n${YELLOW}[2/4] Checking VPS service status...${NC}"
STATUS=$(sshpass -p "$VPS_PASSWORD" ssh -o StrictHostKeyChecking=no "$VPS_USER@$VPS_HOST" "systemctl is-active $SERVICE_NAME")
echo "Current service status: $STATUS"

echo -e "\n${YELLOW}[3/4] Checking deployment directory...${NC}"
sshpass -p "$VPS_PASSWORD" ssh -o StrictHostKeyChecking=no "$VPS_USER@$VPS_HOST" "ls -lh $DEPLOY_PATH/server"

echo -e "\n${YELLOW}[4/4] Checking recent logs...${NC}"
sshpass -p "$VPS_PASSWORD" ssh -o StrictHostKeyChecking=no "$VPS_USER@$VPS_HOST" "sudo journalctl -u $SERVICE_NAME -n 5 --no-pager"

echo -e "\n${GREEN}================================================${NC}"
echo -e "${GREEN}    Deployment Test Complete! ✓${NC}"
echo -e "${GREEN}================================================${NC}"
echo ""
echo "Everything looks good! You can now deploy with:"
echo "  ./deploy-vps.sh"
