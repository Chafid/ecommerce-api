package main

import (
    "database/sql"
    "github.com/gorilla/mux"
    _ "github.com/go-sql-driver/mysql"
    "fmt"
    "log"
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
}

func (a *App) getCart(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid cart ID")
		return
	}

	p := product{ID: id}
	if err := p.getProduct(a.DB); err != nil {
		switch err {
		case sql.ErrNoRows:
			respondWithError(w, http.StatusNotFound, "Cart not found")
		default:
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	respondWithJSON(w, http.StatusOK, p)
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func (a *App) Run(addr string) { }