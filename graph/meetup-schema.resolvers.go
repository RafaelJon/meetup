package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"

	"github.com/RafaelJon/meetup/graph/generated"
	"github.com/RafaelJon/meetup/graph/models"
)

var meetups = []*models.Meetup{
	{
		ID:          "1",
		Name:        "First Meetup",
		Description: "First Description",
		UserID:      "1",
	},
	{
		ID:          "2",
		Name:        "Second Meetup",
		Description: "Second Description",
		UserID:      "2",
	},
}
var users = []*models.User{
	{
		ID:       "1",
		Username: "Babe",
		Email:    "babe@gmail.com",
	},
	{
		ID:       "2",
		Username: "Bibidi",
		Email:    "bibidi@gmail.com",
	},
}

func (r *meetupResolver) User(ctx context.Context, obj *models.Meetup) (*models.User, error) {
	user := new(models.User)

	for _, u := range users {
		if u.ID == obj.UserID {
			user = u
			break
		}
	}

	if user == nil {
		return nil, errors.New("user with id not exists")
	}

	return user, nil
}

func (r *queryResolver) Meetups(ctx context.Context) ([]*models.Meetup, error) {
	return meetups, nil
}

func (r *userResolver) Meetups(ctx context.Context, obj *models.User) ([]*models.Meetup, error) {
	panic("implement me")
}

// Meetup returns generated.MeetupResolver implementation.
func (r *Resolver) Meetup() generated.MeetupResolver { return &meetupResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type meetupResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type userResolver struct{ *Resolver }
