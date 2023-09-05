package server

import (
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type Server struct {
	Router *mux.Router
	DB     *gorm.DB
	deps   Deps
}

func NewServer(router *mux.Router, db *gorm.DB) *Server {
	return &Server{
		Router: router,
		DB:     db,
		deps:   NewDeps(),
	}
}
