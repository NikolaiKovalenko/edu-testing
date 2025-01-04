package sql_db_mock

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestGetUserByID(t *testing.T) {
	t.Parallel()
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	userID := 1
	userName := "John Doe"

	// Setup expected query and result
	rows := sqlmock.NewRows([]string{"id", "name"}).
		AddRow(userID, userName)

	mock.ExpectQuery("SELECT id, name FROM users WHERE id = \\$1").
		WithArgs(userID).
		WillReturnRows(rows)

	// Call the function
	user, err := GetUserByID(db, userID)

	assert.NoError(t, err, "expected no error")
	assert.NotNil(t, user, "expected a valid user object")
	assert.Equal(t, userID, user.ID, "expected user ID to match")
	assert.Equal(t, userName, user.Name, "expected user Name to match")

	// Ensure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

// OUTPUT:
//=== RUN   TestGetUserByID
//--- PASS: TestGetUserByID (0.00s)
//PASS
//
//Process finished with the exit code 0
