package server

import (
	"net/http"

	"digiors/helpers"
)

type serveHandle func(http.ResponseWriter, *http.Request) error

func (fn serveHandle) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := fn(w, r)
	if err != nil {
		helpers.WriteError(w, err)
	}
}

func (s *Server) RegisterHandlers() {
	webPath := "/"

	s.Router.Host("example.com").Schemes("https", "http")

	s.Router.Use(CORS)

	publicAPIRouter := s.Router.PathPrefix(webPath + "api").Subrouter()
	publicAPIRouter.Handle("/stations", serveHandle(s.GetAllStations)).Methods("GET").Name("getAllStations")
	publicAPIRouter.Handle("/stations", serveHandle(s.AddStation)).Methods("POST").Name("addStation")
	publicAPIRouter.Handle("/stations/{name}", serveHandle(s.GetStation)).Methods("GET").Name("getStation")
	publicAPIRouter.Handle("/stations/{name}", serveHandle(s.UpdateStation)).Methods("PUT").Name("updateStation")

	publicAPIRouter.Handle("/stations/{name}/registrations", serveHandle(s.GetAllStationEventRegistrations)).Methods("GET").Name("getAllStationEventRegistrations")
	publicAPIRouter.Handle("/stations/{name}/registrations", serveHandle(s.AddStationEventRegistration)).Methods("POST").Name("addStationEventRegistration")
	publicAPIRouter.Handle("/stations/{name}/registrations/{registration}/users", serveHandle(s.AddStationEventUser)).Methods("POST").Name("addStationEventUser")
}
