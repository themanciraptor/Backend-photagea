package util

import (
	"database/sql"
	"fmt"
	"strings"
	"time"
)

// DateProcessor and associated methods provide an easy way to convert multiple dates returned from sql
type DateProcessor struct {
	unprocessedDates []*sql.NullTime
	destinations     []*time.Time
}

// Add another date to be processed
func (d *DateProcessor) Add(dest *time.Time) *sql.NullTime {
	d.destinations = append(d.destinations, dest)
	d.unprocessedDates = append(d.unprocessedDates, new(sql.NullTime))

	return d.unprocessedDates[len(d.unprocessedDates)-1]
}

// ProcessDates converts a list of integers into a time object
func (d *DateProcessor) ProcessDates() error {
	if len(d.destinations) != len(d.unprocessedDates) {
		return fmt.Errorf("Date Processor: internal error")
	}

	var errors []error
	for i, date := range d.unprocessedDates {
		if date.Valid {
			*d.destinations[i] = date.Time
		}
	}

	if len(errors) > 0 {
		var b strings.Builder
		for _, err := range errors {
			b.WriteString(err.Error())
		}

		return fmt.Errorf(b.String())
	}

	return nil
}
