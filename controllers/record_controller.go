package controllers

import (
	"btc-billionaire/conf"
	"btc-billionaire/models"
	"fmt"
	"net/http"
	"regexp"
	"time"

	"github.com/gin-gonic/gin"
)

type Payload struct {
	Amount   int    `json:"Amount"`
	Datetime string `json:"datetime"`
}

func GetRecords(c *gin.Context) {
	var records []models.Record
	conf.DB.Find(&records)
	c.JSON(http.StatusOK, gin.H{"data": records})
}

func ShowHistory(c *gin.Context) {
	var history []models.Record
	var input models.StartEndTimeRecordInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	startTime, _ := time.Parse(time.RFC3339, input.StartTime)
	endTime, _ := time.Parse(time.RFC3339, input.EndTime)

	conf.DB.Where("datetime BETWEEN ? AND ?", startTime, endTime).Find(&history)
	c.JSON(http.StatusOK, gin.H{"data": history})
}

func CreateRecord(c *gin.Context) {
	var input models.CreateRecordInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := validateNewRecordPayload(input)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	txtTime, _ := time.Parse(time.RFC3339, input.DateTime)
	record := models.Record{Amount: input.Amount, DateTime: txtTime}

	conf.DB.Create(record)
	c.JSON(http.StatusCreated, gin.H{"data": record})
}

func validateNewRecordPayload(p models.CreateRecordInput) error {
	// Check if the Datetime field is in the correct format.
	datetimeRegexp := regexp.MustCompile(`\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}Z`)
	if !datetimeRegexp.MatchString(p.DateTime) {
		return fmt.Errorf("The Datetime field is not in the correct format: %s", p.DateTime)
	}

	// Try to parse the Datetime field as a time.Time.
	_, err := time.Parse(time.RFC3339, p.DateTime)
	if err != nil {
		return fmt.Errorf("The Datetime field is not a valid time: %v", err)
	}

	return nil
}
