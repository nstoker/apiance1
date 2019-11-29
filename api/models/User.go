package models

import (
	"database/sql"
	"errors"
	"fmt"
	"html"
	"log"
	"strings"

	"github.com/badoux/checkmail"
	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
)

// User structure
type User struct {
	ID        uint32       `json:"id"`
	Name      string       `json:"name"`
	Email     string       `json:"email"` // Indexed as lowercase
	Password  string       `json:"password"`
	CreatedAt sql.NullTime `json:"created_at"`
	UpdatedAt sql.NullTime `json:"updated_at"`
}

// Hash user password
func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

// VerifyPassword verifies password
func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

// BeforeSave does before save
func (u *User) BeforeSave() error {
	hashedPassword, err := Hash(u.Password)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

// Prepare prepares
func (u *User) Prepare() {
	u.ID = 0
	u.Name = html.EscapeString(strings.TrimSpace(u.Name))
	u.Email = html.EscapeString(strings.TrimSpace(u.Email))
}

// Validate validates record
func (u *User) Validate(action string) error {
	switch strings.ToLower(action) {
	case "update":
		if u.Password == "" {
			return errors.New("Required Password")
		}
		if u.Email == "" {
			return errors.New("Required Email")
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("Invalid Email")
		}

		return nil
	case "login":
		if u.Password == "" {
			return errors.New("Required Password")
		}
		if u.Email == "" {
			return errors.New("Required Email")
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("Invalid Email")
		}
		return nil

	default:
		if u.Name == "" {
			return errors.New("Required Name")
		}
		if u.Password == "" {
			return errors.New("Required Password")
		}
		if u.Email == "" {
			return errors.New("Required Email")
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("Invalid Email")
		}
		return nil
	}
}

// CreateUser saves user
func (u *User) CreateUser(db *sqlx.DB) (*User, error) {
	hash, err := Hash(u.Password)
	if err != nil {
		return nil, err
	}
	sqlStatement := "INSERT INTO users (email, name, password) VALUES ($1, $2, $3) RETURNING id;"
	err = db.QueryRow(sqlStatement, u.Email, u.Name, hash).Scan(&u.ID)
	if err != nil {
		return nil, fmt.Errorf("CreateUser: %w", err)
	}

	return u, err
}

// FindAllUsers finds all users
func (u *User) FindAllUsers(db *sqlx.DB) (*[]User, error) {
	users := []User{}

	rows, err := db.Query("SELECT id, email, name, created_at, updated_at FROM users ORDER BY id;")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		user := User{}
		err = rows.Scan(&user.ID, &user.Email, &user.Name, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}
	return &users, nil
}

// FindUserByEmail finds a user by email
func (u *User) FindUserByEmail(db *sqlx.DB, email string) (*User, error) {
	sqlStatement := `SELECT id, name, email, created_at, updated_at FROM users WHERE email=$1`
	row := db.QueryRow(sqlStatement, email)
	err := row.Scan(&u.ID, &u.Name, &u.Email, &u.CreatedAt, &u.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return u, nil
}

// FindUserByID finds a user by ... id
func (u *User) FindUserByID(db *sqlx.DB, uid uint32) (*User, error) {
	sqlStatement := `SELECT id, name, email, created_at, updated_at FROM users WHERE id=$1`

	row := db.QueryRow(sqlStatement, uid)

	err := row.Scan(&u.ID, &u.Name, &u.Email, &u.CreatedAt, &u.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return u, nil
}

// UpdateAUser updates a user
func (u *User) UpdateAUser(db *sqlx.DB, uid uint32) (*User, error) {

	// To hash the password
	err := u.BeforeSave()
	if err != nil {
		log.Fatal(err)
	}
	// See https://www.calhoun.io/updating-and-deleting-postgresql-records-using-gos-sql-package/
	sqlStatement := `UPDATE users SET email=$2, name=$3 WHERE id=$4;`
	return u, nil
}

// DeleteAUser deletes a user. Returning the number of rows deleted on success
func (u *User) DeleteAUser(db *sqlx.DB, ID int64) (int64, error) {
	sqlStatement := `DELETE FROM users WHERE id=$1`

	res, err := db.Exec(sqlStatement, ID)
	if err != nil {
		return 0, err
	}

	count, err := res.RowsAffected()
	if err != nil {
		log.Printf("DeleteAUser error: %s", err)
		return 0, err
	}

	return count, nil
}
