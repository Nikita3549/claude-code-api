// Package claude provides a thin wrapper around the Claude Code CLI.
package claude

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"os/exec"
	"strconv"
	"strings"
)

type Client struct{}

func NewClient() *Client {
	return &Client{}
}

func (repo *Client) Run(ctx context.Context, input string, opts ...Option) (*RunResponse, error) {
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

	cmd := exec.CommandContext(ctx, "claude", args...)

	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	cmd.Stdin = strings.NewReader(input)

	out, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("claude run: %s: %w", stderr.String(), err)
	}

	var r RunResponse
	if err := json.Unmarshal(out, &r); err != nil {
		return nil, err
	}
	r.Raw = json.RawMessage(out)
	r.Model = cfg.model

	return &r, nil
}
