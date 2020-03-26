package accountrepo

/* TODO: Some of this functionality can be abstracted out */

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	imagedata "github.com/themanciraptor/Backend-photagea/internal/imagedata/model"
	"github.com/themanciraptor/Backend-photagea/internal/util"
)

// Interface is the interface for user repository interactions
type Interface interface {
	Get(ctx context.Context, accountID int64, imageDataID int64) (*imagedata.Model, error)
	Upload(ctx context.Context, i *imagedata.Model) error
	Delete(ctx context.Context, imageDataID string) error
}

// Repository implements the repo Interface
type Repository struct {
	db *sql.DB
}

// Initialize the repository
func Initialize(db *sql.DB) Interface {
	return &Repository{db: db}
}

const filepath = "images/%d/%d"

// Get gets a single imagedata
func (r *Repository) Get(ctx context.Context, accountID int64, imageDataID int64) (*imagedata.Model, error) {
	row := r.db.QueryRowContext(ctx, "SELECT * FROM imagedata WHERE `ImageDataID`=?;", imageDataID)

	// TODO: REMOVE date processor, and rely on standard sqlnullable types
	d := imagedata.Model{}
	dproc := util.DateProcessor{}
	err := row.Scan(util.AugmentRefList(&dproc, d.ToRefList())...)
	if err != nil {
		return nil, err
	}

	err = dproc.ProcessDates()
	if err != nil {
		return nil, err
	}

	f, err := os.Open(fmt.Sprintf(filepath, accountID, imageDataID))
	if err != nil {
		return nil, err
	}

	d.ImageData = f

	return &d, nil
}

// Upload saves an image to the server
func (r *Repository) Upload(ctx context.Context, i *imagedata.Model) error {
	_, err := r.db.ExecContext(ctx, "INSERT INTO imagedata (`AccountID`, `mimetype`) VALUES (?, ?)", i.AccountID, i.MimeType)

	return err
}

// Delete an image
func (r *Repository) Delete(ctx context.Context, imageDataID string) error {
	conn, err := r.db.Conn(ctx)
	if err != nil {
		return err
	}
	defer conn.Close()

	rows := conn.QueryRowContext(ctx, "UPDATE imagedata SET DateDeleted=NOW() WHERE `ImageDataID`=?;", imageDataID)
	err = rows.Scan()
	if err != nil {
		return err
	}

	return nil
}
