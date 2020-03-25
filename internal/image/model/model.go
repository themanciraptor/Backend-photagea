package image

import "github.com/themanciraptor/Backend-photagea/internal/util"

// Model is container for user data
type Model struct {
	util.BaseModel
	AccountID int64  `json:"-"`
	ImageID   string `json:"-"`
	URL       string `json:"URL"`
}

// ToRefList returns a list of references to make sql queries easier
func (m *Model) ToRefList() []interface{} {
	return []interface{}{&m.ImageID, &m.AccountID, &m.URL, &m.Created, &m.Updated, &m.Deleted}
}

// CopyFrom is a shallow copy of one image model to another
func (m *Model) CopyFrom(other *Model) {
	m.AccountID = other.AccountID
	m.URL = other.URL
	m.ImageID = other.ImageID
	m.Created = other.Created
	m.Deleted = other.Deleted
	m.Updated = other.Updated
}
