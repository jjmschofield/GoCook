package logger

import (
	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
	"go.uber.org/zap"
	"time"
)

func ContextLoggerMiddleware(service string) func(context *gin.Context) {
	return func(context *gin.Context) {
		ctxLogger := WithContext(
			zap.String("service", service),
			zap.String("correlationId", uuid.Must(uuid.NewV4()).String()),
		)

		context.Set("logger", ctxLogger)
	}
}

func RequestLoggerMiddleware(context *gin.Context) {
	logger := context.MustGet("logger").(zap.Logger)

	path := context.Request.URL.Path
	start := time.Now()

	logRequestStart(logger, context, path, start)

	context.Next()

	logRequestEnd(logger, context, path, start)
}

func logRequestStart(logger zap.Logger, context *gin.Context, path string, start time.Time) {
	logger.Info("Request started "+ path, getRequestContext(context, start)...)
}

func logRequestEnd(logger zap.Logger, context *gin.Context,  path string, start time.Time) {
	if len(context.Errors) > 0 {
		for _, e := range context.Errors.Errors() {
			logger.Error(e)
		}
	} else{
		status := context.Writer.Status()
		requestContext := getEndRequestContext(context, start)

		if status > 499 {
			logger.Error("Request failed " + path, requestContext...)
		} else if status > 399{
			logger.Warn("Request rejected " + path, requestContext...)
		} else{
			logger.Info("Request successful " + path, requestContext...)
		}
	}
}

func getRequestContext(context *gin.Context, start time.Time) (fields []zap.Field) {
	fields = []zap.Field{
		zap.Int("status", context.Writer.Status()),
		zap.String("method", context.Request.Method),
		zap.String("path", context.Request.URL.Path),
		zap.String("query", context.Request.URL.RawQuery),
		zap.String("ip", context.ClientIP()),
		zap.String("user-agent", context.Request.UserAgent()),
		zap.String("start", start.String()),
	}
	return fields
}

func getEndRequestContext(context *gin.Context, start time.Time) (fields []zap.Field) {
	fields = getRequestContext(context, start)

	end := time.Now()
	latency := end.Sub(start)
	fields = append(fields, zap.Duration("latency", latency))
	fields = append(fields, zap.String("end", end.String()))

	return fields
}
