package backlogs

import (
	"context"

	"github.com/abekoh/go-saas-ovation/domain/backlog_item"
)

type Usecase interface {
	CreateBacklog(ctx context.Context, param *backlogitem.CreateBacklogItemParam) error
}

type UsecaseImpl struct {
	BacklogItemRepo backlogitem.Repository
}
