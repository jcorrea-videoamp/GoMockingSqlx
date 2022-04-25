package main

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type User struct {
	FirstName string
	LastName  string
	Username  string
	Password  string
}

type UserRepository interface {
	Insert(user User) (int, error)
	GetById(id int) (User, error)
	Get(username, password string) (User, error)
}

type userRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *userRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Insert(user User) (int, error) {
	var id int
	row := r.db.QueryRow("INSERT INTO users (first_name, last_name, username, password) VALUES ($1, $2, $3, $4) RETURNING id",
		user.FirstName, user.LastName, user.Username, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *userRepository) GetById(id int) (User, error) {
	fmt.Println("User id:", id)
	return User{}, nil
}

func (r *userRepository) Get(username, password string) (User, error) {
	return User{Username: username, Password: password}, nil
}

func main() {

}
