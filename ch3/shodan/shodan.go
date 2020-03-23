package shodan

const BaseURL = "https://api.shodan.io"

type Client struct {
	apiKey string
}

func New(apiKeyey string) *Client {
	return &Client{apiKey: apiKeyey}
}
