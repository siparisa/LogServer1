package file_log_service

import (
	"bufio"
	"fmt"
	entity "github.com/siparisa/LogServer/internal/entity"
	pag "github.com/siparisa/LogServer/internal/service/file_log_service/paginator"
	utils "github.com/siparisa/LogServer/internal/service/file_log_service/utils"
	"golang.org/x/net/html/charset"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

const pageSize = 1000

// LogService defines the interface for retrieving log lines.
type LogService interface {
	GetLogLines(query entity.LogQuery) ([]string, error)
	GetLogDir() string
	ListLogFiles() ([]string, error)
}

// FileLogService implements the LogService interface.
type FileLogService struct {
	logDir string
}

// NewFileLogService creates a new instance of FileLogService.
func NewFileLogService(logDir string) LogService {
	return &FileLogService{
		logDir: logDir,
	}
}

// GetLogLines retrieves log lines from the log file based on the given query parameters.
func (s *FileLogService) GetLogLines(query entity.LogQuery) ([]string, error) {
	filePath := fmt.Sprintf("%s/%s", s.logDir, query.Filename)

	file, err := os.Open(filePath)
	if err != nil {
		log.Printf("Failed to open log file: %v", err)
		return nil, fmt.Errorf("failed to open log file: %w", err)
	}
	defer file.Close()

	return s.readLogLines(file, query.Keyword, strconv.Itoa(query.N))
}

// readLogLines reads lines and applies filter and pagination
func (s *FileLogService) readLogLines(file *os.File, keyword string, n string) ([]string, error) {
	// Specify UTF-8 encoding when reading the log file
	utf8Reader, err := charset.NewReader(file, "text/plain; charset=utf-8")
	if err != nil {
		return nil, fmt.Errorf("failed to create UTF-8 reader: %w", err)
	}

	scanner := bufio.NewScanner(utf8Reader)

	paginator := pag.NewPaginator(scanner, keyword, pageSize)
	var lines []string
	for scanner.Scan() {
		line := scanner.Text()

		if keyword != "" && !strings.Contains(line, keyword) {
			continue
		}

		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("failed to read log file: %w", err)
	}

	// Reverse the order of log lines for the first page only
	utils.ReverseLines(lines)

	if n == "" || n == "0" || len(lines) <= paginator.IntVal(n) {
		return lines, nil
	}

	// Get the last 'n' lines directly from the 'lines' slice
	lastNLines := lines[:paginator.IntVal(n)]

	if len(lastNLines) >= pageSize {
		remainingLines := paginator.PaginateLines(pageSize - len(lastNLines))
		lastNLines = append(lastNLines, remainingLines...)
	}

	return lastNLines, nil
}

func (s *FileLogService) GetLogDir() string {
	return s.logDir
}

// ListLogFiles retrieves the list of available log files in the log directory.
func (s *FileLogService) ListLogFiles() ([]string, error) {
	fileInfoList, err := ioutil.ReadDir(s.logDir)
	if err != nil {
		log.Printf("Failed to read log directory: %v", err)
		return nil, fmt.Errorf("failed to read log directory: %w", err)
	}

	var logFiles []string
	for _, fileInfo := range fileInfoList {
		if !fileInfo.IsDir() && filepath.Ext(fileInfo.Name()) == ".log" {
			logFiles = append(logFiles, fileInfo.Name())
		}
	}

	return logFiles, nil
}
