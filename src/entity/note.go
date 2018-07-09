package entity

import "time"

// Note - Holds the Note Data
type Note struct {
	ID        int64     `json:"id"`
	Text      string    `json:"text"`
	CreatedOn time.Time `json:"createdOn"`
	UpdatedOn time.Time `json:"updatedOn"`
}
