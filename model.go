package main
import (
    "database/sql"
    //"errors"
    "fmt"
)
type cart struct {
    ID int `json:"id"`
    CartId int `json:"cart_id"`
    ItemName string `json:"item_name"`
    Total int `json:"total_price"`
    NumItem int `json:"num_of_item"`
}

func (p *cart) getCart(db *sql.DB) ([]cart, error) {
    //fmt.Println(p.ID)
    statement := fmt.Sprintf("SELECT id, cart_id, item_name, total_price, num_of_item FROM cart WHERE cart_id=%d", p.CartId)
    rows, err := db.Query(statement)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    items := []cart{}

    for rows.Next() {
        var item cart
        if err := rows.Scan(&item.ID, &item.CartId, &item.ItemName, &item.Total, &item.NumItem); err != nil {
            return nil, err
        }
        items = append(items, item)
    }

    return items, nil
}
func (p *cart) updateCart(db *sql.DB) error {
    _, err := db.Exec("UPDATE cart SET total_price=?, num_of_item=? WHERE id=? and cart_id=?", p.Total, p.NumItem, p.ID, p.CartId)
    return err
}

func (p *cart) deleteCart(db *sql.DB) error {
    _, err := db.Exec("DELETE FROM cart WHERE id=? and cart_id=?", p.ID, p.CartId)
    fmt.Println(err)
    return err
}
func (p *cart) createCart(db *sql.DB) error {
    statement := fmt.Sprintf("INSERT INTO cart(cart_id, item_name, total_price, num_of_item) VALUES(%d,'%s', %d, %d)", p.CartId, p.ItemName, p.Total, p.NumItem ) 
    fmt.Println(statement)
    _, err := db.Exec(statement)
    
    if err == nil {
        return err
    }

    err = db.QueryRow("SELECT LAST_INSERT_ID()").Scan(&p.ID)
    if err != nil {
        return err
    }
    return nil
}
