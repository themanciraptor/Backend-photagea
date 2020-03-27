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
	Upload(ctx context.Context, accountID int64, mimetype string, imageData []byte) (string, error)
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

// Upload a single image, return the serving url
func (s *Service) Upload(ctx context.Context, accountID int64, mimetype string, imageData []byte) (string, error) {
	return s.repo.Upload(ctx, &imagedata.Model{
		AccountID: accountID,
		MimeType:  mimetype,
		ImageData: imageData,
	})
}
