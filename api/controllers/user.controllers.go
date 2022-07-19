package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/zkdltidchan/dao-manager-api-server/api/models"
	"github.com/zkdltidchan/dao-manager-api-server/api/responses"
)

func (server *Server) GetUsers(w http.ResponseWriter, r *http.Request) {
	// query := r.URL.Query()
	// page_index, present := query["page_index"] //filters=["color", "price", "brand"]
	// if !present || len(page_index) == 0 {
	// 	fmt.Println("filters not present")
	// }
	// size, present := query["size"] //filters=["color", "price", "brand"]
	// if !present || len(size) == 0 {
	// 	fmt.Println("filters not present")
	// }
	userParameter := models.UserParameter{}
	userParameter.PageIndex = 2
	userParameter.Size = 2

	user := models.User{}

	users, err := user.FindAllUsers(server.DB, userParameter)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	// TODO, paramer size, index, total...
	// featchListResp := &responses.FeatchListResponse{
	// 	Total:     userParameter.Count,
	// 	Size:      len(*users),
	// 	PageIndex: 1,
	// 	Data:      &users,
	// }
	fmt.Printf("%v", users)
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
