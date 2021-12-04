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

	parcel := model.Parcel{
		ID: 				data.ID,
		UserID:             data.UserID,
		CarrierID:      	data.CarrierID,
		Status: 			data.Status,
		SourceAddress:      data.SourceAddress,
		DestinationAddress: data.DestinationAddress,
		SourceTime:         data.SourceTime,
		Price:              data.Price,
		CarrierFee:         data.CarrierFee,
		CompanyFee:         data.CompanyFee,
		CreatedAt:          data.CreatedAt,
		UpdatedAt:          data.UpdatedAt,
	}

	message, err := json.Marshal(&parcel)
	if err != nil {
		log.Error().Err(err).Msg("json marshal failed")
		ErrUnprocessableEntityResponse(w, "bad request", err)
		return
	}

	SuccessResponse(w, http.StatusCreated, message)
}
