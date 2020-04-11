package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	VerifyURL = "https://www.google.com/recaptcha/api/siteverify"
)

type Verifier struct {
	client *http.Client
}

type VerifyResponse struct {
	Success     bool      `json:"success"`
	ChallengeTs time.Time `json:"challenge_ts"`
	Hostname    string    `json:"hostname"`
	Score       float64   `json:"score"`
	Action      string    `json:"action"`
	ErrorCodes  []string  `json:"error-codes"`
}

func NewVerifier(client *http.Client) *Verifier {
	return &Verifier{client: client}
}

func (v *Verifier) Verify(ctx context.Context, secret, response, remoteIP string) (resp *VerifyResponse, err error) {
	form := url.Values{}
	form.Add("secret", secret)
	form.Add("response", response)
	if remoteIP != "" {
		form.Add("remoteip", remoteIP)
	}

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		VerifyURL,
		strings.NewReader(form.Encode()),
	)
	if err != nil {
		return nil, err
	}

	req.Header.Set("content-type", "application/x-www-form-urlencoded")

	res, err := v.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	bodyBytes, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		return nil, readErr
	}

	resp = &VerifyResponse{}
	jsonErr := json.Unmarshal(bodyBytes, &resp)
	if jsonErr != nil {
		return nil, jsonErr
	}

	return
}
