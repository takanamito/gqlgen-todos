package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/takanamito/gqlgen-todos/ent"
	"github.com/takanamito/gqlgen-todos/ent/user"
	"github.com/takanamito/gqlgen-todos/graph/generated"
	"github.com/takanamito/gqlgen-todos/graph/model"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	client, err := ent.Open("postgres", "host=localhost port=5432 user=admin dbname=todos password=admin sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	userId, err := strconv.Atoi(input.UserID)
	if err != nil {
		return nil, fmt.Errorf("invalid userId: #{err}", err)
	}
	dbUser, err := client.User.Query().Where(user.ID(userId)).Only(ctx)

	todo, err := client.Todo.Create().
		SetUser(dbUser).
		SetBody(input.Text).
		SetCreatedAt(time.Now()).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating todo: #{err}", err)
	}

	todoOwner, err := todo.QueryUser().Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed query user: #{err}", err)
	}

	user := model.User{
		ID:     strconv.Itoa(todoOwner.ID),
		Name:   todoOwner.Name,
		Age:    todoOwner.Age,
		Gender: todoOwner.Gender,
	}

	return &model.Todo{
		ID:     strconv.Itoa(todo.ID),
		Text:   todo.Body,
		User: &user,
	}, nil
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
		return nil, fmt.Errorf("failed creating user: #{err}", err)
	}

	log.Println("return user: ", user)
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

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
