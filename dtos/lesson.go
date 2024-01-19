package dtos

import "BilimSearch/models"

type LessonItemResponse struct {
	LessonItem models.LessonItem
	Tasks []models.Task
}