package alephalpha

import "net/http"

type SummarizeRequest struct {
	Model    Model    `json:"model" binding:"required"`
	Hosting  string   `json:"hosting,omitempty"`
	Document Document `json:"document" binding:"required" validate:"required,oneof=Docx Text Prompt"`
}

//nolint:tagliatelle
type SummarizeResponse struct {
	ModelVersion string `json:"model_version,omitempty"`
	Summary      string `json:"summary,omitempty"`
}

func (c *Client) Summarize(r SummarizeRequest) (SummarizeResponse, error) {
	return req[SummarizeRequest, SummarizeResponse](c, http.MethodPost, "summarize", r)
}
