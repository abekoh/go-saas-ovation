package postgres

import (
	"context"

	"github.com/abekoh/go-saas-ovation/domain/backlog_item"
)

type BacklogItemRepository struct {
}

func (b BacklogItemRepository) GetOne(ctx context.Context, id backlogitem.ID) (*backlogitem.BacklogItem, error) {
	//TODO implement me
	panic("implement me")
}
