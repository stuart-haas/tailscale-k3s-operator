package tailscale

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"golang.org/x/oauth2/clientcredentials"
	"sigs.k8s.io/controller-runtime/pkg/log"
)

type ClientConfig struct {
	ClientId     string
	ClientSecret string
	OrgName      string
}

type Client struct {
	baseURL string
	config  ClientConfig
}

// Device represents a Tailscale device
type Device struct {
	ID        string    `json:"id"`
	Hostname  string    `json:"hostname"`
	IPAddress string    `json:"ipAddress"`
	Tags      []string  `json:"tags"`
	LastSeen  time.Time `json:"lastSeen"`
}

// NewClient creates a new Tailscale API client
func NewClient(config ClientConfig) *Client {
	log := log.FromContext(context.Background())
	log.Info("Creating Tailscale API client")
	return &Client{
		baseURL: "https://api.tailscale.com/api/v2",
		config:  config,
	}
}

func (c *Client) getHTTPClient() *http.Client {
	oauthConfig := clientcredentials.Config{
		ClientID:     c.config.ClientId,
		ClientSecret: c.config.ClientSecret,
		TokenURL:     fmt.Sprintf("%s/oauth/token", c.baseURL),
	}
	client := oauthConfig.Client(context.Background())
	return client
}

// ListDevices returns all devices in the Tailscale network
func (c *Client) ListDevices(ctx context.Context) ([]Device, error) {
	httpClient := c.getHTTPClient()
	res, err := httpClient.Get(fmt.Sprintf("%s/tailnet/%s/devices", c.baseURL, c.config.OrgName))
	if err != nil {
		return nil, fmt.Errorf("getting devices: %v", err)
	}

	var devices []Device
	if err := json.NewDecoder(res.Body).Decode(&devices); err != nil {
		return nil, fmt.Errorf("decoding response: %w", err)
	}

	return devices, nil
}
