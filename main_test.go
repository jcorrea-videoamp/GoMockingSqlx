package main

import (
	"fmt"
	"regexp"
	"testing"

	sqlxmock "github.com/zhashkevych/go-sqlxmock"
)

func TestUserRepository_Insert(t *testing.T) {
	db, mock, err := sqlxmock.Newx()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	s := NewUserRepository(db)

	tests := []struct {
		name    string
		s       UserRepository
		user    User
		mock    func()
		want    int
		wantErr bool
	}{
		{
			//When everything works as expected
			name: "OK",
			s:    s,
			user: User{
				FirstName: "first_name",
				LastName:  "last_name",
				Username:  "username",
				Password:  "password",
			},
			mock: func() {
				rows := sqlxmock.NewRows([]string{"id"}).AddRow(1)
				mock.ExpectQuery("INSERT INTO users").WithArgs("first_name", "last_name", "username", "password").WillReturnRows(rows)
			},
			want: 1,
		},
		{
			name: "Empty Fields",
			s:    s,
			user: User{
				FirstName: "",
				LastName:  "",
				Username:  "username",
				Password:  "password",
			},
			mock: func() {
				rows := sqlxmock.NewRows([]string{"id"})
				mock.ExpectQuery("INSERT INTO users").WithArgs("first_name", "last_name", "username", "password").WillReturnRows(rows)
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			got, err := tt.s.Insert(tt.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error new = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err == nil && got != tt.want {
				t.Errorf("Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSelectContext(t *testing.T) {
	db, mock, err := sqlxmock.Newx()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	s := NewUserRepository(db)
	rows := sqlxmock.NewRows([]string{"FIRST_NAME", "LAST_NAME", "USERNAME", "PASSWORD"}).AddRow("John", "Doe", "johndoe", "qwerty1234")
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM table WHERE id=15`)).WillReturnRows(rows)
	u, err := s.GetById(15)
	if err != nil {
		fmt.Println("Errrorr!", err)
	}
	fmt.Printf("users: %#v", u)
}
