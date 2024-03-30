package model

import (
	"database/sql"
	"errors"
	"log"
)

type CarDealer struct {
	Id          string `json:"id"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Address     string `json:"address"`
	Coordinates string `json:"coordinates"`
	Country     string `json:"country"`
}

type CarDealerModel struct {
	DB       *sql.DB
	InfoLog  *log.Logger
	ErrorLog *log.Logger
}

var carDealers = []CarDealer{
	{
		Id:      "1",
		Title:   "Porshe Center Almaty",
		Address: "Кульдинский тракт 12/1",
		Country: "Germany",
	},
	{
		Id:      "2",
		Title:   "Exeed",
		Address: "Халлулина 196Б",
		Country: "Chinese",
	},
	{
		Id:      "3",
		Title:   "Cadillac Almaty",
		Address: "Суюнбая 243/2",
		Country: "USA",
	},
	{
		Id:      "4",
		Title:   "Toyota Center Almaty",
		Address: "Суюнбая 151",
		Country: "Japanese",
	},
	{
		Id:      "5",
		Title:   "Bentley Almaty",
		Address: "Суюнбая 100",
		Country: "Great Britain",
	},
}

func GetCarDealers() []CarDealer {
	return carDealers
}

func GetCarDealer(id string) (*CarDealer, error) {
	for _, r := range carDealers {
		if r.Id == id {
			return &r, nil
		}
	}
	return nil, errors.New("Car dealer not found")
}
