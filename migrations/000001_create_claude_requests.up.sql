CREATE TABLE claude_requests (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    request_uuid uuid NOT NULL UNIQUE,
    session_id uuid NOT NULL,
    model text NOT NULL,
    is_error boolean NOT NULL DEFAULT false,
    api_error_status text,
    input_tokens int NOT NULL DEFAULT 0,
    output_tokens int NOT NULL DEFAULT 0,
    total_cost_usd numeric(12, 6) NOT NULL DEFAULT 0,
    duration_ms int NOT NULL DEFAULT 0,
    created_at timestamptz NOT NULL DEFAULT now(),
    raw jsonb NOT NULL
);

CREATE INDEX idx_claude_requests_request_uuid ON claude_requests (request_uuid);
CREATE INDEX idx_claude_requests_session_id ON claude_requests (session_id);
CREATE INDEX idx_claude_requests_is_error ON claude_requests (is_error);
CREATE INDEX idx_claude_requests_created_at ON claude_requests (created_at);
