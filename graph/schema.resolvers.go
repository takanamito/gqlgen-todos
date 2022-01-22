package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/takanamito/gqlgen-todos/domain/entity"
	"github.com/takanamito/gqlgen-todos/domain/repository"
	"github.com/takanamito/gqlgen-todos/ent"
	"github.com/takanamito/gqlgen-todos/graph/generated"
	"github.com/takanamito/gqlgen-todos/graph/model"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	client, err := ent.Open("postgres", "host=localhost port=5432 user=admin dbname=todos password=admin sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	repo := repository.NewTodoRepository(client)
	todo, err := repo.Create(ctx, input)
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
		ID:   strconv.Itoa(todo.ID),
		Text: todo.Body,
		User: &user,
	}, nil
}

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	client, err := ent.Open("postgres", "host=localhost port=5432 user=admin dbname=todos password=admin sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	repo := repository.NewUserRepository(client)
	userEntity := &entity.User{
		Age:    input.Age,
		Name:   input.Name,
		Gender: input.Gender,
	}
	userEntity, err = repo.Create(ctx, userEntity)
	if err != nil {
		return nil, fmt.Errorf("failed creating user: #{err}", err)
	}

	return &model.User{
		ID:     strconv.Itoa(userEntity.ID),
		Name:   userEntity.Name,
		Age:    userEntity.Age,
		Gender: userEntity.Gender,
	}, nil
}

func (r *queryResolver) User(ctx context.Context, id string) (*model.User, error) {
	client, err := ent.Open("postgres", "host=localhost port=5432 user=admin dbname=todos password=admin sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	repo := repository.NewUserRepository(client)
	userModel, err := repo.Find(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to query user: #{err}", err)
	}

	return &model.User{
		ID: strconv.Itoa(userModel.ID),
		Name: userModel.Name,
		Age: userModel.Age,
		Gender: userModel.Gender,
	}, nil
}

func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Todo(ctx context.Context, id string) (*model.Todo, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
