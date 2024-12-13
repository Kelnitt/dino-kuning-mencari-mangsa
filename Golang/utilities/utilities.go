package utilities

import "github.com/gin-gonic/gin"

func FailMess(c *gin.Context, statusCode int, message string, details ...string) {
	// FailMess Return a JSON Response with an Error Message and Optional Details
	response := gin.H{"error": message}

	if len(details) > 0 {
		response["details"] = details[0]
	}

	c.JSON(statusCode, response)
}
