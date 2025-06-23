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
			err := db.QueryRow("SELECT password FROM usuario WHERE email = $1", email).Scan(&dbPassword)
			if err != nil {
				tpl.Execute(w, map[string]string{"Error": "Email não encontrado"})
				return
			}

			if password != dbPassword {
				tpl.Execute(w, map[string]string{"Error": "Senha incorreta"})
				return
			}

			
			http.Redirect(w, r, "/dashboard", http.StatusSeeOther)

		} else if tipo == "instituicao" {
			cnes := r.FormValue("cnes")
			crm := r.FormValue("crm")
			password := r.FormValue("password")

			var dbPassword string
			err := db.QueryRow(`
				SELECT u.password
				FROM usuario u
				JOIN ubs_infos ubs ON u.ref_id = ubs.id
				WHERE ubs.cnes = $1 AND u.cpf = $2
			`, cnes, crm).Scan(&dbPassword)

			if err != nil {
				tpl.Execute(w, map[string]string{"Error": "CNES ou CRM/COREN inválido"})
				return
			}

			if password != dbPassword {
				tpl.Execute(w, map[string]string{"Error": "Senha incorreta"})
				return
			}

			
			http.Redirect(w, r, "/dashboard-instituicao", http.StatusSeeOther)

		} else {
			tpl.Execute(w, map[string]string{"Error": "Tipo de usuário inválido"})
		}

	default:
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
	}
}
