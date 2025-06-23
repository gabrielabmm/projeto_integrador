package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"projeto_integrador_versao_final/internal/models"
)


func LimparMascara(s string) string {
	var builder strings.Builder
	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			builder.WriteByte(s[i])
		}
	}
	return builder.String()
}

func ListarPacientesAPI(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	EnableCORS(w)
	if r.Method == http.MethodOptions {
		return
	}

	rows, err := db.Query("SELECT id, nome_completo, cpf_paciente FROM paciente_infos")
	if err != nil {
		http.Error(w, "Erro ao buscar pacientes", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var pacientes []models.Paciente
	for rows.Next() {
		var p models.Paciente
		err := rows.Scan(&p.ID, &p.Nome, &p.CPF)
		if err == nil {
			pacientes = append(pacientes, p)
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pacientes)
}

func InserirPacienteAPI(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	EnableCORS(w)
	if r.Method == http.MethodOptions {
		return
	}

	var p models.Paciente
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		http.Error(w, "Dados inválidos", http.StatusBadRequest)
		return
	}

	fmt.Printf("Recebido do frontend: %+v\n", p)

	p.CPF = LimparMascara(p.CPF)
	p.CartaoSUS = LimparMascara(p.CartaoSUS)
	p.CEP = LimparMascara(p.CEP)
	p.Telefone = LimparMascara(p.Telefone)
	p.DDD = LimparMascara(p.DDD)

	if len(p.DDD) > 2 {
		p.DDD = p.DDD[:2]
	}
	if len(p.Telefone) > 9 {
		p.Telefone = p.Telefone[:9]
	}
	if len(p.CEP) > 8 {
		p.CEP = p.CEP[:8]
	}

	var codMunicipio sql.NullString
	if p.CodMunicipio != "" {
		codMunicipio = sql.NullString{String: p.CodMunicipio, Valid: true}
	} else {
		codMunicipio = sql.NullString{Valid: false}
	}

	_, err := db.Exec(`INSERT INTO paciente_infos (
        cartao_sus, cpf_paciente, nome_completo, data_nascimento,
        cep, ddd, telefone, nacionalidade, uf,
        raca_cor, escolaridade, nome_mae, nome_social,
        logradouro, numero_residencia, complemento, setor, cod_municipio, ponto_referencia
    ) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16,$17,$18,$19)`,
		p.CartaoSUS, p.CPF, p.Nome, p.DataNasc,
		p.CEP, p.DDD, p.Telefone, p.Nacionalidade, p.UF,
		p.Raca, p.Escolaridade, p.NomeMae, p.NomeSocial,
		p.Logradouro, p.NumeroResidencia, p.Complemento, p.Setor, codMunicipio, p.PontoReferencia,
	)
	if err != nil {
		http.Error(w, fmt.Sprintf("Erro ao inserir paciente: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, "Paciente inserido com sucesso.")
}

func ListarPacientePorID(w http.ResponseWriter, r *http.Request, db *sql.DB, id int) {
	var p models.Paciente
	err := db.QueryRow(`SELECT id, cartao_sus, cpf_paciente, nome_completo, data_nascimento, cep, ddd, telefone, nacionalidade, uf, raca_cor, escolaridade, nome_mae, nome_social, logradouro, numero_residencia, complemento, setor, cod_municipio, ponto_referencia
        FROM paciente_infos WHERE id = $1`, id).Scan(
		&p.ID, &p.CartaoSUS, &p.CPF, &p.Nome, &p.DataNasc, &p.CEP, &p.DDD, &p.Telefone, &p.Nacionalidade, &p.UF, &p.Raca, &p.Escolaridade, &p.NomeMae, &p.NomeSocial, &p.Logradouro, &p.NumeroResidencia, &p.Complemento, &p.Setor, &p.CodMunicipio, &p.PontoReferencia,
	)
	if err != nil {
		http.Error(w, "Paciente não encontrado", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(p)
}

func BuscarPacientePorCartaoSUS(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	EnableCORS(w)
	if r.Method == http.MethodOptions {
		return
	}

	cartao := r.URL.Query().Get("cartao_sus")
	if cartao == "" {
		http.Error(w, "Informe o cartao_sus", http.StatusBadRequest)
		return
	}

	var p models.Paciente
	err := db.QueryRow(`SELECT id, cartao_sus, cpf_paciente, nome_completo, data_nascimento, cep, ddd, telefone, nacionalidade, uf, raca_cor, escolaridade, nome_mae, nome_social, logradouro, numero_residencia, complemento, setor, cod_municipio, ponto_referencia
        FROM paciente_infos WHERE cartao_sus = $1`, cartao).Scan(
		&p.ID, &p.CartaoSUS, &p.CPF, &p.Nome, &p.DataNasc, &p.CEP, &p.DDD, &p.Telefone, &p.Nacionalidade, &p.UF, &p.Raca, &p.Escolaridade, &p.NomeMae, &p.NomeSocial, &p.Logradouro, &p.NumeroResidencia, &p.Complemento, &p.Setor, &p.CodMunicipio, &p.PontoReferencia,
	)
	if err != nil {
		http.Error(w, "Paciente não encontrado", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(p)
}
