package requests

import (
	"context"

	"claude-code-api/pkg/db"
)

type RequestRepository struct {
	db *db.DB
}

func NewRequestRepository(db *db.DB) *RequestRepository {
	return &RequestRepository{
		db,
	}
}

func (repo *RequestRepository) Create(ctx context.Context, r *Request) error {
	return repo.db.WithContext(ctx).Create(r).Error
}
