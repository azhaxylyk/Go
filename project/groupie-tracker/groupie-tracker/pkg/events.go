package pkg

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
	"sync"
)

var (
	BandInfo     []Band
	RelationInfo Relations
	bandInfoMu   sync.RWMutex
)

const (
	artistAPI   = "https://groupietrackers.herokuapp.com/api/artists"
	relationAPI = "https://groupietrackers.herokuapp.com/api/relation"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		ErrorHandler(w, http.StatusNotFound)
		return
	}

	if r.Method != http.MethodGet {
		ErrorHandler(w, http.StatusMethodNotAllowed)
		return
	}

	templates, err := template.ParseGlob("./web/templates/*.html")
	if err != nil {
		log.Println(err)
		ErrorHandler(w, http.StatusInternalServerError)
		return
	}

	err = templates.ExecuteTemplate(w, "index.html", &BandInfo)
	if err != nil {
		log.Println(err)
		ErrorHandler(w, http.StatusInternalServerError)
		return
	}
}

func BandHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/band" {
		ErrorHandler(w, http.StatusNotFound)
		return
	}

	if r.Method != http.MethodGet {
		ErrorHandler(w, http.StatusMethodNotAllowed)
		return
	}

	id := r.URL.Query().Get("id")

	numID, err := strconv.Atoi(id)
	if err != nil {
		log.Println(err)
		ErrorHandler(w, http.StatusInternalServerError)
		return
	}

	if numID > len(BandInfo) {
		log.Println(err)
		ErrorHandler(w, http.StatusNotFound)
		return
	}

	band := BandInfo[numID-1]

	band.Locations = RelationInfo.Index[numID-1].DatesLocations

	templates, err := template.ParseGlob("./web/templates/*.html")
	if err != nil {
		log.Println(err)
		ErrorHandler(w, http.StatusInternalServerError)
		return
	}

	err = templates.ExecuteTemplate(w, "band.html", &band)
	if err != nil {
		log.Println(err)
		ErrorHandler(w, http.StatusInternalServerError)
		return
	}
}

func ErrorHandler(w http.ResponseWriter, statusCode int) {
	w.WriteHeader(statusCode)

	data := struct {
		StatusMsg  string
		StatusCode int
	}{
		"Ooops. Error ",
		statusCode,
	}

	templates, err := template.ParseGlob("./web/templates/*.html")
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	err = templates.ExecuteTemplate(w, "error.html", &data)
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}
