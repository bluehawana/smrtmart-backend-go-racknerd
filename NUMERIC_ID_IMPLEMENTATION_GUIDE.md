# Numeric Product ID Implementation - Complete Guide

**Date:** 2025-10-25
**Status:** ✅ CODE UPDATED - AWAITING SUPABASE MIGRATION & DEPLOYMENT

---

## 🎯 What Was Fixed

You mentioned "many times" that you wanted numeric product IDs instead of UUIDs. This has now been fully implemented!

### Before (UUID):
```json
{
  "id": "550e8400-e29b-41d4-a716-446655440028",
  "name": "Product Name",
  "price": 199.99
}
```

### After (Numeric):
```json
{
  "id": 1,
  "name": "Product Name",
  "price": 199.99
}
```

---

## ✅ Backend Code Changes (COMPLETED)

### 1. **Product Model Updated** (`internal/models/models.go`)
```go
type Product struct {
    ID          uuid.UUID     `json:"-" db:"id"`                    // Hidden from JSON
    NumericID   int           `json:"id" db:"numeric_id"`           // Exposed as "id"
    VendorID    uuid.UUID     `json:"vendor_id" db:"vendor_id"`
    Name        string        `json:"name" db:"name"`
    // ... rest of fields
}
```

### 2. **Repository Updated** (`internal/repository/product_repository.go`)
✅ All SELECT queries now include `numeric_id`
✅ All Scan operations populate `NumericID` field
✅ Create function returns auto-generated `numeric_id`
✅ `GetByNumericID(int)` function exists for numeric lookups

### 3. **Git Commit**
```
Commit: 99e75b2
Message: "Switch products from UUID to numeric IDs for cleaner API responses"
Status: Committed locally
Push: In progress to git@github.com:bluehawana/smrtmart-backend-go-racknerd.git
```

---

## 🔧 REQUIRED: Supabase Database Migration

### Step 1: Run Migration SQL

**File Location:** `/mnt/c/Users/BLUEH/projects/smrmart/heroku-backend/supabase_numeric_id_migration.sql`

**Instructions:**
1. Go to Supabase Dashboard: https://supabase.com/dashboard
2. Select your project
3. Click **SQL Editor** in left sidebar
4. Click **New Query**
5. Copy the contents of `supabase_numeric_id_migration.sql`
6. Paste into SQL Editor
7. Click **Run** (or press Cmd/Ctrl+Enter)

### Step 2: Verify Migration Success

Run this query in Supabase SQL Editor:
```sql
SELECT
    id,
    numeric_id,
    name,
    created_at
FROM products
ORDER BY numeric_id
LIMIT 10;
```

**Expected Result:**
- All products should have `numeric_id` values (1, 2, 3, 4, etc.)
- Newest products get highest numbers
- No NULL values in `numeric_id` column

---

## 🚀 REQUIRED: Deploy to RackNerd VPS

After running the Supabase migration, you need to deploy the updated Go backend.

### Option 1: SSH into VPS and Pull Latest Code

```bash
# SSH into your RackNerd VPS
ssh your-user@api.smrtmart.com

# Navigate to backend directory
cd /path/to/smrtmart-backend

# Pull latest changes from GitHub
git pull origin main

# Rebuild the Go application
go build -o smrtmart-api ./cmd/api

# Restart the service (exact command depends on your setup)
sudo systemctl restart smrtmart-backend
# OR
./smrtmart-api
```

### Option 2: Use Your Existing Deployment Script

If you have a deployment script or CI/CD pipeline:
```bash
# Run your deployment script
./deploy.sh
```

### Verify Deployment

After deployment, test the API:
```bash
curl "https://api.smrtmart.com/api/v1/products?limit=1"
```

**Expected Response:**
```json
{
  "success": true,
  "data": [
    {
      "id": 1,  // ✅ Numeric ID instead of UUID!
      "name": "Product Name",
      "price": 199.99,
      ...
    }
  ]
}
```

---

## 📋 Migration Script Details

The migration script (`supabase_numeric_id_migration.sql`) does the following:

1. ✅ **Adds `numeric_id` column** to products table
2. ✅ **Creates auto-increment sequence** (`product_numeric_id_seq`)
3. ✅ **Populates existing products** with sequential IDs based on creation date
4. ✅ **Sets up auto-increment** for new products
5. ✅ **Adds unique constraint** on `numeric_id`
6. ✅ **Creates index** for fast lookups
7. ✅ **Idempotent** - safe to run multiple times

---

## 🧪 Testing After Migration & Deployment

### Test 1: List Products
```bash
curl "https://api.smrtmart.com/api/v1/products"
```
**Expected:** Products have numeric IDs (1, 2, 3...)

### Test 2: Get Single Product by Numeric ID
```bash
curl "https://api.smrtmart.com/api/v1/products/1"
```
**Expected:** Returns product with ID 1

### Test 3: Create New Product
```bash
curl -X POST "https://api.smrtmart.com/api/v1/products" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Test Product",
    "description": "Testing numeric ID",
    "price": 99.99,
    "category": "test",
    "stock": 10,
    "vendor_id": "YOUR_VENDOR_UUID"
  }'
```
**Expected:** Returns new product with auto-incremented numeric ID

---

## 📝 What Happens Next

### Automatic Behavior After Migration:

1. **Existing Products:** Will have numeric IDs 1, 2, 3, 4... based on creation date
2. **New Products:** Will automatically get next available numeric ID (e.g., 15, 16, 17...)
3. **API Responses:** Will return `"id": 1` instead of `"id": "uuid-string"`
4. **Frontend URLs:** Can use `/products/1` instead of `/products/550e8400-...`

### Database Schema:

```sql
CREATE TABLE products (
    id UUID PRIMARY KEY,              -- Still used internally for relationships
    numeric_id INTEGER UNIQUE NOT NULL DEFAULT nextval('product_numeric_id_seq'),
    vendor_id UUID NOT NULL,
    name VARCHAR NOT NULL,
    -- ... other fields
);
```

---

## ⚠️ Important Notes

### UUID Not Deleted
- UUID `id` column is still the PRIMARY KEY
- Used internally for foreign key relationships (orders, reviews, etc.)
- Just hidden from JSON API responses with `json:"-"` tag

### Frontend Changes
After deployment, your frontend can:
- ✅ Use numeric IDs in URLs: `/products/1`
- ✅ Display cleaner product IDs to users
- ✅ Share easier-to-remember product links

### Backward Compatibility
- The `GetByID(uuid)` function still works for internal use
- The `GetByNumericID(int)` function is used for public API

---

## 🔍 Troubleshooting

### If API Still Returns UUIDs After Deployment

**Cause:** Migration not run on Supabase
**Fix:** Run `supabase_numeric_id_migration.sql` in Supabase SQL Editor

### If API Returns Error "column numeric_id does not exist"

**Cause:** Migration not run on Supabase
**Fix:** Run `supabase_numeric_id_migration.sql` in Supabase SQL Editor

### If New Products Don't Get Numeric IDs

**Cause:** Sequence not set up correctly
**Fix:** Check Step 4 of migration script ran successfully:
```sql
SELECT setval('product_numeric_id_seq', COALESCE((SELECT MAX(numeric_id) FROM products), 0) + 1, false);
```

### If Numeric IDs Have Gaps

**Behavior:** This is normal and expected
**Reason:** Auto-increment sequences can have gaps due to rollbacks or deleted products
**Solution:** No action needed - gaps are acceptable

---

## ✅ Checklist

Complete these steps in order:

- [x] Backend code updated (models.go, product_repository.go)
- [x] Changes committed to Git
- [ ] Changes pushed to GitHub (in progress)
- [ ] Run `supabase_numeric_id_migration.sql` in Supabase SQL Editor
- [ ] Verify migration with SELECT query
- [ ] Pull latest code on RackNerd VPS
- [ ] Rebuild and restart backend service
- [ ] Test API returns numeric IDs: `curl https://api.smrtmart.com/api/v1/products`
- [ ] Test frontend loads products with numeric IDs
- [ ] Update frontend to use numeric IDs in product URLs (optional)

---

## 📞 Next Steps

1. **Run Supabase Migration** - This is the critical step!
2. **Deploy to RackNerd** - Pull latest code and restart service
3. **Test API** - Verify numeric IDs are returned
4. **Test Frontend** - Check products load correctly

After completing these steps, products will use numeric IDs everywhere!

---

**Generated:** 2025-10-25
**Backend Repo:** git@github.com:bluehawana/smrtmart-backend-go-racknerd.git
**Commit:** 99e75b2
