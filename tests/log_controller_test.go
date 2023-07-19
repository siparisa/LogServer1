package tests

import (
	_ "fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/siparisa/LogServer/internal/controller"
	_ "github.com/siparisa/LogServer/internal/entity"
	"github.com/siparisa/LogServer/internal/service/file_log_service"
	"github.com/stretchr/testify/assert"
)

func TestGetLogLinesHandler_HappyPath(t *testing.T) {
	// Create a mock gin context with sample query parameters
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	logService := file_log_service.NewFileLogService("/var/log")
	logController := controller.NewLogController(logService)

	router.GET("/logs", logController.GetLogLines)

	// Perform a GET request with sample query parameters
	req, _ := http.NewRequest("GET", "/logs?filename=test.log&n=10&keyword=error", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assert that the response status code is 200 OK
	assert.Equal(t, http.StatusOK, w.Code)

}

func TestGetLogLinesHandler_InvalidQueryParams(t *testing.T) {
	// Create a mock gin context with invalid query parameters
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	logService := file_log_service.NewFileLogService("/var/log")
	logController := controller.NewLogController(logService)

	router.GET("/logs", logController.GetLogLines)

	// Perform a GET request with invalid query parameters (e.g., invalid 'n' value)
	req, _ := http.NewRequest("GET", "/logs?filename=test.log&n=invalid&keyword=error", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assert that the response status code is 400 Bad Request
	assert.Equal(t, http.StatusBadRequest, w.Code)

}

// Test case for a valid positive integer value for 'N'
func TestGetLogLinesHandler_ValidN(t *testing.T) {
	// Create a mock gin context with sample query parameters
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	logService := file_log_service.NewFileLogService("/var/log")
	logController := controller.NewLogController(logService)

	router.GET("/logs", logController.GetLogLines)

	// Perform a GET request with valid query parameters (e.g., 'n' = 10)
	req, _ := http.NewRequest("GET", "/logs?filename=test.log&n=10&keyword=error", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assert that the response status code is 200 OK
	assert.Equal(t, http.StatusOK, w.Code)
}

// Test case for invalid value for 'N'
func TestGetLogLinesHandler_InvalidN(t *testing.T) {
	// Create a mock gin context with sample query parameters (with invalid 'n' value)
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	logService := file_log_service.NewFileLogService("/var/log")
	logController := controller.NewLogController(logService)

	router.GET("/logs", logController.GetLogLines)

	// Perform a GET request with invalid 'n' value
	req, _ := http.NewRequest("GET", "/logs?filename=test.log&n=invalid&keyword=error", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assert that the response status code is 400 Bad Request
	assert.Equal(t, http.StatusBadRequest, w.Code)
}
