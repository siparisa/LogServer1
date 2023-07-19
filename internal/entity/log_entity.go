package entity

type LogQuery struct {
	Filename string
	Keyword  string
	N        int
}

type LogEntry struct {
	ID        uint   `json:"id"`
	Filename  string `json:"filename"`
	Content   string `json:"content"`
	Timestamp string `json:"timestamp"`
}
