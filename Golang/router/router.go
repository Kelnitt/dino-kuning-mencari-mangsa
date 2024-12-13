package router

import (
	"Jur/controller"

	"github.com/gin-gonic/gin"
)

func SampleRouter(route *gin.Engine) {
	// Sample Router
	sample := route.Group("/sample")
	{
		// Get all samples
		sample.GET("/GetAll", controller.GetAll)

		// Get a specific sample by SampleID
		sample.GET("/GetSample/:SampleID", controller.GetSample)

		// Create a new sample
		sample.POST("/Create", controller.CreateSample)

		// Delete a sample by SampleID
		sample.DELETE("/Delete/:SampleID", controller.DeleteSample)

		// Update Sample
		sample.PUT("/Update/:SampleID", controller.UpdateSample)
	}
}
