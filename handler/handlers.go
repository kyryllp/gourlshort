package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"gourlshort/db"
	"gourlshort/model"
	"net/http"
)

func CreateUrl(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	url := new(model.URL)
	err = json.Unmarshal(body, url)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	db.Save(url.RedirectName, *url)
}

func GetRedirect(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	path := vars["name"]
	db.FindAll()
	url, ok := db.FindBy(path)
	if ok {
		http.Redirect(w, r, url.OriginalUrl, http.StatusFound)
	} else {
		http.NotFound(w, r)
	}
}
