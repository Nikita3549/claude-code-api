# claude-code-api

HTTP API on top of the Claude Code CLI.

The problem it solves: the Claude API requires a separate API key with pay-per-token billing. If you already have a Claude Pro/Max subscription, this service lets you use Claude programmatically through your subscription limits instead — no API key needed. Each HTTP request spawns `claude -p` (headless mode) under the hood, authenticated with your subscription account.

Every request is also logged to Postgres (tokens, cost, duration, raw response) for usage tracking.

## Setup

### 1. Claude auth token

Generate a long-lived OAuth token (valid for 1 year) with your Pro/Max account:

```bash
npm install -g @anthropic-ai/claude-code
claude setup-token
```

Copy the token — it goes into `.env` as `CLAUDE_CODE_OAUTH_TOKEN`. The Claude CLI inside the container picks it up from the environment, no login or config files needed.

### 2. Environment

```bash
cp .env.example .env
```

Fill in the values, including `CLAUDE_CODE_OAUTH_TOKEN` from the previous step.

### 3. Run

```bash
docker compose --profile api up --build
```

## Usage

```bash
curl -X POST http://localhost:8082/claude/ask \
  -H "Content-Type: application/json" \
  -d '{"question": "What is a goroutine?", "prompt": "Answer in one sentence"}'
```

```json
{"answer": "A goroutine is a lightweight thread managed by the Go runtime."}
```

`question` is required, `prompt` is an optional system prompt.
