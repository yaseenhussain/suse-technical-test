package main

import (
        "io/ioutil"
        "net/http"
        "net/http/httptest"
        "strings"
        "testing"
)

func TestLogin(t *testing.T) {
        handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
                // Verify the request method is POST
                if r.Method != http.MethodPost {
                        t.Errorf("expected %s request, got %s", http.MethodPost, r.Method)
                        w.WriteHeader(http.StatusMethodNotAllowed)
                        return
                }

                // Verify the request URL matches the expected login endpoint
                if r.URL.Path != "/v3-public/localProviders/local" || r.URL.Query().Get("action") != "login" {
                        t.Errorf("unexpected request URL: %s", r.URL.Path)
                        w.WriteHeader(http.StatusNotFound)
                        return
                }

                // Read the request body
                body, err := ioutil.ReadAll(r.Body)
                if err != nil {
                        t.Errorf("failed to read request body: %v", err)
                        w.WriteHeader(http.StatusInternalServerError)
                        return
                }

                // Verify the request body contains the expected credentials
                expectedBody := `{
                        "username": "admin",
                        "password": "SuSetechnicaltest"
                }`
                if string(body) != expectedBody {
                        t.Errorf("unexpected request body: %s", body)
                        w.WriteHeader(http.StatusBadRequest)
                        return
                }

                // Write the response body
                responseBody := "success"
                _, err = w.Write([]byte(responseBody))
                if err != nil {
                        t.Errorf("failed to write response body: %v", err)
                }
        })

        // Create a test server with the handler
        server := httptest.NewServer(handler)
        defer server.Close()

        // Prepare the request body
        requestBody := `{
                "username": "admin",
                "password": "SuSetechnicaltest"
        }`
        requestBodyReader := strings.NewReader(requestBody)

        // Send a POST request to the test server
        resp, err := http.Post(server.URL+"/v3-public/localProviders/local?action=login", "application/json", requestBodyReader)
        if err != nil {
                t.Fatalf("failed to send request: %v", err)
        }
        defer resp.Body.Close()

        // Check the response status code
        if resp.StatusCode != http.StatusOK {
                t.Errorf("expected status %v, got %v", http.StatusOK, resp.StatusCode)
        }

        // Read and check the response body
        body, err := ioutil.ReadAll(resp.Body)
        if err != nil {
                t.Fatalf("failed to read response body: %v", err)
        }

        expectedResponseBody := "success"
        if string(body) != expectedResponseBody {
                t.Errorf("expected response body %q, got %q", expectedResponseBody, string(body))
        }
}