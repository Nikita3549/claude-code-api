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

### 3. Notion MCP (optional)

The service can give Claude access to your Notion workspace via the [Notion MCP server](https://github.com/makenotion/notion-mcp-server). It is enabled automatically when `NOTION_TOKEN` is set; if the variable is empty, the service runs without it.

To get a token:

1. Open [notion.so/my-integrations](https://www.notion.so/my-integrations) (you need permission to create integrations in the workspace).
2. Click **New integration**, give it a name (e.g. `claude-code-api`) and pick the workspace.
3. In the integration settings set capabilities: **Read content** at minimum; add **Update/Insert content** only if the service should write to Notion.
4. Copy the **Internal Integration Secret** (starts with `ntn_`) — this is the token.
5. **Required:** share the pages/databases with the integration — open a page in Notion → **•••** menu → **Connections** → select your integration. Child pages inherit access. Without this step the API returns `object not found` even with a valid token.
6. Put the token into `.env` as `NOTION_TOKEN`. It is a secret — don't commit it, keep it in `.env` or a secret store.

### 4. Run

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
