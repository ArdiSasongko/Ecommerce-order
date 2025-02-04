package external

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/ArdiSasongko/Ecommerce-order/internal/config/env"
	"github.com/joho/godotenv"
)

type Response struct {
	Data Data `json:"data"`
}

type Data struct {
	ID          int32  `json:"id"`
	Username    string `json:"username"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	Address     string `json:"address"`
	DoB         string `json:"dob"`
	Password    string `json:"-"`
	Fullname    string `json:"fullname"`
	Role        int32  `json:"Role"`
}

type UserExternal struct {
	httpClient *http.Client
}

func (e *UserExternal) Profile(ctx context.Context, token string) (*Response, error) {
	if err := godotenv.Load(); err != nil {
		return nil, err
	}

	url := env.GetEnvString("USER_BASE_URL", "") + "/profile"

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request :%w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	resp, err := e.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to get profile :%w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body :%w", err)
	}

	log.Println("response body", string(body))

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("user service returned error :%s", string(body))
	}

	var response Response
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, fmt.Errorf("failed to marshalling body :%w", err)
	}

	return &response, nil
}
