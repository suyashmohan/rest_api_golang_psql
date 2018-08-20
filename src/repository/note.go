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
	err := nr.DB.QueryRow(sql, args...).Scan(&dbNote.ID, &dbNote.Text, &dbNote.UserID, &dbNote.CreatedOn, &dbNote.UpdatedOn)

	if err != nil {
		nr.logErr(err)
		return nil
	}

	return &dbNote
}

// ----- Public Methods -----

// New - Create new Note in DB
func (nr *NoteRepository) New(text, userID string) *entity.Note {
	dbNote := nr.runSQL("INSERT INTO notes(text, userid, createdOn, updatedOn) VALUES($1::TEXT, $2::INTEGER, now()::TIMESTAMP, now()::TIMESTAMP) RETURNING *", text, userID)
	return dbNote
}

// Get - Return a Note based on ID
func (nr *NoteRepository) Get(id, userID string) *entity.Note {
	dbNote := nr.runSQL("SELECT * FROM notes WHERE id = $1 AND userid = $2 LIMIT 1", id, userID)
	return dbNote
}

// Update - Update the text for Note
func (nr *NoteRepository) Update(id, userID, text string) *entity.Note {
	dbNote := nr.runSQL("UPDATE notes SET text=$3, updatedon=now()::TIMESTAMP WHERE id=$1 AND userid=$2 RETURNING *", id, userID, text)
	return dbNote
}

// Delete - Delete a record from DB
func (nr *NoteRepository) Delete(id, userID string) *entity.Note {
	dbNote := nr.runSQL("DELETE FROM notes WHERE id=$1 AND userid=$2 RETURNING *", id, userID)
	return dbNote
}
