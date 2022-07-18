package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/zkdltidchan/dao-manager-api-server/api/auth"
	"github.com/zkdltidchan/dao-manager-api-server/api/models"
	"github.com/zkdltidchan/dao-manager-api-server/api/responses"
	"github.com/zkdltidchan/dao-manager-api-server/api/utils/formaterror"
	"golang.org/x/crypto/bcrypt"
)

func (server *Server) Login(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	managerUser := models.ManagerUser{}
	err = json.Unmarshal(body, &managerUser)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	managerUser.Prepare()
	err = managerUser.Validate("login")
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	token, err := server.SignIn(managerUser.Email, managerUser.Password)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusUnprocessableEntity, formattedError)
		return
	}

	responses.JSON(w, http.StatusOK, token)
}

func (server *Server) SignIn(email, password string) (string, error) {

	var err error

	managerUser := models.ManagerUser{}

	err = server.DB.Debug().Model(models.ManagerUser{}).Where("email = ?", email).Take(&managerUser).Error
	if err != nil {
		return "", err
	}
	err = models.VerifyPassword(managerUser.Password, password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}
	return auth.CreateToken(managerUser.ID)
}
