package util

import (
	"fmt"
	"time"
)

// DateProcessor and associated methods provide an easy way to convert multiple dates returned from sql
type DateProcessor struct {
	unprocessedDates []*string
	destinations     []interface{}
}

// RefList returns the list of references to make sql references easier
func (d *DateProcessor) RefList() []*string {
	return d.unprocessedDates
}

// Add another date to be processed
func (d *DateProcessor) Add(dest *time.Time) {
	d.destinations = append(d.destinations, dest)
	d.unprocessedDates = append(d.unprocessedDates, new(string))
}

// ProcessDates converts a list of integers into a time object
func (d *DateProcessor) ProcessDates() error {
	if len(d.destinations != d.unprocessedDates) {
		return fmt.Errorf("Date Processor: internal error")
	}
	
	var errors []error
	for i, var := range d.unprocessedDates {
		d.destinations[i], err := time.Parse("2006-01-02 15:04:05", var)
		if err != nil {
			errors = append(errors, err)
		}
	}

	if len(errors > 0) {
		b strings.Builder
		b.
	}


	return nil
}
