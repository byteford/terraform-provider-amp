package amp

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

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
	err = json.Unmarshal(body, &corelog)
	if err != nil {
		return nil, err
	}
	client.sessionID = corelog.sessionID
	return client, nil
}

func CreateInstance(client Client, obj CreateInstanceObj) error {
	obj.SessionID = client.sessionID
	rb, err := json.Marshal(obj)
	if err != nil {
		return err
	}
	_, err = apiCall(rb, "/API/ADSModule/CreateInstance", http.MethodPost, client)
	if err != nil {
		return err
	}
	return nil
}

func GetInstance(client Client, InstID string) (Instance, error) {
	var inst Instance
	auth := &GetInst{SessionID: client.sessionID, InstanceID: InstID}
	rb, err := json.Marshal(auth)
	if err != nil {
		return inst, err
	}
	body, err := apiCall(rb, "/API/ADSModule/GetInstance", http.MethodPost, client)
	if err != nil {
		return inst, err
	}
	err = json.Unmarshal(body, &inst)
	if err != nil {
		return inst, err
	}
	return inst, nil
}

func apiCall(body []byte, endpoint, method string, client Client) ([]byte, error) {
	url := fmt.Sprintf("%s%s", client.HostURL, endpoint)
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
