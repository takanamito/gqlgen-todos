package repository

import (
	"context"
	"fmt"
	"github.com/takanamito/gqlgen-todos/ent"
	"github.com/takanamito/gqlgen-todos/graph/model"
	"strconv"
	"time"
)

type TodoRepository struct {
	client *ent.Client
}

func NewTodoRepository(client *ent.Client) *TodoRepository {
	return &TodoRepository{
		client: client,
	}
}

func (r *TodoRepository) Create(ctx context.Context, input model.NewTodo) (*ent.Todo, error) {
	userId, err := strconv.Atoi(input.UserID)
	if err != nil {
		return nil, fmt.Errorf("failed converting UserID string to int")
	}

	todoModel, err := r.client.Todo.Create().
		SetUserID(userId).
		SetBody(input.Text).
		SetCreatedAt(time.Now()).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating todo: #{err}", err)
	}

	return todoModel, nil
}