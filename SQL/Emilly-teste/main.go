package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"regexp"
	"strings"
	"time"

	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

// --- MODELOS ---
type Paciente struct {
	CartaoSUS        string `json:"cartao_sus"`
	CPFPaciente      string `json:"cpf_paciente"`
	NomeCompleto     string `json:"nome_completo"`
	DataNascimento   string `json:"data_nascimento"`
	CEP              string `json:"cep"`
	DDD              string `json:"ddd"`
	Telefone         string `json:"telefone"`
	Fixo             string `json:"fixo"`
	EmailPaciente    string `json:"email_paciente"`
	Nacionalidade    string `json:"nacionalidade"`
	UF               string `json:"uf"`
	RacaCor          string `json:"raca_cor"`
	Escolaridade     string `json:"escolaridade"`
	NomeMae          string `json:"nome_mae"`
	NomeSocial       string `json:"nome_social"`
	Logradouro       string `json:"logradouro"`
	NumeroResidencia string `json:"numero_residencia"`
	Complemento      string `json:"complemento"`
	Setor            string `json:"setor"`
	Cidade           string `json:"cidade"`
	CodMunicipio     string `json:"cod_municipio"`
	PontoReferencia  string `json:"ponto_referencia"`
}

// --- CORS ---
func enableCORS(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
}

// --- CONEXÃO COM O BANCO ---
func conectar() (*sql.DB, error) {
	connStr := "host=localhost port=5432 user=postgres password=postgres dbname=ProjetoIntegrador sslmode=disable"
	return sql.Open("postgres", connStr)
}

// --- HANDLER DE CADASTRO DE PACIENTE ---
func cadastrarPacienteHandler(w http.ResponseWriter, r *http.Request) {
	enableCORS(w)
	if r.Method == http.MethodOptions {
		return
	}
	if r.Method != http.MethodPost {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}
	var paciente Paciente
	err := json.NewDecoder(r.Body).Decode(&paciente)
	if err != nil {
		http.Error(w, "Erro ao decodificar JSON", http.StatusBadRequest)
		return
	}

	// Limpeza de máscaras
	paciente.CPFPaciente = limparMascara(paciente.CPFPaciente)
	paciente.CartaoSUS = limparMascara(paciente.CartaoSUS)
	paciente.CEP = limparMascara(paciente.CEP)
	paciente.Telefone = limparMascara(paciente.Telefone)
	paciente.DDD = limparMascara(paciente.DDD)
	paciente.Fixo = limparMascara(paciente.Fixo)
	if len(paciente.DDD) > 2 {
		paciente.DDD = paciente.DDD[:2]
	}
	if len(paciente.Telefone) > 9 {
		paciente.Telefone = paciente.Telefone[:9]
	}
	if len(paciente.CEP) > 8 {
		paciente.CEP = paciente.CEP[:8]
	}
	if len(paciente.Fixo) > 20 {
		paciente.Fixo = paciente.Fixo[:20]
	}

	db, err := conectar()
	if err != nil {
		http.Error(w, "Erro ao conectar ao banco", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	var codMunicipio interface{}
	if paciente.CodMunicipio == "" {
		codMunicipio = nil
	} else {
		codMunicipio = paciente.CodMunicipio
	}

	_, err = db.Exec(`
    INSERT INTO paciente_infos 
    (cartao_sus, cpf_paciente, nome_completo, data_nascimento, cep, ddd, telefone, fixo, email, nacionalidade, uf, raca_cor, escolaridade, nome_mae, nome_social, logradouro, numero_residencia, complemento, setor, cidade, cod_municipio, ponto_referencia)
    VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16,$17,$18,$19,$20,$21,$22)
`,
		paciente.CartaoSUS, paciente.CPFPaciente, paciente.NomeCompleto, paciente.DataNascimento, paciente.CEP, paciente.DDD, paciente.Telefone, paciente.Fixo, paciente.EmailPaciente, paciente.Nacionalidade, paciente.UF, paciente.RacaCor, paciente.Escolaridade, paciente.NomeMae, paciente.NomeSocial, paciente.Logradouro, paciente.NumeroResidencia, paciente.Complemento, paciente.Setor, paciente.Cidade, codMunicipio, paciente.PontoReferencia,
	)
	if err != nil {
		http.Error(w, "Erro ao inserir no banco: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Paciente cadastrado com sucesso!"))
}

// --- HANDLERS DE CÓDIGO POR EMAIL ---
var codigosVerificacao = make(map[string]string)

func enviarCodigoHandler(w http.ResponseWriter, r *http.Request) {
	enableCORS(w)
	if r.Method == http.MethodOptions {
		return
	}
	if r.Method != http.MethodPost {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}
	email := r.FormValue("email")
	if email == "" {
		http.Error(w, "E-mail obrigatório", http.StatusBadRequest)
		return
	}
	rand.Seed(time.Now().UnixNano())
	codigo := rand.Intn(900000) + 100000
	codigoStr := fmt.Sprintf("%06d", codigo)
	codigosVerificacao[email] = codigoStr
	log.Printf("Código de verificação para %s: %s", email, codigoStr)
	w.Write([]byte("Código enviado para seu e-mail!"))
}

func verificarCodigoHandler(w http.ResponseWriter, r *http.Request) {
	enableCORS(w)
	if r.Method == http.MethodOptions {
		return
	}
	if r.Method != http.MethodPost {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}
	email := r.FormValue("email")
	codigo := r.FormValue("codigo")
	if email == "" || codigo == "" {
		http.Error(w, "E-mail e código são obrigatórios", http.StatusBadRequest)
		return
	}
	if codigosVerificacao[email] == codigo {
		w.Write([]byte("Código correto"))
	} else {
		http.Error(w, "Código incorreto", http.StatusUnauthorized)
	}
}

func redefinirSenhaHandler(w http.ResponseWriter, r *http.Request) {
	enableCORS(w)
	if r.Method == http.MethodOptions {
		return
	}
	if r.Method != http.MethodPost {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}
	email := r.FormValue("email")
	novaSenha := r.FormValue("nova_senha")
	if email == "" || novaSenha == "" {
		http.Error(w, "E-mail e nova senha são obrigatórios", http.StatusBadRequest)
		return
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(novaSenha), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Erro ao gerar hash da senha", http.StatusInternalServerError)
		return
	}
	db, err := conectar()
	if err != nil {
		http.Error(w, "Erro ao conectar ao banco", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	res, err := db.Exec("UPDATE paciente_infos SET senha = $1 WHERE email = $2", hash, email)
	if err != nil {
		http.Error(w, "Erro ao atualizar senha: "+err.Error(), http.StatusInternalServerError)
		return
	}
	rows, _ := res.RowsAffected()
	if rows == 0 {
		http.Error(w, "E-mail não encontrado", http.StatusNotFound)
		return
	}
	w.Write([]byte("Senha redefinida com sucesso!"))
}

// --- HANDLER DE SAC (VALIDAÇÃO DE EMAIL) ---
type SacForm struct {
	Email    string `json:"email"`
	Mensagem string `json:"mensagem"`
}

func isValidEmail(email string) bool {
	emailRegex := regexp.MustCompile(`^[^\s@]+@[^\s@]+\.[^\s@]+$`)
	return emailRegex.MatchString(email)
}

func handleSacForm(w http.ResponseWriter, r *http.Request) {
	enableCORS(w)
	if r.Method == http.MethodOptions {
		return
	}
	if r.Method != http.MethodPost {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	var form SacForm
	err := json.NewDecoder(r.Body).Decode(&form)
	if err != nil {
		http.Error(w, "Erro ao ler dados", http.StatusBadRequest)
		return
	}

	if !isValidEmail(form.Email) {
		http.Error(w, "Email inválido", http.StatusBadRequest)
		return
	}

	if form.Mensagem == "" {
		http.Error(w, "Mensagem não pode estar vazia", http.StatusBadRequest)
		return
	}

	log.Printf("SAC recebido: Email=%s | Mensagem=%s\n", form.Email, form.Mensagem)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Mensagem recebida com sucesso"))
}

// --- UTILITÁRIO ---
func limparMascara(str string) string {
	return strings.Map(func(r rune) rune {
		if r >= '0' && r <= '9' {
			return r
		}
		return -1
	}, str)
}

// --- MAIN ---
func main() {
	http.HandleFunc("/enviar-codigo", enviarCodigoHandler)
	http.HandleFunc("/verificar-codigo", verificarCodigoHandler)
	http.HandleFunc("/redefinir-senha", redefinirSenhaHandler)
	http.HandleFunc("/pacientes", cadastrarPacienteHandler)
	http.HandleFunc("/api/sac", handleSacForm)

	log.Println("Servidor rodando em http://localhost:3000")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
