package imagedata

import (
	"os"

	"github.com/themanciraptor/Backend-photagea/internal/util"
)

// TODO: Add permissions

// Model is container for user data
type Model struct {
	util.BaseModel
	AccountID   int64    `json:"-"`
	ImageDataID int64    `json:"ImageDataID"`
	MimeType    string   `json:"mimetype"`
	ImageData   *os.File `json:"-"`
}

// ToRefList returns a list of references to make sql queries easier
func (m *Model) ToRefList() []interface{} {
	return []interface{}{&m.AccountID, &m.ImageDataID, &m.Created, &m.Updated, &m.Deleted}
}
