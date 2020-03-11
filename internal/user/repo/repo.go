package userrepo

import (
	"context"
	"database/sql"

	user "github.com/themanciraptor/Backend-photagea/internal/user/model"
	"github.com/themanciraptor/Backend-photagea/internal/util"
)

// Interface is the interface for user repository interactions
type Interface interface {
	Get(context.Context, string) (*user.Model, error)
	Update(context.Context, *user.Model) error
	Create(context.Context, *user.Model) error
	Delete(context.Context, string) error
}

// Repository implements the repo Interface
type Repository struct {
	db *sql.DB
}

// Initialize the repository
func Initialize(db *sql.DB) Interface {
	return &Repository{db: db}
}

// Get gets a single user
func (r *Repository) Get(ctx context.Context, UserID string) (*user.Model, error) {
	conn, err := r.db.Conn(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	rows := conn.QueryRowContext(ctx, "SELECT * FROM User2 WHERE `UserID`=?;", UserID)

	u := user.Model{}
	dproc := util.DateProcessor{}
	err = rows.Scan(util.AugmentRefList(&dproc, u.ToRefList())...)
	if err != nil {
		return nil, err
	}

	err = dproc.ProcessDates()
	if err != nil {
		return nil, err
	}

	return &u, nil
}

// Update attempts to update a single row in the user table
func (r *Repository) Update(ctx context.Context, User *user.Model) error {
	// TODO: finish figuring out how to make updates

	conn, err := r.db.Conn(ctx)
	if err != nil {
		return err
	}
	defer conn.Close()

	rows := conn.QueryRowContext(ctx, "UPDATE User2 SET ", User.UserID)

	u := user.Model{}
	dproc := util.DateProcessor{}
	err = rows.Scan(util.AugmentRefList(&dproc, u.ToRefList())...)
	if err != nil {
		return err
	}

	err = dproc.ProcessDates()
	if err != nil {
		return err
	}

	return nil
}

// Create a user
func (r *Repository) Create(ctx context.Context, User *user.Model) error {
	conn, err := r.db.Conn(ctx)
	if err != nil {
		return err
	}
	defer conn.Close()

	rows := conn.QueryRowContext(ctx, "INSERT INTO User2 (`UserID`, `Alias`, `FirstName`, `LastName`, `Email`) VALUES ( ?, ?, ?, ?, ?)", User.ToRefList()[:5]...)

	err = rows.Scan()
	if err != nil && err != sql.ErrNoRows {
		return err
	}

	return nil
}

// Delete a user
func (r *Repository) Delete(ctx context.Context, UserID string) error {
	conn, err := r.db.Conn(ctx)
	if err != nil {
		return err
	}
	defer conn.Close()

	rows := conn.QueryRowContext(ctx, "Update User2 SET DateDeleted=NOW() WHERE `UserID`=?;", UserID)
	err = rows.Scan()
	if err != nil {
		return err
	}

	return nil
}
