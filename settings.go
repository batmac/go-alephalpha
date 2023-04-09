package alephalpha

import "net/http"

//nolint:tagliatelle
type Settings struct {
	ID                    int    `json:"id"`
	Email                 string `json:"email"`
	Role                  string `json:"role"`
	CreditsRemaining      any    `json:"credits_remaining"`
	InvoiceAllowed        bool   `json:"invoice_allowed"`
	OutOfCreditsThreshold any    `json:"out_of_credits_threshold"`
	TermsOfServiceVersion string `json:"terms_of_service_version"`
}

func (c *Client) QuerySettings() (Settings, error) {
	return req[any, Settings](c, http.MethodGet, "users/me", nil)
}
