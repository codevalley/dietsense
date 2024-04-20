package logging

import (
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// Log is the exported Logger
var Log *logrus.Logger

// Setup initializes the logger
func Setup() {
	Log = logrus.New()
	Log.SetOutput(os.Stdout)
	Log.SetLevel(logrus.InfoLevel) // Default level is Info, change it according to need

	// Set formatter, you can change to JSON if you prefer structured logs
	Log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
}

// GinLogger returns a gin.HandlerFunc (middleware) that logs requests using Logrus
func GinLogger(logger *logrus.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Starting time
		start := time.Now()

		// Process request
		c.Next()

		// End time
		latency := time.Since(start)
		status := c.Writer.Status()
		logger.WithFields(logrus.Fields{
			"status":  status,
			"latency": latency,
			"ip":      c.ClientIP(),
			"method":  c.Request.Method,
			"path":    c.Request.URL.Path,
		}).Info("request details")
	}
}
