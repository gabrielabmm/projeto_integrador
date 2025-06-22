package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	_ "github.com/lib/pq"
)

type Paciente struct {
	ID               int    `json:"id"`
	CartaoSUS        string `json:"cartao_sus"`
	CPF              string `json:"cpf"`
	Nome             string `json:"nome"`
	DataNasc         string `json:"data_nasc"`
	CEP              string `json:"cep"`
	DDD              string `json:"ddd"`
	Telefone         string `json:"telefone"`
	Nacionalidade    string `json:"nacionalidade"`
	UF               string `json:"uf"`
	Raca             string `json:"raca_cor"`
	Escolaridade     string `json:"escolaridade"`
	NomeMae          string `json:"nome_mae"`
	NomeSocial       string `json:"nome_social"`
	Logradouro       string `json:"logradouro"`
	NumeroResidencia string `json:"numero_residencia"`
	Complemento      string `json:"complemento"`
	Setor             string `json:"setor"`
	CodMunicipio     string `json:"cod_municipio"`
	PontoReferencia  string `json:"ponto_referencia"`
}

func conectar() (*sql.DB, error) {
	connStr := "host=localhost port=5432 user=postgres password=postgres dbname=ProjetoIntegrador sslmode=disable"
	return sql.Open("postgres", connStr)
}

func enableCORS(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
}

func limparMascara(str string) string {
	return strings.Map(func(r rune) rune {
		if r >= '0' && r <= '9' {
			return r
		}
		return -1
	}, str)
}

func inserirPacienteAPI(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	enableCORS(w)
	if r.Method == http.MethodOptions {
		return
	}

	var p Paciente
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		http.Error(w, "Dados inválidos", http.StatusBadRequest)
		return
	}

	fmt.Printf("Recebido do frontend: %+v\n", p)

	p.CPF = limparMascara(p.CPF)
	p.CartaoSUS = limparMascara(p.CartaoSUS)
	p.CEP = limparMascara(p.CEP)
	p.Telefone = limparMascara(p.Telefone)
	p.DDD = limparMascara(p.DDD)
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

func listarPacientesAPI(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	enableCORS(w)
	if r.Method == http.MethodOptions {
		return
	}

	rows, err := db.Query("SELECT id, nome_completo, cpf_paciente FROM paciente_infos")
	if err != nil {
		http.Error(w, "Erro ao buscar pacientes", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var pacientes []Paciente
	for rows.Next() {
		var p Paciente
		err := rows.Scan(&p.ID, &p.Nome, &p.CPF)
		if err == nil {
			pacientes = append(pacientes, p)
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pacientes)
}

func listarPacientePorID(w http.ResponseWriter, r *http.Request, db *sql.DB, id int) {
	var p Paciente
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


func buscarPacientePorCartaoSUS(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	enableCORS(w)
	if r.Method == http.MethodOptions {
		return
	}

	cartao := r.URL.Query().Get("cartao_sus")
	if cartao == "" {
		http.Error(w, "Informe o cartao_sus", http.StatusBadRequest)
		return
	}

	var p Paciente
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

type ExameCitopatologico struct {
	PacienteID             int    `json:"paciente_id"`
	MotivoExame            string `json:"motivo_exame"`
	PrimeiraVezExame       string `json:"primeira_vez_exame"`
	UsaDIU                 string `json:"usa_diu"`
	UsaAnticoncepcional    string `json:"usa_anticoncepcional"`
	EstaGestante           string `json:"esta_gestante"`
	UsaHormonio            string `json:"usa_hormonio"`
	JaFezRadioterapia      string `json:"ja_fez_radioterapia"`
	DataUltimaMenstruacao  string `json:"data_ultima_menstruacao"`
	EstaMenopausa          string `json:"esta_menopausa"`
	TeveCorrimento         string `json:"teve_corrimento"`
	TeveSangramentoAnormal string `json:"teve_sangramento_anormal"`
}

func inserirExameCitopatologicoAPI(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	enableCORS(w)
	if r.Method == http.MethodOptions {
		return
	}

	var e ExameCitopatologico
	if err := json.NewDecoder(r.Body).Decode(&e); err != nil {
		http.Error(w, "Dados inválidos", http.StatusBadRequest)
		return
	}

	_, err := db.Exec(`INSERT INTO exame_citopatologico (
        paciente_id, motivo_exame, primeira_vez_exame, usa_diu, usa_anticoncepcional, esta_gestante, usa_hormonio, ja_fez_radioterapia, data_ultima_menstruacao, esta_menopausa, teve_corrimento, teve_sangramento_anormal
    ) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12)`,
		e.PacienteID, e.MotivoExame, e.PrimeiraVezExame, e.UsaDIU, e.UsaAnticoncepcional, e.EstaGestante, e.UsaHormonio, e.JaFezRadioterapia, e.DataUltimaMenstruacao, e.EstaMenopausa, e.TeveCorrimento, e.TeveSangramentoAnormal,
	)
	if err != nil {
		http.Error(w, fmt.Sprintf("Erro ao inserir exame: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, "Exame inserido com sucesso")
}

func main() {
	db, err := conectar()
	if err != nil {
		log.Fatal("Erro ao conectar ao banco:", err)
	}
	defer db.Close()

	http.HandleFunc("/pacientes", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			listarPacientesAPI(w, r, db)
		case http.MethodPost:
			inserirPacienteAPI(w, r, db)
		case http.MethodOptions:
			enableCORS(w)
		default:
			http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		}
	})

	
	http.HandleFunc("/pacientes/busca", func(w http.ResponseWriter, r *http.Request) {
		buscarPacientePorCartaoSUS(w, r, db)
	})

	http.HandleFunc("/pacientes/", func(w http.ResponseWriter, r *http.Request) {
		enableCORS(w)
		if r.Method == http.MethodOptions {
			return
		}
		idStr := strings.TrimPrefix(r.URL.Path, "/pacientes/")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "ID inválido", http.StatusBadRequest)
			return
		}
		listarPacientePorID(w, r, db, id)
	})

	http.HandleFunc("/exame_citopatologico", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			inserirExameCitopatologicoAPI(w, r, db)
		case http.MethodOptions:
			enableCORS(w)
		default:
			http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		}
	})

	fmt.Println("Servidor rodando em http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
