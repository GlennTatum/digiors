package server

import (
	"digiors/helpers"
	"digiors/models"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gosimple/slug"
)

func (s *Server) GetAllStations(w http.ResponseWriter, r *http.Request) error {
	// TODO Add pagination parameter
	var stations []models.Station
	err := s.DB.
		Limit(10).
		Preload("Registrations.Users").
		Find(&stations).Error
	if err != nil {
		return err
	}

	b, err := json.Marshal(stations)
	if err != nil {
		return err
	}
	w.Header().Set("content-type", "application/json")
	w.Write(b)

	return nil
}

func (s *Server) GetStation(w http.ResponseWriter, r *http.Request) error {
	v := mux.Vars(r)

	var station models.Station
	err := s.DB.
		Preload("Registrations.Users").
		First(&station, "slug = ?", v["name"]).Error
	if err != nil {
		return err
	}

	b, err := json.Marshal(station)
	if err != nil {
		return err
	}
	w.Header().Set("content-type", "application/json")
	w.Write(b)

	return nil
}

func (s *Server) AddStation(w http.ResponseWriter, r *http.Request) error {

	var station models.Station

	err := helpers.Decode(r, &station)
	if err != nil {
		return err
	}

	station.Slug = slug.Make(station.Name)

	err = s.DB.Create(&station).Error
	if err != nil {
		return err
	}

	u, _ := s.Router.Get("getStation").URL("name", fmt.Sprint(station.Name))
	helpers.WriteCreated(w, u)

	return nil
}

func (s *Server) UpdateStation(w http.ResponseWriter, r *http.Request) error {

	v := mux.Vars(r)

	var station models.Station

	err := s.DB.First(&station, "slug = ?", v["name"]).Error
	if err != nil {
		return err
	}

	err = helpers.Decode(r, &station)
	if err != nil {
		return err
	}

	err = s.DB.Save(&station).Error
	if err != nil {
		return err
	}

	return nil
}

func (s *Server) GetAllStationEventRegistrations(w http.ResponseWriter, r *http.Request) error {
	var registrations []models.Registration

	s.DB.Limit(10).Preload("Users").Find(&registrations)

	b, err := json.Marshal(registrations)
	if err != nil {
		return err
	}

	w.Write(b)

	return nil
}

func (s *Server) AddStationEventRegistration(w http.ResponseWriter, r *http.Request) error {

	v := mux.Vars(r)

	var registration models.Registration
	err := helpers.Decode(r, &registration)
	if err != nil {
		return err
	}
	registration.DateBegin_fmt = helpers.Mil_to_Civ(registration.DateBegin)
	registration.DateEnd_fmt = helpers.Mil_to_Civ(registration.DateEnd)

	var station models.Station
	err = s.DB.First(&station, "slug = ?", v["name"]).Error
	if err != nil {
		return err
	}

	s.DB.Model(&station).
		Association("Registrations").
		Append(&registration)

	w.WriteHeader(http.StatusCreated)

	return nil
}

func (s *Server) AddStationEventUser(w http.ResponseWriter, r *http.Request) error {

	v := mux.Vars(r)

	var registration models.Registration
	err := s.DB.First(&registration, "id = ?", v["registration"]).Error
	if err != nil {
		return err
	}

	var anon models.AnonUser
	err = helpers.Decode(r, &anon)
	if err != nil {
		return err
	}

	s.DB.Model(&registration).
		Association("Users").
		Append([]models.AnonUser{anon})

	w.WriteHeader(http.StatusCreated)

	return nil
}
