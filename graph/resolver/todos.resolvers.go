package resolver

import (
	"context"
	"crypto/rand"
	"math/big"

	"github.com/ericgrandt/gqlgen-example/graph/generated"
	"github.com/ericgrandt/gqlgen-example/graph/model"
	"github.com/ericgrandt/gqlgen-example/middleware"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (model.Todo, error) {
	randNumber, _ := rand.Int(rand.Reader, big.NewInt(10000))
	todo := model.Todo{
		Text:   input.Text,
		ID:     randNumber.String(),
		UserID: input.UserID,
	}

	stmt, err := r.db.Prepare("INSERT INTO todo(id, value, user_id) VALUES (?, ?, ?)")
	if err != nil {
		panic(err)
	}

	_, err = stmt.Exec(todo.ID, todo.Text, todo.UserID)
	if err != nil {
		panic(err)
	}

	return todo, nil
}

func (r *queryResolver) Todos(ctx context.Context, pageSize int32, pageNum int32) ([]model.Todo, error) {
	stmt, _ := r.db.Prepare("SELECT * FROM todo WHERE user_id = ? LIMIT ? OFFSET ?")
	defer stmt.Close()

	user := middleware.GetUserContext(ctx)
	rows, _ := stmt.Query(user.ID, pageSize, pageSize*(pageNum-1))
	var todos []model.Todo
	for rows.Next() {
		var todo model.Todo
		rows.Scan(&todo.ID, &todo.UserID, &todo.Text)
		todos = append(todos, todo)
	}
	return todos, nil
}

func (r *todoResolver) User(ctx context.Context, obj *model.Todo) (model.User, error) {
	return r.userData.GetUser(obj.UserID)
}

func (r *Resolver) Todo() generated.TodoResolver { return &todoResolver{r} }

type todoResolver struct{ *Resolver }
