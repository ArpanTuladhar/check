package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.41

import (
	"context"
	"errors"
	"fmt"

	"strconv"

	graph "github.com/88labs/andpad-engineer-training/2023/Gopi/backend/internal/handler/graph/generated"
	"github.com/88labs/andpad-engineer-training/2023/Gopi/backend/internal/handler/graph/model"
	usecase_input "github.com/88labs/andpad-engineer-training/2023/Gopi/backend/internal/usecase/input"
)

// CreateTodo is the resolver for the createTodo field.
func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	todo, err := r.todoCreator.CreateTodo(ctx, &usecase_input.TodoCreator{Text: input.Text})
	if err != nil {
		return nil, errors.New("error")
	}

	return &model.Todo{ID: todo.ID.String(), Text: todo.Text, User: &model.User{ID: strconv.Itoa((int)(todo.UserID))}}, nil
}

// Todos is the resolver for the todos field.
func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	panic(fmt.Errorf("not implemented: Todos - todos"))
}

// Mutation returns graph.MutationResolver implementation.
func (r *Resolver) Mutation() graph.MutationResolver { return &mutationResolver{r} }

// Query returns graph.QueryResolver implementation.
func (r *Resolver) Query() graph.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
