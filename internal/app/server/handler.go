package server

import (
	"CreateParcelApi/internal/app/model"
	"encoding/json"
	"errors"
	"github.com/rs/zerolog/log"
	"google.golang.org/protobuf/types/known/timestamppb"
	"net/http"
	"time"
)

func (s *server) newParcel(w http.ResponseWriter, r *http.Request) {
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
		log.Error().Err(err).Msg("proto marshal failed")
		ErrUnprocessableEntityResponse(w, "bad request", err)
		return
	}

	if err != nil {
		if errors.Is(err, model.ErrInvalid) {
			ErrInvalidEntityResponse(w, "invalid parcel", err)
			return
		}
		log.Error().Err(err).Msgf("[parcel] failed to create parcel Error: %v", err)
		ErrInternalServerResponse(w, "failed to create parcel", err)
		return
	}

	SuccessResponse(w, http.StatusCreated, parcel)
}
