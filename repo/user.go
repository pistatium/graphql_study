package repo

import (
	"context"
	"fmt"
	"github.com/pistatium/graphql_study/graph/model"
)

type UserRepo interface {
	GetUser(ctx context.Context, id string) (*model.User, error)
}

type DummyUserRepo struct {
	Users []*model.User
}

func (d DummyUserRepo) GetUser(ctx context.Context, id string) (*model.User, error) {
	for _, user := range d.Users {
		if user.ID == id {
			return user, nil
		}
	}
	return nil, fmt.Errorf("not found")
}

func NewDummyUserRepo() UserRepo {
	return &DummyUserRepo{
		Users: []*model.User{
			{
				ID:   "1",
				Name: "User1",
			},
			{
				ID:   "2",
				Name: "User2",
			},
		},
	}
}
