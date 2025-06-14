package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func main() {
	connStr := "host=localhost port=5432 user=postgres password=postgres dbname=ProjetoIntegrador sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Erro ao conectar:", err)
	}
	defer db.Close()

	// Testa a conexão
	err = db.Ping()
	if err != nil {
		log.Fatal("Erro ao testar conexão:", err)
	}

	sqlFile, err := os.ReadFile("14.06.sql")
	if err != nil {
		log.Fatal("Erro ao ler o arquivo:", err)
	}

	_, err = db.Exec(string(sqlFile))
	if err != nil {
		log.Fatal("Erro ao executar SQL:", err)
	}

	fmt.Println("Script SQL executado com sucesso.")
}
