package repository

import (
	"context"
	"fmt"
	"github.com/takanamito/gqlgen-todos/domain/entity"
	"github.com/takanamito/gqlgen-todos/ent"
	"github.com/takanamito/gqlgen-todos/ent/user"
	"strconv"
)

type UserRepository struct {
	client *ent.Client
}

func NewUserRepository(client *ent.Client) *UserRepository {
	return &UserRepository{
		client: client,
	}
}

func (r *UserRepository) Find(ctx context.Context, id string) (*ent.User, error) {
	userId, err := strconv.Atoi(id)
	if err != nil {
		return nil, fmt.Errorf("failed to convert userId: #{err}", err)
	}
	userModel, err := r.client.User.Query().Where(user.ID(userId)).Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed query user: #{err}", err)
	}
	return userModel, nil
}

func (r *UserRepository) Create(ctx context.Context, user *entity.User) (*entity.User, error) {
	userModel, err := r.client.User.Create().
		SetAge(user.Age).
		SetName(user.Name).
		SetGender(user.Gender).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating user: #{err}", err)
	}

	return &entity.User{
		ID: userModel.ID,
		Age: userModel.Age,
		Name: userModel.Name,
		Gender: userModel.Gender,
	}, nil
}
