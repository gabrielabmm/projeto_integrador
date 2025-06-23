package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"projeto_integrador_versao_final/internal/handlers"
	"projeto_integrador_versao_final/database"
)

func main() {
	db, err := database.Conectar()
	if err != nil {
		log.Fatal("Erro ao conectar ao banco:", err)
	}
	defer db.Close()

	http.HandleFunc("/pacientes", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handlers.ListarPacientesAPI(w, r, db)
		case http.MethodPost:
			handlers.InserirPacienteAPI(w, r, db)
		case http.MethodOptions:
			handlers.EnableCORS(w)
		default:
			http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/pacientes/busca", func(w http.ResponseWriter, r *http.Request) {
		handlers.BuscarPacientePorCartaoSUS(w, r, db)
	})

	http.HandleFunc("/pacientes/", func(w http.ResponseWriter, r *http.Request) {
		handlers.EnableCORS(w)
		if r.Method == http.MethodOptions {
			return
		}
		idStr := strings.TrimPrefix(r.URL.Path, "/pacientes/")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "ID inválido", http.StatusBadRequest)
			return
		}
		handlers.ListarPacientePorID(w, r, db, id)
	})

	 handlers.SetupAgendamentoRoutes()
	 handlers.RegistrarRotasAPI()
	 handlers.RegistrarRotasCodigoEmail()


	http.HandleFunc("/exame_citopatologico", func(w htt/p.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			handlers.InserirExameCitopatologicoAPI(w, r, db)
		case http.MethodOptions:
			handlers.EnableCORS(w)
		default:
			http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		}
	})

	fmt.Println("Servidor rodando em http://localhost:3000")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
