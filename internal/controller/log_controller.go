package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/siparisa/LogServer/internal/controller/helper/request"
	"github.com/siparisa/LogServer/internal/controller/helper/response"
	entity "github.com/siparisa/LogServer/internal/entity"
	"github.com/siparisa/LogServer/internal/service/file_log_service"
	"os"
	"strconv"
)

// LogController handles incoming HTTP requests related to log files.
type LogController struct {
	logService file_log_service.LogService
}

// NewLogController creates a new instance of LogController with the given LogService.
// It takes a LogService as a parameter and returns a pointer to the newly created LogController.
func NewLogController(logService file_log_service.LogService) *LogController {
	return &LogController{
		logService: logService,
	}
}

// GetLogLines handles the HTTP request to retrieve log lines from a log file.
func (c *LogController) GetLogLines(ctx *gin.Context) {

	var qp request.GetLogsQueryParams
	if err := ctx.ShouldBindQuery(&qp); err != nil {
		fmt.Println("Error while parsing query parameters:", err)
		response.BadRequest(ctx, "Invalid query parameters", err.Error())
		return
	}

	// Validate and process 'N' parameter
	n := 0
	if qp.N != "" {
		var err error
		n, err = strconv.Atoi(qp.N)
		if err != nil || n <= 0 {
			fmt.Println("Invalid value for parameter 'n':", err)
			response.BadRequest(ctx, "Invalid value for parameter 'n'", "Parameter 'n' must be a positive integer")
			return
		}
	}

	query := entity.LogQuery{
		Filename: qp.Filename,
		Keyword:  qp.Keyword,
		N:        n,
	}

	// Check if the file exists
	filePath := fmt.Sprintf("%s/%s", c.logService.GetLogDir(), query.Filename)
	if _, err := os.Stat(filePath); err != nil {
		if os.IsNotExist(err) {
			fmt.Println("Log file not found:", err)
			response.NotFound(ctx, "Log file not found")
		} else {
			fmt.Println("Failed to check file existence:", err)
			response.InternalServerError(ctx, "Failed to check file existence", err.Error())
		}
		return
	}

	lines, err := c.logService.GetLogLines(query)
	if err != nil {
		fmt.Println("Failed to get log lines:", err)
		response.InternalServerError(ctx, "Failed to get log lines", err.Error())
		return
	}

	if len(lines) == 0 {
		response.OK(ctx, gin.H{"log_lines": []string{}})
	} else {
		response.OK(ctx, gin.H{"log_lines": lines})
	}

}

// ListLogFiles handles the HTTP request to retrieve the list of available log files in the /var/log directory.
// It sends the list of log files as a JSON response.
func (c *LogController) ListLogFiles(ctx *gin.Context) {

	logFiles, err := c.logService.ListLogFiles()
	if err != nil {
		response.InternalServerError(ctx, "Failed to get list of log files", err.Error())
		return
	}

	response.OK(ctx, gin.H{"log_files": logFiles})

}
