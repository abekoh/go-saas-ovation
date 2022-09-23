package postgres

import (
	"context"
	"database/sql"

	"github.com/abekoh/go-saas-ovation/domain/backlog_item"
	"github.com/abekoh/go-saas-ovation/sqlc"
	"github.com/google/uuid"
)

type BacklogItemRepository struct {
	db sql.DB
}

func (b *BacklogItemRepository) GetOne(ctx context.Context, id backlogitem.ID) (*backlogitem.BacklogItem, error) {
	queries := sqlc.New(&b.db)
	res, err := queries.GetOneBacklogItem(ctx, uuid.UUID(id))
	if err != nil {
		return nil, err
	}
	return convert(res), nil
}

func convert(from sqlc.BacklogItem) *backlogitem.BacklogItem {
	return &backlogitem.BacklogItem{
		ID:      backlogitem.ID(from.ID),
		Type:    from.Type,
		Summary: from.Summary,
		StoryPoint: func() *int {
			if from.StoryPoint.Valid {
				r := int(from.StoryPoint.Int32)
				return &r
			} else {
				return nil
			}
		}(),
	}
}
