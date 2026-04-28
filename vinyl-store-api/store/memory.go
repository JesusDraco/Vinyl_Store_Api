package store

import (
	"sync"
	"vinyl-store-api/models"
)

var Albums = []models.Album{
	{
		ID:     "1",
		Title:  "Blue Train",
		Artist: "John Coltrane",
		Price:  56.99,
	},
	{
		ID:     "2",
		Title:  "Time Out",
		Artist: "Dave Brubeck",
		Price:  37.99,
	},
	{
		ID:     "3",
		Title:  "Flying Beagle",
		Artist: "Himiko Kikuchi",
		Price:  69.99,
	},
}

var Users = map[string]string{
	"admin": "1234",
	"user":  "pass",
}

var Tokens = make(map[string]string)

var Mutex sync.Mutex
