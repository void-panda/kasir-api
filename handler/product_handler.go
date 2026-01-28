package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"kasir-api/model"
	"kasir-api/service"
)

// GetProducts godoc
// @Summary Get all products
// @Description Mengambil semua data produk
// @Tags products
// @Accept json
// @Produce json
// @Success 200 {object} model.Response
// @Router /api/products [get]
func GetProducts(w http.ResponseWriter, r *http.Request) {
	products := service.GetAllProducts()
	model.Success(w, http.StatusOK, "successfully get products", products)
}

// GetProductByID godoc
// @Summary Get product by ID
// @Description Mengambil produk berdasarkan ID
// @Tags products
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Success 200 {object} model.Response
// @Failure 404 {object} model.Response
// @Router /api/products/{id} [get]
func GetProductByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		model.Error(w, http.StatusBadRequest, "Invalid Product ID")
		return
	}

	product := service.GetProductByID(id)
	if product == nil {
		model.Error(w, http.StatusNotFound, "product not found")
		return
	}

	model.Success(w, http.StatusOK, "successfully get product", product)
}

// CreateProduct godoc
// @Summary Create new product
// @Description Menambahkan produk baru
// @Tags products
// @Accept json
// @Produce json
// @Param product body model.Product true "Product Data" SchemaExample({"name":"Product Name","price":10000,"stock":5})
// @Success 201 {object} model.Response
// @Router /api/products [post]
func CreateProduct(w http.ResponseWriter, r *http.Request) {
	var newProduct model.Product
	if err := json.NewDecoder(r.Body).Decode(&newProduct); err != nil {
		model.Error(w, http.StatusBadRequest, "Invalid Request")
		return
	}

	createdProduct := service.CreateProduct(newProduct)
	model.Success(w, http.StatusCreated, "successfully added product", createdProduct)
}

// UpdateProduct godoc
// @Summary Update product
// @Description Update produk berdasarkan ID
// @Tags products
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Param product body model.Product true "Product Data"
// @Success 200 {object} model.Response
// @Failure 404 {object} model.Response
// @Router /api/products/{id} [put]
func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		model.Error(w, http.StatusBadRequest, "Invalid Product ID")
		return
	}

	var updateProduct model.Product
	if err := json.NewDecoder(r.Body).Decode(&updateProduct); err != nil {
		model.Error(w, http.StatusBadRequest, "Invalid Request")
		return
	}

	updated := service.UpdateProduct(id, updateProduct)
	if updated == nil {
		model.Error(w, http.StatusNotFound, "product not found")
		return
	}

	model.Success(w, http.StatusOK, "successfully updated product", updated)
}

// DeleteProduct godoc
// @Summary Delete product
// @Description Menghapus produk berdasarkan ID
// @Tags products
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Success 200 {object} model.Response
// @Failure 404 {object} model.Response
// @Router /api/products/{id} [delete]
func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		model.Error(w, http.StatusBadRequest, "Invalid Product ID")
		return
	}

	if !service.DeleteProduct(id) {
		model.Error(w, http.StatusNotFound, "product not found")
		return
	}

	model.Success(w, http.StatusOK, "successfully deleted product", nil)
}
