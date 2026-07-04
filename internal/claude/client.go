// Package claude provides a thin wrapper around the Claude Code CLI.
package claude

import (
	"context"
	"encoding/json"
	"fmt"
	"os/exec"
	"path/filepath"
	"strings"

	"claude-code-api/internal/configs"
)

type Client struct {
	conf *configs.Config
}

func NewClient(conf *configs.Config) *Client {
	return &Client{conf: conf}
}

func (client *Client) Run(ctx context.Context, input string, opts ...Option) (*RunResponse, error) {
	cfg := defaultOptions()

	notionToken := client.conf.NotionToken
	if notionToken != "" {
		opts = append(opts, WithAllowedTools("mcp__notion__*"))

		configPath := filepath.Join(".mcp-configs", "notion-mcp.json")
		opts = append(opts, WithMcpConfig(configPath))
	}

	for _, opt := range opts {
		opt(cfg)
	}

	args := []string{
		"-p",
		"--system-prompt", cfg.systemPrompt,
		"--output-format", cfg.outputFormat,
		"--permission-mode", cfg.permissionMode,
		"--model", cfg.model,
	}
	if cfg.allowedTools != "" {
		args = append(
			args, "--allowedTools", cfg.allowedTools,
		)
	}
	if cfg.mcpConfig != "" {
		args = append(
			args, "--mcp-config", cfg.mcpConfig,
		)
	}

	cmd := exec.CommandContext(ctx, "claude", args...)

	cmd.Stdin = strings.NewReader(input)

	out, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("claude run: %s: %w", string(out), err)
	}

	var r RunResponse
	if err := json.Unmarshal(out, &r); err != nil {
		return nil, err
	}
	r.Raw = json.RawMessage(out)
	r.Model = cfg.model

	return &r, nil
}
