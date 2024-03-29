package modeltests

import (
	"fmt"
	"log"
	"testing"

	"github.com/nstoker/apiance1/api/models"
	"gopkg.in/go-playground/assert.v1"
)

func TestFindAllUsers(t *testing.T) {

	err := refreshTable("users")
	if err != nil {
		t.Fatal(err)
	}

	seededUsers, err := seedUsers()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("seededUsers: %+v", seededUsers)

	users, err := userInstance.FindAllUsers(server.DB)
	if err != nil {
		t.Errorf("TestUpdateUser error getting the users: %v\n", err)
		return
	}
	assert.Equal(t, len(*users), 2)
}

func TestSaveUser(t *testing.T) {

	err := refreshTable("users")
	if err != nil {
		log.Fatal(err)
	}
	newUser := models.User{
		ID:       1,
		Email:    fmt.Sprintf("%s@example.com", models.GenKsuid()),
		Name:     "test",
		Password: "password",
	}
	savedUser, err := newUser.CreateUser(server.DB)
	if err != nil {
		t.Errorf("TestSaveError error getting the users: %v\n", err)
		return
	}
	assert.Equal(t, newUser.ID, savedUser.ID)
	assert.Equal(t, newUser.Email, savedUser.Email)
	assert.Equal(t, newUser.Name, savedUser.Name)
}

func TestGetUserByID(t *testing.T) {

	err := refreshTable("users")
	if err != nil {
		log.Fatal(err)
	}

	user, err := seedOneUser()
	if err != nil {
		log.Fatalf("cannot seed users table: %v", err)
	}
	foundUser, err := userInstance.FindUserByID(server.DB, user.ID)
	if err != nil {
		t.Errorf("TestGetUserById error getting one user: %v\n", err)
		return
	}
	assert.Equal(t, foundUser.ID, user.ID)
	assert.Equal(t, foundUser.Email, user.Email)
	assert.Equal(t, foundUser.Name, user.Name)
}

func TestUpdateAUser(t *testing.T) {

	err := refreshTable("users")
	if err != nil {
		log.Fatal(err)
	}

	user, err := seedOneUser()
	if err != nil {
		log.Fatalf("Cannot seed user: %v\n", err)
	}

	userUpdate := models.User{
		ID:       1,
		Name:     "modiUpdate",
		Email:    fmt.Sprintf("%s@example.com", models.GenKsuid()),
		Password: "password",
	}
	updatedUser, err := userUpdate.UpdateAUser(server.DB, user.ID)
	if err != nil {
		t.Errorf("TestUpdateAUser error updating the user: %v\n", err)
		return
	}
	assert.Equal(t, updatedUser.ID, userUpdate.ID)
	assert.Equal(t, updatedUser.Email, userUpdate.Email)
	assert.Equal(t, updatedUser.Name, userUpdate.Name)
}

func TestDeleteAUser(t *testing.T) {

	err := refreshTable("users")
	if err != nil {
		log.Fatal(err)
	}

	user, err := seedOneUser()

	if err != nil {
		log.Fatalf("Cannot seed user: %v\n", err)
	}

	isDeleted, err := userInstance.DeleteAUser(server.DB, int64(user.ID))
	if err != nil {
		t.Errorf("TestDeleteUser error updating the user: %v\n", err)
		return
	}

	if isDeleted != 1 {
		t.Errorf("expected one record to be deleted, got %d", isDeleted)
	}
}
