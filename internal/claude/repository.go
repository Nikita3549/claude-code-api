// Package claude provides a thin wrapper around the Claude Code CLI.
package claude

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"os/exec"
	"strconv"
	"strings"
)

type ClaudeRepository struct{}

func NewClaudeRepository() *ClaudeRepository {
	return &ClaudeRepository{}
}

func (repo *ClaudeRepository) Run(ctx context.Context, input string, opts ...Option) (*RunResponse, error) {
	cfg := defaultOptions()
	for _, opt := range opts {
		opt(cfg)
	}

	args := []string{
		"-p",
		"--system-prompt", cfg.systemPrompt,
		"--output-format", cfg.outputFormat,
		"--max-turns", strconv.Itoa(cfg.maxTurns),
		"--tools", cfg.tools,
		"--permission-mode", cfg.permissionMode,
		"--model", cfg.model,
	}

	cmd := exec.Command("claude", args...)

	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	cmd.Stdin = strings.NewReader(input)

	out, err := cmd.Output()
	if err != nil {
		return nil, errors.New(string(out))
	}

	var r RunResponse
	if err := json.Unmarshal(out, &r); err != nil {
		return nil, err
	}

	return &r, nil
}
