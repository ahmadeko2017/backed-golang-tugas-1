# Deployment Instructions for VM/Production

## Method 1: Auto-Seed (Current Default)
The seeder runs automatically on startup and checks if data exists.

### On VM:
```bash
# Build the application
go build -o main cmd/api/main.go

# Run with auto-seed (controlled by env var)
SEED_DATA=true ./main
```

### First Time Setup:
1. App starts
2. Database creates tables
3. Seeder checks if data exists
4. If empty, inserts sample data
5. If has data, skips

## Method 2: Manual Seed Control

### Skip Seeding:
```bash
# Don't set SEED_DATA or set it to false
./main
# OR
SEED_DATA=false ./main
```

### Force Re-Seed:
```bash
# 1. Stop application
# 2. Delete database
rm tugas1.db

# 3. Start with SEED_DATA=true
SEED_DATA=true ./main
```

## Docker Deployment

### Dockerfile already configured
The existing `Dockerfile` works as-is.

### Run with Docker:
```bash
# Build image
docker build -t category-api .

# Run with auto-seed
docker run -e SEED_DATA=true -p 8080:8080 category-api

# Run without seed
docker run -p 8080:8080 category-api
```

## Production Best Practices

### Option 1: Keep Auto-Seed (Small Apps)
- Good for: Development, staging, small apps
- Seeder checks if data exists before inserting
- No duplicate data

### Option 2: Disable Auto-Seed (Large Apps)
- Good for: Production with real data
- Use migrations or manual SQL for production data
- Set `SEED_DATA=false` or don't set it

### Option 3: Use Migration Tools
For production, consider:
- SQL migration files
- Database backup/restore
- Admin panel for data management
