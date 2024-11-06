package handlers

import "go_final_project/pkg/models"


type TaskScheduler interface {
	SaveTask(*models.Task) (int64, error)
	GetTasks(string) ([]*models.Task, error)
	GetTaskByID(string) (*models.Task, error)
	UpdateTask(*models.Task) error
	MarkTaskCompleted(string) error
	DeleteTask(string) error
}