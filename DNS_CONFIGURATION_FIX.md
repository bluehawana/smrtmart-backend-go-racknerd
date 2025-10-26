# DNS Configuration Fix for smrtmart.com

## 🚨 Current Issue

The `smrtmart.com` A record is pointing to your VPS instead of Vercel, causing:
- ❌ Frontend not loading (404 error)
- ❌ No HTTPS/SSL on main domain
- ❌ Vercel can't serve the frontend

## ✅ Correct DNS Configuration

### In Cloudflare DNS Settings:

**1. Delete the Current A Record**
```
❌ DELETE THIS:
Type: A
Name: smrtmart.com (or @)
IPv4: 107.175.235.220
```

**2. Add CNAME Record for Root Domain**
```
✅ ADD THIS:
Type: CNAME
Name: smrtmart.com (or @)
Target: cname.vercel-dns.com
Proxy Status: DNS only (gray cloud)
TTL: Auto
```

**Note:** Cloudflare supports CNAME flattening, so you CAN use CNAME for the root domain.

**3. Keep the WWW CNAME** (Already Correct ✅)
```
✅ KEEP THIS:
Type: CNAME
Name: www
Target: ae44a10e580bbed8.vercel-dns-017.com
Proxy Status: DNS only
```

**4. Add A Records for API Subdomain**
```
✅ ADD THIS:
Type: A
Name: api
IPv4: 107.175.235.220
Proxy Status: DNS only (or Proxied if you want Cloudflare protection)
TTL: Auto
```

## 📋 Complete DNS Configuration Summary

After the fix, your DNS should look like this:

| Type  | Name         | Target/Value                          | Purpose          |
|-------|--------------|---------------------------------------|------------------|
| CNAME | smrtmart.com | cname.vercel-dns.com                  | Frontend (Vercel)|
| CNAME | www          | ae44a10e580bbed8.vercel-dns-017.com  | Frontend (Vercel)|
| A     | api          | 107.175.235.220                       | Backend API (VPS)|

**Keep your existing records:**
- MX records (for email)
- TXT records (for SPF, DKIM, Vercel verification)

## 🔧 Step-by-Step Fix

### In Cloudflare Dashboard:

1. **Go to DNS settings**
   - Select your domain `smrtmart.com`
   - Click **DNS** tab

2. **Delete the root A record**
   - Find the A record with name `smrtmart.com` or `@`
   - Click **Edit** → **Delete**

3. **Add CNAME for root domain**
   - Click **Add record**
   - Type: `CNAME`
   - Name: `@` (or `smrtmart.com`)
   - Target: `cname.vercel-dns.com`
   - Proxy status: **DNS only** (gray cloud icon)
   - TTL: Auto
   - Click **Save**

4. **Add A record for API subdomain**
   - Click **Add record**
   - Type: `A`
   - Name: `api`
   - IPv4 address: `107.175.235.220`
   - Proxy status: **DNS only** (or Proxied for DDoS protection)
   - TTL: Auto
   - Click **Save**

5. **Wait for DNS propagation** (usually 5-10 minutes with Cloudflare)

## 🧪 Verify the Fix

After making the changes, wait 5-10 minutes, then test:

```bash
# Test main domain (should show Vercel frontend)
curl -I https://smrtmart.com

# Test www subdomain (should show Vercel frontend)
curl -I https://www.smrtmart.com

# Test API subdomain (should show backend)
curl https://api.smrtmart.com/api/v1/health
```

**Expected Results:**
- ✅ `https://smrtmart.com` → Vercel frontend (200 OK)
- ✅ `https://www.smrtmart.com` → Vercel frontend (200 OK)
- ✅ `https://api.smrtmart.com/api/v1/health` → Backend API (200 OK)

## 🌐 Domain Architecture

```
smrtmart.com (CNAME → Vercel)
    ↓
Vercel Frontend (Next.js)
    ↓
Calls API at: https://api.smrtmart.com/api/v1/*
    ↓
api.smrtmart.com (A → 107.175.235.220)
    ↓
VPS Backend (Go API on port 8080)
```

## 🔐 SSL/HTTPS Handling

- **Frontend** (`smrtmart.com`, `www.smrtmart.com`): Vercel handles SSL automatically
- **Backend** (`api.smrtmart.com`): Already has Let's Encrypt SSL via nginx

## ⚠️ Important Notes

1. **CNAME Flattening**: Cloudflare supports CNAME at the root domain level (unlike traditional DNS). This is the recommended approach for Vercel.

2. **Alternative Method** (if CNAME doesn't work):
   - Vercel provides A records you can use instead
   - Check Vercel dashboard for the specific IP addresses
   - Contact Vercel support for their recommended DNS setup

3. **Proxy Status**:
   - For Vercel records: Use **DNS only** (gray cloud)
   - For API subdomain: You can use **Proxied** (orange cloud) for DDoS protection

4. **Propagation Time**:
   - Cloudflare: Usually 5-10 minutes
   - Global propagation: Up to 24-48 hours (but usually much faster)

## 🔍 Troubleshooting

### Issue: "CNAME not allowed at root domain"

If Cloudflare doesn't allow CNAME at root, use Vercel's A records:

1. Go to Vercel Dashboard → Your Project → Settings → Domains
2. Click on `smrtmart.com`
3. Vercel will show you specific A records to use
4. Add those A records in Cloudflare

### Issue: SSL certificate errors

- Wait 10-15 minutes after DNS changes
- Vercel automatically provisions SSL certificates
- Check Vercel dashboard for certificate status

### Issue: Still seeing 404

1. Check DNS propagation: `nslookup smrtmart.com`
2. Clear browser cache
3. Try incognito/private window
4. Verify Vercel deployment is live in Vercel dashboard

## 📞 Quick Reference

**Vercel DNS Documentation:**
https://vercel.com/docs/projects/domains/add-a-domain

**Cloudflare CNAME Flattening:**
https://developers.cloudflare.com/dns/cname-flattening/

**Check DNS Propagation:**
https://dnschecker.org

---

**After fixing DNS, your domains will work like this:**
- `https://smrtmart.com` → Vercel frontend ✅
- `https://www.smrtmart.com` → Vercel frontend ✅
- `https://api.smrtmart.com` → VPS backend ✅
