package userrepo

import (
	"context"
	"database/sql"
	"fmt"

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
	DateCreated := []uint8{}
	DateUpdated := []uint8{}
	err = rows.Scan(&u.UserID, &u.Alias, &u.FirstName, &u.LastName, &u.Email, &DateCreated, &DateUpdated)
	if err != nil {
		return nil, err
	}

	u.Created = util.IntsToTime(DateCreated)
	u.Updated = util.IntsToTime(DateUpdated)

	fmt.Println("for good")

	return &u, nil
}

// Update gets a single user
func (r *Repository) Update(ctx context.Context, User *user.Model) error {
	return nil
}

// Create gets a single user
func (r *Repository) Create(ctx context.Context, User *user.Model) error {
	return nil
}

// Delete gets a single user
func (r *Repository) Delete(ctx context.Context, UserID string) error {
	return nil
}
