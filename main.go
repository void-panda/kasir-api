package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"kasir-api/database"
	_ "kasir-api/docs"
	"kasir-api/handler"
	"kasir-api/repositories"
	"kasir-api/service"

	"github.com/spf13/viper"
	httpSwagger "github.com/swaggo/http-swagger"
)

type Config struct {
	Port   string `mapstructure:"PORT"`
	DBConn string `mapstructure:"DB_CONN"`
}

// @title Kasir API
// @version 1.0
// @host localhost:8080
// @BasePath /

func main() {
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if _, err := os.Stat(".env"); err == nil {
		viper.SetConfigFile(".env")
		_ = viper.ReadInConfig()
	}

	config := Config{
		Port:   viper.GetString("PORT"),
		DBConn: viper.GetString("DB_CONN"),
	}

	db, err := database.InitDB(config.DBConn)
	if err != nil {
		log.Fatal("Failed to initialize database:", err)
	}
	defer db.Close()

	// Run migrations
	if err := database.RunMigrations(db, "migrations"); err != nil {
		log.Fatal("Failed to run migrations:", err)
	}

	// Dependency Injection
	// Repositories
	categoryRepo := repositories.NewCategoryRepository(db)
	productRepo := repositories.NewProductRepository(db)

	// Services
	categoryService := service.NewCategoryService(categoryRepo)
	productService := service.NewProductService(productRepo)

	// Handlers
	categoryHandler := handler.NewCategoryHandler(categoryService)
	productHandler := handler.NewProductHandler(productService)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("API ready!"))
	})

	// Register routes - Products
	http.HandleFunc("GET /api/products", productHandler.GetAll)
	http.HandleFunc("GET /api/products/{id}", productHandler.GetByID)
	http.HandleFunc("POST /api/products", productHandler.Create)
	http.HandleFunc("PUT /api/products/{id}", productHandler.Update)
	http.HandleFunc("DELETE /api/products/{id}", productHandler.Delete)

	// Register routes - Categories
	http.HandleFunc("GET /api/categories", categoryHandler.HandleCategories)
	http.HandleFunc("GET /api/categories/{id}", categoryHandler.HandleCategoryByID)
	http.HandleFunc("POST /api/categories", categoryHandler.HandleCategories)
	http.HandleFunc("PUT /api/categories/{id}", categoryHandler.HandleCategoryByID)
	http.HandleFunc("DELETE /api/categories/{id}", categoryHandler.HandleCategoryByID)

	// Swagger
	http.HandleFunc("/swagger/", httpSwagger.WrapHandler)

	addr := "0.0.0.0:" + config.Port
	fmt.Println("Server running di", addr)

	err = http.ListenAndServe(addr, nil)
	if err != nil {
		fmt.Println("gagal running server", err)
	}
}
