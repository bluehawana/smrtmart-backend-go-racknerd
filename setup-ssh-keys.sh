#!/bin/bash

# SSH Key Setup Script for Secure VPS Deployment

set -e

VPS_HOST="107.175.235.220"
VPS_USER="harvad"
VPS_PASSWORD="11"

GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m'

echo -e "${GREEN}================================================${NC}"
echo -e "${GREEN}    SSH Key Setup for VPS Deployment${NC}"
echo -e "${GREEN}================================================${NC}"

# Check if SSH key already exists
if [ -f "$HOME/.ssh/id_rsa" ]; then
    echo -e "${YELLOW}SSH key already exists at ~/.ssh/id_rsa${NC}"
    read -p "Do you want to use the existing key? (y/n): " use_existing

    if [ "$use_existing" != "y" ]; then
        echo "Please backup your existing key and run this script again."
        exit 1
    fi
else
    echo -e "\n${YELLOW}[1/3] Generating new SSH key...${NC}"
    ssh-keygen -t rsa -b 4096 -C "deployment@smrtmart" -f "$HOME/.ssh/id_rsa" -N ""
    echo -e "${GREEN}✓ SSH key generated${NC}"
fi

# Copy public key to VPS
echo -e "\n${YELLOW}[2/3] Copying public key to VPS...${NC}"

# Check if sshpass is available
if command -v sshpass &> /dev/null; then
    sshpass -p "$VPS_PASSWORD" ssh-copy-id -o StrictHostKeyChecking=no "$VPS_USER@$VPS_HOST"
else
    echo "Please enter your VPS password when prompted:"
    ssh-copy-id "$VPS_USER@$VPS_HOST"
fi

echo -e "${GREEN}✓ Public key copied to VPS${NC}"

# Test SSH connection
echo -e "\n${YELLOW}[3/3] Testing SSH connection...${NC}"
if ssh -o BatchMode=yes -o StrictHostKeyChecking=no "$VPS_USER@$VPS_HOST" "echo 'SSH key authentication working!'" 2>/dev/null; then
    echo -e "${GREEN}✓ SSH key authentication is working!${NC}"

    echo -e "\n${GREEN}================================================${NC}"
    echo -e "${GREEN}    Setup Complete! 🎉${NC}"
    echo -e "${GREEN}================================================${NC}"
    echo ""
    echo "You can now deploy without passwords using:"
    echo "  ./deploy-vps-secure.sh"
    echo ""
    echo "For GitHub Actions, add this as a secret:"
    echo "  SSH_PRIVATE_KEY: $(cat ~/.ssh/id_rsa | base64 -w 0)"
else
    echo -e "${RED}✗ SSH key authentication failed!${NC}"
    echo "Please check the VPS SSH configuration."
    exit 1
fi
