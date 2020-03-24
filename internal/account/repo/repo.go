package accountrepo

/* TODO: Some of this functionality can be abstracted out */

import (
	"context"
	"database/sql"
	"log"

	account "github.com/themanciraptor/Backend-photagea/internal/account/model"
	"github.com/themanciraptor/Backend-photagea/internal/util"
	"golang.org/x/crypto/bcrypt"
)

// Interface is the interface for user repository interactions
type Interface interface {
	Get(ctx context.Context, email string, password string) (*account.Model, error) // May remove this function
	Update(ctx context.Context, a *account.Model) error
	Create(ctx context.Context, a *account.Model) error
	Delete(ctx context.Context, accountID string) error
}

// Repository implements the repo Interface
type Repository struct {
	db *sql.DB
}

// Initialize the repository
func Initialize(db *sql.DB) Interface {
	return &Repository{db: db}
}

// Get gets a single account
func (r *Repository) Get(ctx context.Context, email string, password string) (*account.Model, error) {
	conn, err := r.db.Conn(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	rows := conn.QueryRowContext(ctx, "SELECT * FROM account WHERE `Email`=?;", email)

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

	err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
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
func (r *Repository) Create(ctx context.Context, a *account.Model) error {
	conn, err := r.db.Conn(ctx)
	if err != nil {
		return err
	}
	defer conn.Close()

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(a.Password), 12)
	if err != nil {
		log.Printf("Unable to hash password: %s", err)
	}

	rows := conn.QueryRowContext(ctx, "INSERT INTO account (`Email`, `Password`) VALUES ( ?, ? )", &a.Email, &hashedPassword)

	err = rows.Scan()
	if err != nil && err != sql.ErrNoRows {
		return err
	}

	return nil
}

// Delete a account
func (r *Repository) Delete(ctx context.Context, accountID string) error {
	conn, err := r.db.Conn(ctx)
	if err != nil {
		return err
	}
	defer conn.Close()

	rows := conn.QueryRowContext(ctx, "Update account SET DateDeleted=NOW() WHERE `AccountID`=?;", accountID)
	err = rows.Scan()
	if err != nil {
		return err
	}

	return nil
}
