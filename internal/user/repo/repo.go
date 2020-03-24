package userrepo

import (
	"context"
	"database/sql"

	user "github.com/themanciraptor/Backend-photagea/internal/user/model"
	"github.com/themanciraptor/Backend-photagea/internal/util"
)

// Interface is the interface for user repository interactions
type Interface interface {
	Get(context.Context, int64) (*user.Model, error)
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
func (r *Repository) Get(ctx context.Context, AccountID int64) (*user.Model, error) {
	conn, err := r.db.Conn(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	rows := conn.QueryRowContext(ctx, "SELECT * FROM User WHERE `AccountID`=?;", AccountID)

	// TODO: REMOVE date processor, and rely on standard sqlnullable types
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
	qb := util.InitUpdateQueryBuilder("User").
		Add("Alias", User.Alias).
		Add("FirstName", User.FirstName).
		Add("LastName", User.LastName).
		AddFilter("AccountID", User.AccountID)

	_, err := qb.ExecuteQuery(ctx, r.db)

	return err
}

// Create a user
func (r *Repository) Create(ctx context.Context, User *user.Model) error {
	conn, err := r.db.Conn(ctx)
	if err != nil {
		return err
	}
	defer conn.Close()

	rows := conn.QueryRowContext(ctx, "INSERT INTO User (`Alias`, `FirstName`, `LastName`, `AccountID`) VALUES ( ?, ?, ?, ?)", User.ToRefList()[1:5]...)

	err = rows.Scan()
	if err != nil && err != sql.ErrNoRows {
		return err
	}

	return nil
}

// Delete a user
func (r *Repository) Delete(ctx context.Context, AccountID string) error {
	conn, err := r.db.Conn(ctx)
	if err != nil {
		return err
	}
	defer conn.Close()

	rows := conn.QueryRowContext(ctx, "Update User SET DateDeleted=NOW() WHERE `AccountID`=?;", AccountID)
	err = rows.Scan()
	if err != nil {
		return err
	}

	return nil
}
