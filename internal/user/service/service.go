package userservice

import (
	"context"

	user "github.com/themanciraptor/Backend-photagea/internal/user/model"
	userrepo "github.com/themanciraptor/Backend-photagea/internal/user/repo"
)

// Interface is the service interface
type Interface interface {
	Get(context.Context, int64) (*user.Model, error)
	Create(context.Context, int64, string, string, string) error
	Update(context.Context, int64, string, string, string) error
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
func (u *UserService) Get(ctx context.Context, AccountID int64) (*user.Model, error) {
	return u.repo.Get(ctx, AccountID)
}

// Create a user
func (u *UserService) Create(ctx context.Context, AccountID int64, Alias string, FirstName string, LastName string) error {
	return u.repo.Create(ctx, &user.Model{
		AccountID: AccountID,
		Alias:     Alias,
		FirstName: FirstName,
		LastName:  LastName,
	})
}

// Update a user
func (u *UserService) Update(ctx context.Context, AccountID int64, Alias string, FirstName string, LastName string) error {
	return u.repo.Create(ctx, &user.Model{
		AccountID: AccountID,
		Alias:     Alias,
		FirstName: FirstName,
		LastName:  LastName,
	})
}
