package utils

import "regexp"

type HTTPAnalyzer struct{}

func NewHTTPAnalyzer() *HTTPAnalyzer {
	return &HTTPAnalyzer{}
}

func (h *HTTPAnalyzer) IsHTTP(req []byte) bool {
	return regexp.MustCompile(`HTTP\/1\.1`).Match(req)
}
