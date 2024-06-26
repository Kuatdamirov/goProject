package model

import (
	"database/sql"
	"log"
	"os"
)

type Models struct {
	Cars       CarModel
	CarDealers CarDealerModel
}

func NewModels(db *sql.DB) Models {
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	return Models{
		Cars: CarModel{
			DB:       db,
			InfoLog:  infoLog,
			ErrorLog: errorLog,
		},
		CarDealers: CarDealerModel{
			DB:       db,
			InfoLog:  infoLog,
			ErrorLog: errorLog,
		},
	}
}
