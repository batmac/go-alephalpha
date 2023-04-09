package alephalpha

import "net/http"

//nolint:tagliatelle
type AvailableModel struct {
	Name                    string   `json:"name"`
	Description             string   `json:"description"`
	MaxContextSize          int      `json:"max_context_size"`
	Hostings                []string `json:"hostings"`
	ImageSupport            bool     `json:"image_support"`
	QASupport               bool     `json:"qa_support"`
	SummarizationSupport    bool     `json:"summarization_support"`
	EmbeddingTypes          []string `json:"embedding_types"`
	MaximumCompletionTokens *int     `json:"maximum_completion_tokens"`
}

func (c *Client) ListAvailableModels() ([]AvailableModel, error) {
	var err error
	if len(c.availableModels) == 0 {
		c.availableModels, err = req[any, []AvailableModel](c, http.MethodGet, "models_available", nil)
	}
	return c.availableModels, err
}

type Model string

const (
	ModelLuminousExtended       Model = "luminous-extended"
	ModelLuminousBase           Model = "luminous-base"
	ModelLuminousSupreme        Model = "luminous-supreme"
	ModelLuminousSupremeControl Model = "luminous-supreme-control"
)

/* [
  {
    "name": "luminous-extended",
    "description": "Multilingual model trained on English, German, French, Spanish and Italian",
    "max_context_size": 2048,
    "hostings": [
      "aleph-alpha"
    ],
    "image_support": true,
    "qa_support": false,
    "summarization_support": false,
    "embedding_types": [],
    "maximum_completion_tokens": null
  },
  {
    "name": "luminous-base",
    "description": "Multilingual model trained on English, German, French, Spanish and Italian",
    "max_context_size": 2048,
    "hostings": [
      "aleph-alpha"
    ],
    "image_support": true,
    "qa_support": false,
    "summarization_support": false,
    "embedding_types": [
      "symmetric_128",
      "asymmetric_128_document",
      "asymmetric_128_query",
      "symmetric",
      "asymmetric_document",
      "asymmetric_query"
    ],
    "maximum_completion_tokens": null
  },
  {
    "name": "luminous-supreme",
    "description": "Multilingual model trained on English, German, French, Spanish and Italian",
    "max_context_size": 2048,
    "hostings": [
      "aleph-alpha"
    ],
    "image_support": false,
    "qa_support": false,
    "summarization_support": false,
    "embedding_types": [],
    "maximum_completion_tokens": 1990
  },
  {
    "name": "luminous-supreme-control",
    "description": "A variant of the Luminous-Supreme model that is optimized for downstream task performance",
    "max_context_size": 2048,
    "hostings": [
      "aleph-alpha"
    ],
    "image_support": false,
    "qa_support": false,
    "summarization_support": false,
    "embedding_types": [],
    "maximum_completion_tokens": 1990
  }
] */
