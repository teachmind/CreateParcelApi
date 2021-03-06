package server

import (
	"CreateParcelApi/internal/app/model"
	"github.com/unrolled/render"
	"net/http"
)

const (
	codeInvalidErr        = "INVALID"
	codeNotFoundErr       = "NOT FOUND"
	codeInternalServerErr = "SERVER_ERROR"
)

var renderer = render.New(render.Options{})

func SuccessResponse(w http.ResponseWriter, httpStatusCode int, data interface{}) {
	renderer.JSON(w, httpStatusCode, model.GenericResponse{
		Success: true,
		Errors:  nil,
		Data:    data,
	})
}

func ErrUnprocessableEntityResponse(w http.ResponseWriter, title string, err error) {
	errorResponse(w, http.StatusUnprocessableEntity, codeInvalidErr, title, err)
}

func ErrInvalidEntityResponse(w http.ResponseWriter, title string, err error) {
	errorResponse(w, http.StatusBadRequest, codeInvalidErr, title, err)
}

func ErrNotFoundResponse(w http.ResponseWriter, title string, err error) {
	errorResponse(w, http.StatusNotFound, codeNotFoundErr, title, err)
}

func ErrInternalServerResponse(w http.ResponseWriter, title string, err error) {
	errorResponse(w, http.StatusInternalServerError, codeInternalServerErr, title, err)
}

func errorResponse(w http.ResponseWriter, httpStatusCode int, code string, title string, err error) {
	renderer.JSON(w, httpStatusCode, model.GenericResponse{
		Success: false,
		Data:    nil,
		Errors: []model.ErrorDetailsResponse{
			{
				Code:     code,
				Message:  err.Error(),
				Title:    title,
				Severity: "error",
			},
		},
	})
}
