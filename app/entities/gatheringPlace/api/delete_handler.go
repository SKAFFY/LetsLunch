package api

import (
	"cmd/app/entities/gatheringPlace/usecases"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/gofrs/uuid/v5"
	"net/http"
)

type DeleteGatheringPlaceHandler struct {
	useCase *usecases.DeleteGatheringPlaceUseCase
}

func NewDeleteUserHandler(useCase *usecases.DeleteGatheringPlaceUseCase) *DeleteGatheringPlaceHandler {
	return &DeleteGatheringPlaceHandler{useCase: useCase}
}

func (handler *DeleteGatheringPlaceHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	id := chi.URLParam(request, "placeID")

	uuidID, err := uuid.FromString(id)
	if err != nil {
		marshaledError, _ := json.Marshal(err)

		writer.WriteHeader(http.StatusBadRequest)
		writer.Write(marshaledError)
		return
	}

	err = handler.useCase.Handle(request.Context(), uuidID)
	if err != nil {
		marshaledError, _ := json.Marshal(err)

		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write(marshaledError)
		return
	}

	writer.WriteHeader(http.StatusOK)
}
