package controllers

import (
	// "log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
	"github.com/zkdltidchan/dao-manager-api-server/api/models"
	"github.com/zkdltidchan/dao-manager-api-server/api/responses"
)

var decoder = schema.NewDecoder()

func (server *Server) GetUsers(w http.ResponseWriter, r *http.Request) {
	var userParameter models.UserParameter
	decoder.Decode(&userParameter, r.URL.Query())
	// // for handle error query,
	// err := decoder.Decode(&userParameter, r.URL.Query())
	// if err != nil {
	// 	responses.ERROR(w, http.StatusInternalServerError, err)
	// 	return
	// } else {
	// 	log.Println("GET parameters Page Index : ", userParameter.PageIndex)
	// 	log.Println("GET parameters Size: ", userParameter.Size)
	// }
	user := models.User{}
	users, err := user.FindAllUsers(server.DB, userParameter)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, users)
}

func (server *Server) GetUser(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	uid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	user := models.User{}
	userGotten, err := user.FindUserByID(server.DB, uint32(uid))
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	responses.JSON(w, http.StatusOK, userGotten)
}
