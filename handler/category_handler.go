package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"kasir-api/model"
	"kasir-api/service"
)

type CategoryHandler struct {
	service *service.CategoryService
}

func NewCategoryHandler(service *service.CategoryService) *CategoryHandler {
	return &CategoryHandler{service: service}
}

// HandleCategories - GET /api/categories, POST /api/categories
func (h *CategoryHandler) HandleCategories(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.GetAll(w, r)
	case http.MethodPost:
		h.Create(w, r)
	default:
		model.Error(w, http.StatusMethodNotAllowed, "Method not allowed")
	}
}

// HandleCategoryByID - GET/PUT/DELETE /api/categories/{id}
func (h *CategoryHandler) HandleCategoryByID(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.GetByID(w, r)
	case http.MethodPut:
		h.Update(w, r)
	case http.MethodDelete:
		h.Delete(w, r)
	default:
		model.Error(w, http.StatusMethodNotAllowed, "Method not allowed")
	}
}

// GetAll godoc
// @Summary Get all categories
// @Description Mengambil semua data kategori
// @Tags categories
// @Accept json
// @Produce json
// @Success 200 {object} model.Response
// @Router /api/categories [get]
func (h *CategoryHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	categories, err := h.service.GetAll()
	if err != nil {
		model.Error(w, http.StatusInternalServerError, err.Error())
		return
	}
	model.Success(w, http.StatusOK, "successfully get categories", categories)
}

// GetByID godoc
// @Summary Get category by ID
// @Description Mengambil kategori berdasarkan ID
// @Tags categories
// @Accept json
// @Produce json
// @Param id path int true "Category ID"
// @Success 200 {object} model.Response
// @Failure 404 {object} model.Response
// @Router /api/categories/{id} [get]
func (h *CategoryHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		model.Error(w, http.StatusBadRequest, "Invalid Category ID")
		return
	}

	category, err := h.service.GetByID(id)
	if err != nil {
		model.Error(w, http.StatusNotFound, err.Error())
		return
	}

	model.Success(w, http.StatusOK, "successfully get category", category)
}

// Create godoc
// @Summary Create new category
// @Description Menambahkan kategori baru
// @Tags categories
// @Accept json
// @Produce json
// @Param category body model.Category true "Category Data" SchemaExample({"name":"Category Name","description":"Category Description"})
// @Success 201 {object} model.Response
// @Router /api/categories [post]
func (h *CategoryHandler) Create(w http.ResponseWriter, r *http.Request) {
	var category model.Category
	if err := json.NewDecoder(r.Body).Decode(&category); err != nil {
		model.Error(w, http.StatusBadRequest, "Invalid Request")
		return
	}

	err := h.service.Create(&category)
	if err != nil {
		model.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	model.Success(w, http.StatusCreated, "successfully added category", category)
}

// Update godoc
// @Summary Update category
// @Description Update kategori berdasarkan ID
// @Tags categories
// @Accept json
// @Produce json
// @Param id path int true "Category ID"
// @Param category body model.Category true "Category Data"
// @Success 200 {object} model.Response
// @Failure 404 {object} model.Response
// @Router /api/categories/{id} [put]
func (h *CategoryHandler) Update(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		model.Error(w, http.StatusBadRequest, "Invalid Category ID")
		return
	}

	var category model.Category
	if err := json.NewDecoder(r.Body).Decode(&category); err != nil {
		model.Error(w, http.StatusBadRequest, "Invalid Request")
		return
	}

	category.ID = id
	err = h.service.Update(&category)
	if err != nil {
		model.Error(w, http.StatusNotFound, err.Error())
		return
	}

	model.Success(w, http.StatusOK, "successfully updated category", category)
}

// Delete godoc
// @Summary Delete category
// @Description Menghapus kategori berdasarkan ID
// @Tags categories
// @Accept json
// @Produce json
// @Param id path int true "Category ID"
// @Success 200 {object} model.Response
// @Failure 404 {object} model.Response
// @Router /api/categories/{id} [delete]
func (h *CategoryHandler) Delete(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		model.Error(w, http.StatusBadRequest, "Invalid Category ID")
		return
	}

	err = h.service.Delete(id)
	if err != nil {
		model.Error(w, http.StatusNotFound, err.Error())
		return
	}

	model.Success(w, http.StatusOK, "successfully deleted category", nil)
}
