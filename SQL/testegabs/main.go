package main

import (
        "bytes"
        "io"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
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
	Setor            string `json:"setor"`
	CodMunicipio     string `json:"cod_municipio"`
	PontoReferencia  string `json:"ponto_referencia"`
}

func conectar() (*sql.DB, error) {
	connStr := "host=localhost port=5432 user=postgres password=postgres dbname=ProjetoIntegrador sslmode=disable"
	return sql.Open("postgres", connStr)
}

func uploadProfileImageBLOB(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseMultipartForm(10 << 20) // Limite de 10 MB
	if err != nil {
		http.Error(w, "Erro ao parsear formulário: "+err.Error(), http.StatusBadRequest)
		return
	}

	file, handler, err := r.FormFile("profile_image") // "profile_image" do HTML
	if err != nil {
		http.Error(w, "Erro ao obter arquivo: "+err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()buffer := bytes.NewBuffer(nil)
	if _, err := io.Copy(buffer, file); err != nil {
		http.Error(w, "Erro ao ler arquivo: "+err.Error(), http.StatusInternalServerError)
		return
	}

	imageData := buffer.Bytes() 

// 2. Identificar o paciente (exemplo: ID 1. Em um app real, viria de uma sessão/token)
	pacienteID := 1

_, err = db.Exec("UPDATE pacientes SET imagem_perfil = $1 WHERE id = $2", imageData, pacienteID)
	if err != nil {
		http.Error(w, "Erro ao salvar imagem no banco de dados: "+err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Imagem de perfil atualizada com sucesso (armazenada como BLOB)!")
}

func getProfileImageBLOB(w http.ResponseWriter, r *http.Request) {
	// Você precisará de alguma forma de identificar o paciente
	// Para este exemplo, vamos buscar o paciente com ID 1
	pacienteID := 1

	var imageData []byte
	err := db.QueryRow("SELECT imagem_perfil FROM pacientes WHERE id = $1", pacienteID).Scan(&imageData)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Paciente não encontrado ou sem imagem de perfil.", http.StatusNotFound)
			return
		}
		http.Error(w, "Erro ao buscar imagem: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if len(imageData) == 0 {
		http.Error(w, "Imagem de perfil não encontrada.", http.StatusNotFound)
		return
	}

contentType := http.DetectContentType(imageData)
	if contentType == "application/octet-stream" {
		
		w.Header().Set("Content-Type", "image/png") // Ou image/jpeg, dependendo do que você espera
	} else {
		w.Header().Set("Content-Type", contentType)
	}

	w.Header().Set("Content-Length", fmt.Sprintf("%d", len(imageData)))
	w.Write(imageData)
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

	// Tratar cod_municipio para aceitar NULL
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

func main() {
http.HandleFunc("/upload-profile-image-blob", uploadProfileImageBLOB)
	http.HandleFunc("/get-profile-image-blob", getProfileImageBLOB) // Endpoint para buscar a imagem

	fmt.Println("Servidor Go rodando na porta :8080 (BLOB Mode)")
	http.ListenAndServe(":8080", nil)

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

	fmt.Println("Servidor rodando em http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
