package accountservice

import (
	"context"

	account "github.com/themanciraptor/Backend-photagea/internal/account/model"
	accountrepo "github.com/themanciraptor/Backend-photagea/internal/account/repo"
	user "github.com/themanciraptor/Backend-photagea/internal/user/model"
	userrepo "github.com/themanciraptor/Backend-photagea/internal/user/repo"
)

// Interface is the service interface
type Interface interface {
	Get(ctx context.Context, AccountID int64) (*account.Model, error)
	Create(ctx context.Context, AccountID int64, Email string) error
	Update(ctx context.Context, AccountID int64, Email string) error
}

// Service implements account service interface
type Service struct {
	repo accountrepo.Interface
}

// Initialize a new User Service
func Initialize(r userrepo.Interface) Interface {
	return &Service{repo: r}
}

// Get a user
func (u *Service) Get(ctx context.Context, AccountID int64) (*account.Model, error) {
	return u.repo.Get(ctx, AccountID)
}

// Create a user
func (u *Service) Create(ctx context.Context, AccountID int64, Email string) error {
	return u.repo.Create(ctx, &user.Model{
		AccountID: AccountID,
		Email:     Email,
	})
}

// Update a user
func (u *Service) Update(ctx context.Context, AccountID int64, Email string) error {
	return u.repo.Update(ctx, &user.Model{
		AccountID: AccountID,
		Email:     Email,
	})
}
