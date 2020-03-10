package user

import "github.com/themanciraptor/Backend-photagea/internal/util"

// Model is container for user data
type Model struct {
	util.BaseModel
	UserID    int64
	Alias     string
	FirstName string
	LastName  string
	Email     string
}
