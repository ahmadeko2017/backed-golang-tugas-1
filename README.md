# Tugas 1 - Category API

This is a RESTful API implementation for managing Categories, built using Go, Gin functionality, and GORM with SQLite.

## 🚀 Features

- **CRUD Operations**: Create, Read, Update, and Delete categories.
- **Database**: SQLite for easy local development (no additional setup required).
- **Documentation**: Integrated Swagger UI.
- **Rate Limiting**: Protected endpoints with a rate limit of 10 requests per second.

## 🛠️ Tech Stack

- **Language**: [Go](https://go.dev/)
- **Framework**: [Gin](https://github.com/gin-gonic/gin)
- **ORM**: [GORM](https://gorm.io/)
- **Database Driver**: [glebarez/sqlite](https://github.com/glebarez/sqlite) (Pure Go implementation)
- **Docs**: [Swag](https://github.com/swaggo/swag)

## 📦 Installation & Running

1.  **Clone the repository** (if you haven't already):
    ```bash
    git clone <repository-url>
    cd <repository-folder>
    ```

2.  **Install Dependencies**:
    ```bash
    go mod tidy
    ```

3.  **Run the Application**:
    ```bash
    go run cmd/api/main.go
    ```

    The server will start on port `8080`.

## 🚀 Deployment (Release Mode)

To run the application in **Release Mode** (Production), follow these steps:

### 1. Build the Binary
Compiling the code into a single executable file makes it faster and easier to deploy.

```bash
go build -o app.exe cmd/api/main.go
```

### 2. Set Environment Variables
Set the `GIN_MODE` environment variable to `release` to disable debug logs and optimize performance.

**PowerShell (Windows):**
```powershell
$env:GIN_MODE="release"
./app.exe
```

**Command Prompt (cmd):**
```cmd
set GIN_MODE=release
app.exe
```

**Linux / Mac:**
```bash
export GIN_MODE=release
./app.exe
```

> [!NOTE]
> In release mode, the verbose debug logs from Gin will be hidden, and the server will be optimized for performance.

## 📖 API Documentation

Once the server is running, you can access the interactive Swagger documentation at:

👉 **[http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html)**

## 🔗 Endpoints

| Method | Endpoint | Description |
| :--- | :--- | :--- |
| `GET` | `/categories` | Get all categories |
| `POST` | `/categories` | Create a new category |
| `GET` | `/categories/:id` | Get details of a specific category |
| `PUT` | `/categories/:id` | Update a category |
| `DELETE` | `/categories/:id` | Delete a category |

## 🧪 Testing

You can test the endpoints using **Swagger UI**, **Postman**, or **curl**.

**Example (Create Category):**
```bash
curl -X POST http://localhost:8080/categories \
  -H "Content-Type: application/json" \
  -d '{"name": "Electronics", "description": "Gadgets and devices"}'
```
