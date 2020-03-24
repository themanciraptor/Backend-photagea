package account

import "github.com/themanciraptor/Backend-photagea/internal/util"

// Model is container for user data
type Model struct {
	util.BaseModel
	AccountID int64  `json:"-"`
	Email     string `json:"Email"`
}
