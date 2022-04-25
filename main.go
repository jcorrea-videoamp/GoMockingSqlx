package main

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/jmoiron/sqlx"
)

type User struct {
	FirstName string `db:"FIRST_NAME"`
	LastName  string `db:"LAST_NAME"`
	Username  string `db:"USERNAME"`
	Password  string `db:"PASSWORD"`
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
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Minute)
	defer cancel()
	queryUser := []User{}
	query := "SELECT * FROM table WHERE id=" + strconv.Itoa(id)
	err := r.db.SelectContext(ctx, &queryUser, query)
	if err != nil {
		fmt.Println("Ohhh", err)
		return User{}, err
	}
	return queryUser[0], nil
}

func (r *userRepository) Get(username, password string) (User, error) {
	return User{Username: username, Password: password}, nil
}

func main() {

}

