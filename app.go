package main

import (
    "database/sql"
    "github.com/gorilla/mux"
    _ "github.com/go-sql-driver/mysql"
    "fmt"
    "log"
	"net/http"
	"strconv"
	"encoding/json"

)

type App struct {
    Router *mux.Router
    DB     *sql.DB
}

func (a *App) Initialize(user, password, dbname string) {
    connectionString := fmt.Sprintf("%s:%s@/%s", user, password, dbname)
    var err error
    a.DB, err = sql.Open("mysql", connectionString)
    if err != nil {
        log.Fatal(err)
    }
    a.Router = mux.NewRouter()
    a.initializeRoutes()
}

func (a *App) getCart(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	cart_id, err := strconv.Atoi(vars["cart_id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid cart ID")
		return
	}

	p := cart{CartId: cart_id}
	items, err := p.getCart(a.DB);
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			respondWithError(w, http.StatusNotFound, "Cart not found")
		default:
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	respondWithJSON(w, http.StatusOK, items)
}

func (a *App) createCart(w http.ResponseWriter, r *http.Request) {
	var p cart
	decoder := json.NewDecoder(r.Body)
	fmt.Println(r)
	if err := decoder.Decode(&p); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	if err := p.createCart(a.DB); err != nil {

		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}


	respondWithJSON(w, http.StatusCreated, p)
}

func (a *App) updateCart(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid cart ID")
		return
	}

	var p cart
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&p); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid resquest payload")
		return
	}
	defer r.Body.Close()
	p.ID = id

	if err := p.updateCart(a.DB); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, p)
}

func (a *App) deleteCart(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	fmt.Println(r)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid item ID")
		return
	}

	cart_id, err := strconv.Atoi(vars["cart_id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid cart ID")
		return
	}
	p := cart{ID: id, CartId: cart_id}

	if err := p.deleteCart(a.DB); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}


func (a *App) initializeRoutes() {
	a.Router.HandleFunc("/cart", a.createCart).Methods("POST")
	a.Router.HandleFunc("/cart/{cart_id}", a.getCart).Methods("GET")
	a.Router.HandleFunc("/cart/{cart_id}", a.updateCart).Methods("PUT")
	a.Router.HandleFunc("/cart/{cart_id}/{id}", a.deleteCart).Methods("DELETE")
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	fmt.Println(message)
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func (a *App) Run(addr string) { 
	log.Fatal(http.ListenAndServe(":8000", a.Router))
}