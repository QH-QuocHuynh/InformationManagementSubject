package repository

import (
	"context"
	"fmt"

	"iow/go-backend/domain"
	"iow/go-backend/mongo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type taskRepository struct {
	database   mongo.Database
	collection string
}

func NewTaskRepository(db mongo.Database, collection string) domain.TaskRepository {
	return &taskRepository{
		database:   db,
		collection: collection,
	}
}

// Create is a function to create a new task
// It will return the created task
// If the request is invalid, it will return 400
// If there is an error in the database, it will return 500
func (tr *taskRepository) Create(c context.Context, task *domain.Task) error {
	collection := tr.database.Collection(tr.collection)

	// if task.ID is empty, generate new ObjectID
	if task.ID == primitive.NilObjectID {
		task.ID = primitive.NewObjectID()
	}

	_, err := collection.InsertOne(c, task)

	return err
}

// Fetch is a function to get all tasks
// It will return all tasks
// If the request is invalid, it will return 400
// If there is an error in the database, it will return 500
func (tr *taskRepository) Fetch(c context.Context, query domain.TaskQuery) ([]domain.Task, error) {
	collection := tr.database.Collection(tr.collection)

	skip := (query.Page - 1) * query.PageSize

	// Use the `i` option in the regex for case-insensitive search
	filter := bson.M{
		"$or": []bson.M{
			{"notes": bson.M{"$regex": query.Search, "$options": "i"}},
			{"code": bson.M{"$regex": query.Search, "$options": "i"}},
		},
	}

	var tasks []domain.Task
	findOptions := options.Find().
		SetSkip(int64(skip)).
		SetLimit(int64(query.PageSize)).
		SetSort(bson.M{"date": -1}).
		SetProjection(bson.M{"notes": 1, "code": 1, "date": 1, "amount": 1})

	cursor, err := collection.Find(c, filter, findOptions)
	if err != nil {
		return nil, err
	}

	err = cursor.All(c, &tasks)
	return tasks, err
}

// GetByID is a function to get a task by ID
// It will return the task
// If the task is not found, it will return 404
// If there is an error in the database, it will return 500
func (tr *taskRepository) GetByID(c context.Context, taskID string) (*domain.Task, error) {
	collection := tr.database.Collection(tr.collection)

	var task domain.Task

	taskObjectID, err := primitive.ObjectIDFromHex(taskID)
	if err != nil {
		return nil, err
	}

	err = collection.FindOne(c, bson.M{"_id": taskObjectID}).Decode(&task)
	if err != nil {
		return nil, fmt.Errorf("task not found")
	}

	return &task, nil
}

// Delete is a function to delete a task
// It will delete a task based on the ID
// If the task is not found, it will return 404
// If there is an error in the database, it will return 500
func (tr *taskRepository) Delete(c context.Context, taskID string) error {
	collection := tr.database.Collection(tr.collection)

	taskObjectID, err := primitive.ObjectIDFromHex(taskID)
	if err != nil {
		return err
	}

	_, err = collection.DeleteOne(c, bson.M{"_id": taskObjectID})
	return err
}
