package route

import (
	"time"

	"iow/go-backend/api/middleware"
	"iow/go-backend/bootstrap"
	"iow/go-backend/mongo"

	"github.com/gin-gonic/gin"
)

func Setup(env *bootstrap.Env, timeout time.Duration, db mongo.Database, gin *gin.Engine) {

	// Middleware server
	gin.Use(middleware.ServerMiddleware())

	apiRouter := gin.Group("api")

	// All Public APIs
	publicRouter := apiRouter.Group("")
	{
		// Home
		NewHomeRoute(env, timeout, db, publicRouter)

		//Task
		NewTaskRouter(env, timeout, db, publicRouter)
	}
}
