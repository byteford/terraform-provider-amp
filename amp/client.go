package amp

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Client struct {
	HostURL    string
	HTTPClient *http.Client
	sessionID  string
}
type login struct {
	username   string
	password   string
	token      string
	rememberMe bool
}

type coreLoging struct {
	sessionID string
}

func NewClient(host, username, password string, httpclient *http.Client) (*Client, error) {
	client := &Client{
		HTTPClient: httpclient,
		HostURL:    host,
		sessionID:  "",
	}
	log := login{username: "admin", password: "WdzWRQJLHkreUjRz7Zb8EL4C", token: "", rememberMe: false}
	rb, err := json.Marshal(log)
	if err != nil {
		return nil, err
	}
	body, err := apiCall(rb, "/API/Core/Login", http.MethodGet, *client)
	if err != nil {
		return nil, err
	}
	var corelog coreLoging
	err = json.Unmarshal(body, corelog)
	if err != nil {
		return nil, err
	}
	client.sessionID = corelog.sessionID
	return client, nil
}

func apiCall(body []byte, endpoint, method string, client Client) ([]byte, error) {
	url := fmt.Sprintf("%s/api/states/%s", client.HostURL, endpoint)
	req, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	res, err := client.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status: %d", res.StatusCode)
	}
	defer res.Body.Close()
	resbody, err := ioutil.ReadAll(res.Body)
	return resbody, nil
}
