package verifysphere_sdk

type Client struct {
	baseURL string
	params  Request
}

func NewClient(baseURL string) *Client {
	return &Client{
		baseURL: baseURL,
	}
}
