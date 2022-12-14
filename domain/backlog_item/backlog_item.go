package backlogitem

import (
	"context"

	subtask "github.com/abekoh/go-saas-ovation/domain/sub_task"
	"github.com/google/uuid"
)

type ID uuid.UUID

func (i ID) String() string {
	return uuid.UUID(i).String()
}

type Type string

const (
	Story Type = "STORY"
	Bug        = "BUG"
)

type StoryPoint int

type CreateBacklogItemParam struct {
	Type       Type
	Summary    string
	StoryPoint *int
}

func (p CreateBacklogItemParam) New() (*BacklogItem, error) {
	item := &BacklogItem{
		ID:         ID(uuid.New()),
		Type:       p.Type,
		Summary:    p.Summary,
		StoryPoint: p.StoryPoint,
	}
	if err := item.Validate(); err != nil {
		return nil, err
	}
	return item, nil
}

type BacklogItem struct {
	ID         ID     `json:"id"`
	Type       Type   `json:"type"`
	Summary    string `json:"summary"`
	StoryPoint *int   `json:"storyPoint"`
}

func (bi BacklogItem) Validate() error {
	return nil
}

type BacklogItemList []BacklogItem

type Repository interface {
	Create(ctx context.Context, item *BacklogItem) error
	GetOne(ctx context.Context, id ID) (*BacklogItem, error)
}

type BacklogItemQuery struct {
}

type BacklogItemNode struct {
	BacklogItem
	subtask.SubTaskNodeList
}

type BacklogItemNodeList []BacklogItemNode
