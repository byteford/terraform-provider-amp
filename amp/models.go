package amp

import "net/http"

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

type GetInst struct {
	SessionID  string `json:"SESSIONID"`
	InstanceID string `json:"InstanceID"`
}

type CreateInstanceObj struct {
	AutoConfigure     bool
	FriendlyName      string
	Module            string
	SessionID         string `json:"SESSIONID"`
	TargetADSInstance string
	NewInstanceId     string
	InstanceName      string
	PortNumber        int32
	ProvisionSettings map[string]string
}

type Instance struct {
	InstanceID   string `json:"InstanceID"`
	TargetID     string `json:"TargetID"`
	FriendlyName string
	InstanceName string
	Module       string
	Port         int32
	Running      bool
}
