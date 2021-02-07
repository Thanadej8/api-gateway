package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/Thanadej8/api-gateway/graph/generated"
	"github.com/Thanadej8/api-gateway/graph/model"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (model.Todo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Todos(ctx context.Context) ([]model.Todo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	var user = new(model.User)
	user.ID = "1"
	user.Name = "Thanadej Phadtong"

	usersV2 := []*model.User{user, user, user}
	// panic(fmt.Errorf("not implemented"))

	return usersV2, nil
}

func (r *queryResolver) TestExtends(ctx context.Context) (*model.Test, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *userResolver) Name(ctx context.Context, obj *model.User) (string, error) {
	return obj.Name, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type userResolver struct{ *Resolver }
