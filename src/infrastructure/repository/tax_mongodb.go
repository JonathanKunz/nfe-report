package repository

import (
	"database/sql"
)

type TaxMongoDB struct {
	db *sql.DB
}

func NewTaxMongoDB() *TaxMongoDB {
	return nil
}

func (r *TaxMongoDB) Create() string {
	return "nil"
}

func Get() string {
	return "nil"
}

func Update() string {
	return "nil"
}

func Search() string {
	return "nil"
}

func List() string {
	return "nil"
}

func (r *TaxMongoDB) Delete() error {
	return nil
}
