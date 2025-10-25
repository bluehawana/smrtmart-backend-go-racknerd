# VPS Environment Setup Guide

## Initial VPS Configuration

### 1. Configure VPS User for Passwordless Sudo (Optional but Recommended)

This allows the deployment script to restart services without password prompts.

```bash
# SSH into your VPS
ssh harvad@107.175.235.220

# Edit sudoers file
sudo visudo

# Add this line at the end (allows harvad to run systemctl without password)
harvad ALL=(ALL) NOPASSWD: /bin/systemctl start smrtmart, /bin/systemctl stop smrtmart, /bin/systemctl restart smrtmart, /bin/systemctl status smrtmart, /usr/bin/journalctl

# Save and exit (Ctrl+X, then Y, then Enter)
```

### 2. Set Up Environment Variables

```bash
# SSH into VPS
ssh harvad@107.175.235.220

# Navigate to deployment directory
cd /opt/smrtmart

# Create .env file from example
# Either copy from local:
scp .env.example harvad@107.175.235.220:/opt/smrtmart/.env

# Or create manually:
nano /opt/smrtmart/.env
```

**Required Environment Variables:**

```bash
# Database (UPDATE THESE!)
DB_HOST=your-database-host
DB_PORT=5432
DB_USER=your-db-user
DB_PASSWORD=your-secure-password
DB_NAME=smrtmart_db
DB_SSLMODE=require

# Server
PORT=8080
GIN_MODE=release
CORS_ORIGINS=https://yourdomain.com

# Security
JWT_SECRET=generate-a-long-random-string-here

# Stripe (if using payments)
STRIPE_SECRET_KEY=sk_live_your_key
STRIPE_WEBHOOK_SECRET=whsec_your_secret
```

### 3. Set Correct Permissions

```bash
# Set ownership
sudo chown -R harvad:harvad /opt/smrtmart

# Protect .env file
chmod 600 /opt/smrtmart/.env

# Make binary executable
chmod +x /opt/smrtmart/server

# Create log directory if it doesn't exist
sudo mkdir -p /var/log/smrtmart
sudo chown harvad:harvad /var/log/smrtmart
```

### 4. Verify Service Configuration

```bash
# Check if service exists
systemctl status smrtmart

# If not, create it:
sudo nano /etc/systemd/system/smrtmart.service
```

Paste the following:

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
StartLimitInterval=200
StartLimitBurst=5
StandardOutput=append:/var/log/smrtmart/stdout.log
StandardError=append:/var/log/smrtmart/stderr.log
SyslogIdentifier=smrtmart

[Install]
WantedBy=multi-user.target
```

```bash
# Reload systemd
sudo systemctl daemon-reload

# Enable service to start on boot
sudo systemctl enable smrtmart

# Start service
sudo systemctl start smrtmart

# Check status
sudo systemctl status smrtmart
```

### 5. Configure Firewall (if applicable)

```bash
# Allow SSH
sudo ufw allow 22/tcp

# Allow HTTP/HTTPS
sudo ufw allow 80/tcp
sudo ufw allow 443/tcp

# Allow application port (if not behind reverse proxy)
sudo ufw allow 8080/tcp

# Enable firewall
sudo ufw enable
```

### 6. Set Up Reverse Proxy (Optional - Recommended for Production)

#### Using Nginx:

```bash
# Install Nginx
sudo apt update
sudo apt install nginx -y

# Create site configuration
sudo nano /etc/nginx/sites-available/smrtmart
```

Paste:

```nginx
server {
    listen 80;
    server_name yourdomain.com www.yourdomain.com;

    location / {
        proxy_pass http://localhost:8080;
        proxy_http_version 1.1;
        proxy_set_header Upgrade $http_upgrade;
        proxy_set_header Connection 'upgrade';
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto $scheme;
        proxy_cache_bypass $http_upgrade;
    }

    # API specific settings
    location /api/ {
        proxy_pass http://localhost:8080;
        proxy_http_version 1.1;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto $scheme;

        # Increase timeouts for long requests
        proxy_connect_timeout 60s;
        proxy_send_timeout 60s;
        proxy_read_timeout 60s;
    }
}
```

```bash
# Enable site
sudo ln -s /etc/nginx/sites-available/smrtmart /etc/nginx/sites-enabled/

# Test configuration
sudo nginx -t

# Reload Nginx
sudo systemctl reload nginx
```

#### Add SSL with Let's Encrypt:

```bash
# Install Certbot
sudo apt install certbot python3-certbot-nginx -y

# Get certificate
sudo certbot --nginx -d yourdomain.com -d www.yourdomain.com

# Auto-renewal is set up automatically
# Test renewal:
sudo certbot renew --dry-run
```

## Database Setup

### PostgreSQL on VPS:

```bash
# Install PostgreSQL
sudo apt install postgresql postgresql-contrib -y

# Create database and user
sudo -u postgres psql

# Inside PostgreSQL prompt:
CREATE DATABASE smrtmart_db;
CREATE USER smrtmart_user WITH ENCRYPTED PASSWORD 'your_secure_password';
GRANT ALL PRIVILEGES ON DATABASE smrtmart_db TO smrtmart_user;
\q

# Update .env file with these credentials
```

### Remote Database (Recommended):

Use a managed database service like:
- Supabase (PostgreSQL)
- AWS RDS
- DigitalOcean Managed Database
- Railway
- Render

Update your `.env` file with the connection string provided.

## Health Checks

### Verify Everything Works:

```bash
# Check service status
sudo systemctl status smrtmart

# Test API endpoint
curl http://localhost:8080/api/v1/health

# Check logs
sudo journalctl -u smrtmart -f

# Check database connection
# From your Go app logs, you should see "Database connected successfully"
```

## Monitoring Setup (Optional)

### Basic Monitoring:

```bash
# Install monitoring tools
sudo apt install htop iotop nethogs -y

# Monitor in real-time
htop
```

### Log Rotation:

```bash
# Create logrotate config
sudo nano /etc/logrotate.d/smrtmart
```

Paste:

```
/var/log/smrtmart/*.log {
    daily
    rotate 14
    compress
    delaycompress
    notifempty
    create 0644 harvad harvad
    sharedscripts
    postrotate
        systemctl reload smrtmart > /dev/null 2>&1 || true
    endscript
}
```

## Security Checklist

- [ ] Changed default SSH port (optional)
- [ ] Disabled root login via SSH
- [ ] Set up SSH key authentication
- [ ] Configured firewall (UFW)
- [ ] Set secure file permissions on .env
- [ ] Using SSL/TLS for API (HTTPS)
- [ ] Database uses strong password
- [ ] JWT_SECRET is long and random
- [ ] Regular security updates enabled
- [ ] Fail2ban configured (optional)
- [ ] Backups automated

## Backup Automation

```bash
# Create backup script
nano ~/backup-smrtmart.sh
```

Paste:

```bash
#!/bin/bash
BACKUP_DIR="/home/harvad/backups"
DATE=$(date +%Y%m%d_%H%M%S)

mkdir -p $BACKUP_DIR

# Backup binary and configs
tar -czf $BACKUP_DIR/smrtmart-$DATE.tar.gz /opt/smrtmart

# Backup database (if local)
sudo -u postgres pg_dump smrtmart_db > $BACKUP_DIR/db-$DATE.sql

# Keep only last 7 days
find $BACKUP_DIR -name "smrtmart-*.tar.gz" -mtime +7 -delete
find $BACKUP_DIR -name "db-*.sql" -mtime +7 -delete

echo "Backup completed: $DATE"
```

```bash
# Make executable
chmod +x ~/backup-smrtmart.sh

# Add to crontab (daily at 2 AM)
crontab -e

# Add this line:
0 2 * * * /home/harvad/backup-smrtmart.sh
```

## Troubleshooting Common Issues

### Issue: Service won't start

```bash
# Check detailed logs
sudo journalctl -u smrtmart -xe

# Check if binary is executable
ls -la /opt/smrtmart/server

# Test binary directly
cd /opt/smrtmart
./server
```

### Issue: Database connection fails

```bash
# Verify environment variables
cat /opt/smrtmart/.env | grep DB_

# Test database connection
psql -h $DB_HOST -U $DB_USER -d $DB_NAME
```

### Issue: Permission denied errors

```bash
# Fix ownership
sudo chown -R harvad:harvad /opt/smrtmart

# Fix permissions
chmod 755 /opt/smrtmart
chmod 600 /opt/smrtmart/.env
chmod +x /opt/smrtmart/server
```

## Ready to Deploy?

Once you've completed the setup:

```bash
# From your local machine
cd /mnt/c/Users/BLUEH/projects/smrmart/heroku-backend

# Test deployment
./test-deployment.sh

# Deploy!
./deploy-vps.sh
```

---

**Need Help?** Check the main deployment guide: `VPS_DEPLOYMENT_GUIDE.md`
