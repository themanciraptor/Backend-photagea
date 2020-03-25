package imagerepo

/* TODO: Some of this functionality can be abstracted out */

import (
	"context"
	"database/sql"

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
	conn, err := r.db.Conn(ctx)
	if err != nil {
		return err
	}
	defer conn.Close()

	rows := conn.QueryRowContext(ctx, "INSERT INTO image (`URL`, `AccountID`) VALUES ( ?, ?)", a.ToRefList()[1:3]...)

	err = rows.Scan()
	if err != nil && err != sql.ErrNoRows {
		return err
	}

	return nil
}

// Delete a image
func (r *Repository) Delete(ctx context.Context, accountID string) error {
	conn, err := r.db.Conn(ctx)
	if err != nil {
		return err
	}
	defer conn.Close()

	rows := conn.QueryRowContext(ctx, "Update image SET DateDeleted=NOW() WHERE `accountID`=?;", accountID)
	err = rows.Scan()
	if err != nil {
		return err
	}

	return nil
}

// List returns a list of images for the account
func (r *Repository) List(ctx context.Context, accountID int64, limit int, cursor int64) ([]*image.Model, int64, error) {
	rows, err := r.db.QueryContext(ctx, "SELECT * FROM image WHERE accountID=? limit ? offset ?", accountID, limit, cursor)
	if err != nil {
		return nil, 0, err
	}

	images := make([]*image.Model, limit)
	// temp := image.Model{}

	err = rows.Scan(images[0].ToRefList()...)
	if err != nil {
		return nil, 0, err
	}

	for i := 1; rows.Next(); i++ {
		rows.Scan(images[i].ToRefList()...)
	}

	return images, cursor + int64(len(images)), nil
}