package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.46

import (
	"context"
	"fmt"
	"time"

	"github.com/ausimity/review_gqlgen/graph/model"
)

var reviews = []*model.Review{
	{
		ID:      1,
		Body:    "Apollo may only be the 3rd US human spaceflight program, but my ride was 1st class!",
		Product: &model.Product{ID: "p1"},
	},
	{
		ID:      2,
		Body:    "What a ride - we went up AND down! A++!",
		Product: &model.Product{ID: "p2"},
	},
	{
		ID:      3,
		Body:    "Nobody beats Apollo when it comes to command and service modules - my time in lunar orbit is going to be hard to beat. Highly recommended!",
		Product: &model.Product{ID: "p3"},
	},
	{
		ID:      4,
		Body:    "Thank you Apollo for a wonderful time, and even more wonderful photos - I got my best far side of the moon shots yet!",
		Product: &model.Product{ID: "p4"},
	},
}

// Review is the resolver for the review field.
func (r *queryResolver) Review(ctx context.Context, id int) (*model.Review, error) {
	for _, review := range reviews {
		if review.ID == id {
			return review, nil
		}
	}
	return nil, fmt.Errorf("review not found")
}

// ReviewAdded is the resolver for the reviewAdded field.
func (r *subscriptionResolver) ReviewAdded(ctx context.Context) (<-chan *model.Review, error) {
	ch := make(chan *model.Review)

	go func() {
		defer close(ch)
		for _, review := range reviews {
			ch <- review
			time.Sleep(3 * time.Second)
		}
	}()
	return ch, nil
}

// Add the missing type resolver for __resolveReference
func (r *queryResolver) __resolveReference(ctx context.Context, reference *model.Review) (*model.Review, error) {
	for _, review := range reviews {
		if review.ID == reference.ID {
			return review, nil
		}
	}
	return nil, fmt.Errorf("review not found")
}

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

// Subscription returns SubscriptionResolver implementation.
func (r *Resolver) Subscription() SubscriptionResolver { return &subscriptionResolver{r} }

type queryResolver struct{ *Resolver }
type subscriptionResolver struct{ *Resolver }