package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ServerList(c *gin.Context) {

	// Making Request to Scrapper API
	resp, err := http.Post(fmt.Sprintf("%s/system/servers/list", GetAPIURL()), "application/json", nil)
	if err != nil {
		log.Fatal("Error making request to scrapper API")
	}
	defer resp.Body.Close()

	// Decoding
	var sl ServerListing
	err = json.NewDecoder(resp.Body).Decode(&sl.ServerList)
	if err != nil {
		log.Fatal("Error Decoding JSON")
	}

	// Response to Client
	c.JSON(http.StatusOK, sl)
}

func ActivateServer(c *gin.Context) {
	// Get Request Body
	var reqBody ActivateServerRequest
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	// Making Request to Activate Server
	data, _ := json.Marshal(reqBody)
	resp, err := http.Post(fmt.Sprintf("%s/system/servers/activate", GetAPIURL()), "application/json", bytes.NewBuffer(data))
	if err != nil {
		log.Fatal("Error making request to scrapper API")
	}
	defer resp.Body.Close()

	// Decoding Response
	var message SuccessResponse
	err = json.NewDecoder(resp.Body).Decode(&message)
	if err != nil {
		log.Fatal("Error Decoding JSON")
	}

	// Response to Client
	c.JSON(http.StatusOK, message)
}

func Redeploy(c *gin.Context) {
	// Making Request to Scrapper API
	resp, err := http.Post(fmt.Sprintf("%s/system/servers/redeploy", GetAPIURL()), "application/json", nil) //BAD REQUEST NOT SURE WHY
	if err != nil {
		log.Fatal("Error making request to scrapper API")
	}

	// Decoding
	var message SuccessResponse
	err = json.NewDecoder(resp.Body).Decode(&message)
	if err != nil {
		log.Fatal("Error Decoding")
	}

	// Response to Client
	c.JSON(http.StatusOK, message)
}

func Active(c *gin.Context) {
	// Making Request to Scrapper API
	resp, err := http.Post(fmt.Sprintf("%s/system/servers/active-list", GetAPIURL()), "application/json", nil)
	if err != nil {
		log.Fatal("Error making request to scrapper API")
	}

	// Decoding
	var message ActiveList
	err = json.NewDecoder(resp.Body).Decode(&message)
	if err != nil {
		log.Fatal("Error Decoding")
	}

	// Response to Client
	c.JSON(http.StatusOK, message)
}

func Details(c *gin.Context) {
	// Get Request Body
	var reqBody DetailRequest
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	// Making Request to Scrapper API
	data, _ := json.Marshal(reqBody)
	resp, err := http.Post(fmt.Sprintf("%s/system/logs/detail", GetAPIURL()), "application/json", bytes.NewBuffer(data))
	if err != nil {
		log.Fatal("Error making request to scrapper API")
	}

	// Reading
	details, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error Reading the Response Body")
	}

	// Response to Client
	c.Data(200, "text/plain", details)
}
