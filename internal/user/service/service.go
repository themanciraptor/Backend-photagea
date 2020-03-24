package userservice

import (
	"context"

	user "github.com/themanciraptor/Backend-photagea/internal/user/model"
	userrepo "github.com/themanciraptor/Backend-photagea/internal/user/repo"
)

// Interface is the service interface
type Interface interface {
	Get(context.Context, int64) (*user.Model, error)
	Create(context.Context, int64, string, string, string) (*user.Model, error)
	Update(context.Context, int64, string, string, string) (*user.Model, error)
}

// UserService implements user service interface
type UserService struct {
	repo userrepo.Interface
}

// Initialize a new User Service
func Initialize(r userrepo.Interface) Interface {
	return &UserService{repo: r}
}

// Get a user
func (u *UserService) Get(ctx context.Context, UserID int64) (*user.Model, error) {
	return u.repo.Get(ctx, UserID)
}

// Create a user
func (u *UserService) Create(ctx context.Context, UserID int64, Alias string, FirstName string, LastName string) (*user.Model, error) {
	return nil, nil
}

// Update a user
func (u *UserService) Update(ctx context.Context, UserID int64, Alias string, FirstName string, LastName string) (*user.Model, error) {
	return nil, nil
}
