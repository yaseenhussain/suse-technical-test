ackage main

import (
    "testing"
    "github.com/rancher/go-rancher/client"
)

func TestLogin(t *testing.T) {
    // Create a new Rancher client.
    client, err := client.NewClient("https://localhost/v3-public/localProviders/local?action=login", "admin", "SuSetechnicaltest")
    if err != nil {
        t.Errorf("Error creating Rancher client: %v", err)
    }

    // Login to Rancher.
    err = client.Login()
    if err != nil {
        t.Errorf("Error logging in to Rancher: %v", err)
    }

    // Check if the user is logged in.
    if !client.IsLoggedIn() {
        t.Errorf("User is not logged in")
    }
}