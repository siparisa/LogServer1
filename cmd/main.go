package main

import (
	"github.com/gin-gonic/gin"
	"github.com/siparisa/LogServer/internal/controller"
	"github.com/siparisa/LogServer/internal/service/file_log_service"
)

func main() {

	logService := file_log_service.NewFileLogService("/var/log")
	logController := controller.NewLogController(logService)

	router := gin.Default()

	router.GET("/logs", logController.GetLogLines)
	router.GET("/log-files", logController.ListLogFiles)

	router.Run(":8080")
}
