package query

import (
	"context"

	backlogitem "github.com/abekoh/go-saas-ovation/domain/backlog_item"
)

type Queries interface {
	GetBacklogItems(ctx context.Context, query *backlogitem.BacklogItemQuery) (backlogitem.BacklogItemNodeList, error)
}
