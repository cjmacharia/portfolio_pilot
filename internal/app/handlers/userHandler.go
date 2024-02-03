package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/cjmacharia/portfolioPilot/internal/app/service"
	"github.com/cjmacharia/portfolioPilot/internal/domain"
	"github.com/cjmacharia/portfolioPilot/internal/utils"
	"net/http"
)

type UserHandler struct {
	UserService service.UserService
}

func (uh UserHandler) UserSignUpHandler(w http.ResponseWriter, r *http.Request) {
	u := &domain.User{}
	err := json.NewDecoder(r.Body).Decode(u)
	if err != nil {
		http.Error(w, utils.ErrInvalidPayload.Error(), http.StatusBadRequest)
		return
	}
	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		http.Error(w, utils.ErrHashPassword.Error(), http.StatusInternalServerError)
		return
	}
	u.Password = hashedPassword
	userId, err := uh.UserService.UserSignUp(u)
	if err != nil {
		if err == utils.ErrDuplicateEmail {
			http.Error(w, utils.ErrDuplicateEmail.Error(), http.StatusConflict)
			return
		}
		fmt.Println(err, "err")
		http.Error(w, utils.ErrDatabaseRequest.Error(), http.StatusInternalServerError)
		return
	}
	token, err := utils.GenerateToken(&userId)
	if err != nil {
		http.Error(w, utils.ErrTokenGenerate.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(token)
}

func (uh UserHandler) LoginHandler(w http.ResponseWriter, r *http.Request) {
	u := &domain.User{}
	err := json.NewDecoder(r.Body).Decode(&u)
	savedUser := domain.User{}
	if err != nil {
		http.Error(w, utils.ErrInvalidPayload.Error(), http.StatusBadRequest)
		return
	}
	savedUser, err = uh.UserService.GetUserByIDService(u.Email)
	checkPass := utils.ComparePasswords(savedUser.Password, u.Password)
	if !checkPass {
		http.Error(w, utils.ErrAuthenticate.Error(), http.StatusUnauthorized)
		return
	}
	token, err := utils.GenerateToken(&savedUser.UserID)
	if err != nil {
		http.Error(w, utils.ErrTokenGenerate.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(token)
}
