package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"kasir-api/model"
)

// ============== PRODUCTS HANDLERS ==============

// GetProducts godoc
// @Summary Get all products
// @Description Mengambil semua data produk
// @Tags products
// @Accept json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/products [get]
func GetProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(model.Response{
		Status:  "ok",
		Message: "successfully get products",
		Data:    GetAllProducts(),
	})
}

// GetProductByID godoc
// @Summary Get product by ID
// @Description Mengambil produk berdasarkan ID
// @Tags products
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Success 200 {object} map[string]interface{}
// @Failure 404 {object} map[string]string
// @Router /api/products/{id} [get]
func GetProductByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	idStr := strings.TrimPrefix(r.URL.Path, "/api/products/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Product ID", http.StatusBadRequest)
		return
	}

	product := GetProduct(id)
	if product == nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(model.Response{
			Status:  "error",
			Message: "product not found",
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(model.Response{
		Status:  "ok",
		Message: "successfully get product",
		Data:    product,
	})
}

// CreateProduct godoc
// @Summary Create new product
// @Description Menambahkan produk baru
// @Tags products
// @Accept json
// @Produce json
// @Param product body model.Product true "Product Data" SchemaExample({"name":"Product Name","price":10000,"stock":5})
// @Success 201 {object} map[string]interface{}
// @Router /api/products [post]
func CreateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var newProduct model.Product
	err := json.NewDecoder(r.Body).Decode(&newProduct)
	if err != nil {
		http.Error(w, "Invalid Request", http.StatusBadRequest)
		return
	}

	createdProduct := AddProduct(newProduct)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(model.Response{
		Status:  "ok",
		Message: "successfully added product",
		Data:    createdProduct,
	})
}

// UpdateProduct godoc
// @Summary Update product
// @Description Update produk berdasarkan ID
// @Tags products
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Param product body model.Product true "Product Data"
// @Success 200 {object} map[string]interface{}
// @Failure 404 {object} map[string]string
// @Router /api/products/{id} [put]
func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	idStr := strings.TrimPrefix(r.URL.Path, "/api/products/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Product ID", http.StatusBadRequest)
		return
	}

	var updateProduct model.Product
	err = json.NewDecoder(r.Body).Decode(&updateProduct)
	if err != nil {
		http.Error(w, "Invalid Request", http.StatusBadRequest)
		return
	}

	updated := UpdateProductData(id, updateProduct)
	if updated == nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(model.Response{
			Status:  "error",
			Message: "product not found",
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(model.Response{
		Status:  "ok",
		Message: "successfully updated product",
		Data:    updated,
	})
}

// DeleteProduct godoc
// @Summary Delete product
// @Description Menghapus produk berdasarkan ID
// @Tags products
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Success 200 {object} map[string]interface{}
// @Failure 404 {object} map[string]string
// @Router /api/products/{id} [delete]
func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	idStr := strings.TrimPrefix(r.URL.Path, "/api/products/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Product ID", http.StatusBadRequest)
		return
	}

	if !DeleteProductData(id) {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(model.Response{
			Status:  "error",
			Message: "product not found",
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(model.Response{
		Status:  "ok",
		Message: "successfully deleted product",
	})
}

// ============== CATEGORIES HANDLERS ==============

// GetCategories godoc
// @Summary Get all categories
// @Description Mengambil semua data kategori
// @Tags categories
// @Accept json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/categories [get]
func GetCategories(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(model.Response{
		Status:  "ok",
		Message: "successfully get categories",
		Data:    GetAllCategories(),
	})
}

// GetCategoryByID godoc
// @Summary Get category by ID
// @Description Mengambil kategori berdasarkan ID
// @Tags categories
// @Accept json
// @Produce json
// @Param id path int true "Category ID"
// @Success 200 {object} map[string]interface{}
// @Failure 404 {object} map[string]string
// @Router /api/categories/{id} [get]
func GetCategoryByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	idStr := strings.TrimPrefix(r.URL.Path, "/api/categories/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Category ID", http.StatusBadRequest)
		return
	}

	category := GetCategory(id)
	if category == nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(model.Response{
			Status:  "error",
			Message: "category not found",
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(model.Response{
		Status:  "ok",
		Message: "successfully get category",
		Data:    category,
	})
}

// CreateCategory godoc
// @Summary Create new category
// @Description Menambahkan kategori baru
// @Tags categories
// @Accept json
// @Produce json
// @Param category body model.Category true "Category Data" SchemaExample({"name":"Category Name","description":"Category Description"})
// @Success 201 {object} map[string]interface{}
// @Router /api/categories [post]
func CreateCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var newCategory model.Category
	err := json.NewDecoder(r.Body).Decode(&newCategory)
	if err != nil {
		http.Error(w, "Invalid Request", http.StatusBadRequest)
		return
	}

	createdCategory := AddCategory(newCategory)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(model.Response{
		Status:  "ok",
		Message: "successfully added category",
		Data:    createdCategory,
	})
}

// UpdateCategory godoc
// @Summary Update category
// @Description Update kategori berdasarkan ID
// @Tags categories
// @Accept json
// @Produce json
// @Param id path int true "Category ID"
// @Param category body model.Category true "Category Data"
// @Success 200 {object} map[string]interface{}
// @Failure 404 {object} map[string]string
// @Router /api/categories/{id} [put]
func UpdateCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	idStr := strings.TrimPrefix(r.URL.Path, "/api/categories/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Category ID", http.StatusBadRequest)
		return
	}

	var updateCategory model.Category
	err = json.NewDecoder(r.Body).Decode(&updateCategory)
	if err != nil {
		http.Error(w, "Invalid Request", http.StatusBadRequest)
		return
	}

	updated := UpdateCategoryData(id, updateCategory)
	if updated == nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(model.Response{
			Status:  "error",
			Message: "category not found",
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(model.Response{
		Status:  "ok",
		Message: "successfully updated category",
		Data:    updated,
	})
}

// DeleteCategory godoc
// @Summary Delete category
// @Description Menghapus kategori berdasarkan ID
// @Tags categories
// @Accept json
// @Produce json
// @Param id path int true "Category ID"
// @Success 200 {object} map[string]interface{}
// @Failure 404 {object} map[string]string
// @Router /api/categories/{id} [delete]
func DeleteCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	idStr := strings.TrimPrefix(r.URL.Path, "/api/categories/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Category ID", http.StatusBadRequest)
		return
	}

	if !DeleteCategoryData(id) {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(model.Response{
			Status:  "error",
			Message: "category not found",
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(model.Response{
		Status:  "ok",
		Message: "successfully deleted category",
	})
}

// ============== HEALTH CHECK ==============

// HealthCheck godoc
// @Summary Health check
// @Description Check if server is running
// @Tags health
// @Accept json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /health [get]
func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(model.Response{
		Status:  "ok",
		Message: "server is running",
	})
}
