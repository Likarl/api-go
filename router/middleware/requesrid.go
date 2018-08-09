package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
)

func RequestId() gin.HandlerFunc {
	return func(c *gin.Context) {
		// check for incoming header , use it if exists
		requestId := c.Request.Header.Get("X-Request-Id")

		// create request id with uuid4
		if requestId == "" {
			u4, _ := uuid.NewV4()
			requestId = u4.String()
		}

		// expose it for use in the application
		c.Set("X-Request-Id", requestId)

		// set X-request-Id header
		c.Writer.Header().Set("X-Request-Id", requestId)
		c.Next()
	}

}
