# SmrtMart VPS Deployment Guide

## Overview

This guide covers automated deployment of the SmrtMart Go backend to your VPS at `107.175.235.220`.

## Quick Start

### Option 1: Deploy with Password (Quick & Easy)

```bash
# Deploy immediately using existing binary
./deploy-vps.sh
```

### Option 2: Deploy with SSH Keys (More Secure)

```bash
# One-time setup
./setup-ssh-keys.sh

# Then deploy without passwords
./deploy-vps-secure.sh
```

## Deployment Scripts

### 1. `deploy-vps.sh` - Password-based Deployment

**Features:**
- ✅ Builds Go binary locally (Linux AMD64)
- ✅ Creates backup of current deployment
- ✅ Zero-downtime deployment
- ✅ Automatic rollback on failure
- ✅ Service health checks
- ✅ Deploys static files and templates

**Usage:**
```bash
./deploy-vps.sh
```

**Requirements:**
- `sshpass` installed: `sudo apt install sshpass`
- Go 1.24+ installed locally (for building)

### 2. `deploy-vps-secure.sh` - SSH Key-based Deployment

Same features as above, but uses SSH key authentication instead of passwords.

**Usage:**
```bash
# First time only - setup SSH keys
./setup-ssh-keys.sh

# Deploy
./deploy-vps-secure.sh
```

### 3. `test-deployment.sh` - Pre-flight Check

Tests your deployment setup without making changes.

**Usage:**
```bash
./test-deployment.sh
```

**Checks:**
- SSH connectivity
- Service status
- Deployment directory permissions
- Recent logs

## GitHub Actions CI/CD

### Setup Instructions

1. **Add Repository Secrets** (Settings → Secrets and variables → Actions):

   ```
   VPS_HOST = 107.175.235.220
   VPS_USER = harvad
   VPS_PASSWORD = your_password
   ```

2. **Enable GitHub Actions:**
   - Push this repository to GitHub
   - Actions will run automatically on push to `main` or `dev`

3. **Manual Deployment:**
   - Go to Actions tab
   - Select "Deploy to VPS" workflow
   - Click "Run workflow"

### Workflow Features

- ✅ Automated testing on every push
- ✅ Builds and tests before deployment
- ✅ Deploys only on `main` or `dev` branches
- ✅ Environment-specific deployments
- ✅ Automatic rollback on failure
- ✅ Deployment logs and verification

## VPS Configuration

### Current Setup

**Service:** `smrtmart.service` (systemd)
**Location:** `/opt/smrtmart/`
**Binary:** `server`
**Port:** 8080
**User:** `harvad`
**Logs:** `/var/log/smrtmart/` and systemd journal

### Service Configuration File

```ini
[Unit]
Description=SmartMart Go Backend API
After=network.target

[Service]
Type=simple
User=harvad
Group=harvad
WorkingDirectory=/opt/smrtmart
EnvironmentFile=/opt/smrtmart/.env
ExecStart=/opt/smrtmart/server
Restart=always
RestartSec=10
StandardOutput=append:/var/log/smrtmart/stdout.log
StandardError=append:/var/log/smrtmart/stderr.log

[Install]
WantedBy=multi-user.target
```

## Useful Commands

### Local Development

```bash
# Build for local testing
make build

# Run locally
make run

# Run tests
make test

# Build for Linux VPS
GOOS=linux GOARCH=amd64 go build -o server ./cmd/server
```

### VPS Management

```bash
# SSH into VPS
ssh harvad@107.175.235.220

# Check service status
sudo systemctl status smrtmart

# View live logs
sudo journalctl -u smrtmart -f

# Restart service
sudo systemctl restart smrtmart

# Stop service
sudo systemctl stop smrtmart

# Start service
sudo systemctl start smrtmart

# View recent logs (last 50 lines)
sudo journalctl -u smrtmart -n 50 --no-pager

# View error logs
sudo journalctl -u smrtmart -p err -n 20

# Check which binary is running
ps aux | grep server
```

### Quick Checks

```bash
# Test API endpoint
curl http://107.175.235.220:8080/api/v1/health

# Check if port is listening
netstat -tulpn | grep 8080

# View environment variables
cat /opt/smrtmart/.env

# List recent backups
ls -lth /opt/smrtmart/*.backup.* | head -5
```

## Deployment Process Flow

```
1. Build Application (GOOS=linux GOARCH=amd64)
   ↓
2. Create Backup (server → server.backup.TIMESTAMP)
   ↓
3. Stop Service (systemctl stop smrtmart)
   ↓
4. Upload New Binary
   ↓
5. Upload Static Files/Templates
   ↓
6. Set Permissions (chmod +x)
   ↓
7. Start Service (systemctl start smrtmart)
   ↓
8. Health Check (systemctl is-active smrtmart)
   ↓
9. Success! ✅  OR  Rollback ↻
```

## Rollback Procedure

### Automatic Rollback

The deployment script automatically rolls back if the service fails to start.

### Manual Rollback

```bash
# SSH into VPS
ssh harvad@107.175.235.220

# List available backups
ls -lth /opt/smrtmart/*.backup.*

# Restore from backup (replace TIMESTAMP)
sudo systemctl stop smrtmart
cp /opt/smrtmart/server.backup.TIMESTAMP /opt/smrtmart/server
chmod +x /opt/smrtmart/server
sudo systemctl start smrtmart

# Verify
sudo systemctl status smrtmart
```

## Troubleshooting

### Issue: Build Fails Locally

```bash
# Install Go dependencies
go mod download
go mod tidy

# Try building again
go build -o server ./cmd/server
```

### Issue: SSH Connection Fails

```bash
# Test connection
ssh harvad@107.175.235.220

# If password fails, check VPS SSH configuration
# Ensure PasswordAuthentication is enabled in /etc/ssh/sshd_config
```

### Issue: Service Won't Start

```bash
# Check detailed logs
ssh harvad@107.175.235.220 'sudo journalctl -u smrtmart -n 100 --no-pager'

# Check if port is already in use
ssh harvad@107.175.235.220 'netstat -tulpn | grep 8080'

# Check environment file
ssh harvad@107.175.235.220 'cat /opt/smrtmart/.env'

# Test binary directly
ssh harvad@107.175.235.220 'cd /opt/smrtmart && ./server'
```

### Issue: Permission Denied

```bash
# Fix binary permissions
ssh harvad@107.175.235.220 'chmod +x /opt/smrtmart/server'

# Fix ownership
ssh harvad@107.175.235.220 'sudo chown harvad:harvad /opt/smrtmart/server'
```

### Issue: Database Connection Fails

```bash
# Check database credentials in .env
ssh harvad@107.175.235.220 'grep DB_ /opt/smrtmart/.env'

# Test database connection
ssh harvad@107.175.235.220 'psql -h <host> -U <user> -d <database> -c "SELECT 1;"'
```

## Security Best Practices

### 1. Use SSH Keys (Recommended)

```bash
# Run setup script
./setup-ssh-keys.sh

# Use secure deployment
./deploy-vps-secure.sh
```

### 2. GitHub Actions Secrets

For GitHub Actions with SSH keys:

```bash
# Generate SSH key pair
ssh-keygen -t rsa -b 4096 -f ~/.ssh/deploy_key

# Add public key to VPS
ssh-copy-id -i ~/.ssh/deploy_key.pub harvad@107.175.235.220

# Add private key to GitHub Secrets
cat ~/.ssh/deploy_key | base64 -w 0
# Copy output and add as SSH_PRIVATE_KEY secret
```

### 3. Environment Variables

Never commit `.env` files! Use GitHub Secrets or VPS environment files.

### 4. Regular Updates

```bash
# Keep VPS updated
ssh harvad@107.175.235.220 'sudo apt update && sudo apt upgrade -y'
```

## Performance Monitoring

### Resource Usage

```bash
# Check CPU and memory
ssh harvad@107.175.235.220 'top -bn1 | grep server'

# Check disk space
ssh harvad@107.175.235.220 'df -h /opt/smrtmart'

# Check running processes
ssh harvad@107.175.235.220 'ps aux | grep server'
```

### Log Management

```bash
# Rotate old logs
ssh harvad@107.175.235.220 'sudo journalctl --vacuum-time=7d'

# Check log size
ssh harvad@107.175.235.220 'du -sh /var/log/smrtmart/'
```

## Backup Strategy

### Automated Backups

Every deployment creates a timestamped backup:
- Location: `/opt/smrtmart/server.backup.YYYYMMDD_HHMMSS`
- Retention: Manual cleanup recommended (keep last 5-10)

### Manual Backup

```bash
# Create backup
ssh harvad@107.175.235.220 'cp /opt/smrtmart/server /opt/smrtmart/server.backup.manual.$(date +%Y%m%d)'

# Backup entire directory
ssh harvad@107.175.235.220 'tar -czf /tmp/smrtmart-backup-$(date +%Y%m%d).tar.gz /opt/smrtmart'

# Download backup locally
scp harvad@107.175.235.220:/tmp/smrtmart-backup-*.tar.gz ./backups/
```

### Cleanup Old Backups

```bash
# List backups by size
ssh harvad@107.175.235.220 'ls -lhS /opt/smrtmart/*.backup.*'

# Remove backups older than 30 days
ssh harvad@107.175.235.220 'find /opt/smrtmart -name "*.backup.*" -mtime +30 -delete'

# Keep only last 5 backups
ssh harvad@107.175.235.220 'cd /opt/smrtmart && ls -t *.backup.* | tail -n +6 | xargs rm -f'
```

## Advanced Usage

### Custom Build Flags

Edit `deploy-vps.sh`:

```bash
# Add build optimization
GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o server ./cmd/server
```

### Deploy to Multiple Environments

```bash
# Create environment-specific configs
cp deploy-vps.sh deploy-staging.sh
cp deploy-vps.sh deploy-production.sh

# Edit VPS_HOST in each file
```

### Pre/Post Deployment Hooks

Add to deployment scripts:

```bash
# Before deployment
echo "Running pre-deployment checks..."
./scripts/pre-deploy.sh

# After deployment
echo "Running post-deployment tasks..."
./scripts/post-deploy.sh
```

## Support & Resources

- **Documentation:** This file
- **Issues:** https://github.com/bluehawana/smrtmart-backend-go-racknerd/issues
- **VPS Provider:** RackNerd
- **Backend Repository:** https://github.com/bluehawana/smrtmart-backend-go-racknerd

## Checklist for New Deployments

- [ ] Code changes committed to Git
- [ ] Tests passing locally (`make test`)
- [ ] Environment variables updated on VPS (if needed)
- [ ] Database migrations run (if needed)
- [ ] Backup created
- [ ] Service health verified after deployment
- [ ] API endpoints tested
- [ ] Logs checked for errors

## Version History

| Date | Version | Changes |
|------|---------|---------|
| 2025-10-25 | 1.0 | Initial deployment automation setup |

---

**Last Updated:** October 25, 2025
**Maintained By:** SmrtMart Team
