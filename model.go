package main
import (
    "database/sql"
    "errors"
)
type cart struct {
    ID    int    `json:"id"`
    TrxID  string `json:"trx_id"`
    ItemName string    `json:"item_name"`
    ItemID string    `json:"item_id"`
    Total int    `json:"total_price"`
    NumItem int `json:"num_of_item"`
}

func (u *cart) updateItemCart(db *sql.DB) error {
    return errors.New("Not implemented")
}
func (u *cart) deleteItemCart(db *sql.DB) error {
    return errors.New("Not implemented")
}
func (u *cart) insertItemCart(db *sql.DB) error {
    return errors.New("Not implemented")
}
func getItemCart(db *sql.DB, start, count int) ([]cart, error) {
    return nil, errors.New("Not implemented")
}