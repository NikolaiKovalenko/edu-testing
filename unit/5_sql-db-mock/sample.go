package sql_db_mock

import (
	"database/sql"
	"fmt"
)

type User struct {
	ID   int
	Name string
}

func GetUserByID(db *sql.DB, userID int) (*User, error) {
	query := "SELECT id, name FROM users WHERE id = $1"
	row := db.QueryRow(query, userID)

	var user User
	if err := row.Scan(&user.ID, &user.Name); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user not found")
		}
		return nil, err
	}
	return &user, nil
}
