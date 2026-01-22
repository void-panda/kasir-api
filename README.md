# KASIR API (GOLANG)

## Description
Aplikasi kasir sederhana yang dibangun menggunakan Go dengan arsitektur bersih (Clean Architecture). API menyediakan endpoint untuk mengelola produk dan kategori.

## Fitur
- Get all products / categories
- Get product / category by ID
- Create product / category
- Update product / category by ID
- Delete product / category by ID
- Health check endpoint
- API documentation dengan Swagger/OpenAPI

## Tech Stack
- **Language:** Go 1.25+
- **Framework:** Standard library `net/http`
- **Documentation:** Swagger/OpenAPI via swaggo
- **Data Storage:** In-memory (mock data)

## Project Structure
```
kasir-api/
├── main.go              # Entry point, route registration
├── go.mod              # Go module definition
├── model/              # Data structures
│   ├── product.go      # Product struct
│   ├── category.go     # Category struct
│   └── response.go     # Response wrapper struct
├── handler/            # Business logic & HTTP handlers
│   ├── handlers.go     # HTTP request handlers with Swagger annotations
│   ├── product_service.go   # Product CRUD operations
│   └── category_service.go  # Category CRUD operations
└── docs/               # Swagger documentation (auto-generated)
    ├── docs.go
    ├── swagger.json
    └── swagger.yaml
```

## Installation

### 1. Clone Repository
```bash
git clone <repository-url>
cd kasir-api
```

### 2. Install Dependencies
```bash
go mod download
```

### 3. Install Swagger Tools (Optional, for documentation generation)
```bash
go install github.com/swaggo/swag/cmd/swag@latest
go get -u github.com/swaggo/http-swagger
```

### 4. Generate Swagger Documentation (Optional)
```bash
swag init
```

## Running the Application

### Development Mode
```bash
go run main.go
```

### Build Binary
```bash
go build -o kasir-api main.go
```

### Run Binary
```bash
./kasir-api
```

Server akan berjalan pada `http://localhost:8080`

## Environment Variables
- `PORT` - Port server (default: 8080)

```bash
PORT=3000 go run main.go
```

## API Endpoints

### Health Check
- `GET /health` - Check if server is running

### Products
- `GET /api/products` - Get all products
- `GET /api/products/{id}` - Get product by ID
- `POST /api/products` - Create new product
- `PUT /api/products/{id}` - Update product by ID
- `DELETE /api/products/{id}` - Delete product by ID

### Categories
- `GET /api/categories` - Get all categories
- `GET /api/categories/{id}` - Get category by ID
- `POST /api/categories` - Create new category
- `PUT /api/categories/{id}` - Update category by ID
- `DELETE /api/categories/{id}` - Delete category by ID

### Swagger Documentation
- `GET /swagger/` - Swagger UI
- `GET /swagger/doc.json` - OpenAPI specification

## API Documentation

### Product Model
```go
type Product struct {
	ID    int    `json:"id"`        // Auto-generated
	Name  string `json:"name"`      // Required
	Price int    `json:"price"`     // Required
	Stock int    `json:"stock"`     // Required
}
```

### Category Model
```go
type Category struct {
	ID          int    `json:"id"`          // Auto-generated
	Name        string `json:"name"`        // Required
	Description string `json:"description"` // Required
}
```

### Response Format
```go
type Response struct {
	Status  string      `json:"status"`            // "ok" or "error"
	Message string      `json:"message"`           // Response message
	Data    interface{} `json:"data,omitempty"`   // Response data (optional)
}
```

## Example Requests

### Get All Products
```bash
curl http://localhost:8080/api/products
```

### Get Product by ID
```bash
curl http://localhost:8080/api/products/1
```

### Create Product
```bash
curl -X POST http://localhost:8080/api/products \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Product Name",
    "price": 10000,
    "stock": 5
  }'
```

### Update Product
```bash
curl -X PUT http://localhost:8080/api/products/1 \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Updated Name",
    "price": 15000,
    "stock": 10
  }'
```

### Delete Product
```bash
curl -X DELETE http://localhost:8080/api/products/1
```

### Create Category
```bash
curl -X POST http://localhost:8080/api/categories \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Category Name",
    "description": "Category Description"
  }'
```

## Code Architecture

### main.go
- Application entry point
- Route registration menggunakan `http.HandleFunc()`
- Server startup dengan port configuration
- Imports: handler package untuk business logic

### handler/ Package
- `handlers.go` - HTTP request handlers dengan Swagger/OpenAPI annotations
- `product_service.go` - Product business logic (GetAllProducts, GetProduct, AddProduct, UpdateProductData, DeleteProductData)
- `category_service.go` - Category business logic (GetAllCategories, GetCategory, AddCategory, UpdateCategoryData, DeleteCategoryData)

### model/ Package
- Data structures untuk Product, Category, dan Response
- JSON serialization/deserialization

## Development Notes

### Adding New Endpoints
1. Create handler function in `handler/handlers.go`
2. Add Swagger/OpenAPI annotations as godoc comments
3. Register route in `main.go`
4. Run `swag init` to regenerate documentation

### Modifying Business Logic
1. Update logic in `handler/product_service.go` atau `handler/category_service.go`
2. No need to restart if using hot reload tools
3. Rebuild binary: `go build -o kasir-api main.go`

### Swagger Annotations
Gunakan format godoc di atas setiap handler function:
```go
// FunctionName godoc
// @Summary Brief description
// @Description Detailed description
// @Tags tag-name
// @Accept json
// @Produce json
// @Param paramName body Model true "Description"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Router /api/endpoint [method]
func FunctionName(w http.ResponseWriter, r *http.Request) {
	// Implementation
}
```

## License
MIT
