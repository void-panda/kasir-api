package service

import "kasir-api/model"

var tasks = []model.Task{
	{ID: 1, Title: "Task 1", Description: "Description for Task 1", Status: model.TaskStatusCompleted},
	{ID: 2, Title: "Task 2", Description: "Description for Task 2", Status: model.TaskStatusInProgress},
	{ID: 3, Title: "Task 3", Description: "Description for Task 3", Status: model.TaskStatusCompleted},
	{ID: 4, Title: "Task 4", Description: "Description for Task 4", Status: model.TaskStatusPending},
	{ID: 5, Title: "Task 5", Description: "Description for Task 5", Status: model.TaskStatusInProgress},
}

func GetAllTasks() []model.Task {
	return tasks
}

func GetTaskByID(id int) *model.Task {
	for i, task := range tasks {
		if task.ID == id {
			return &tasks[i]
		}
	}
	return nil
}

func CreateTask(task model.Task) model.Task {
	task.ID = len(tasks) + 1
	tasks = append(tasks, task)
	return task
}

func UpdateTask(id int, updatedTask model.Task) *model.Task {
	for i := range tasks {
		if tasks[i].ID == id {
			if updatedTask.Title != "" {
				tasks[i].Title = updatedTask.Title
			}
			if updatedTask.Description != "" {
				tasks[i].Description = updatedTask.Description
			}
			if updatedTask.Status != "" {
				tasks[i].Status = updatedTask.Status
			}
			return &tasks[i]
		}
	}
	return nil
}

func DeleteTask(id int) bool {
	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			return true
		}
	}
	return false
}
