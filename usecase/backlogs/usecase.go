package backlogs

import (
	"context"

	backlogitem "github.com/abekoh/go-saas-ovation/domain/backlog_item"
)

type Usecase interface {
	CreateBacklog(ctx context.Context, item *backlogitem.BacklogItem) error
}

type UsecaseImpl struct {
	BacklogItemRepo backlogitem.Repository
}
