# SmrtMart VPS Backend - Current Configuration

**Last Verified:** October 25, 2025
**VPS IP:** 107.175.235.220
**OS:** Debian 12
**User:** harvad

---

## 📍 Deployment Location

```
/opt/smrtmart/
```

**Contents:**
- `server` - Go binary (30MB)
- `.env` - Environment variables
- `static/` - Static files
- `templates/` - Templates
- `*.sql` - Database migration files

**Not a Git Repository:** Files are deployed directly (no git tracking on VPS)

---

## 🚀 Running Services

### Backend API (Go)
- **Port:** 8080
- **Process ID:** 157741 (running)
- **Binary:** `/opt/smrtmart/server`
- **Service:** `smrtmart.service` (systemd)
- **Status:** Active and running
- **Access:** https://api.smrtmart.com

### Other Services
- **Port 8081:** Java (Jenkins)
- **Port 80/443:** Nginx (reverse proxy)
- **Port 22:** SSH

---

## 🔧 Service Management

### Systemd Service: `smrtmart`

**Service File:** `/etc/systemd/system/smrtmart.service`

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

### Common Commands

```bash
# Check status
ssh harvad@107.175.235.220 'systemctl status smrtmart'

# Restart service
ssh harvad@107.175.235.220 'sudo systemctl restart smrtmart'

# Stop service
ssh harvad@107.175.235.220 'sudo systemctl stop smrtmart'

# Start service
ssh harvad@107.175.235.220 'sudo systemctl start smrtmart'

# View logs
ssh harvad@107.175.235.220 'sudo journalctl -u smrtmart -f'

# View recent errors
ssh harvad@107.175.235.220 'sudo journalctl -u smrtmart -p err -n 20'
```

---

## 🌐 Nginx Configuration

### API Subdomain: api.smrtmart.com

**Config:** `/etc/nginx/sites-available/smrtmart-api`

```nginx
server {
    server_name api.smrtmart.com;

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

    listen 443 ssl;
    ssl_certificate /etc/letsencrypt/live/api.smrtmart.com/fullchain.pem;
    ssl_certificate_key /etc/letsencrypt/live/api.smrtmart.com/privkey.pem;
    include /etc/letsencrypt/options-ssl-nginx.conf;
    ssl_dhparam /etc/letsencrypt/ssl-dhparams.pem;
}

server {
    if ($host = api.smrtmart.com) {
        return 301 https://$host$request_uri;
    }

    listen 80;
    server_name api.smrtmart.com;
    return 404;
}
```

**SSL Certificate:** Let's Encrypt (auto-renewal configured)

---

## 🔐 Environment Variables

**Location:** `/opt/smrtmart/.env`

**Key Variables:**
- `DB_HOST` - Database host
- `DB_PORT` - Database port (5432)
- `DB_USER` - Database user
- `DB_PASSWORD` - Database password
- `DB_NAME` - Database name
- `STRIPE_SECRET_KEY` - Stripe live key (configured)
- `STRIPE_WEBHOOK_SECRET` - Webhook secret (configured)
- `PORT` - Server port (8080)
- `CORS_ORIGINS` - Allowed origins

**Security:** File permissions set to 600 (owner read/write only)

---

## 📊 Current Status

### Health Check
```bash
curl https://api.smrtmart.com/api/v1/health
```

**Expected Response:**
```json
{
  "status": "healthy",
  "service": "SmrtMart API v1",
  "version": "1.0.0"
}
```

### Stripe Integration
- ✅ Checkout endpoint working
- ✅ Webhook endpoint configured
- ✅ Live keys active

### Database
- ✅ Connected to Supabase PostgreSQL
- ✅ 21 products loaded
- ✅ All tables migrated

---

## 🚀 Deployment Process

### Current Method: Manual Deployment

Since `/opt/smrtmart/` is not a git repository, deployments are done by:

1. Build locally on your machine
2. Copy binary to VPS via SCP
3. Restart service

### Automated Deployment Available

**Using deployment scripts from repository:**

```bash
# From local machine
cd /mnt/c/Users/BLUEH/projects/smrmart/heroku-backend

# Deploy with password
./deploy-vps.sh

# Or deploy with SSH keys (more secure)
./setup-ssh-keys.sh  # One-time setup
./deploy-vps-secure.sh
```

**Scripts will:**
1. Build Go binary for Linux
2. Create backup of current deployment
3. Stop service
4. Upload new binary
5. Restart service
6. Verify deployment

---

## 📝 Maintenance Tasks

### Regular Checks

```bash
# Check disk space
ssh harvad@107.175.235.220 'df -h'

# Check memory usage
ssh harvad@107.175.235.220 'free -h'

# Check service uptime
ssh harvad@107.175.235.220 'systemctl status smrtmart'

# Check running processes
ssh harvad@107.175.235.220 'ps aux | grep server'

# Check open ports
ssh harvad@107.175.235.220 'netstat -tulpn | grep LISTEN'
```

### Log Management

```bash
# View logs in real-time
ssh harvad@107.175.235.220 'sudo journalctl -u smrtmart -f'

# View last 100 lines
ssh harvad@107.175.235.220 'sudo journalctl -u smrtmart -n 100'

# View errors only
ssh harvad@107.175.235.220 'sudo journalctl -u smrtmart -p err'

# Clean old logs (older than 7 days)
ssh harvad@107.175.235.220 'sudo journalctl --vacuum-time=7d'
```

### Database Backups

The database is on Supabase, which handles automatic backups.

**Binary Backups:**
- Deployment scripts create timestamped backups: `server.backup.YYYYMMDD_HHMMSS`
- Located in `/opt/smrtmart/`

```bash
# List backups
ssh harvad@107.175.235.220 'ls -lth /opt/smrtmart/*.backup.*'

# Restore from backup (if needed)
ssh harvad@107.175.235.220 'sudo systemctl stop smrtmart && \
  cp /opt/smrtmart/server.backup.TIMESTAMP /opt/smrtmart/server && \
  sudo systemctl start smrtmart'
```

---

## 🔒 Security Notes

1. **Firewall:** Configure UFW or use Cloudflare proxy
2. **SSH:** Password authentication currently enabled (consider SSH keys only)
3. **SSL:** Let's Encrypt certificates auto-renew
4. **Environment:** Secrets stored in `.env` file (not in git)
5. **Database:** Hosted on Supabase (separate from VPS)

---

## 📞 Support Commands

### Quick Diagnostics

```bash
# Full system check
ssh harvad@107.175.235.220 '
  echo "=== System Info ===";
  uname -a;
  echo "";
  echo "=== Service Status ===";
  systemctl status smrtmart;
  echo "";
  echo "=== Port Status ===";
  netstat -tulpn | grep 8080;
  echo "";
  echo "=== Recent Logs ===";
  sudo journalctl -u smrtmart -n 10 --no-pager
'
```

### Emergency Restart

```bash
ssh harvad@107.175.235.220 'sudo systemctl restart smrtmart && sleep 2 && systemctl status smrtmart'
```

---

## 🔗 Related Documentation

- `VPS_DEPLOYMENT_GUIDE.md` - Deployment automation
- `VPS_SETUP_GUIDE.md` - Initial VPS setup
- `DEPLOYMENT_QUICK_REF.md` - Quick commands
- `STRIPE_WEBHOOK_GUIDE.md` - Stripe configuration

---

**Last Updated:** October 25, 2025
**Maintained By:** SmrtMart Team
**VPS Provider:** RackNerd
