/*Package util contains basica features all models should contain*/
package util

import "time"

// BaseModel contains common fields for all models
type BaseModel struct {
	Created time.Time
	Updated time.Time
	Deleted time.Time
}

// AugmentRefList replaces time.Time objects with strings
func AugmentRefList(d *DateProcessor, refs []interface{}) []interface{} {
	for i, ref := range refs {
		if t, ok := ref.(*time.Time); ok {
			refs[i] = d.Add(t)
		}
	}

	return refs
}
