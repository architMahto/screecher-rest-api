package middlewares

import (
	"bytes"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type LogResponseWriter struct {
	http.ResponseWriter
	StatusCode int
	buffer     bytes.Buffer
}

func NewLogResponseWriter(res http.ResponseWriter) *LogResponseWriter {
	return &LogResponseWriter{ResponseWriter: res}
}

func (writer *LogResponseWriter) WriteHeader(statusCode int) {
	writer.StatusCode = statusCode
	writer.ResponseWriter.WriteHeader(statusCode)
}

func (writer *LogResponseWriter) Write(body []byte) (int, error) {
	writer.buffer.Write(body)
	return writer.ResponseWriter.Write(body)
}

type LoggingMiddleware struct {
	logger *log.Logger
}

func NewLoggingMiddleware(loggger *log.Logger) *LoggingMiddleware {
	return &LoggingMiddleware{logger: loggger}
}

func (middleware *LoggingMiddleware) Func() mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
			startTime := time.Now()

			logResponseWriter := NewLogResponseWriter(res)
			next.ServeHTTP(logResponseWriter, req)
			statusCode := strconv.FormatInt(int64(logResponseWriter.StatusCode), 10)
			statusText := http.StatusText(logResponseWriter.StatusCode)

			middleware.logger.Printf(
				"%s - %s %s %s %s %s",
				req.Proto,
				req.Method,
				statusCode,
				statusText,
				time.Since(startTime).String(),
				req.URL.RequestURI(),
			)
		})
	}
}
