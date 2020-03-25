package imageservice

import (
	"context"

	image "github.com/themanciraptor/Backend-photagea/internal/image/model"
	imagerepo "github.com/themanciraptor/Backend-photagea/internal/image/repo"
)

// Interface for the image service
type Interface interface {
	List(ctx context.Context, accountID int64, limit int, cursor int64) ([]*image.Model, int64, error)
	Create(ctx context.Context, URL string, accountID int64) error
}

// Service is the image service implementation
type Service struct {
	repo imagerepo.Interface
}

// Initialize an instance of the image service
func Initialize(r imagerepo.Interface) Interface {
	return &Service{repo: r}
}

// List images for an account
func (s *Service) List(ctx context.Context, accountID int64, limit int, cursor int64) ([]*image.Model, int64, error) {
	return s.repo.List(ctx, accountID, limit, cursor)
}

// Create an image
func (s *Service) Create(ctx context.Context, URL string, accountID int64) error {
	return s.repo.Create(ctx, &image.Model{
		AccountID: accountID,
		URL:       URL,
	})
}
