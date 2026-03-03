//struct
//store pgxpool.Pool
//crud methods(functions)
//repositry is responible for taling to the db
//executing quires
//runing models.User

package repository

import (
	"api/models"
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository struct {
	DB *pgxpool.Pool
}

func NewUser(pool *pgxpool.Pool) *UserRepository {
	//return struct with pool
	return &UserRepository{
		DB: pool,
	}
}
func (r *UserRepository) CreateUser(user models.User) error {
	//query
	query := "INSERT INTO users(name, email)VALUES($1 , $2)"

	_, err := r.DB.Exec(
		context.Background(), query, user.Name, user.Email,
	)
	return err

}

func (r *UserRepository) GetUsers() ([]models.User, error) {

	query := "SELECT id, name, email FROM users"

	rows, err := r.DB.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User

	for rows.Next() {
		var user models.User

		err := rows.Scan(&user.ID, &user.Name, &user.Email)
		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (r *UserRepository) GetUserByID(id int) (models.User, error) {
	query := "SELECT id , name,email FROM users WHERE id = $1"
	var user models.User
	err := r.DB.QueryRow(context.Background(), query, id).Scan(&user.ID, &user.Name, &user.Email)

	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (r *UserRepository) DeleteUser(id int) error {

	query := "DELETE FROM users WHERE id = $1"

	cmdTag, err := r.DB.Exec(context.Background(), query, id)
	if err != nil {
		return err
	}

	if cmdTag.RowsAffected() == 0 {
		return fmt.Errorf("user not found")
	}

	return nil
}
