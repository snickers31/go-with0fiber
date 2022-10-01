// ./app/queries/user_query.go
package queries

import (
	"fmt"

	"github.com/snickers31/go-with-fiber/app/models"
)

func (q *DBQueries) CreateUser(u *models.User) error {

	query := `INSERT INTO users (id,username,password,created_at) VALUES ($1,$2,$3,$4)`

	result, err := q.Exec(query, u.ID, u.Username, u.Password, u.CreatedAt)
	if err != nil {
		return err
	}

	fmt.Println("Ini dari query")
	fmt.Println(result)

	return nil
}

func (q *DBQueries) UserCredential(u *models.User) (*models.User, error) {
	user := models.User{}
	query := `SELECT id,username,password FROM users WHERE username = $1`

	err := q.Get(&user, query, u.Username)
	if err != nil {
		return &user, fmt.Errorf("Unknown account")
	}

	return &user, nil

}

func (q *DBQueries) GetUsers() ([]models.User, error) {
	users := []models.User{}

	query := `SELECT * FROM users where isDeleted = false`

	err := q.Select(&users, query)
	if err != nil {
		return users, err
	}

	return users, err

}
