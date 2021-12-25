package server

import (
	"CreateParcelApi/internal/app/service"
	"context"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
	"net/http"
)

type server struct {
	listenAddress        string
	http                 *http.Server
	publisherService 	 service.PublisherService
}

func NewServer(port string, publisherSvc service.PublisherService) *server {
	s := &server{
		listenAddress: port,
		publisherService: publisherSvc,
	}
	s.http = &http.Server{
		Addr: port,
		Handler: s.route(),
	}
	return s
}

func (s *server) route() *mux.Router {
	r := mux.NewRouter()
	//apiRoute := r.PathPrefix("/api/v1").Subrouter()
	r.Methods(http.MethodGet).Path("/ping").HandlerFunc(s.pingHandler)
	r.Methods(http.MethodPost).Path("/api/v1/parcel").HandlerFunc(s.createParcel)
	r.Methods(http.MethodPost).Path("/api/v1/parcel/{id}/accept").HandlerFunc(s.parcelCarrierAccept)
	
	return r
}

func (s *server) Run() error {
	log.Info().Msgf("start listen server in %s", s.listenAddress)

	if err := s.http.ListenAndServe(); err != nil {
		if err != http.ErrServerClosed {
			log.Error().Err(err).Msg("unexpected error while running server")
			return s.Shutdown()
		}
	}
	return nil
}

func (s *server) Shutdown() error {
	log.Info().Msg("shutting down server")
	if err := s.http.Shutdown(context.Background()); err != nil {
		if err != http.ErrServerClosed {
			return err
		}
	}
	return nil
}

func (s *server) pingHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"success": "successfully pinged"}`))
}
