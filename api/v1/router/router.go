
package router

import (
	"github.com/gin-gonic/gin"
	"github.com/0x30c4/GoPasteBin/api/v1/handlers"
	"github.com/0x30c4/GoPasteBin/internal/env"
	"fmt"
	"os"
	"time"
	"io"
)

func Initialize() *gin.Engine {

	// Check if the release mode is on or not
	if env.RELEASE_MODE {
		gin.SetMode(gin.ReleaseMode)
	}

    gin.DisableConsoleColor()

	gin.SetMode(gin.DebugMode)
    // Logging to a file.
    f, _ := os.Create(env.LOG_FILE)

    gin.DefaultWriter = io.MultiWriter(f)

	router := gin.New()

	// Defining loging parameters 
  	router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
    	return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
        	param.ClientIP,
        	param.TimeStamp.Format(time.RFC1123),
        	param.Method,
        	param.Path,
        	param.Request.Proto,
        	param.StatusCode,
        	param.Latency,
        	param.Request.UserAgent(),
        	param.ErrorMessage,
    	)
  	}))

  	router.Use(gin.Recovery())

	// https://blog.devgenius.io/golang-apis-a-skeleton-for-your-future-projects-a082dc4d681
	router.GET("/", handlers.Index)
	router.GET("/:ID", handlers.GetData)
	router.POST("/", handlers.PostData)

	return router
}


func ServeRouter() {
	router := Initialize()
	router.Run(env.HOST + ":" + env.PORT)
}
