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
	Email     string       `json:"email"`
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
	// u.CreatedAt = time.Now() // TODO: Fix me
	// u.UpdatedAt = time.Now()
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

// SaveUser saves user
func (u *User) SaveUser(db *sqlx.DB) (*User, error) {

	var err error
	err = fmt.Errorf("Not Implement")
	// err = db.Debug().Create(&u).Error
	// if err != nil {
	// 	return &User{}, err
	// }
	return u, err
}

// FindAllUsers finds all users
func (u *User) FindAllUsers(db *sqlx.DB) (*[]User, error) {
	var err error
	users := []User{}
	err = fmt.Errorf("Not Implemented")
	// err = db.Debug().Model(&User{}).Limit(100).Find(&users).Error
	// if err != nil {
	// 	return &[]User{}, err
	// }
	return &users, err
}

// FindUserByID finds a user by ... id
func (u *User) FindUserByID(db *sqlx.DB, uid uint32) (*User, error) {
	var err error
	err = fmt.Errorf("Not Implemented")
	// err = db.Debug().Model(User{}).Where("id = ?", uid).Take(&u).Error
	// if err != nil {
	// 	return &User{}, err
	// }
	// if gorm.IsRecordNotFoundError(err) {
	// 	return &User{}, errors.New("User Not Found")
	// }
	return u, err
}

// UpdateAUser updates a user
func (u *User) UpdateAUser(db *sqlx.DB, uid uint32) (*User, error) {

	// To hash the password
	err := u.BeforeSave()
	if err != nil {
		log.Fatal(err)
	}
	// db = db.Debug().Model(&User{}).Where("id = ?", uid).Take(&User{}).UpdateColumns(
	// 	map[string]interface{}{
	// 		"password":  u.Password,
	// 		"Name":      u.Name,
	// 		"email":     u.Email,
	// 		"update_at": time.Now(),
	// 	},
	// )
	// if db.Error != nil {
	// 	return &User{}, db.Error
	// }
	// // This is the display the updated user
	// // err = db.Debug().Model(&User{}).Where("id = ?", uid).Take(&u).Error
	// if err != nil {
	// 	return &User{}, err
	// }
	return u, fmt.Errorf("Not Implemented")
}

// DeleteAUser deletes a user
func (u *User) DeleteAUser(db *sqlx.DB, uid uint32) (int64, error) {

	// db = db.Debug().Model(&User{}).Where("id = ?", uid).Take(&User{}).Delete(&User{})

	// if db.Error != nil {
	// 	return 0, db.Error
	// }
	return 0, fmt.Errorf("Not Implemented")
}