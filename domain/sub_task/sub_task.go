package subtask

import (
	"context"

	"github.com/abekoh/go-saas-ovation/domain/backlog_item"
	"github.com/google/uuid"
)

type ID uuid.UUID

type SubTask struct {
	ID            ID             `json:"id"`
	Summary       string         `json:"summary"`
	StoryPoint    *int           `json:"storyPoint"`
	BackLogItemID backlogitem.ID `json:"backLogItemId"`
}

type SubTaskList []SubTask

type Repository interface {
	GetOne(ctx context.Context, id ID) (*SubTask, error)
}

type SubTaskNode struct {
	SubTask
}

type SubTaskNodeList []SubTaskNode
