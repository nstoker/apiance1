package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/nstoker/apiance1/api/models"
	"github.com/nstoker/apiance1/api/responses"
	"github.com/nstoker/apiance1/api/utils/formaterror"
)

// Login logs a user in
func (server *Server) Login(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	user := models.User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	// if err == nil {
	// 	responses.ERROR(w, http.StatusTeapot, fmt.Errorf("User %+v", user))
	// 	return
	// }

	user.Prepare()
	err = user.Validate("login")
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	token, err := server.SignIn(user.Email, user.Password)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusUnprocessableEntity, formattedError)
		return
	}
	responses.JSON(w, http.StatusOK, token)
}

// SignIn signs in a user
func (server *Server) SignIn(email, password string) (string, error) {

	var err error
	err = fmt.Errorf("Server:SignIn Not Implemented")
	// user := models.User{}

	// err = server.DB.Debug().Model(models.User{}).Where("email = ?", email).Take(&user).Error
	// if err != nil {
	// 	return "", err
	// }
	// err = models.VerifyPassword(user.Password, password)
	// if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
	// 	return "", err
	// }
	// return auth.CreateToken(user.ID)

	return "", err
}
