package route

import (
	"time"

	"iow/go-backend/api/controller"
	"iow/go-backend/bootstrap"
	"iow/go-backend/domain"
	"iow/go-backend/mongo"
	"iow/go-backend/repository"
	"iow/go-backend/usecase"

	"github.com/gin-gonic/gin"
)

func NewTaskRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	tr := repository.NewTaskRepository(db, domain.CollectionTask)
	tc := &controller.TaskController{
		TaskUsecase: usecase.NewTaskUsecase(tr, timeout),
	}
	group.GET("/task", tc.Fetch)
	group.POST("/task", tc.Create)
	group.GET("/task/:id", tc.GetByID)
	group.GET("/task/import", tc.Import)
	group.DELETE("/task/:id", tc.Delete)
}
