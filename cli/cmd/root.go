package cmd

import (
	"digiors/models"
	"digiors/server"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/spf13/cobra"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func init() {
	rootCmd.AddCommand(serverCmd)
}

var rootCmd = &cobra.Command{
	Use:   "digiors",
	Short: "The digital platform for Oyster Research Station Management.",
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Help()
		os.Exit(0)
	},
}

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Run in server mode",
	Run: func(cmd *cobra.Command, args []string) {
		runService()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func runService() {

	server := NewService()

	var router http.Handler = server.Router

	router = handlers.ProxyHeaders(router)
	http.Handle("/", router)

	err := http.ListenAndServe("127.0.0.1:8000", router)
	if err != nil {
		log.Panic(err)
	}
}

func NewService() *server.Server {

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Panic("failed to connect to database")
	}

	models.Migrate(db)

	r := mux.NewRouter()
	s := server.NewServer(r, db)
	s.RegisterHandlers()

	return s
}
