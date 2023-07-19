package paginator

import (
	"bufio"
	"log"
	"strconv"
	"strings"
)

type Paginator struct {
	scanner  *bufio.Scanner
	keyword  string
	pageSize int
}

// NewPaginator creates a new Paginator instance.
func NewPaginator(scanner *bufio.Scanner, keyword string, pageSize int) *Paginator {
	return &Paginator{
		scanner:  scanner,
		keyword:  keyword,
		pageSize: pageSize,
	}
}

// PaginateLines paginates the lines based on the given number of lines to show.
func (p *Paginator) PaginateLines(numLines int) []string {
	var lines []string
	var buffer []string

	for p.scanner.Scan() {
		line := p.scanner.Text()

		if p.keyword != "" && !strings.Contains(line, p.keyword) {
			continue
		}

		buffer = append(buffer, line)
		if len(buffer) > numLines {
			buffer = buffer[1:]
		}
	}

	if len(buffer) < numLines {
		lines = buffer
	} else {
		lines = make([]string, numLines)
		copy(lines, buffer[len(buffer)-numLines:])
	}

	return lines
}

// IntVal is a helper function to convert a string to an integer.
func (p *Paginator) IntVal(s string) int {
	val, err := strconv.Atoi(s)
	if err != nil {
		// Log an error if the conversion fails
		log.Printf("Error converting string to integer: %s", err.Error())
		return 0
	}
	return val
}
