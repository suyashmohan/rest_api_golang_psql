package entity

import "time"

// Note - Holds the Note Data
type Note struct {
	ID        int64     `json:"id"`
	UserID    int64     `json:"userId"`
	Text      string    `json:"text"`
	CreatedOn time.Time `json:"createdOn"`
	UpdatedOn time.Time `json:"updatedOn"`
}
