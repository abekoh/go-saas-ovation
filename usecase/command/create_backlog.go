package command

import (
	"context"

	"github.com/abekoh/go-saas-ovation/domain/backlog_item"
)

func (u CommandsImpl) CreateBacklog(ctx context.Context, param *backlogitem.CreateBacklogItemParam) (*backlogitem.BacklogItem, error) {
	item, err := param.New()
	if err != nil {
		return nil, err
	}
	if err := u.BacklogItemRepo.Create(ctx, item); err != nil {
		return nil, err
	}
	return item, nil
}
