package claude

import "encoding/json"

type CacheCreation struct {
	Ephemeral1hInputTokens int `json:"ephemeral_1h_input_tokens"`
	Ephemeral5mInputTokens int `json:"ephemeral_5m_input_tokens"`
}

type ServerToolUse struct {
	WebSearchRequests int `json:"web_search_requests"`
	WebFetchRequests  int `json:"web_fetch_requests"`
}

type Iteration struct {
	Type                     string        `json:"type"`
	InputTokens              int           `json:"input_tokens"`
	OutputTokens             int           `json:"output_tokens"`
	CacheReadInputTokens     int           `json:"cache_read_input_tokens"`
	CacheCreationInputTokens int           `json:"cache_creation_input_tokens"`
	CacheCreation            CacheCreation `json:"cache_creation"`
}

type Usage struct {
	InputTokens              int           `json:"input_tokens"`
	CacheCreationInputTokens int           `json:"cache_creation_input_tokens"`
	CacheReadInputTokens     int           `json:"cache_read_input_tokens"`
	OutputTokens             int           `json:"output_tokens"`
	ServerToolUse            ServerToolUse `json:"server_tool_use"`
	ServiceTier              string        `json:"service_tier"`
	CacheCreation            CacheCreation `json:"cache_creation"`
	InferenceGeo             string        `json:"inference_geo"`
	Iterations               []Iteration   `json:"iterations"`
	Speed                    string        `json:"speed"`
}

type ModelUsage struct {
	InputTokens              int     `json:"inputTokens"`
	OutputTokens             int     `json:"outputTokens"`
	CacheReadInputTokens     int     `json:"cacheReadInputTokens"`
	CacheCreationInputTokens int     `json:"cacheCreationInputTokens"`
	WebSearchRequests        int     `json:"webSearchRequests"`
	CostUSD                  float64 `json:"costUSD"`
	ContextWindow            int     `json:"contextWindow"`
	MaxOutputTokens          int     `json:"maxOutputTokens"`
}

type RunResponse struct {
	Type              string                `json:"type"`
	Subtype           string                `json:"subtype"`
	IsError           bool                  `json:"is_error"`
	APIErrorStatus    *string               `json:"api_error_status"`
	DurationMs        int                   `json:"duration_ms"`
	DurationAPIMs     int                   `json:"duration_api_ms"`
	TTFTMs            int                   `json:"ttft_ms"`
	TTFTStreamMs      int                   `json:"ttft_stream_ms"`
	TimeToRequestMs   int                   `json:"time_to_request_ms"`
	NumTurns          int                   `json:"num_turns"`
	Result            string                `json:"result"`
	StopReason        string                `json:"stop_reason"`
	SessionID         string                `json:"session_id"`
	TotalCostUSD      float64               `json:"total_cost_usd"`
	Usage             Usage                 `json:"usage"`
	ModelUsage        map[string]ModelUsage `json:"modelUsage"`
	PermissionDenials []string              `json:"permission_denials"`
	TerminalReason    string                `json:"terminal_reason"`
	FastModeState     string                `json:"fast_mode_state"`
	UUID              string                `json:"uuid"`
	Raw               json.RawMessage       `json:"raw"`
	Model             string                `json:"model"`
}
