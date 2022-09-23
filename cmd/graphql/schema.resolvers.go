package main

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/abekoh/go-saas-ovation/domain/backlog_item"
	"github.com/abekoh/go-saas-ovation/graphql/generated"
	"github.com/abekoh/go-saas-ovation/graphql/model"
)

// CreateBacklog is the resolver for the createBacklog field.
func (r *mutationResolver) CreateBacklog(ctx context.Context, input model.NewBacklogItem) (*model.BacklogItem, error) {
	param := backlogitem.CreateBacklogItemParam{
		Type:       input.Type,
		Summary:    input.Summary,
		StoryPoint: input.StoryPoint,
	}
	item, err := r.backlogsUsecase.CreateBacklog(ctx, &param)
	if err != nil {
		return nil, err
	}
	return &model.BacklogItem{
		ID:         item.ID.String(),
		Type:       item.Type,
		Summary:    item.Summary,
		StoryPoint: item.StoryPoint,
	}, nil
}

// BacklogItems is the resolver for the backlogItems field.
func (r *queryResolver) BacklogItems(ctx context.Context) ([]*model.BacklogItem, error) {
	panic(fmt.Errorf("not implemented: BacklogItems - backlogItems"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
