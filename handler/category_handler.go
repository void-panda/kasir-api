package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"kasir-api/model"
	"kasir-api/service"
)

// GetCategories godoc
// @Summary Get all categories
// @Description Mengambil semua data kategori
// @Tags categories
// @Accept json
// @Produce json
// @Success 200 {object} model.Response
// @Router /api/categories [get]
func GetCategories(w http.ResponseWriter, r *http.Request) {
	categories := service.GetAllCategories()
	model.Success(w, http.StatusOK, "successfully get categories", categories)
}

// GetCategoryByID godoc
// @Summary Get category by ID
// @Description Mengambil kategori berdasarkan ID
// @Tags categories
// @Accept json
// @Produce json
// @Param id path int true "Category ID"
// @Success 200 {object} model.Response
// @Failure 404 {object} model.Response
// @Router /api/categories/{id} [get]
func GetCategoryByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		model.Error(w, http.StatusBadRequest, "Invalid Category ID")
		return
	}

	category := service.GetCategoryByID(id)
	if category == nil {
		model.Error(w, http.StatusNotFound, "category not found")
		return
	}

	model.Success(w, http.StatusOK, "successfully get category", category)
}

// CreateCategory godoc
// @Summary Create new category
// @Description Menambahkan kategori baru
// @Tags categories
// @Accept json
// @Produce json
// @Param category body model.Category true "Category Data" SchemaExample({"name":"Category Name","description":"Category Description"})
// @Success 201 {object} model.Response
// @Router /api/categories [post]
func CreateCategory(w http.ResponseWriter, r *http.Request) {
	var newCategory model.Category
	if err := json.NewDecoder(r.Body).Decode(&newCategory); err != nil {
		model.Error(w, http.StatusBadRequest, "Invalid Request")
		return
	}

	createdCategory := service.CreateCategory(newCategory)
	model.Success(w, http.StatusCreated, "successfully added category", createdCategory)
}

// UpdateCategory godoc
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
func UpdateCategory(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		model.Error(w, http.StatusBadRequest, "Invalid Category ID")
		return
	}

	var updateCategory model.Category
	if err := json.NewDecoder(r.Body).Decode(&updateCategory); err != nil {
		model.Error(w, http.StatusBadRequest, "Invalid Request")
		return
	}

	updated := service.UpdateCategory(id, updateCategory)
	if updated == nil {
		model.Error(w, http.StatusNotFound, "category not found")
		return
	}

	model.Success(w, http.StatusOK, "successfully updated category", updated)
}

// DeleteCategory godoc
// @Summary Delete category
// @Description Menghapus kategori berdasarkan ID
// @Tags categories
// @Accept json
// @Produce json
// @Param id path int true "Category ID"
// @Success 200 {object} model.Response
// @Failure 404 {object} model.Response
// @Router /api/categories/{id} [delete]
func DeleteCategory(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		model.Error(w, http.StatusBadRequest, "Invalid Category ID")
		return
	}

	if !service.DeleteCategory(id) {
		model.Error(w, http.StatusNotFound, "category not found")
		return
	}

	model.Success(w, http.StatusOK, "successfully deleted category", nil)
}
