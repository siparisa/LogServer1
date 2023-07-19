package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// BadRequest sends a Bad Request response with the provided error message and details.
func BadRequest(ctx *gin.Context, message string, details string) {
	ctx.JSON(http.StatusBadRequest, gin.H{
		"error":   message,
		"details": details,
	})
}

// InternalServerError sends an Internal Server Error response with the provided error message and details.
func InternalServerError(ctx *gin.Context, message string, details string) {
	ctx.JSON(http.StatusInternalServerError, gin.H{
		"error":   message,
		"details": details,
	})
}

// OK sends a success response with the provided data.
func OK(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, data)
}

// NotFound sends a Not Found response with the provided error message.
func NotFound(ctx *gin.Context, message string) {
	ctx.JSON(http.StatusNotFound, gin.H{
		"error": message,
	})
}
