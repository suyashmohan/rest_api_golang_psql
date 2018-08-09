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

// ----- Private Methods -----

// Throw Fatal Error
func (nr *NoteRepository) logErr(err error) {
	log.Println("[Note Repository] ", err)
}

// Run SQL
func (nr *NoteRepository) runSQL(sql string, args ...interface{}) *entity.Note {
	dbNote := entity.Note{}
	err := nr.DB.QueryRow(sql, args...).Scan(&dbNote.ID, &dbNote.Text, &dbNote.CreatedOn, &dbNote.UpdatedOn)

	if err != nil {
		nr.logErr(err)
		return nil
	}

	return &dbNote
}

// ----- Public Methods -----

// New - Create new Note in DB
func (nr *NoteRepository) New(text string) *entity.Note {
	dbNote := nr.runSQL("INSERT INTO notes(text, createdOn, updatedOn) VALUES($1::TEXT, now()::TIMESTAMP, now()::TIMESTAMP) RETURNING *", text)
	return dbNote
}

// Get - Return a Note based on ID
func (nr *NoteRepository) Get(id string) *entity.Note {
	dbNote := nr.runSQL("SELECT * FROM notes WHERE id = $1 LIMIT 1", id)
	return dbNote
}

// Update - Update the text for Note
func (nr *NoteRepository) Update(id, text string) *entity.Note {
	dbNote := nr.runSQL("UPDATE notes SET text=$2, updatedon=now()::TIMESTAMP WHERE id=$1 RETURNING *", id, text)
	return dbNote
}

// Delete - Delete a record from DB
func (nr *NoteRepository) Delete(id string) *entity.Note {
	dbNote := nr.runSQL("DELETE FROM notes WHERE id=$1 RETURNING *", id)
	return dbNote
}
