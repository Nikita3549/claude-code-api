// Package requests provides requests domain logic
package requests

import (
	"encoding/json"
	"time"
)

type Request struct {
	ID             string          `gorm:"column:id;type:uuid;primaryKey;default:gen_random_uuid()"`
	RequestUUID    string          `gorm:"column:request_uuid;not null;uniqueIndex"`
	SessionID      string          `gorm:"column:session_id;not null;index"`
	Model          string          `gorm:"column:model;not null"`
	IsError        bool            `gorm:"column:is_error;not null;default:false"`
	APIErrorStatus *string         `gorm:"column:api_error_status"`
	InputTokens    int             `gorm:"column:input_tokens;not null;default:0"`
	OutputTokens   int             `gorm:"column:output_tokens;not null;default:0"`
	TotalCostUSD   float64         `gorm:"column:total_cost_usd;not null;default:0"`
	DurationMS     int             `gorm:"column:duration_ms;not null;default:0"`
	CreatedAt      time.Time       `gorm:"column:created_at;not null;autoCreateTime"`
	Raw            json.RawMessage `gorm:"column:raw;not null;type:jsonb"`
}

func (Request) TableName() string {
	return "claude_requests"
}
