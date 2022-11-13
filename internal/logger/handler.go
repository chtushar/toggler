package logger

import (
	"net"
	"net/http"
	"strings"
	"time"

	"go.uber.org/zap"
)

type timer interface {
	// Now returns the current time
	Now() time.Time
	// Since returns the time passed since the given time
	Since(time.Time) time.Duration
}

// realClock save request times
type realClock struct{}

// Now wraps time.Now() to return current time
func (rc *realClock) Now() time.Time {
	return time.Now()
}

// Since returns the duration since the given time
func (rc *realClock) Since(t time.Time) time.Duration {
	return time.Since(t)
}

// Handler is a wrapper to store the logger
type Handler struct {
	log     *zap.Logger
	handler http.Handler
	clock   timer
}

// NewHandler creates a new instance of Handler.
func NewHandler(log *zap.Logger) func(h http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return &Handler{
			log:     log,
			handler: h,
			clock:   &realClock{},
		}
	}
}

type loggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func newLoggingResponseWriter(w http.ResponseWriter) *loggingResponseWriter {
	return &loggingResponseWriter{w, http.StatusOK}
}

// WriteHeader writes a status code to http response
func (lw *loggingResponseWriter) WriteHeader(code int) {
	lw.statusCode = code
	lw.ResponseWriter.WriteHeader(code)
}

//func (lw *loggingResponseWriter) Write(b []byte) (int, error) {
//	return lw.ResponseWriter.Write(b)
//}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	requestStartTime := h.clock.Now()
	logEntry := h.log.With(
		zap.String("IP", h.getRealIP(r)),
		zap.String("timestamp", requestStartTime.Format("02/Jan/2006:15:04:05 -0700")),
		zap.String("method", r.Method),
		zap.String("path", r.URL.Path),
		zap.String("protocol", r.Proto),
	)

	lw := newLoggingResponseWriter(w)

	h.handler.ServeHTTP(lw, r)

	requestStopTime := h.clock.Since(requestStartTime)

	logEntry.Info(
		"HTTP request received",
		zap.Int("status", lw.statusCode),
		zap.Duration("processing_time", requestStopTime),
	)
}

// getRealIP - returns real IP from http request
func (h *Handler) getRealIP(r *http.Request) string {
	var ip string

	remoteAddr := r.RemoteAddr

	if ip = r.Header.Get("X-Forwarded-For"); ip != "" {
		remoteAddr = strings.Split(ip, ", ")[0]
	} else if ip = r.Header.Get("X-Real-IP"); ip != "" {
		remoteAddr = ip
	} else {
		var err error
		remoteAddr, _, err = net.SplitHostPort(remoteAddr)
		if err != nil {
			h.log.Error("failed to find ip address", zap.Error(err))
		}
	}

	return remoteAddr
}
