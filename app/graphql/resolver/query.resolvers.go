package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	graphql1 "github.com/pangaunn/gqlgen-apollo-tutorial/app/graphql"
	"github.com/pangaunn/gqlgen-apollo-tutorial/app/graphql/model"
)

func (r *queryResolver) Launches(ctx context.Context, pageSize *int, after *string) (*model.LaunchConnection, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Launch(ctx context.Context, id string) (*model.Launch, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Me(ctx context.Context) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

// Query returns graphql1.QueryResolver implementation.
func (r *Resolver) Query() graphql1.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
