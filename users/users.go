package users

import (
	"database/sql"
	"fmt"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"

	"ep-backend/database"
)

type User struct {
	Username        string `json:"username"`
	Password        string `json:"password"`
	ProfilePhotoURL string `json:"profile_photo_url"`
}

type Credential struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// HashPassword hashes the given password using bcrypt
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// CheckPasswordHash compares the provided password with the hashed password
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// Creates a new user in the database. The password is hashed before being stored.
func CreateUser(user User) error {
	db, err := database.ConnectDB()
	if err != nil {
		return fmt.Errorf("failed to connect to the database: %v", err)
	}

	hashedPassword, err := HashPassword(user.Password)
	if err != nil {
		return fmt.Errorf("failed to hash password: %v", err)
	}
	user.Password = string(hashedPassword)

	_, err = db.Exec("INSERT INTO engineerpro.users (username, password, profile_photo_url) VALUES (?, ?, ?)",
		user.Username, user.Password, user.ProfilePhotoURL)
	if err != nil {
		// Check if the user already exists
		if strings.Contains(err.Error(), "1062") {
			return fmt.Errorf("username already exists")
		}
		// Handle other errors
		return fmt.Errorf("error creating user: %v", err)
	}

	return err
}

func AuthenticateUser(credential Credential) error {
	db, err := database.ConnectDB()
	if err != nil {
		return fmt.Errorf("failed to connect to the database: %v", err)
	}

	var user User
	row := db.QueryRow("SELECT username, password, profile_photo_url FROM engineerpro.users WHERE username = ?", credential.Username)
	err = row.Scan(&user.Username, &user.Password, &user.ProfilePhotoURL)
	if err != nil {
		// Handle the case where the user is not found
		if err == sql.ErrNoRows {
			return fmt.Errorf("user not found")
		}
		// Handle other errors
		return fmt.Errorf("error authenticating user: %v", err)
	}

	if !CheckPasswordHash(credential.Password, user.Password) {
		return fmt.Errorf("wrong password")
	}

	return nil
}
