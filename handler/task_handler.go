package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"kasir-api/model"
	"kasir-api/service"
)

// GetAllTasks godoc
// @Summary Get all tasks
// @Description Mengambil semua data tugas
// @Tags tasks
// @Accept json
// @Produce json
// @Success 200 {object} model.Response
// @Router /api/tasks [get]
func GetAllTasks(w http.ResponseWriter, r *http.Request) {
	tasks := service.GetAllTasks()
	model.Success(w, http.StatusOK, "successfully get tasks", tasks)
}

// GetTaskByID godoc
// @Summary Get task by ID
// @Description Mengambil tugas berdasarkan ID
// @Tags tasks
// @Accept json
// @Produce json
// @Param id path int true "Task ID"
// @Success 200 {object} model.Response
// @Failure 404 {object} model.Response
// @Router /api/tasks/{id} [get]
func GetTaskByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		model.Error(w, http.StatusBadRequest, "Invalid Task ID")
		return
	}

	task := service.GetTaskByID(id)
	if task == nil {
		model.Error(w, http.StatusNotFound, "task not found")
		return
	}

	model.Success(w, http.StatusOK, "task found", task)
}

// CreateTask godoc
// @Summary Create new task
// @Description Membuat tugas baru
// @Tags tasks
// @Accept json
// @Produce json
// @Param task body model.Task true "Task Data" SchemaExample({"title":"Task Title","description":"Task Description", "status": "in_progress"})
// @Success 201 {object} model.Response
// @Router /api/tasks [post]
func CreateTask(w http.ResponseWriter, r *http.Request) {
	var newTask model.Task
	if err := json.NewDecoder(r.Body).Decode(&newTask); err != nil {
		model.Error(w, http.StatusBadRequest, "Invalid Request")
		return
	}

	createdTask := service.CreateTask(newTask)
	model.Success(w, http.StatusCreated, "New Task created", createdTask)
}

// UpdateTask godoc
// @Summary Update task
// @Description Update tugas berdasarkan ID
// @Tags tasks
// @Accept json
// @Produce json
// @Param id path int true "Task ID"
// @Param task body model.Task true "Task Data"
// @Success 200 {object} model.Response
// @Failure 404 {object} model.Response
// @Router /api/tasks/{id} [patch]
func UpdateTask(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		model.Error(w, http.StatusBadRequest, "Invalid Task ID")
		return
	}

	var editTask model.Task
	if err := json.NewDecoder(r.Body).Decode(&editTask); err != nil {
		model.Error(w, http.StatusBadRequest, "Invalid Request")
		return
	}

	updated := service.UpdateTask(id, editTask)
	if updated == nil {
		model.Error(w, http.StatusNotFound, "Task not found")
		return
	}

	model.Success(w, http.StatusOK, "successfully update task", updated)
}

// DeleteTask godoc
// @Summary Delete task
// @Description Menghapus tugas berdasarkan ID
// @Tags tasks
// @Accept json
// @Produce json
// @Param id path int true "Task ID"
// @Success 200 {object} model.Response
// @Failure 404 {object} model.Response
// @Router /api/tasks/{id} [delete]
func DeleteTask(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		model.Error(w, http.StatusBadRequest, "Invalid Task ID")
		return
	}

	if !service.DeleteTask(id) {
		model.Error(w, http.StatusNotFound, "task not found")
		return
	}

	model.Success(w, http.StatusOK, "successfully deleted task", nil)
}

// HealthCheck godoc
// @Summary Health check
// @Description Check if server is running
// @Tags health
// @Accept json
// @Produce json
// @Success 200 {object} model.Response
// @Router /health [get]
func HealthCheck(w http.ResponseWriter, r *http.Request) {
	model.Success(w, http.StatusOK, "server is running", nil)
}
