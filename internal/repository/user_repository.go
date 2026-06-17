package repository

import (
	"database/sql"

	"github.com/akash/go-user-api/internal/models"
)

func CreateUser(
	db *sql.DB,
	req models.CreateUserRequest,
) (*models.User, error) {

	query := `
	INSERT INTO users(name,dob)
	VALUES($1,$2)
	RETURNING id,name,dob
	`

	user := &models.User{}

	err := db.QueryRow(
		query,
		req.Name,
		req.DOB,
	).Scan(
		&user.ID,
		&user.Name,
		&user.DOB,
	)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func GetUserByID(
	db *sql.DB,
	id int,
) (*models.User, error) {

	query := `
	SELECT id,name,dob
	FROM users
	WHERE id=$1
	`

	user := &models.User{}

	err := db.QueryRow(
		query,
		id,
	).Scan(
		&user.ID,
		&user.Name,
		&user.DOB,
	)

	if err != nil {
		return nil, err
	}

	return user, nil
}
func GetAllUsers(db *sql.DB) ([]models.User, error) {

	query := `
	SELECT id,name,dob
	FROM users
	ORDER BY id
	`

	rows, err := db.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var users []models.User

	for rows.Next() {

		var user models.User

		err := rows.Scan(
			&user.ID,
			&user.Name,
			&user.DOB,
		)

		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}
func UpdateUser(
	db *sql.DB,
	id int,
	req models.CreateUserRequest,
) (*models.User, error) {

	query := `
	UPDATE users
	SET name=$1,dob=$2
	WHERE id=$3
	RETURNING id,name,dob
	`

	user := &models.User{}

	err := db.QueryRow(
		query,
		req.Name,
		req.DOB,
		id,
	).Scan(
		&user.ID,
		&user.Name,
		&user.DOB,
	)

	if err != nil {
		return nil, err
	}

	return user, nil
}
func DeleteUser(
	db *sql.DB,
	id int,
) error {

	query := `
	DELETE FROM users
	WHERE id=$1
	`

	_, err := db.Exec(
		query,
		id,
	)

	return err
}
