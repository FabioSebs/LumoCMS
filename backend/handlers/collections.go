package handlers

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
)

type ServerListing struct {
	ServerList []Server
}

type Server struct {
	ID            int       `json:"id"`
	ServerCode    string    `json:"server_code"`
	ServerURL     string    `json:"server_url"`
	ScrapURL      string    `json:"scrap_url"`
	RegisterURL   string    `json:"register_url"`
	UnregisterURL string    `json:"unregister_url"`
	Nodes         int       `json:"nodes"`
	Enabled       bool      `json:"enabled"`
	CreatedAt     time.Time `json:"CreatedAt"`
	UpdatedAt     time.Time `json:"UpdatedAt"`
}

type ActivateServerRequest struct {
	Code    string `json:"server_code"`
	Enabled bool   `json:"enabled"`
}

type SuccessResponse struct {
	Message string `json:"message"`
}

type ActiveList struct {
	Index int64   `json:"index"`
	List  []AList `json:"list"`
}

type AList struct {
	CreatedAt     string `json:"CreatedAt"`
	UpdatedAt     string `json:"UpdatedAt"`
	Enabled       bool   `json:"enabled"`
	ID            int64  `json:"id"`
	Nodes         int64  `json:"nodes"`
	RegisterURL   string `json:"register_url"`
	ScrapURL      string `json:"scrap_url"`
	ServerCode    string `json:"server_code"`
	ServerURL     string `json:"server_url"`
	UnregisterURL string `json:"unregister_url"`
}

type DetailRequest struct {
	Name string `json:"name"`
	Type int    `json:"type"`
}

func GetAPIURL() string {
	// Env Vars
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error Loading Env Variables")
	}
	API_URL := os.Getenv("API_URL")
	return API_URL
}
