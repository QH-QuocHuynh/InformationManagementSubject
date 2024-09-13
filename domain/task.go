package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionTask = "tasks"
)

type TaskQuery struct {
	Page     int    `form:"page,default=1" binding:"gte=1"`
	PageSize int    `form:"pageSize,default=200000" binding:"gte=1"`
	Search   string `form:"search"`
}

type TaskRequest struct {
	Amount int    `form:"amount" json:"amount"`
	Notes  string `form:"notes" json:"notes"`
	Code   string `form:"code" json:"code"`
	Date   string `form:"date" json:"date"`
}

type Task struct {
	ID     primitive.ObjectID `bson:"_id" json:"id"`
	Amount int                `form:"amount" binding:"required" json:"amount"`
	Notes  string             `form:"notes" json:"notes"`
	Code   string             `form:"code" json:"code"`
	Date   string             `form:"date" json:"date"`
}

type TaskRepository interface {
	Create(c context.Context, task *Task) error
	Fetch(c context.Context, query TaskQuery) ([]Task, error)
	GetByID(c context.Context, GetByID string) (*Task, error)
	Delete(c context.Context, GetByID string) error
}

type TaskUsecase interface {
	Create(c context.Context, task *Task) error
	Fetch(c context.Context, query TaskQuery) ([]Task, error)
	GetByID(c context.Context, GetByID string) (*Task, error)
	Delete(c context.Context, GetByID string) error
}
