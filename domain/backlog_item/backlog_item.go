package backlogitem

import (
	"context"

	"github.com/google/uuid"
)

type ID uuid.UUID

type Type string

const (
	Story Type = "STORY"
	Bug        = "BUG"
)

type StoryPoint int

type BacklogItem struct {
	ID         ID     `json:"id"`
	Type       Type   `json:"type"`
	Summary    string `json:"summary"`
	StoryPoint *int   `json:"storyPoint"`
}

type Repository interface {
	GetOne(ctx context.Context, id ID) (*BacklogItem, error)
}
