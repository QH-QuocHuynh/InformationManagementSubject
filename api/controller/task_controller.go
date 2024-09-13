package controller

import (
	"encoding/json"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"

	"iow/go-backend/domain"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TaskController struct {
	TaskUsecase domain.TaskUsecase
}

// Create is a function to create a new task
// It will return the created task
// If the request is invalid, it will return 400
func (tc *TaskController) Create(c *gin.Context) {
	var request domain.TaskRequest

	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	// Create a new task using the validated fields
	task := domain.Task{
		ID:     primitive.NewObjectID(),
		Amount: request.Amount,
		Notes:  request.Notes,
		Code:   request.Code,
		Date:   request.Date,
	}

	err = tc.TaskUsecase.Create(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, task)
}

// Fetch is a function to get all tasks
// It will return all tasks
// If the request is invalid, it will return 400
func (u *TaskController) Fetch(c *gin.Context) {
	// Bind the query to the TaskQuery struct
	var query domain.TaskQuery
	if err := c.ShouldBindQuery(&query); err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	tasks, err := u.TaskUsecase.Fetch(c, query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, tasks)
}

// GetByID is a function to get task by ID
// It will return a task based on the ID
// If the task is not found, it will return 404
func (u *TaskController) GetByID(c *gin.Context) {

	taskID := c.Param("id")

	task, err := u.TaskUsecase.GetByID(c, taskID)
	if err != nil {
		c.JSON(http.StatusNotFound, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, task)
}

// Import is a function to import data from json file
// and save it to the database
// This function is only for development purpose
// It will return ok if the import is successful
func (u *TaskController) Import(c *gin.Context) {
	timeStart := time.Now()

	// Open the json file
	jsonFile, err := os.Open("transactions.json")
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: "Cannot open file: " + err.Error()})
		return
	}
	defer jsonFile.Close()

	// Decode the json file
	var tasks []domain.TaskRequest
	decoder := json.NewDecoder(jsonFile)
	if err := decoder.Decode(&tasks); err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: "Error decoding JSON: " + err.Error()})
		return
	}

	// Create a new task for each task in the json file
	var wg sync.WaitGroup
	errCh := make(chan error, len(tasks))
	for _, task := range tasks {
		wg.Add(1)
		go func(task domain.TaskRequest) {
			defer wg.Done()

			t := domain.Task{
				ID:     primitive.NewObjectID(),
				Amount: task.Amount,
				Notes:  task.Notes,
				Code:   task.Code,
				Date:   task.Date,
			}

			if err := u.TaskUsecase.Create(c, &t); err != nil {
				errCh <- err
			}
		}(task)
	}

	// Wait for all tasks to be created
	wg.Wait()
	close(errCh)

	// Todo: return list of errors and success
	if len(errCh) > 0 {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: "Error creating tasks"})
		return
	}

	// Return success message
	timeEnd := time.Now()
	elapsed := timeEnd.Sub(timeStart)
	message := "Created " + strconv.Itoa(len(tasks)) + " tasks in " + elapsed.String()

	// return total task create success
	c.JSON(http.StatusOK, message)
}

// Delete is a function to delete a task
// It will delete a task based on the ID
// If the task is not found, it will return 404
func (u *TaskController) Delete(c *gin.Context) {
	taskID := c.Param("id")

	err := u.TaskUsecase.Delete(c, taskID)
	if err != nil {
		c.JSON(http.StatusNotFound, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, "ok")
}
