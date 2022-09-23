package query

import (
	"context"

	backlogitem "github.com/abekoh/go-saas-ovation/domain/backlog_item"
)

type Queries interface {
	GetBacklogs(ctx context.Context, param *backlogitem.CreateBacklogItemParam) (*backlogitem.BacklogItem, error)
}

type QueriesImpl struct {
}
