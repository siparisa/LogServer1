package request

// GetLogsQueryParams is the request query params for getting a list of services
type GetLogsQueryParams struct {
	Filename string `form:"filename" example:"Log file Name"`
	Keyword  string `form:"keyword" example:"Specific Keyword"`
	N        string `form:"n" example:"10"`
}
