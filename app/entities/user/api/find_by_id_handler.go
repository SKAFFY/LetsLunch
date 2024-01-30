package api

import (
	domain "cmd/app/entities/user"
	usecases "cmd/app/entities/user/usecases"
	"cmd/pkg/errors"
	"encoding/json"
	"github.com/gofrs/uuid/v5"
	"github.com/gorilla/mux"
	"net/http"
)

type JsonFindUserByIdRespond struct {
	User *domain.User `json:"data"`
}

type FindUserByIdHandler struct {
	useCase *usecases.FindUserByIdUseCase
}

func NewFindUserByIdHandler(useCase *usecases.FindUserByIdUseCase) *FindUserByIdHandler {
	return &FindUserByIdHandler{useCase: useCase}
}

func (handler *FindUserByIdHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	id, ok := mux.Vars(request)["id"]
	if !ok {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	uuidID, err := uuid.FromString(id)
	if err != nil {
		customError := errors.NewError(err)
		marshledError, _ := json.Marshal(customError)

		writer.WriteHeader(http.StatusBadRequest)
		writer.Write(marshledError)
		return
	}

	users, err := handler.useCase.Handle(request.Context(), uuidID)
	if err != nil {
		customError := errors.NewError(err)
		marshaledError, _ := json.Marshal(customError)

		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write(marshaledError)
		return
	}

	response := JsonFindUserByIdRespond{User: users.User}

	marshaledResponse, err := json.Marshal(response)
	if err != nil {
		customError := errors.NewError(err)
		marshaledError, _ := json.Marshal(customError)

		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write(marshaledError)
		return
	}

	writer.WriteHeader(http.StatusOK)
	writer.Write(marshaledResponse)
}
