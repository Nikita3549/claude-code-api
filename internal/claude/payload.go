package claude

type AskClaudeRequest struct {
	Question string `json:"question" validate:"required"`
	Prompt   string `json:"prompt"`
}

type AskClaudeResponse struct {
	Answer string `json:"answer"`
}
