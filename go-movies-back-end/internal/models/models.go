package models

import "time"

type Movie struct{
	ID int `json:id`
	Title int `json:title`
	ReleaseDate time.Time `json:release_date`
}