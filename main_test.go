//This file is obsolote as the test is not relevant to the current version of the api

package main
import (
    "os"
    "log"
    "testing"
    "net/http"
    "net/http/httptest"
    "encoding/json"
    "bytes"
    "fmt"
)
var a App
func TestMain(m *testing.M) {
    a = App{}
    a.Initialize("ecom", "ecom123", "ecommerce_example")
    ensureTableExists()
    code := m.Run()
    clearTable()
    os.Exit(code)
}
func ensureTableExists() {
    if _, err := a.DB.Exec(tableCreationQuery); err != nil {
        log.Fatal(err)
    }
}
func clearTable() {
    a.DB.Exec("DELETE FROM cart")
    a.DB.Exec("ALTER TABLE cart AUTO_INCREMENT = 1")
}

func TestEmptyTable(t *testing.T) {
    clearTable()
    req, _ := http.NewRequest("GET", "/cart", nil)
    response := executeRequest(req)
    checkResponseCode(t, http.StatusOK, response.Code)
    fmt.Println(response)
    if body := response.Body.String(); body != "" {
        t.Errorf("Expected nil. Got %s", body)
    }
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
    rr := httptest.NewRecorder()
    a.Router.ServeHTTP(rr, req)

    return rr
}
func checkResponseCode(t *testing.T, expected, actual int) {
    if expected != actual {
        t.Errorf("Expected response code %d. Got %d\n", expected, actual)
    }
}

func TestGetNonExistentTrx(t *testing.T) {
    clearTable()
    req, _ := http.NewRequest("GET", "/cart/1", nil)
    response := executeRequest(req)
    checkResponseCode(t, http.StatusNotFound, response.Code)
    var m map[string]string
    json.Unmarshal(response.Body.Bytes(), &m)
    if m["error"] != "Cart not found" {
        t.Errorf("Expected the 'error' key of the response to be set to 'Trx not found'. Got '%s'", m["error"])
    }
}

func addItemCart(item_name string, num_of_item int, total_price int) {

    statement := fmt.Sprintf("INSERT INTO cart(item_name, num_of_item, total_price) VALUES('%s', %d, %d)", item_name, num_of_item, total_price)
    fmt.Println(statement)
    a.DB.Exec(statement)
}

func TestCreateTrx(t *testing.T) {
    clearTable()
    //var payloadString = '`{"item_name": "' + item_name + '", "total_price": ' + total_price + ', "num_of_item":' + num_of_item + "}`"
    //fmt.Println(payloadString)
    payload := []byte(`{"item_name":"laptop/handphone","total_price":17000000, "num_of_item":2}`)
    req, _ := http.NewRequest("POST", "/cart", bytes.NewBuffer(payload))

    //fmt.Println(req)
    response := executeRequest(req)
    //fmt.Println(response)
    checkResponseCode(t, http.StatusCreated, response.Code)
    var m map[string]interface{}
    json.Unmarshal(response.Body.Bytes(), &m)
    if m["item_name"] != "laptop/handphone" {
        t.Errorf("Expected item name to be 'laptop'. Got '%v'", m["item_name"])
    }
    // the id is compared to 1.0 because JSON unmarshaling converts numbers to
    // floats, when the target is a map[string]interface{}
}

const tableCreationQuery = `
CREATE TABLE IF NOT EXISTS cart
(
    id INT AUTO_INCREMENT PRIMARY KEY,
    item_name VARCHAR(50) NOT NULL,
    total_price int not null, 
    num_of_item int not null
)`
