package server

import (
	"CreateParcelApi/internal/app/model"
	"encoding/json"
	"github.com/rs/zerolog/log"
	"net/http"
)

func (s *server) createParcel(w http.ResponseWriter, r *http.Request) {
	var data model.Parcel

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		ErrUnprocessableEntityResponse(w, "Decode Error", err)
		return
	}

	if err := data.ValidateParcelInput(); err != nil {
		ErrInvalidEntityResponse(w, "Invalid Input", err)
		return
	}

	message, err := json.Marshal(data)
	if err != nil {
		log.Error().Err(err).Msg("json marshal failed")
		ErrUnprocessableEntityResponse(w, "bad request", err)
		return
	}

	err = s.publisherService.Push(message)
	if err != nil {
		log.Error().Err(err).Msg("failed to publish message")
		ErrInternalServerResponse(w, "Failed to publish message", err)
		return
	}

	SuccessResponse(w, http.StatusOK, "your parcel is being created withing some moments")
}
