package main

import (
	"btc-billionaire/conf"
	"btc-billionaire/controllers"
	"btc-billionaire/models"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func TestNewRecord(t *testing.T) {
	conf.ConnectDatabase()

	gin.SetMode(gin.TestMode)

	r := SetUpRouter()
	r.POST("/records", controllers.CreateRecord)

	dt, _ := time.Parse(time.RFC3339, "2011-10gf-05T18+00:00")

	record := models.Record{
		Amount:   1000,
		DateTime: dt,
	}
	jsonValue, _ := json.Marshal(record)
	req, _ := http.NewRequest("POST", "/records", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestShowHistory(t *testing.T) {
	conf.ConnectDatabase()
	gin.SetMode(gin.TestMode)

	r := SetUpRouter()
	r.GET("/showHistory", controllers.ShowHistory)

	startTime := "2018-10-05T18:48:02+00:00"
	endTime := "2019-10-05T18:48:02+00:00"

	record := models.StartEndTimeRecordInput{
		StartTime: startTime,
		EndTime:   endTime,
	}

	jsonValue, _ := json.Marshal(record)
	fmt.Println(jsonValue)
	req, _ := http.NewRequest("GET", "/showHistory", bytes.NewBuffer(jsonValue))

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

}
