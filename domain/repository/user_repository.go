package repository

import (
	"context"
	"fmt"
	"github.com/takanamito/gqlgen-todos/domain/entity"
	"github.com/takanamito/gqlgen-todos/ent"
)

type UserRepository struct {
	client *ent.Client
}

func NewUserRepository(client *ent.Client) *UserRepository {
	return &UserRepository{
		client: client,
	}
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
