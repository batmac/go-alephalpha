package alephalpha

import (
	"net/http"
	"time"
)

//nolint:tagliatelle
type RecentUsage struct {
	CreateTimestamp      time.Time `json:"create_timestamp"`
	ModelName            string    `json:"model_name"`
	RequestType          string    `json:"request_type"`
	TokenCountPrompt     int       `json:"token_count_prompt"`
	ImageCountPrompt     int       `json:"image_count_prompt"`
	TokenCountCompletion int       `json:"token_count_completion"`
	DurationMillis       int       `json:"duration_millis"`
	Credits              float64   `json:"credits"`
}

func (c *Client) QueryRecentUsage() ([]RecentUsage, error) {
	return req[any, []RecentUsage](c, http.MethodGet, "users/me/requests", nil)
}
