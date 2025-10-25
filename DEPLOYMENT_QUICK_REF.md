# SmrtMart Deployment - Quick Reference

## 🚀 Deploy Now

```bash
./deploy-vps.sh
```

## 📋 Common Commands

### Deploy
```bash
./deploy-vps.sh              # Deploy with password
./deploy-vps-secure.sh       # Deploy with SSH keys (after setup)
./test-deployment.sh         # Test before deploying
```

### VPS Management
```bash
# Check status
ssh harvad@107.175.235.220 'sudo systemctl status smrtmart'

# View logs (live)
ssh harvad@107.175.235.220 'sudo journalctl -u smrtmart -f'

# Restart service
ssh harvad@107.175.235.220 'sudo systemctl restart smrtmart'

# View recent errors
ssh harvad@107.175.235.220 'sudo journalctl -u smrtmart -p err -n 20'
```

### Quick Checks
```bash
# Test API
curl http://107.175.235.220:8080/api/v1/health

# Check running process
ssh harvad@107.175.235.220 'ps aux | grep server'

# Check port
ssh harvad@107.175.235.220 'netstat -tulpn | grep 8080'
```

## 🔄 Rollback

```bash
# List backups
ssh harvad@107.175.235.220 'ls -lth /opt/smrtmart/*.backup.*'

# Restore (replace TIMESTAMP)
ssh harvad@107.175.235.220 'sudo systemctl stop smrtmart && \
  cp /opt/smrtmart/server.backup.TIMESTAMP /opt/smrtmart/server && \
  sudo systemctl start smrtmart'
```

## 🔧 Troubleshooting

### Service won't start
```bash
ssh harvad@107.175.235.220 'sudo journalctl -u smrtmart -n 50'
```

### Database issues
```bash
ssh harvad@107.175.235.220 'cat /opt/smrtmart/.env | grep DB_'
```

### Permission issues
```bash
ssh harvad@107.175.235.220 'ls -la /opt/smrtmart/server'
ssh harvad@107.175.235.220 'chmod +x /opt/smrtmart/server'
```

## 📊 Monitoring

```bash
# CPU/Memory usage
ssh harvad@107.175.235.220 'top -bn1 | head -20'

# Disk space
ssh harvad@107.175.235.220 'df -h'

# Service uptime
ssh harvad@107.175.235.220 'systemctl status smrtmart | grep Active'
```

## 🔐 First Time Setup

1. **Setup SSH Keys (Optional)**
   ```bash
   ./setup-ssh-keys.sh
   ```

2. **Configure GitHub Secrets** (for CI/CD)
   - `VPS_HOST`: 107.175.235.220
   - `VPS_USER`: harvad
   - `VPS_PASSWORD`: your_password

3. **Setup VPS Environment**
   - See `VPS_SETUP_GUIDE.md` for details

## 📁 Project Structure

```
heroku-backend/
├── deploy-vps.sh              # Main deployment script
├── deploy-vps-secure.sh       # Secure deployment (SSH keys)
├── setup-ssh-keys.sh          # SSH setup helper
├── test-deployment.sh         # Pre-deployment tests
├── VPS_DEPLOYMENT_GUIDE.md    # Full documentation
├── VPS_SETUP_GUIDE.md         # VPS configuration guide
├── .github/workflows/
│   └── deploy.yml             # GitHub Actions CI/CD
└── cmd/server/main.go         # Application entry point
```

## 🌐 URLs

- **API Endpoint**: http://107.175.235.220:8080
- **Health Check**: http://107.175.235.220:8080/api/v1/health
- **GitHub Repo**: https://github.com/bluehawana/smrtmart-backend-go-racknerd

## 💡 Tips

- Always run `./test-deployment.sh` before deploying
- Keep at least 5 backup files
- Check logs after every deployment
- Use SSH keys for better security
- Enable GitHub Actions for automated deployments

## 📚 Documentation

- **Full Guide**: `VPS_DEPLOYMENT_GUIDE.md`
- **VPS Setup**: `VPS_SETUP_GUIDE.md`
- **Main README**: `README.md`

## 🆘 Emergency Contacts

- **GitHub Issues**: https://github.com/bluehawana/smrtmart-backend-go-racknerd/issues
- **VPS Provider**: RackNerd Support

---

**Last Updated**: October 25, 2025
