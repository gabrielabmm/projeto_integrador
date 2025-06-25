package main

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

var db *sql.DB
var tpl *template.Template

func conectar() (*sql.DB, error) {
	connStr := "host=localhost port=5432 user=postgres password=postgres dbname=ProjetoIntegrador sslmode=disable"
	return sql.Open("postgres", connStr)
}

func init() {

	var err error
	tpl, err = template.ParseFiles("inicio.html")
	if err != nil {
		log.Fatal("Erro ao carregar templates: ", err)
	}
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		tpl.Execute(w, nil)

	case "POST":
		tipo := r.FormValue("tipo")

		if tipo == "beneficiario" {
			email := r.FormValue("email")
			password := r.FormValue("password")

			var dbPassword string
			var refID int

			err := db.QueryRow("SELECT password, ref_id FROM usuario WHERE email = $1 AND tipo = 'paciente'", email).Scan(&dbPassword, &refID)
			if err != nil {
				tpl.Execute(w, map[string]string{"Error": "Email não encontrado ou tipo inválido"})
				return
			}

			if password != dbPassword {
				tpl.Execute(w, map[string]string{"Error": "Senha incorreta"})
				return
			}

			var paciente struct {
				NomeCompleto   string
				CPF            string
				CartaoSUS      string
				DataNascimento string
				CEP            string
			}

			err = db.QueryRow(`
				SELECT nome_completo, cpf_paciente, cartao_sus, TO_CHAR(data_nascimento, 'DD/MM/YYYY'), cep
				FROM paciente_infos WHERE id = $1
			`, refID).Scan(&paciente.NomeCompleto, &paciente.CPF, &paciente.CartaoSUS, &paciente.DataNascimento, &paciente.CEP)

			if err != nil {
				log.Println("Erro ao buscar dados do paciente:", err)
				tpl.Execute(w, map[string]string{"Error": "Erro ao carregar dados do paciente"})
				return
			}

			err = tpl.ExecuteTemplate(w, "inicio.html", map[string]interface{}{
				"Paciente": paciente,
			})
			if err != nil {
				log.Println("Erro ao renderizar template:", err)
			}

		} else if tipo == "instituicao" {
			cnes := r.FormValue("cnes")
			crm := r.FormValue("crm")
			password := r.FormValue("password")

			var dbPassword string
			err := db.QueryRow(`
				SELECT u.password
				FROM usuario u
				JOIN ubs_infos ubs ON u.ref_id = ubs.id
				WHERE ubs.cnes = $1 AND u.cpf = $2 AND u.tipo = 'profissional'
			`, cnes, crm).Scan(&dbPassword)

			if err != nil {
				tpl.Execute(w, map[string]string{"Error": "CNES ou CRM/COREN inválido"})
				return
			}

			if password != dbPassword {
				tpl.Execute(w, map[string]string{"Error": "Senha incorreta"})
				return
			}

			http.Redirect(w, r, "iniciomedico.html", http.StatusSeeOther)

		} else {
			tpl.Execute(w, map[string]string{"Error": "Tipo de usuário inválido"})
		}

	default:
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
	}

}
