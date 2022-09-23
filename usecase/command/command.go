package command

import (
	"context"

	"github.com/abekoh/go-saas-ovation/domain/backlog_item"
)

type Commands interface {
	CreateBacklog(ctx context.Context, param *backlogitem.CreateBacklogItemParam) (*backlogitem.BacklogItem, error)
}

type CommandsImpl struct {
	BacklogItemRepo backlogitem.Repository
}
