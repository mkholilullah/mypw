package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/roamercodes/mypw/domain"
	"github.com/roamercodes/mypw/usecase"
)

type UserHandler struct {
	UseCase *usecase.UserUsecase
}

func NewUserHandler(uc *usecase.UserUsecase) *UserHandler {
	return &UserHandler{
		UseCase: uc,
	}
}

func (h *UserHandler) GetUserById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) // get url variable
	idParam := vars["id"]

	id, err := strconv.Atoi(idParam) 
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	user, err := h.UseCase.GetUserById(int(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user domain.User
	json.NewDecoder(r.Body).Decode(&user)

	err := h.UseCase.CreateUser(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}