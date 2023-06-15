package main

import (
	"testing"

	rancher "github.com/rancher/rancher-sdk-go"
)

func TestLogin(t *testing.T) {
	// Create a new Rancher API client.
	apiClient, err := rancher.NewRancherClient(&rancher.ClientOpts{
		URL:       "https://localhost/v3-public/localProviders/local?action=login",
		AccessKey: "admin",
		SecretKey: "SuSetechnicaltest",
	})
	if err != nil {
		t.Fatalf("Error creating Rancher API client: %v", err)
	}

	// Login to Rancher and obtain an API token.
	_, err = apiClient.Login(&rancher.AuthOpts{
		Username: "admin",
		Password: "SuSetechnicaltest",
	})
	if err != nil {
		t.Fatalf("Error logging in to Rancher: %v", err)
	}

	// Check if the user is authenticated.
	if !apiClient.IsAuthActive() {
		t.Fatalf("User is not authenticated")
	}
}