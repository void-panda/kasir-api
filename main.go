package main

import (
	"fmt"
	"net/http"
	"os"

	_ "kasir-api/docs"
	"kasir-api/handler"

	httpSwagger "github.com/swaggo/http-swagger"
)

// @title Kasir API
// @version 1.0
// @host localhost:8080
// @BasePath /

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("API ready!"))
	})

	// Register routes
	http.HandleFunc("GET /api/products", handler.GetProducts)
	http.HandleFunc("GET /api/products/{id}", handler.GetProductByID)
	http.HandleFunc("POST /api/products", handler.CreateProduct)
	http.HandleFunc("PUT /api/products/{id}", handler.UpdateProduct)
	http.HandleFunc("DELETE /api/products/{id}", handler.DeleteProduct)

	http.HandleFunc("GET /api/categories", handler.GetCategories)
	http.HandleFunc("GET /api/categories/{id}", handler.GetCategoryByID)
	http.HandleFunc("POST /api/categories", handler.CreateCategory)
	http.HandleFunc("PUT /api/categories/{id}", handler.UpdateCategory)
	http.HandleFunc("DELETE /api/categories/{id}", handler.DeleteCategory)

	http.HandleFunc("GET /api/tasks", handler.GetAllTasks)
	http.HandleFunc("GET /api/tasks/{id}", handler.GetTaskByID)
	http.HandleFunc("POST /api/tasks", handler.CreateTask)
	http.HandleFunc("PATCH /api/tasks/{id}", handler.UpdateTask)
	http.HandleFunc("DELETE /api/tasks/{id}", handler.DeleteTask)

	http.HandleFunc("GET /health", handler.HealthCheck)

	// Swagger
	http.HandleFunc("/swagger/", httpSwagger.WrapHandler)

	// Get port from environment variable or use default
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Printf("🚀 server running on port %s\n", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		panic("gagal running server")
	}
}
