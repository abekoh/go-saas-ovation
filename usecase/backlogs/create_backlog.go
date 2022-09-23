package backlogs

import (
	"context"

	"github.com/abekoh/go-saas-ovation/domain/backlog_item"
)

func (u UsecaseImpl) CreateBacklog(ctx context.Context, param backlogitem.CreateBacklogItemParam) error {
	item, err := param.New()
	if err != nil {
		return err
	}
	if err := u.BacklogItemRepo.Create(ctx, item); err != nil {
		return err
	}
	return nil
}
