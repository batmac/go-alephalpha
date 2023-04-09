//nolint:tagliatelle
package alephalpha

import "net/http"

type CompletionRequest struct {
	Model         Model       `json:"model" binding:"required"`
	Prompt        interface{} `json:"prompt" binding:"required"`
	MaximumTokens int         `json:"maximum_tokens" binding:"required"`

	Hosting           string      `json:"hosting,omitempty"`
	MinimumTokens     int         `json:"minimum_tokens,omitempty"`
	Echo              bool        `json:"echo,omitempty"`
	Temperature       float64     `json:"temperature,omitempty"`
	TopK              int         `json:"top_k,omitempty"`
	TopP              float64     `json:"top_p,omitempty"`
	PresencePenalty   float64     `json:"presence_penalty,omitempty"`
	FrequencyPenalty  float64     `json:"frequency_penalty,omitempty"`
	SequencePenalty   float64     `json:"sequence_penalty,omitempty"`
	SequencePenLength int         `json:"sequence_penalty_min_length,omitempty"`
	RepPenIncPrompt   bool        `json:"repetition_penalties_include_prompt,omitempty"`
	RepPenIncComp     bool        `json:"repetition_penalties_include_completion,omitempty"`
	UseMulPresPen     bool        `json:"use_multiplicative_presence_penalty,omitempty"`
	UseMulFreqPen     bool        `json:"use_multiplicative_frequency_penalty,omitempty"`
	UseMulSeqPen      bool        `json:"use_multiplicative_sequence_penalty,omitempty"`
	PenaltyBias       string      `json:"penalty_bias,omitempty"`
	PenaltyExceptions []string    `json:"penalty_exceptions,omitempty"`
	PenExclIncStopSeq bool        `json:"penalty_exceptions_include_stop_sequences,omitempty"`
	BestOf            int         `json:"best_of,omitempty"`
	N                 int         `json:"n,omitempty"`
	LogitBias         interface{} `json:"logit_bias,omitempty"`
	LogProbs          int         `json:"log_probs,omitempty"`
	StopSequences     []string    `json:"stop_sequences,omitempty"`
	Tokens            bool        `json:"tokens,omitempty"`
	RawCompletion     bool        `json:"raw_completion,omitempty"`
	DisableOptim      bool        `json:"disable_optimizations,omitempty"`
	CompBiasIncl      []string    `json:"completion_bias_inclusion,omitempty"`
	CompBiasInclFirst bool        `json:"completion_bias_inclusion_first_token_only,omitempty"`
	CompBiasExcl      []string    `json:"completion_bias_exclusion,omitempty"`
	CompBiasExclFirst bool        `json:"completion_bias_exclusion_first_token_only,omitempty"`
	ContextualCtrlThr float64     `json:"contextual_control_threshold,omitempty"`
	CtrlLogAdditive   bool        `json:"control_log_additive,omitempty"`
}

type CompletionResponse struct {
	ModelVersion    string       `json:"model_version"`
	Completions     []Completion `json:"completions"`
	OptimizedPrompt any          `json:"optimized_prompt,omitempty"`
}

type Completion struct {
	/* 	LogProbs         *LogProbs `json:"log_probs,omitempty"` */
	Completion       string   `json:"completion"`
	RawCompletion    *string  `json:"raw_completion,omitempty"`
	CompletionTokens []string `json:"completion_tokens,omitempty"`
	FinishReason     string   `json:"finish_reason"`
}

/* type LogProbs struct {
	Tokens []map[string]float64 `json:"tokens"`
} */

/*
	 type OptimizedPromptItem struct {
		Data any

			 	Type string
				Data []byte
	}
*/
func (c *Client) Completion(r CompletionRequest) (CompletionResponse, error) {
	return req[CompletionRequest, CompletionResponse](c, http.MethodPost, "complete", r)
}
