package accountrepo

import (
	"context"
	"database/sql"

	account "github.com/themanciraptor/Backend-photagea/internal/user/model"
	"github.com/themanciraptor/Backend-photagea/internal/util"
)

// Interface is the interface for user repository interactions
type Interface interface {
	Get(context.Context, int64) (*account.Model, error)
	Update(context.Context, *account.Model) error
	Create(context.Context, *account.Model) error
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
func (r *Repository) Get(ctx context.Context, AccountID int64) (*account.Model, error) {
	conn, err := r.db.Conn(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	rows := conn.QueryRowContext(ctx, "SELECT * FROM account WHERE `AccountID`=?;", AccountID)

	// TODO: REMOVE date processor, and rely on standard sqlnullable types
	u := account.Model{}
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

// Update attempts to update a single row in the account table
func (r *Repository) Update(ctx context.Context, account *account.Model) error {
	qb := util.InitUpdateQueryBuilder("account").
		Add("Email", account.Email).
		AddFilter("AccountID", account.AccountID)

	_, err := qb.ExecuteQuery(ctx, r.db)

	return err
}

// Create a account
func (r *Repository) Create(ctx context.Context, account *account.Model) error {
	conn, err := r.db.Conn(ctx)
	if err != nil {
		return err
	}
	defer conn.Close()

	rows := conn.QueryRowContext(ctx, "INSERT INTO account (`Email`) VALUES ( ? )", &account.Email)

	err = rows.Scan()
	if err != nil && err != sql.ErrNoRows {
		return err
	}

	return nil
}

// Delete a account
func (r *Repository) Delete(ctx context.Context, AccountID string) error {
	conn, err := r.db.Conn(ctx)
	if err != nil {
		return err
	}
	defer conn.Close()

	rows := conn.QueryRowContext(ctx, "Update account SET DateDeleted=NOW() WHERE `AccountID`=?;", AccountID)
	err = rows.Scan()
	if err != nil {
		return err
	}

	return nil
}
