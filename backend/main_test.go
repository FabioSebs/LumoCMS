package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestServerList(t *testing.T) {
	// Make a GET Request
	router := gin.Default()
	req, _ := http.NewRequest("GET", "/v1/serverlist", nil)
	res := httptest.NewRecorder()
	router.ServeHTTP(res, req)

	status := res.Code
	if status != http.StatusOK {
		t.Error(res)
	}
	assert.Equal(t, http.StatusOK, res.Code)
}

func TestActivate(t *testing.T) {
	// JSON DATA
	values := map[string]interface{}{"server_code": "azure_02", "enabled": true}
	jsonValue, _ := json.Marshal(values)
	payload := bytes.NewBuffer(jsonValue)

	// Make a POST Request
	router := gin.Default()
	req, _ := http.NewRequest("POST", "/v1/activate", payload)
	res := httptest.NewRecorder()
	router.ServeHTTP(res, req)

	status := res.Code
	if status != http.StatusOK {
		t.Error(res)
	}
	assert.Equal(t, http.StatusOK, res.Code)
}

func TestRedeploy(t *testing.T) {
	// Make a GET Request
	router := gin.Default()
	req, _ := http.NewRequest("GET", "/v1/redeploy", nil)
	res := httptest.NewRecorder()
	router.ServeHTTP(res, req)

	status := res.Code
	if status != http.StatusOK {
		t.Error(res)
	}
	assert.Equal(t, http.StatusOK, res.Code)
}

func TestActiveList(t *testing.T) {
	// Make a GET Request
	router := gin.Default()
	req, _ := http.NewRequest("GET", "/v1/active_list", nil)
	res := httptest.NewRecorder()
	router.ServeHTTP(res, req)

	status := res.Code
	if status != http.StatusOK {
		t.Error(res)
	}
	assert.Equal(t, http.StatusOK, res.Code)
}

func TestDetails(t *testing.T) {
	// JSON DATA
	values := map[string]interface{}{"name": "", "type": -1}
	jsonValue, _ := json.Marshal(values)
	payload := bytes.NewBuffer(jsonValue)

	// Make a POST Request
	router := gin.Default()
	req, _ := http.NewRequest("POST", "/v1/details", payload)
	res := httptest.NewRecorder()
	router.ServeHTTP(res, req)

	status := res.Code
	if status != http.StatusOK {
		t.Error(res)
	}
	assert.Equal(t, http.StatusOK, res.Code)
}
