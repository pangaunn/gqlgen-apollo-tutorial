package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	graphql1 "github.com/pangaunn/gqlgen-apollo-tutorial/app/graphql"
	"github.com/pangaunn/gqlgen-apollo-tutorial/app/graphql/model"
)

func (r *mutationResolver) BookTrips(ctx context.Context, launchIds []*string) (*model.TripUpdateResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CancelTrip(ctx context.Context, launchID string) (*model.TripUpdateResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) Login(ctx context.Context, email *string) (*string, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns graphql1.MutationResolver implementation.
func (r *Resolver) Mutation() graphql1.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
