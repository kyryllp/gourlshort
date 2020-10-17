package main

import (
	"database/sql"
	"encoding/json"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"gourlshort/db"
	"gourlshort/model"
	"log"
	"net/http"
	"time"
)

type App struct {
	Router *mux.Router
	DB     *sql.DB
	Cache  map[time.Time]*http.Request
}

func (a *App) Initialize(user, password, dbname string) {
	var err error
	a.DB, err = db.InitializeConnection(user, password, dbname)
	if err != nil {
		log.Fatal(err)
	}

	// kyryll: really feel like this is the wrong way to do it, unfortunately, I couldn't think on a better way
	a.Cache = make(map[time.Time]*http.Request)
	a.Router = mux.NewRouter()

	a.initializeRoutes()
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// kyryll: pass the name(yahoo or google) of the new redirect url, not the path (/yahoo or /google)
func (a *App) createUrl(w http.ResponseWriter, r *http.Request) {
	var url model.URL
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&url); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if url.RedirectName == "" || url.OriginalUrl == "" {
		// kyryll: for some reason the first if cause didn't work, I added this one
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	defer r.Body.Close()

	_, err := db.SaveUrl(a.DB, url)
	if err != nil {
		respondWithError(w, http.StatusNotAcceptable, "Duplicate url found")
		return
	}

	respondWithJSON(w, http.StatusCreated, url)
}

func (a *App) getRedirect(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]

	url := db.GetUrl(a.DB, name)
	if url.OriginalUrl != "" {
		a.cacheRequest(r)
		http.Redirect(w, r, url.OriginalUrl, http.StatusFound)
	} else {
		respondWithError(w, http.StatusNotFound, "No shortened url with this name found.")
		return
	}
}

func (a *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(":8080", a.Router))
}

func (a *App) cacheRequest(req *http.Request) {
	a.Cache[time.Now()] = req
}

func (a *App) initializeRoutes() {
	a.Router.HandleFunc("/urls/create", a.createUrl).Methods("POST")
	a.Router.HandleFunc("/urls/{name}", a.getRedirect).Methods("GET")
}
