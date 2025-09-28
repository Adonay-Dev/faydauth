package faydauth

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type FaydaClient struct {
	BaseURL string      
	Client  *http.Client 
}

func NewFaydaClient(baseURL string) *FaydaClient {
	return &FaydaClient{
		BaseURL: baseURL,
		Client: &http.Client{
			Timeout: 10 * time.Second, 
		},
	}
}

func (f *FaydaClient) Authenticate(sessionID, authCode, csrfToken string) (string, error) {
	body := map[string]string{
		"session_id": sessionID,
		"auth_code":  authCode,
		"csrf_token": csrfToken,
	}

	data, err := json.Marshal(body)
	if err != nil {
		return "", fmt.Errorf("failed to marshal request body: %w", err)
	}

	resp, err := f.Client.Post(f.BaseURL+"/authenticate", "application/json", bytes.NewBuffer(data))
	if err != nil {
		return "", fmt.Errorf("failed to call Fayda API: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("Fayda API returned status %d", resp.StatusCode)
	}

	var res struct {
		Data struct {
			UserInfo struct {
				Sub string `json:"sub"`
			} `json:"user_info"`
		} `json:"data"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return "", fmt.Errorf("failed to decode Fayda API response: %w", err)
	}

	return res.Data.UserInfo.Sub, nil
}
