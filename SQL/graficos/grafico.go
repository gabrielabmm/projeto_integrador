package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

func conect() (*sql.DB, error) {
	connStr := "host=localhost port=5432 user=postgres password=postgres dbname=ProjetoIntegrador sslmode=disable"
	return sql.Open("postgres", connStr)
}


func primeiroExameHandler(w http.ResponseWriter, r *http.Request) {
	db, err := conect()

	if err != nil {
		http.Error(w, "Erro ao conectar ao banco", http.StatusInternalServerError)
		return
	}

	defer db.Close()

	var rs [][int]
	err = db.QueryRow("SELECT 1").Scan(&rs)

	if err != nil {
		log.Fatal("Erro na consulta:", err)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(rs)
}

func racaCorHandler() {
	db, err := conect()

	if err != nil {
		http.Error(w, "Erro ao conectar ao banco", http.StatusInternalServerError)
		return
	}

	defer db.Close()

	var rs [][int]
	err = db.QueryRow("SELECT 1").Scan(&rs)

	if err != nil {
		log.Fatal("Erro na consulta:", err)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(rs)
}

func resultadosHandler() {
	db, err := conect()

	if err != nil {
		http.Error(w, "Erro ao conectar ao banco", http.StatusInternalServerError)
		return
	}

	defer db.Close()

	var rs [][int]
	err = db.QueryRow("SELECT 1").Scan(&rs)

	if err != nil {
		log.Fatal("Erro na consulta:", err)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(rs)
}

func idadeHandler() {
	db, err := conect()

	if err != nil {
		http.Error(w, "Erro ao conectar ao banco", http.StatusInternalServerError)
		return
	}

	defer db.Close()

	var rs [][int]
	err = db.QueryRow("SELECT 1").Scan(&rs)

	if err != nil {
		log.Fatal("Erro na consulta:", err)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(rs)
}

func escolaridadeHandler() {
	db, err := conect()

	if err != nil {
		http.Error(w, "Erro ao conectar ao banco", http.StatusInternalServerError)
		return
	}

	defer db.Close()

	var rs [][int]
	err = db.QueryRow("SELECT 1").Scan(&rs)

	if err != nil {
		log.Fatal("Erro na consulta:", err)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(rs)
}

func outrasInfosHandler() {
	db, err := conect()

	if err != nil {
		http.Error(w, "Erro ao conectar ao banco", http.StatusInternalServerError)
		return
	}

	defer db.Close()

	var rs [][int]
	err = db.QueryRow("SELECT 1").Scan(&rs)

	if err != nil {
		log.Fatal("Erro na consulta:", err)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(rs)
}

func graphics() {

http.HandleFunc("/motivo/rastreamento", motivoRastreioHandler)
http.HandleFunc("/motivo/rastreamento", motivoRastreioHandler)
http.HandleFunc("/motivo/repetição", motivoRepeticaoHandler)
http.HandleFunc("/motivo/seguimento", motivoSeguimentoHandler)
http.HandleFunc("/primeiro-exame", primeiroExameHandler)
http.HandleFunc("/raça-cor", racaCorHandler)
http.HandleFunc("/resultados", resultadosHandler)
http.HandleFunc("/idade", idadeHandler)
http.HandleFunc("/escolaridade", escolaridadeHandler)
http.HandleFunc("/outras-infomações", outrasInfosHandler)
}
