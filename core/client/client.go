package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/optiop/next_century_sdk/core/schema"
)

// next century meters API endpoint
const (
	Login = "/Login"

	// GET /api/Properties/:id/DailyReads/?date&from&to // &from=%s&to=%s is optional
	DailyReads = "/api/Properties/%s/DailyReads/"

	// TODO: add more API endpoints
)

type client struct {
	apiURL string

	apiKey string
}

func New(apiURL string) *client {
	loginUrl, err := url.JoinPath(apiURL, Login)
	if err != nil {
		panic(err)
	}

	apiKey := genNewToken(loginUrl)

	return &client{
		apiURL: apiURL,
		apiKey: apiKey,
	}
}

func genNewToken(loginUrl string) string {
	email := os.Getenv("NCM_EMAIL")
	password := os.Getenv("NCM_PASS")

	postData := fmt.Sprintf(`{"email": "%s", "password": "%s"}`, email, password)

	resp, err := http.Post(loginUrl, "application/json", strings.NewReader(postData))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body := struct {
		Token string `json:"token"`
	}{}

	if err = json.NewDecoder(resp.Body).Decode(&body); err != nil {
		panic(err)
	}

	return body.Token
}

func (c *client) GetDailyReads(propertyID string, timeReq schema.TimeRequest) ([]*schema.MeterData, error) {
	DailyReadsUrl, err := url.JoinPath(c.apiURL, fmt.Sprintf(DailyReads, propertyID))
	if err != nil {
		return nil, err
	}

	if timeReq.Date.IsZero() {
		return nil, fmt.Errorf("Date is required")
	}

	DailyReadsUrl += fmt.Sprintf("?date=%s", timeReq.Date.Format("2006-01-02"))

	if !timeReq.From.IsZero() {
		DailyReadsUrl += fmt.Sprintf("&from=%s", timeReq.From.Format("2006-01-02"))
	}

	if !timeReq.To.IsZero() {
		DailyReadsUrl += fmt.Sprintf("&to=%s", timeReq.To.Format("2006-01-02"))
	}

	req, err := http.NewRequest("GET", DailyReadsUrl, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("authorization", c.apiKey)
	req.Header.Set("version", "2")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// if status is StatusUnauthorized gen new token
	if resp.StatusCode == http.StatusUnauthorized {
		loginUrl, err := url.JoinPath(c.apiURL, Login)
		if err != nil {
			panic(err)
		}
		c.apiKey = genNewToken(loginUrl)
		return c.GetDailyReads(propertyID, timeReq)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Failed to get daily reads: %s", resp.Status)
	}

	meterData := []*schema.MeterData{}
	if err = json.NewDecoder(resp.Body).Decode(&meterData); err != nil {
		return nil, err
	}

	return meterData, nil
}
