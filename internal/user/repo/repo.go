package userrepo

import (
	"database/sql"

	user "github.com/themanciraptor/Backend-photagea/internal/user/model"
)

// Interface is the interface for user repository interactions
type Interface interface {
	Get(string) (string, error)
	Update(user.Model) error
	Create(user.Model) error
	Delete(string) error
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
func (r *Repository) Get(UserID string) (string, error) {
	return "user", nil
}

// Update gets a single user
func (r *Repository) Update(User user.Model) error {
	return nil
}

// Create gets a single user
func (r *Repository) Create(User user.Model) error {
	return nil
}

// Delete gets a single user
func (r *Repository) Delete(UserID string) error {
	return nil
}
