package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/takanamito/gqlgen-todos/ent"
	"github.com/takanamito/gqlgen-todos/ent/user"

	_ "github.com/lib/pq"
)

func main() {
	client, err := ent.Open("postgres","host=localhost port=5432 user=admin dbname=todos password=admin sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	CreateTodos(context.Background(), client)
}

func CreateUser(ctx context.Context, client *ent.Client) (*ent.User, error) {
	u, err := client.User.
		Create().
		SetAge(30).
		SetName("a8m").
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating user: %w", err)
	}
	log.Println("user was created: ", u)
	return u, nil
}

func QueryUser(ctx context.Context, client *ent.Client) (*ent.User, error) {
	u, err := client.User.
		Query().
		Where(user.Name("a8m")).
		// ユーザーが見つからない場合、`Only`は失敗する
		// あるいは、1人以上のユーザーが返却される
		Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed querying user: %w", err)
	}
	log.Println("user returned: ", u)
	return u, nil
}

func CreateTodos(ctx context.Context, client *ent.Client) (*ent.User, error) {
	todo, err := client.Todo.
		Create().
		SetBody("掃除をする").
		SetCreatedAt(time.Now()).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating todo: %w", err)
	}
	log.Println("todo was created: ", todo)

	todo2, err := client.Todo.
		Create().
		SetBody("買い物に行く").
		SetCreatedAt(time.Now()).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating todo: %w", err)
	}
	log.Println("todo was created: ", todo2)

	takanamito, err := client.User.
		Create().
		SetAge(30).
		SetName("takanamito").
		AddTodos(todo, todo2).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating user: %w", err)
	}
	log.Println("user was created: ", takanamito)
	return takanamito, nil
}
