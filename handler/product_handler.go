package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"kasir-api/model"
	"kasir-api/service"
)

type ProductHandler struct {
	service *service.ProductService
}

func NewProductHandler(service *service.ProductService) *ProductHandler {
	return &ProductHandler{service: service}
}

// GetAll godoc
// @Summary Get all products
// @Description Mengambil semua data produk
// @Tags products
// @Accept json
// @Produce json
// @Success 200 {object} model.Response
// @Router /api/products [get]
func (h *ProductHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	products, err := h.service.GetAll()
	if err != nil {
		model.Error(w, http.StatusInternalServerError, err.Error())
		return
	}
	model.Success(w, http.StatusOK, "successfully get products", products)
}

// GetByID godoc
// @Summary Get product by ID
// @Description Mengambil produk berdasarkan ID dengan informasi kategori
// @Tags products
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Success 200 {object} model.Response
// @Failure 404 {object} model.Response
// @Router /api/products/{id} [get]
func (h *ProductHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		model.Error(w, http.StatusBadRequest, "Invalid Product ID")
		return
	}

	// Menggunakan GetByIDWithCategory untuk mendapatkan data produk beserta nama kategori
	product, err := h.service.GetByIDWithCategory(id)
	if err != nil {
		model.Error(w, http.StatusNotFound, err.Error())
		return
	}

	model.Success(w, http.StatusOK, "successfully get product", product)
}

// Create godoc
// @Summary Create new product
// @Description Menambahkan produk baru
// @Tags products
// @Accept json
// @Produce json
// @Param product body model.Product true "Product Data" SchemaExample({"name":"Product Name","price":10000,"stock":5})
// @Success 201 {object} model.Response
// @Router /api/products [post]
func (h *ProductHandler) Create(w http.ResponseWriter, r *http.Request) {
	var product model.Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		model.Error(w, http.StatusBadRequest, "Invalid Request")
		return
	}

	err := h.service.Create(&product)
	if err != nil {
		model.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	model.Success(w, http.StatusCreated, "successfully added product", product)
}

// Update godoc
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
func (h *ProductHandler) Update(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		model.Error(w, http.StatusBadRequest, "Invalid Product ID")
		return
	}

	var product model.Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		model.Error(w, http.StatusBadRequest, "Invalid Request")
		return
	}

	product.ID = id
	err = h.service.Update(&product)
	if err != nil {
		model.Error(w, http.StatusNotFound, err.Error())
		return
	}

	model.Success(w, http.StatusOK, "successfully updated product", product)
}

// Delete godoc
// @Summary Delete product
// @Description Menghapus produk berdasarkan ID
// @Tags products
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Success 200 {object} model.Response
// @Failure 404 {object} model.Response
// @Router /api/products/{id} [delete]
func (h *ProductHandler) Delete(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		model.Error(w, http.StatusBadRequest, "Invalid Product ID")
		return
	}

	err = h.service.Delete(id)
	if err != nil {
		model.Error(w, http.StatusNotFound, err.Error())
		return
	}

	model.Success(w, http.StatusOK, "successfully deleted product", nil)
}
