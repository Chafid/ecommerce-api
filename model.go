package main
import (
    "database/sql"
    "errors"
)
type cart struct {
    ID    int    `json:"id"`
    ItemName string    `json:"item_name"`
    Total int    `json:"total_price"`
    NumItem int `json:"num_of_item"`
}

func (p *cart) getItemCart(db *sql.DB) error {
    return db.QueryRow("SELECT item_name, total_price, num_of_item FROM cart WHERE id=$1", p.ID).Scan(&p.ItemName, &p.Total, &p.NumItem)
}
func (u *cart) updateItemCart(db *sql.DB) error {
    _, err := db.Exec("UPDATE cart SET item_name=$1, total_price=$2, num_of_item=$3 WHERE id=$4", p.ItemName, p.TotalPrice, p.NumItem, p.ID)
    return err
}
func (u *cart) deleteItemCart(db *sql.DB) error {
    _, err := db.Exec("DELETE FROM cart WHERE id=$1", p.ID)
    return err
}
func (u *cart) insertItemCart(db *sql.DB) error {
    err := db.QueryRow("INSERT INTO cart(item_name, total_price, num_of_item) VALUES($1, $2, $3) RETURNING id", p.ItemName, p.TotalPrice, p.NumItem ).Scan(&p.ID)
    if err != nil {
        return err
    }

    return nil
}
