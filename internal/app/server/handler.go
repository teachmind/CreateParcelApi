package server

import (
	"CreateParcelApi/internal/app/model"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
	"net/http"
	"strconv"
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

func (s *server) parcelCarrierAccept(w http.ResponseWriter, r *http.Request) {
	var data model.CarrierRequest
	vars := mux.Vars(r)

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		ErrUnprocessableEntityResponse(w, "Decode Error", err)
		return
	}
	parcelID, err := strconv.Atoi(vars["id"])
	if err != nil {
		ErrInvalidEntityResponse(w, "Invalid Parcel ID", err)
		return
	}
	data.ParcelID = parcelID

	// validating input credentials for parcel request
	if err := data.ValidateCarrierId(); err != nil {
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

	SuccessResponse(w, http.StatusOK, "your parcel accept request is being created withing some moments")
}
