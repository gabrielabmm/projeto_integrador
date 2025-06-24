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

			// Buscar os dados da ficha do paciente
			var paciente struct {
				NomeCompleto   string
				Email          string
				CPF            string
				DataNascimento string
			}

			err = db.QueryRow(`
				SELECT nome_completo, email, cpf_paciente, TO_CHAR(data_nascimento, 'DD/MM/YYYY') 
				FROM paciente_infos 
				WHERE id = $1
			`, refID).Scan(&paciente.NomeCompleto, &paciente.Email, &paciente.CPF, &paciente.DataNascimento)

			if err != nil {
				log.Println("Erro ao buscar dados do paciente:", err)
				tpl.Execute(w, map[string]string{"Error": "Erro ao carregar dados do paciente"})
				return
			}

			// Redirecionar com dados para o dashboard
			err = tpl.ExecuteTemplate(w, "dashboard.html", map[string]interface{}{
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

			http.Redirect(w, r, "/dashboard-instituicao", http.StatusSeeOther)

		} else {
			tpl.Execute(w, map[string]string{"Error": "Tipo de usuário inválido"})
		}

	default:
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
	}
}
