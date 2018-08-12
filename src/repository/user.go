package repository

import (
	"database/sql"
	"log"

	"../entity"
)

// UserRepository - Repository for Database Layer of User
type UserRepository struct {
	DB *sql.DB
}

// ----- Private Methods -----

// Throw Fatal Error
func (ur *UserRepository) logErr(err error) {
	log.Println("[User Repository] ", err)
}

// Run SQL
func (ur *UserRepository) runSQL(sql string, args ...interface{}) *entity.User {
	dbUser := entity.User{}
	err := ur.DB.QueryRow(sql, args...).Scan(&dbUser.ID, &dbUser.Username, &dbUser.Password, &dbUser.CreatedOn, &dbUser.UpdatedOn)

	if err != nil {
		ur.logErr(err)
		return nil
	}

	return &dbUser
}

// ----- Public Methods -----

// New - Create new User in DB
func (ur *UserRepository) New(username, password string) *entity.User {
	dbUser := ur.runSQL("INSERT INTO users(username, password, createdOn, updatedOn) VALUES($1::TEXT, $2::TEXT, now()::TIMESTAMP, now()::TIMESTAMP) RETURNING *", username, password)
	return dbUser
}

// Get - Get a User
func (ur *UserRepository) Get(username string) *entity.User {
	dbUser := ur.runSQL("SELECT * FROM users WHERE username = $1::TEXT LIMIT 1", username)
	return dbUser
}
