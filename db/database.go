package db

import (
	"gourlshort/model"
)

var database = make(map[string]model.URL)

func FindAll() []model.URL {
	items := make([]model.URL, 0, len(database))
	for _, v := range database {
		items = append(items, v)
	}

	return items
}

func FindBy(key string) (model.URL, bool) {
	url, ok := database[key]

	return url, ok
}

func Save(key string, item model.URL) {
	database[key] = item
}

func Remove(key string) {
	delete(database, key)
}
