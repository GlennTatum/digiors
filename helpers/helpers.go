package helpers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

func WriteError(w http.ResponseWriter, err error) {
	// check all types of errors db, not found etc
	http.Error(w, "error", 404)
}

func WriteCreated(w http.ResponseWriter, uri *url.URL) {
	w.Header().Set("Location", BaseURI(uri))
	w.WriteHeader(http.StatusCreated)
}

func BaseURI(resource *url.URL) string {
	return URL().String() + resource.String()
}

func URL() *url.URL {
	r := mux.Route{}
	// TODO Fetch host url from ENV setting
	u, _ := r.Schemes("https").Host("example.com").URL()
	return u
}

func Decode(r *http.Request, v interface{}) error {
	err := json.NewDecoder(r.Body).Decode(&v)
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}

func Mil_to_Civ(date string) string {

	t, _ := time.Parse(time.RFC3339, date)
	t = t.Local()

	var suffix string

	year, month, day := t.Date()
	h := t.Hour()
	m := t.Minute()

	if h < 12 {
		suffix = "AM"
	}
	if h == 12 {
		suffix = "PM"
	}
	if h > 12 {
		suffix = "PM"
		h = h - 12
	}

	if m == 0 {
		s := fmt.Sprintf("%v %v %v %v:%v0%v", year, month, day, h, m, suffix)
		return s
	}

	s := fmt.Sprintf("%v %v %v %v:%v%v", year, month, day, h, m, suffix)
	return s
}
