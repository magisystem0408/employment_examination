package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"graphql-server/db"
	"graphql-server/graph/generated"
	"graphql-server/graph/model"

	"github.com/vektah/gqlparser/v2/gqlerror"
	"golang.org/x/crypto/bcrypt"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	userName, err := r.validateJWT(ctx)
	if err != nil {
		return nil, err
	}

	if userName != input.UserName {
		return nil, gqlerror.Errorf("ummatch userName(%s)", input.UserName)
	}

	newTodo := &db.ToDo{
		UserName: input.UserName,
		Text:     input.Text,
	}
	if err := newTodo.Save(r.store); err != nil {
		return nil, err
	}

	return &model.Todo{
		TodoID: newTodo.TodoID,
		Text:   newTodo.Text,
		Done:   newTodo.Done,
	}, nil
}

func (r *mutationResolver) DoneTodo(ctx context.Context, input model.DoneTodo) (*model.Todo, error) {
	userName, err := r.validateJWT(ctx)
	if err != nil {
		return nil, err
	}

	user := db.User{Name: userName}
	if err := user.Get(r.store); err != nil {
		return nil, err
	}

	var todo *model.Todo
	for i, td := range user.ToDos {
		if td.TodoID == input.TodoID {
			todo = &model.Todo{
				TodoID: td.TodoID,
				Text:   td.Text,
				Done:   td.Done,
				User:   &model.User{Name: userName},
			}

			user.ToDos[i].Done = true
		}
	}

	if todo == nil {
		return nil, gqlerror.Errorf("todoID(%s) not found", input.TodoID)
	}

	if err := user.Save(r.store); err != nil {
		return nil, err
	}

	return todo, nil
}

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.AuthenticatePayload, error) {
	user := &db.User{Name: input.Name}
	if err := user.Get(r.store); !errors.Is(err, db.ErrKeyNotFound) {
		if err == nil {
			return nil, gqlerror.Errorf("name(%s) already taken", input.Name)
		}
		return nil, err
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(input.Password), 12)
	if err != nil {
		return nil, err
	}

	newUser := &db.User{Name: input.Name, HashedPassword: string(hash)}
	if err := newUser.Save(r.store); err != nil {
		return nil, err
	}

	jwt, err := r.createJWT(newUser.Name)
	if err != nil {
		return nil, err
	}

	return &model.AuthenticatePayload{
		UserName: newUser.Name,
		Jwt:      jwt,
	}, nil
}

func (r *mutationResolver) Authenticate(ctx context.Context, input model.AuthenticateInput) (*model.AuthenticatePayload, error) {
	user := &db.User{Name: input.Name}
	if err := user.Get(r.store); err != nil {
		if errors.Is(err, db.ErrKeyNotFound) {
			return nil, gqlerror.Errorf("user(%s) not found", input.Name)
		}
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), []byte(input.Password)); err != nil {
		return nil, err
	}

	jwt, err := r.createJWT(user.Name)
	if err != nil {
		return nil, err
	}

	return &model.AuthenticatePayload{
		UserName: user.Name,
		Jwt:      jwt,
	}, nil
}

func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	userName, err := r.validateJWT(ctx)
	if err != nil {
		return nil, err
	}

	user := db.User{Name: userName}
	if err := user.Get(r.store); err != nil {
		return nil, err
	}

	todos := make([]*model.Todo, len(user.ToDos))
	for i, todo := range user.ToDos {
		todos[i] = &model.Todo{
			TodoID: todo.TodoID,
			Text:   todo.Text,
			Done:   todo.Done,
			User:   &model.User{Name: userName},
		}
	}

	return todos, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
