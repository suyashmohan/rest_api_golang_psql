package repository

import (
	"database/sql"
	"log"

	"../entity"
)

// NoteRepository - Repository for Database Layer of Notes
type NoteRepository struct {
	DB *sql.DB
}

// New - Create new Note in DB
func (nr *NoteRepository) New(text string) *entity.Note {
	dbNote := entity.Note{}
	row := nr.DB.QueryRow("INSERT INTO notes(text, createdOn, updatedOn) VALUES($1::TEXT, now()::TIMESTAMP, now()::TIMESTAMP) RETURNING *", text)
	err := row.Scan(&dbNote.ID, &dbNote.Text, &dbNote.CreatedOn, &dbNote.UpdatedOn)

	if err != nil {
		log.Fatal(err)
	}

	return &dbNote
}

// Get - Return a Note based on ID
func (nr *NoteRepository) Get(id string) *entity.Note {
	dbNote := entity.Note{}
	row := nr.DB.QueryRow("SELECT * FROM notes WHERE id = $1 LIMIT 1", id)
	err := row.Scan(&dbNote.ID, &dbNote.Text, &dbNote.CreatedOn, &dbNote.UpdatedOn)

	if err == nil {
		return &dbNote
	}

	return nil
}

// Update - Update the text for Note
func (nr *NoteRepository) Update(id, text string) *entity.Note {
	dbNote := entity.Note{}
	row := nr.DB.QueryRow("UPDATE notes SET text=$2, updatedon=now()::TIMESTAMP WHERE id=$1 RETURNING *", id, text)
	err := row.Scan(&dbNote.ID, &dbNote.Text, &dbNote.CreatedOn, &dbNote.UpdatedOn)
	if err == nil {
		return &dbNote
	}

	return nil
}

// Delete - Delete a record from DB
func (nr *NoteRepository) Delete(id string) *entity.Note {
	dbNote := entity.Note{}
	row := nr.DB.QueryRow("DELETE FROM notes WHERE id=$1 RETURNING *", id)
	err := row.Scan(&dbNote.ID, &dbNote.Text, &dbNote.CreatedOn, &dbNote.UpdatedOn)
	if err == nil {
		return &dbNote
	}

	return nil
}
