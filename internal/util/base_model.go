/*Package util contains basica features all models should contain*/
package util

import "time"

// BaseModel contains common fields for all models
type BaseModel struct {
	Created time.Time
	Updated time.Time
}
