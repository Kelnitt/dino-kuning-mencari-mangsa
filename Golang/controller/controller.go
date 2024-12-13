package controller

import (
	"Jur/config"
	"Jur/entities"
	"Jur/utilities"
	"log"

	"github.com/gin-gonic/gin"
)

func MainHallo(c *gin.Context) {
	// Main Page on Sample System
	c.JSON(200, gin.H{"result": "Hallo ! Simple Application !"})
}

func GetAll(c *gin.Context) {
	// Get All Sample Table
	database, errtable := config.GalaSetup()

	if errtable != nil {
		utilities.FailMess(c, 500, "Fail ! Connection Unavailable !")
		return
	}

	var samplesresult []entities.SampleTabler

	result := database.Find(&samplesresult)

	if result.Error != nil {
		utilities.FailMess(c, 500, "Fail to Get Sample Data", result.Error.Error())
		return
	}

	if result.RowsAffected == 0 {
		utilities.FailMess(c, 404, "Data Unavailable !")
		return
	}

	c.JSON(200, samplesresult)
}

func GetSample(c *gin.Context) {
	// Get Sample ID
	SampleID := c.Param("SampleID")

	if SampleID == "" {
		utilities.FailMess(c, 400, "Sample ID is Missing !")
		return
	}

	database, errtable := config.GalaSetup()

	if errtable != nil {
		utilities.FailMess(c, 500, "Fail ! Connection Unavailable !", errtable.Error())
		return
	}

	var sample entities.SampleTabler

	if GetFail := database.First(&sample, SampleID).Error; GetFail != nil {
		if GetFail.Error() == "record not found" {
			utilities.FailMess(c, 404, "Sample Unavailable !")
		}
		return
	}

	c.JSON(200, sample)
}

func CreateSample(c *gin.Context) {
	// Create Sample Data
	var samples []entities.SampleTabler

	if json_error := c.ShouldBindJSON(&samples); json_error != nil {
		utilities.FailMess(c, 400, "Fail to Bind JSON", json_error.Error())
		return
	}

	database, err := config.GalaSetup()

	if err != nil {
		log.Fatalf("Fail to Connect to Database : %v", err)
		utilities.FailMess(c, 500, "Database Connection Fail !")
		return
	}

	// Create Batches Insert to Table
	if create_err := database.CreateInBatches(samples, len(samples)).Error; create_err != nil {
		utilities.FailMess(c, 500, "Fail to Create Multiple Sample !")
	}

	c.JSON(200, gin.H{"result": "Create Sample ID !"})
}

func DeleteSample(c *gin.Context) {
	// Delete Sample
	SampleID := c.Param("SampleID")

	if SampleID == "" {
		utilities.FailMess(c, 400, "Sample ID is Missing !")
		return
	}

	database, errtable := config.GalaSetup()

	if errtable != nil {
		utilities.FailMess(c, 500, "Fail ! Connection Unavailable !", errtable.Error())
		return
	}

	var sample entities.SampleTabler

	if err := database.First(&sample, SampleID).Error; err != nil {
		// If no record is found, return a "not found" error
		utilities.FailMess(c, 404, "Sample Unavailable !")
		return
	}

	if err := database.Delete(&entities.SampleTabler{}, SampleID).Error; err != nil {
		utilities.FailMess(c, 500, "Error deleting sample: "+err.Error())
		return
	}

	c.JSON(200, gin.H{"message": "Delete Success !"})
}

func UpdateSample(c *gin.Context) {
	// Update Sample by SampleID
	SampleID := c.Param("SampleID")

	if SampleID == "" {
		utilities.FailMess(c, 400, "Sample ID is Missing !")
		return
	}

	// Bind JSON to SampleTabler struct
	var UpdateSampleData entities.SampleTabler

	if json_error := c.ShouldBindJSON(&UpdateSampleData); json_error != nil {
		utilities.FailMess(c, 400, "Fail to Bind JSON", json_error.Error())
		return
	}

	// Connect to database
	database, err := config.GalaSetup()

	if err != nil {
		utilities.FailMess(c, 500, "Fail ! Connection Unavailable !", err.Error())
		return
	}

	// Check if the sample exists
	var existingSample entities.SampleTabler

	if err := database.First(&existingSample, SampleID).Error; err != nil {
		if err.Error() == "record not found" {
			utilities.FailMess(c, 404, "Sample Unavailable !")
		}
		return
	}

	// Perform the update
	if err := database.Model(&existingSample).Updates(UpdateSampleData).Error; err != nil {
		utilities.FailMess(c, 500, "Fail to Update Sample", err.Error())
		return
	}

	c.JSON(200, gin.H{"message": "Update Success !"})
}
