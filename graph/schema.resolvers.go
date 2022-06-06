package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"

	"github.com/99designs/gqlgen/graphql"
	"github.com/graph/gqlgen-todos/graph/generated"
	"github.com/graph/gqlgen-todos/graph/model"
	context2 "github.com/graph/gqlgen-todos/router/context"
)

func (r *mutationResolver) SingleUpload(ctx context.Context, file graphql.Upload) (bool, error) {
	f, _ := os.OpenFile("text.txt", os.O_RDWR | os.O_CREATE, 0755)
	b, _ := ioutil.ReadAll(file.File)
	f.Write(b)
	return true, nil
}

func (r *mutationResolver) CreatePeople(ctx context.Context, input model.NewPeople) (*model.People, error) {
	todo := &model.People{
		Text: input.Text,
		ID:   fmt.Sprintf("T%d", rand.Int()),
		User: &model.User{ID: input.UserID, Name: "user " + input.UserID},
	}
	r.peoples = append(r.peoples, todo)
	return todo, nil
}

func (r *queryResolver) Peoples(ctx context.Context) ([]*model.People, error) {
	// 通过gin context 拿到参数
	gc, err := context2.GinContextFromContext(ctx)
	if err != nil {
		return nil, err
	}
	id := gc.Query("id")
	fmt.Println(id)

	return r.peoples, err
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
