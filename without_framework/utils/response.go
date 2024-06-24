package utils

import (
	"fmt"
	"net/http"
)

type ResponseWriter struct {
	Status int
	Body   string
	Time   int64
	http.ResponseWriter
}

// Converts http.ResponseWriter into *utils.ResponseWriter
func NewResponseWriter(w http.ResponseWriter) *ResponseWriter {
	return &ResponseWriter{ResponseWriter: w}
}

func (w *ResponseWriter) WriteHeader(code int) {
	w.Status = code
	w.ResponseWriter.WriteHeader(code)
}

func (w *ResponseWriter) Write(body []byte) (int, error) {
	w.Body = string(body)
	return w.ResponseWriter.Write(body)
}

func (w *ResponseWriter) String() string {
	out := fmt.Sprintf("status %d (took %dms)", w.Status, w.Time)
	if w.Body != "" {
		out = fmt.Sprintf("%s\n\tresponse: %s", out, w.Body)
	}
	return out
}
