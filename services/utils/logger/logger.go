package logger

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

// LoggerKey specifies the key used to store the logger in context
const LoggerKey = "logger"

// NewLogrusLogger initializes a new logrus logger
func NewLogrusLogger() *logrus.Logger {
	logrusLogger := logrus.New()
	logrusLogger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
	logrusLogger.SetLevel(logrus.DebugLevel)
	return logrusLogger
}

// InitLoggerMiddleware middleware to inject logrus logger and log internal server errors
func InitLoggerMiddleware() fiber.Handler {
	logrusLogger := NewLogrusLogger()
	return func(c *fiber.Ctx) error {
		// Add the logger entry to the context
		entry := logrusLogger.WithFields(logrus.Fields{
			"request_id": c.GetRespHeader("X-Request-ID"),
			"uri":        c.OriginalURL(),
			"method":     c.Method(),
		})
		c.Locals(LoggerKey, entry)

		// Process the request
		err := c.Next()
		// Log request and response if status code is 500 Internal Server Error
		if c.Response().StatusCode() >= http.StatusInternalServerError {
			reqBody := c.Request().Body()
			resBody := c.Response().Body()
			stackTrace := fmt.Sprintf("%+v", err)
			entry.WithFields(logrus.Fields{
				"request_body":  string(reqBody),
				"status_code":   c.Response().StatusCode(),
				"response_body": string(resBody),
				"stackTrace":    stackTrace,
			}).Error("Internal Server Error")
		}

		return err
	}
}

// GetLogger retrieves the logger from the context
func GetLogger(c *fiber.Ctx) *logrus.Entry {
	logger, ok := c.Locals(LoggerKey).(*logrus.Entry)
	if !ok {
		logger = logrus.NewEntry(logrus.StandardLogger())
	}
	return logger
}
