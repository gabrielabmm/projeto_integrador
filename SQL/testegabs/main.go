package main

import (
	"bytes" 
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	_ "github.com/lib/pq" 
)


type Paciente struct {
	ID                 int    `json:"id"`
	CartaoSUS          string `json:"cartao_sus"`
	CPF                string `json:"cpf"`
	Nome               string `json:"nome"`
	DataNasc           string `json:"data_nasc"`
	CEP                string `json:"cep"`
	DDD                string `json:"ddd"`
	Telefone           string `json:"telefone"`
	Nacionalidade      string `json:"nacionalidade"`
	UF                 string `json:"uf"`
	Raca               string `json:"raca_cor"`
	Escolaridade       string `json:"escolaridade"`
	NomeMae            string `json:"nome_mae"`
	NomeSocial         string `json:"nome_social"`
	Logradouro         string `json:"logradouro"`
	NumeroResidencia   string `json:"numero_residencia"`
	Complemento        string `json:"complemento"`
	Setor              string `json:"setor"`
	CodMunicipio       string `json:"cod_municipio"`
	PontoReferencia    string `json:"ponto_referencia"`
}

var globalDB *sql.DB 

func conectar() (*sql.DB, error) {
	connStr := "host=localhost port=5432 user=postgres password=postgres dbname=ProjetoIntegrador sslmode=disable"
	return sql.Open("postgres", connStr)
}

func enableCORS(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT") 
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


func uploadProfileImageBLOB(w http.ResponseWriter, r *http.Request) {
	enableCORS(w) 
	if r.Method != http.MethodPost {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseMultipartForm(10 << 20) 
	if err != nil {
		http.Error(w, "Erro ao parsear formulário: "+err.Error(), http.StatusBadRequest)
		return
	}

	file, handler, err := r.FormFile("profile_image") 
	if err != nil {
		http.Error(w, "Erro ao obter arquivo: "+err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()

	buffer := bytes.NewBuffer(nil)
	if _, err := io.Copy(buffer, file); err != nil {
		http.Error(w, "Erro ao ler arquivo: "+err.Error(), http.StatusInternalServerError)
		return
	}

	imageData := buffer.Bytes()


	pacienteIDStr := r.FormValue("paciente_id")
	if pacienteIDStr == "" {
		http.Error(w, "ID do paciente não fornecido.", http.StatusBadRequest)
		return
	}
	pacienteID := 0
	_, err = fmt.Sscanf(pacienteIDStr, "%d", &pacienteID)
	if err != nil || pacienteID == 0 {
		http.Error(w, "ID do paciente inválido.", http.StatusBadRequest)
		return
	}

	
	_, err = globalDB.Exec("UPDATE paciente_infos SET imagem_perfil = $1 WHERE id = $2", imageData, pacienteID)
	if err != nil {
		http.Error(w, "Erro ao salvar imagem no banco de dados: "+err.Error(), http.StatusInternalServerError)
		return
	}

	log.Printf("Imagem de perfil do paciente ID %d atualizada com sucesso.", pacienteID)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Imagem de perfil atualizada com sucesso (armazenada como BLOB)!")
}


func getProfileImageBLOB(w http.ResponseWriter, r *http.Request) {
	enableCORS(w) 

	
	pacienteIDStr := r.URL.Query().Get("id")
	if pacienteIDStr == "" {
		http.Error(w, "ID do paciente não fornecido na URL.", http.StatusBadRequest)
		return
	}
	pacienteID := 0
	_, err := fmt.Sscanf(pacienteIDStr, "%d", &pacienteID)
	if err != nil || pacienteID == 0 {
		http.Error(w, "ID do paciente inválido na URL.", http.StatusBadRequest)
		return
	}

	var imageData []byte
	
	err = globalDB.QueryRow("SELECT imagem_perfil FROM paciente_infos WHERE id = $1", pacienteID).Scan(&imageData)
	if err != nil {
		if err == sql.ErrNoRows {
			
			log.Printf("Paciente ID %d não encontrado ou sem imagem de perfil.", pacienteID)
			
			http.Redirect(w, r, "/default_avatar.png", http.StatusFound)
			return
		}
		log.Printf("Erro ao buscar imagem para paciente ID %d: %v", pacienteID, err)
		http.Error(w, "Erro ao buscar imagem: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if len(imageData) == 0 {
		log.Printf("Imagem de perfil vazia para paciente ID %d. Redirecionando para avatar padrão.", pacienteID)
		http.Redirect(w, r, "/default_avatar.png", http.StatusFound)
		return
	}

	
	contentType := http.DetectContentType(imageData)
	
	if contentType == "application/octet-stream" {
		log.Printf("Não foi possível detectar o Content-Type para paciente ID %d. Assumindo image/jpeg.", pacienteID)
		contentType = "image/jpeg"
	}

	w.Header().Set("Content-Type", contentType)
	w.Header().Set("Content-Length", fmt.Sprintf("%d", len(imageData)))
	w.Write(imageData)
}


func inserirPacienteAPI(w http.ResponseWriter, r *http.Request) { 
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

	_, err := globalDB.Exec(`INSERT INTO paciente_infos (
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


func listarPacientesAPI(w http.ResponseWriter, r *http.Request) { 
	enableCORS(w)
	if r.Method == http.MethodOptions {
		return
	}

	
	rows, err := globalDB.Query("SELECT id, nome_completo, cpf_paciente FROM paciente_infos")
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
		} else {
			log.Printf("Erro ao escanear paciente na listagem: %v", err)
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pacientes)
}

func main() {
	var err error
	globalDB, err = conectar() 
	if err != nil {
		log.Fatal("Erro ao conectar ao banco:", err)
	}
	defer globalDB.Close()

	
	http.HandleFunc("/pacientes", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			listarPacientesAPI(w, r)
		case http.MethodPost:
			inserirPacienteAPI(w, r)
		case http.MethodOptions:
			enableCORS(w)
		default:
			http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		}
	})

	
	http.HandleFunc("/upload-profile-image", uploadProfileImageBLOB)
	http.HandleFunc("/get-profile-image-blob", getProfileImageBLOB)


	http.Handle("/default_avatar.png", http.FileServer(http.Dir(".")))

	fmt.Println("Servidor rodando em http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
