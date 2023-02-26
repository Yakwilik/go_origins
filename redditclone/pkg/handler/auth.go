package handler

import (
	"encoding/json"
	"gitlab.com/vk-go/lectures-2022-2/06_databases/99_hw/redditclone/structs"
	"net/http"
)

type signInInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type interfaceMap map[string]interface{}

func (h *Handler) register(writer http.ResponseWriter, request *http.Request) {
	user := signInInput{}
	err := json.NewDecoder(request.Body).Decode(&user)
	if err != nil {
		JSONResponse(writer, http.StatusBadRequest, interfaceMap{"error": err})
		return
	}
	newUser := structs.User{
		ID:       0,
		Username: user.Username,
		Password: user.Password,
	}
	token, err := h.services.CreateUser(newUser)
	if err != nil {
		resp := structs.NewErrorResponse()
		Err := structs.NewErrorMap()
		Err["location"] = "body"
		Err["param"] = "username"
		Err["value"] = newUser.Username
		Err["msg"] = Err["value"] + " already  exists"
		Err["err"] = err.Error()
		resp.Err = append(resp.Err, Err)
		JSONResponse(writer, http.StatusUnprocessableEntity, resp)
		return
	}
	JSONResponse(writer, http.StatusOK, interfaceMap{"token": token})
}

func (h *Handler) login(writer http.ResponseWriter, request *http.Request) {
	user := signInInput{}
	err := json.NewDecoder(request.Body).Decode(&user)
	if err != nil {
		JSONResponse(writer, http.StatusBadRequest, interfaceMap{"error": err})
		return
	}
	token, err := h.services.Authorization.GenerateToken(user.Username, user.Password)
	if err != nil {
		JSONResponse(writer, http.StatusBadRequest, interfaceMap{"message": "invalid credentials"})
		return
	}
	JSONResponse(writer, http.StatusOK, interfaceMap{"token": token})
}
