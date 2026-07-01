package claude

import (
	"context"
	"log"
	"time"

	"claude-code-api/internal/requests"
)

type ClaudeService struct {
	requestRepository *requests.RequestRepository
	client            *Client
}

type ClaudeServiceDeps struct {
	RequestRepository *requests.RequestRepository
	ClaudeClient      *Client
}

func NewClaudeService(deps *ClaudeServiceDeps) *ClaudeService {
	return &ClaudeService{
		requestRepository: deps.RequestRepository,
		client:            deps.ClaudeClient,
	}
}

func (s *ClaudeService) Ask(ctx context.Context, body AskClaudeRequest) (string, error) {
	var args []Option

	if body.Prompt != "" {
		args = append(args, WithSystemPrompt(body.Prompt))
	}

	result, err := s.client.Run(ctx, body.Question, args...)
	if err != nil {
		return "", err
	}

	request := &requests.Request{
		RequestUUID:    result.UUID,
		SessionID:      result.SessionID,
		Model:          result.Model,
		IsError:        result.IsError,
		APIErrorStatus: result.APIErrorStatus,
		InputTokens:    result.Usage.InputTokens,
		OutputTokens:   result.Usage.OutputTokens,
		TotalCostUSD:   result.TotalCostUSD,
		DurationMS:     result.DurationMs,
		CreatedAt:      time.Time{},
		Raw:            result.Raw,
	}
	err = s.requestRepository.Create(ctx, request)
	if err != nil {
		log.Println(err)
	}

	return result.Result, nil
}
