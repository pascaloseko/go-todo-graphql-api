package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/pascaloseko/go-todo-graphql-api/graph/generated"
	"github.com/pascaloseko/go-todo-graphql-api/graph/model"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.TodoInput) (*model.Todo, error) {
	todo := model.Todo{
		Text: input.Text,
	}
	err := r.DB.Create(&todo).Error
	if err != nil {
		return nil, err
	}
	return &todo, nil
}

func (r *mutationResolver) UpdateTodo(ctx context.Context, todoID int, input model.TodoInput) (*model.Todo, error) {
	updatedTodo := model.Todo{
		ID:        todoID,
		Text:      input.Text,
		Completed: *input.Completed,
	}
	r.DB.Save(&updatedTodo)
	return &updatedTodo, nil
}

func (r *mutationResolver) DeleteTodo(ctx context.Context, todoID int) (bool, error) {
	err := r.DB.Where("id = ?", todoID).Delete(&model.Todo{}).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	var todos []*model.Todo
	r.DB.Find(&todos)
	return todos, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
