package account

import "github.com/themanciraptor/Backend-photagea/internal/util"

// Model is container for user data
type Model struct {
	util.BaseModel
	AccountID int64  `json:"-"`
	Email     string `json:"Email"`
}

// ToRefList returns a list of references to make sql queries easier
func (m *Model) ToRefList() []interface{} {
	return []interface{}{&m.AccountID, &m.Email, &m.Created, &m.Updated, &m.Deleted}
}
