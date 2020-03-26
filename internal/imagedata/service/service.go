package imagedataservice

import (
	"context"

	imagedata "github.com/themanciraptor/Backend-photagea/internal/imagedata/model"
	imagedatarepo "github.com/themanciraptor/Backend-photagea/internal/imagedata/repo"
)

// Interface for the image service
type Interface interface {
	Get(ctx context.Context, accountID int64, imageDataID int64) (*imagedata.Model, error)
	// Delete()
	// Upload()
}

// Service is the image service implementation
type Service struct {
	repo imagedatarepo.Interface
}

// Initialize an instance of the image service
func Initialize(r imagedatarepo.Interface) Interface {
	return &Service{repo: r}
}

// Get an image
func (s *Service) Get(ctx context.Context, accountID int64, imageDataID int64) (*imagedata.Model, error) {
	return s.repo.Get(ctx, accountID, imageDataID)
}
