package controller

import (
	"encoding/json"
	"net/http"

	usecase "github.com/reangeline/micro_saas/internal/domain/contract/usecase/user"
	"github.com/reangeline/micro_saas/internal/dto"
)

type Error struct {
	Message string `json:"message"`
}

type UserController struct {
	createUserUseCase usecase.CreateUserUseCaseInterface
}

func NewUserController(
	createUserUseCase usecase.CreateUserUseCaseInterface,
) *UserController {
	return &UserController{
		createUserUseCase: createUserUseCase,
	}
}

func (u *UserController) CreateUserRest(w http.ResponseWriter, r *http.Request) {
	var user dto.UserInput
	err := json.NewDecoder(r.Body).Decode(&user)
	defer r.Body.Close()

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		error := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}

	ctx := r.Context()
	err = u.createUserUseCase.Execute(ctx, &user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
