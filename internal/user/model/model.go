package user

import "github.com/themanciraptor/Backend-photagea/internal/util"

// Model is container for user data
type Model struct {
	util.BaseModel
	UserID    int64
	AccountID int64
	Alias     string
	FirstName string
	LastName  string
}

// ToRefList returns a list of references to make sql queries easier
func (m *Model) ToRefList() []interface{} {
	return []interface{}{&m.UserID, &m.Alias, &m.FirstName, &m.LastName, &m.AccountID, &m.Created, &m.Updated, &m.Deleted}
}
