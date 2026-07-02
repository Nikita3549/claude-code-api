package claude

import (
	"log"
	"net/http"

	"claude-code-api/pkg/req"
	"claude-code-api/pkg/res"
)

type ClaudeHandler struct {
	claudeService *ClaudeService
}

func NewClaudeHandler(router *http.ServeMux, claudeService *ClaudeService) {
	handler := &ClaudeHandler{
		claudeService: claudeService,
	}

	router.HandleFunc("POST /claude/ask", handler.Ask())
}

func (h *ClaudeHandler) Ask() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[AskClaudeRequest](w, r)
		if err != nil {
			return
		}

		answer, err := h.claudeService.Ask(r.Context(), *body)
		if err != nil {
			log.Println("unexpected error: " + err.Error())
			res.JSON(w, http.StatusInternalServerError, "Internal Server Error")
			return
		}

		response := AskClaudeResponse{
			Answer: answer,
		}

		res.JSON(w, http.StatusOK, response)
	}
}
