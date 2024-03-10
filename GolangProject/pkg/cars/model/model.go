package model

import (
	"database/sql"
	"log"
	"os"
)

type Models struct {
	cars       carModel
	carDealers carDealerModel
}

func NewModels(db *sql.DB) Models {
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	return Models{
		cars: carModel{
			DB:       db,
			InfoLog:  infoLog,
			ErrorLog: errorLog,
		},
		carDealers: carDealerModel{
			DB:       db,
			InfoLog:  infoLog,
			ErrorLog: errorLog,
		},
	}
}
