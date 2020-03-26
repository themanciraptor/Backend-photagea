package imagerepo

/* TODO: Some of this functionality can be abstracted out */

import (
	"context"
	"database/sql"
	"fmt"

	image "github.com/themanciraptor/Backend-photagea/internal/image/model"
)

// Interface is the interface for user repository interactions
type Interface interface {
	Create(ctx context.Context, i *image.Model) error
	Delete(ctx context.Context, imageID string) error
	List(ctx context.Context, accountID int64, limit int, cursor int64) ([]*image.Model, int64, error)
}

// Repository implements the repo Interface
type Repository struct {
	db *sql.DB
}

// Initialize the repository
func Initialize(db *sql.DB) Interface {
	return &Repository{db: db}
}

// Create a image
func (r *Repository) Create(ctx context.Context, a *image.Model) error {
	_, err := r.db.ExecContext(ctx, "INSERT INTO image (`AccountID`, `URL`) VALUES ( ?, ?)", a.ToRefList()[1:3]...)

	return err
}

// Delete a image
func (r *Repository) Delete(ctx context.Context, accountID string) error {
	_, err := r.db.ExecContext(ctx, "Update image SET DateDeleted=NOW() WHERE `accountID`=?;", accountID)
	return err
}

// List returns a list of images for the account
func (r *Repository) List(ctx context.Context, accountID int64, limit int, cursor int64) ([]*image.Model, int64, error) {
	rows, err := r.db.QueryContext(ctx, "SELECT * FROM image WHERE accountID=? limit ? offset ?", accountID, limit, cursor)
	if err != nil {
		return nil, 0, err
	}

	images := make([]*image.Model, 0, limit)
	for i := 0; rows.Next(); i++ {
		fmt.Println("Getting a row")
		images = append(images, new(image.Model))
		rows.Scan(images[i].ToRefList()...)
		if err != nil {
			return nil, 0, err
		}
	}

	fmt.Println("images: ", images[0].URL, images[0].AccountID)

	return images, cursor + int64(len(images)), nil
}
