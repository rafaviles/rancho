package main 

import (
	"database/sql"
	"errors"
)

type product struct {
    ID    int     `json:"id"`
    Name  string  `json:"name"`
    Price float64 `json:"price"`
}


func (p *product) getProdcut(db *sql.DB) error {
	return errors.New("Not implemented")
}

func (p *product) updateProdcut(db *sql.DB) error {
	return errors.New("Not implemented")
}

func (p *product) deleteProdcut(db *sql.DB) error {
	return errors.New("Not implemented")
}

func (p *product) createProdcut(db *sql.DB) error {
	return errors.New("Not implemented")
}

//fetch al products
func getProdcuts(db *sql.DB, start, count int) ([]product, error) {
	return nil, errors.New("Not implemented")
}