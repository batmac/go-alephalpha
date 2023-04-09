package alephalpha

import "net/http"

//nolint:tagliatelle
type QARequest struct {
	Hosting    *string    `json:"hosting,omitempty"`
	Query      string     `json:"query"`
	Documents  []Document `json:"documents"`
	MaxAnswers int        `json:"max_answers"`
}

//nolint:tagliatelle
type QAResponse struct {
	ModelVersion string     `json:"model_version"`
	Answers      []QAAnswer `json:"answers"`
}

type QAAnswer struct {
	Answer   string  `json:"answer"`
	Score    float64 `json:"score"`
	Evidence string  `json:"evidence"`
}

func (c *Client) QA(r QARequest) (QAResponse, error) {
	return req[QARequest, QAResponse](c, http.MethodPost, "qa", r)
}
