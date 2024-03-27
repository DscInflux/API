package middlewares

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

// CordxResponse represents the response from the Cordx API.
type CordxResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    struct {
		URL string `json:"url"`
	} `json:"data"`
}

// UploadImage uploads an image to the cordx.lol API.
func UploadImage(name string, base64Image string) (*CordxResponse, error) {
	// Prepare the request body
	formData := strings.NewReader("sharex=" + ParseBase64(base64Image))

	// Create a new HTTP client
	client := &http.Client{}

	// Create a POST request to the Cordx API
	req, err := http.NewRequest("POST", "https://cordx.lol/api/upload/sharex", formData)
	if err != nil {
		return nil, err
	}

	// Set request headers
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("userid", os.Getenv("CordxUserID"))
	req.Header.Set("secret", os.Getenv("CordxSecret"))

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Parse the JSON response
	cordxResponse := CordxResponse{}
	err = json.Unmarshal(body, &cordxResponse)
	if err != nil {
		return nil, err
	}

	return &cordxResponse, nil
}

// ParseBase64 extracts the base64-encoded image data from a data URL.
func ParseBase64(base64 string) string {
	parts := strings.Split(base64, ",")
	if len(parts) == 2 {
		return parts[1]
	}
	return ""
}
