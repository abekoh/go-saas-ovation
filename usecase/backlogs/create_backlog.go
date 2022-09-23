package backlogs

import (
	"context"

	backlogitem "github.com/abekoh/go-saas-ovation/domain/backlog_item"
)

func (u UsecaseImpl) CreateBacklog(ctx context.Context, item *backlogitem.BacklogItem) error {
	if err := item.Validate(); err != nil {
		return err
	}
	if err := u.BacklogItemRepo.Create(ctx, item); err != nil {
		return err
	}
	return nil
}
