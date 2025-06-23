package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"html/template"
	"net/http"
)

var db *sql.DB
var tpl *template.Template

func main() {
	var err error
	db, err = sql.Open("postgres", "host=localhost port=5432 user=postgres password=postgres dbname=ProjetoIntegrador sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	
	tpl = template.Must(template.ParseFiles("static/login.html"))

	
	http.HandleFunc("/login", loginHandler)

	
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	fmt.Println("Servidor rodando em http://localhost:8080/login")
	http.ListenAndServe(":8080", nil)
}


func loginHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		tpl.Execute(w, nil)
	case "POST":
		email := r.FormValue("email")
		password := r.FormValue("password")

		var dbPassword string
		err := db.QueryRow("SELECT password FROM usuario WHERE email=$1", email).Scan(&dbPassword)
		if err != nil {
			tpl.Execute(w, map[string]string{"Error": "Email não encontrado"})
			return
		}

		if password != dbPassword {
			tpl.Execute(w, map[string]string{"Error": "Senha incorreta"})
			return
		}

		http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
	default:
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
	}
}
