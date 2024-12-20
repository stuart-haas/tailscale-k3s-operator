package tailscale

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// Client represents a Tailscale API client
type Client struct {
    httpClient   *http.Client
    baseURL      string
    clientID     string
    clientSecret string
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
func NewClient(clientID, clientSecret string) *Client {
    return &Client{
        httpClient: &http.Client{
            Timeout: time.Second * 10,
        },
        baseURL:      "https://api.tailscale.com/api/v2",
        clientID:     clientID,
        clientSecret: clientSecret,
    }
}

// ListDevices returns all devices in the Tailscale network
func (c *Client) ListDevices(ctx context.Context) ([]Device, error) {
    req, err := http.NewRequestWithContext(ctx, "GET", fmt.Sprintf("%s/devices", c.baseURL), nil)
    if err != nil {
        return nil, fmt.Errorf("creating request: %w", err)
    }

    req.SetBasicAuth(c.clientID, c.clientSecret)

    resp, err := c.httpClient.Do(req)
    if err != nil {
        return nil, fmt.Errorf("executing request: %w", err)
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
    }

    var devices []Device
    if err := json.NewDecoder(resp.Body).Decode(&devices); err != nil {
        return nil, fmt.Errorf("decoding response: %w", err)
    }

    return devices, nil
}