package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type UserStore struct {
	db *sql.DB
}

type User struct {
	ID        int
	Username  string
	CreatedAt time.Time
}

type UserHandler struct {
	userStore UserService
}

type UserService interface {
	GetUserByID(id int) (*User, error)
	DeleteUserByID(id int) error
}

func main() {

	db, err := sql.Open("sqlite3", "./some.db")

	if err != nil {
		log.Fatal(err)
	}

	userStore := NewUserStore(db)
	userService := NewUserHandler(userStore)

	log.Println("Starting the DI example web server...")

	http.HandleFunc("/", userService.registerUser)
	log.Fatal(http.ListenAndServe(":3000", nil))
}

func NewUserStore(db *sql.DB) *UserStore {
	return &UserStore{
		db: db,
	}
}

func NewUserHandler(user UserService) *UserHandler {
	return &UserHandler{userStore: user}
}

func (us *UserStore) GetUserByID(id int) (*User, error) {
	// not impl
	return nil, nil
}

func (us *UserStore) DeleteUserByID(id int) error {
	// not impl
	return nil
}

func (uh *UserHandler) registerUser(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	fmt.Fprintf(w, "Hello from the DI example web server")
	// not impl
}
