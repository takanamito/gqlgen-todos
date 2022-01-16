package entity

import (
	"github.com/takanamito/gqlgen-todos/ent"
)

func BodyLength(model *ent.Todo) int {
	return len(model.Body)
}
