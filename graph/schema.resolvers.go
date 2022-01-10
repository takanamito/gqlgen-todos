package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"log"

	"github.com/takanamito/gqlgen-todos/ent"
	"github.com/takanamito/gqlgen-todos/graph/generated"
	"github.com/takanamito/gqlgen-todos/graph/model"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	client, err := ent.Open("postgres", "host=localhost port=5432 user=admin dbname=todos password=admin sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	user, err := client.User.Create().
		SetAge(input.Age).
		SetName(input.Name).
		SetGender(input.Gender).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating user: #{err}")
	}

	return &model.User{
		ID:     string(user.ID),
		Name:   user.Name,
		Age:    user.Age,
		Gender: user.Gender,
	}, nil
}

func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *todoResolver) User(ctx context.Context, obj *model.Todo) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Todo returns generated.TodoResolver implementation.
func (r *Resolver) Todo() generated.TodoResolver { return &todoResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type todoResolver struct{ *Resolver }
