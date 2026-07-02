# claude-code-api

HTTP API on top of the Claude Code CLI.

The problem it solves: the Claude API requires a separate API key with pay-per-token billing. If you already have a Claude Pro/Max subscription, this service lets you use Claude programmatically through your subscription limits instead — no API key needed. Each HTTP request spawns `claude -p` (headless mode) under the hood, authenticated with your subscription account.

Every request is also logged to Postgres (tokens, cost, duration, raw response) for usage tracking.

## Setup

### 1. Claude CLI credentials

The Docker image bakes your local Claude CLI auth into the container:

```bash
npm install -g @anthropic-ai/claude-code
claude   # log in with your Pro/Max account, then exit

mkdir claude-config
cp -r ~/.claude claude-config/.claude
cp ~/.claude.json claude-config/.claude.json
```

### 2. Environment

```bash
cp .env.example .env
```

Fill in the values 

### 3. Run

```bash
docker compose --profile api up --build
```

## Usage

```bash
curl -X POST http://localhost:8080/claude/ask \
  -H "Content-Type: application/json" \
  -d '{"question": "What is a goroutine?", "prompt": "Answer in one sentence"}'
```

```json
{"answer": "A goroutine is a lightweight thread managed by the Go runtime."}
```

`question` is required, `prompt` is an optional system prompt.
