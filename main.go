package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"

	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

func init() {
    rand.Seed(time.Now().UnixNano())
}

type Paciente struct {
	ID               int    `json:"id"`
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

type ExameCitopatologico struct {
	PacienteID             int            `json:"paciente_id"`
	MotivoExame            string         `json:"motivo_exame"`
	PrimeiraVezExame       string         `json:"primeira_vez_exame"`
	UsaDIU                 string         `json:"usa_diu"`
	UsaAnticoncepcional    string         `json:"usa_anticoncepcional"`
	EstaGestante           string         `json:"esta_gestante"`
	UsaHormonio            string         `json:"usa_hormonio"`
	JaFezRadioterapia      string         `json:"ja_fez_radioterapia"`
	DataUltimaMenstruacao string         `json:"data_ultima_menstruacao"`
	EstaMenopausa          string         `json:"esta_menopausa"`
	TeveCorrimento         string         `json:"teve_corrimento"`
	TeveSangramentoAnormal string         `json:"teve_sangramento_anormal"`
}

type UBSInfoRequest struct {
	UF            string `json:"uf"`
	Protocolo     string `json:"protocolo"`
	Cnes          string `json:"cnes"`
	Unidade       string `json:"unidade"`
	MunicipiosUbs string `json:"municipios_ubs"`
	Prontuario    string `json:"prontuario"`
}

type FichaCitopatologicaPayload struct {
	UBSInfoData UBSInfoRequest      `json:"ubs_info_data"`
	ExameData   ExameCitopatologico `json:"exame_data"`
}

type SacForm struct {
	Email    string `json:"email"`
	Mensagem string `json:"mensagem"`
}

func limparMascara(str string) string {
	return strings.Map(func(r rune) rune {
		if r >= '0' && r <= '9' {
			return r
		}
		return -1
	}, str)
}

func isValidEmail(email string) bool {
	emailRegex := regexp.MustCompile(`^[^\s@]+@[^\s@]+\.[^\s@]+$`)
	return emailRegex.MatchString(email)
}

func conectar() (*sql.DB, error) {
	connStr := "host=localhost port=5432 user=postgres password=postgres dbname=ProjetoIntegrador sslmode=disable"
	return sql.Open("postgres", connStr)
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, X-Requested-With, Accept, Origin, Authorization")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func inserirPacienteHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}
	var paciente Paciente
	err := json.NewDecoder(r.Body).Decode(&paciente)
	if err != nil {
		http.Error(w, "Erro ao decodificar JSON: "+err.Error(), http.StatusBadRequest)
		return
	}

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

	var dataNascimentoTime time.Time
	if paciente.DataNascimento != "" {
		dataNascimentoTime, err = time.Parse("2006-01-02", paciente.DataNascimento)
		if err != nil {
			http.Error(w, "Formato de data de nascimento inválido (esperado YYYY-MM-DD): "+err.Error(), http.StatusBadRequest)
			return
		}
	} else {
		http.Error(w, "Data de nascimento é obrigatória", http.StatusBadRequest)
		return
	}

	var codMunicipio sql.NullString
	if paciente.CodMunicipio != "" {
		codMunicipio = sql.NullString{String: paciente.CodMunicipio, Valid: true}
	} else {
		codMunicipio = sql.NullString{Valid: false}
	}

	_, err = db.Exec(`
		INSERT INTO paciente_infos
		(cartao_sus, cpf_paciente, nome_completo, data_nascimento, cep, ddd, telefone, fixo, email, nacionalidade, uf, raca_cor, escolaridade, nome_mae, nome_social, logradouro, numero_residencia, complemento, setor, cidade, cod_municipio, ponto_referencia)
		VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16,$17,$18,$19,$20,$21,$22)
	`,
		paciente.CartaoSUS, paciente.CPFPaciente, paciente.NomeCompleto, dataNascimentoTime,
		paciente.CEP, paciente.DDD, paciente.Telefone, paciente.Fixo, paciente.EmailPaciente, paciente.Nacionalidade, paciente.UF, paciente.RacaCor, paciente.Escolaridade, paciente.NomeMae, paciente.NomeSocial, paciente.Logradouro, paciente.NumeroResidencia, paciente.Complemento, paciente.Setor, paciente.Cidade, codMunicipio, paciente.PontoReferencia,
	)
	if err != nil {
		http.Error(w, "Erro ao inserir no banco: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Paciente cadastrado com sucesso!"))
}

func listarPacientesAPI(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	rows, err := db.Query("SELECT id, nome_completo, cpf_paciente FROM paciente_infos")
	if err != nil {
		http.Error(w, "Erro ao buscar pacientes: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var pacientes []Paciente
	for rows.Next() {
		var p Paciente
		err := rows.Scan(&p.ID, &p.NomeCompleto, &p.CPFPaciente)
		if err == nil {
			pacientes = append(pacientes, p)
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pacientes)
}

func listarPacientePorID(w http.ResponseWriter, r *http.Request, db *sql.DB, id int) {
	var p Paciente
	var codMunicipio sql.NullString
	var dataNascimentoDB time.Time

	err := db.QueryRow(`SELECT id, cartao_sus, cpf_paciente, nome_completo, data_nascimento, cep, ddd, telefone, nacionalidade, uf, raca_cor, escolaridade, nome_mae, nome_social, logradouro, numero_residencia, complemento, setor, cidade, cod_municipio, ponto_referencia, fixo, email
		FROM paciente_infos WHERE id = $1`, id).Scan(
		&p.ID, &p.CartaoSUS, &p.CPFPaciente, &p.NomeCompleto, &dataNascimentoDB,
		&p.CEP, &p.DDD, &p.Telefone, &p.Nacionalidade, &p.UF, &p.RacaCor, &p.Escolaridade, &p.NomeMae, &p.NomeSocial, &p.Logradouro, &p.NumeroResidencia, &p.Complemento, &p.Setor, &p.Cidade, &codMunicipio, &p.PontoReferencia, &p.Fixo, &p.EmailPaciente,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Paciente não encontrado", http.StatusNotFound)
		} else {
			http.Error(w, "Erro ao buscar paciente: "+err.Error(), http.StatusInternalServerError)
		}
		return
	}

	p.DataNascimento = dataNascimentoDB.Format("2006-01-02")

	if codMunicipio.Valid {
		p.CodMunicipio = codMunicipio.String
	} else {
		p.CodMunicipio = ""
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(p)
}

func buscarPacientePorCartaoSUS(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	cartao := r.URL.Query().Get("cartao_sus")
	if cartao == "" {
		http.Error(w, "Informe o cartao_sus", http.StatusBadRequest)
		return
	}

	var p Paciente
	var codMunicipio sql.NullString
	var dataNascimentoDB time.Time

	err := db.QueryRow(`SELECT id, cartao_sus, cpf_paciente, nome_completo, data_nascimento, cep, ddd, telefone, nacionalidade, uf, raca_cor, escolaridade, nome_mae, nome_social, logradouro, numero_residencia, complemento, setor, cidade, cod_municipio, ponto_referencia, fixo, email
		FROM paciente_infos WHERE cartao_sus = $1`, cartao).Scan(
		&p.ID, &p.CartaoSUS, &p.CPFPaciente, &p.NomeCompleto, &dataNascimentoDB,
		&p.CEP, &p.DDD, &p.Telefone, &p.Nacionalidade, &p.UF, &p.RacaCor, &p.Escolaridade, &p.NomeMae, &p.NomeSocial, &p.Logradouro, &p.NumeroResidencia, &p.Complemento, &p.Setor, &p.Cidade, &codMunicipio, &p.PontoReferencia, &p.Fixo, &p.EmailPaciente,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Paciente não encontrado", http.StatusNotFound)
		} else {
			http.Error(w, "Erro ao buscar paciente: "+err.Error(), http.StatusInternalServerError)
		}
		return
	}

	p.DataNascimento = dataNascimentoDB.Format("2006-01-02")

	if codMunicipio.Valid {
		p.CodMunicipio = codMunicipio.String
	} else {
		p.CodMunicipio = ""
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(p)
}

func inserirExameCitopatologicoAPI(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	var e ExameCitopatologico
	var err error

	if err = json.NewDecoder(r.Body).Decode(&e); err != nil {
		http.Error(w, "Dados inválidos: "+err.Error(), http.StatusBadRequest)
		return
	}

	var dataUltimaMenstruacaoDB sql.NullTime
	if e.DataUltimaMenstruacao != "" {
		var parsedTime time.Time
		parsedTime, err = time.Parse("2006-01-02", e.DataUltimaMenstruacao)
		if err != nil {
			log.Printf("Erro ao parsear data da última menstruação '%s': %v", e.DataUltimaMenstruacao, err)
			http.Error(w, "Formato de data da última menstruação inválido (esperado ISO 8601 -YYYY-MM-DD)", http.StatusBadRequest)
			return
		}
		dataUltimaMenstruacaoDB = sql.NullTime{Time: parsedTime, Valid: true}
	}


	_, err = db.Exec(`INSERT INTO exame_citopatologico (
		paciente_id, motivo_exame, primeira_vez_exame, usa_diu, usa_anticoncepcional, esta_gestante, usa_hormonio, ja_fez_radioterapia, data_ultima_menstruacao, esta_menopausa, teve_corrimento, teve_sangramento_anormal
	) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12)`,
		e.PacienteID, e.MotivoExame, e.PrimeiraVezExame, e.UsaDIU, e.UsaAnticoncepcional, e.EstaGestante, e.UsaHormonio, e.JaFezRadioterapia, dataUltimaMenstruacaoDB, e.EstaMenopausa, e.TeveCorrimento, e.TeveSangramentoAnormal,
	)
	if err != nil {
		http.Error(w, fmt.Sprintf("Erro ao inserir exame: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, "Exame inserido com sucesso")
}

func submitFichaCitopatologicaHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	var payload FichaCitopatologicaPayload
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		http.Error(w, "Erro ao decodificar JSON do payload: "+err.Error(), http.StatusBadRequest)
		return
	}

	ubsCnesLimpo := limparMascara(payload.UBSInfoData.Cnes)
	ubsProtocoloLimpo := limparMascara(payload.UBSInfoData.Protocolo)

	_, err = db.Exec(`
		INSERT INTO ubs_infos (cnes, unidade, municipios_ubs, estado, uf, protocolo)
		VALUES ($1, $2, $3, $4, $5, $6)
		ON CONFLICT (protocolo) DO UPDATE SET
			cnes = EXCLUDED.cnes,
			unidade = EXCLUDED.unidade,
			municipios_ubs = EXCLUDED.municipios_ubs,
			estado = EXCLUDED.estado,
			uf = EXCLUDED.uf;
	`, ubsCnesLimpo, payload.UBSInfoData.Unidade, payload.UBSInfoData.MunicipiosUbs, payload.UBSInfoData.UF, payload.UBSInfoData.UF, ubsProtocoloLimpo)
	if err != nil {
		log.Printf("Erro ao inserir/atualizar UBS: %v", err)
		http.Error(w, "Erro ao salvar informações da UBS: "+err.Error(), http.StatusInternalServerError)
		return
	}

	var idProfissional sql.NullInt32

	var dataUltimaMenstruacaoDB sql.NullTime
	if payload.ExameData.DataUltimaMenstruacao != "" {
		parsedTime, err := time.Parse("2006-01-02", payload.ExameData.DataUltimaMenstruacao)
		if err != nil {
			log.Printf("Erro ao parsear data da última menstruação '%s': %v", payload.ExameData.DataUltimaMenstruacao, err)
			http.Error(w, "Formato de data da última menstruação inválido (esperado ISO 8601 -YYYY-MM-DD)", http.StatusBadRequest)
			return
		}
		dataUltimaMenstruacaoDB = sql.NullTime{Time: parsedTime, Valid: true}
	}
	
	_, err = db.Exec(`
		INSERT INTO exame_citopatologico (
			id_profissional, paciente_id, motivo_exame, primeira_vez_exame, usa_diu,
			usa_anticoncepcional, esta_gestante, usa_hormonio, ja_fez_radioterapia,
			data_ultima_menstruacao, esta_menopausa, teve_corrimento, teve_sangramento_anormal
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)
	`,
		idProfissional,
		payload.ExameData.PacienteID,
		payload.ExameData.MotivoExame,
		payload.ExameData.PrimeiraVezExame,
		payload.ExameData.UsaDIU,
		payload.ExameData.UsaAnticoncepcional,
		payload.ExameData.EstaGestante,
		payload.ExameData.UsaHormonio,
		payload.ExameData.JaFezRadioterapia,
		dataUltimaMenstruacaoDB,
		payload.ExameData.EstaMenopausa,
		payload.ExameData.TeveCorrimento,
		payload.ExameData.TeveSangramentoAnormal,
	)
	if err != nil {
		log.Printf("Erro ao inserir exame citopatológico: %v", err)
		http.Error(w, "Erro ao salvar exame citopatológico: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, "Ficha citopatológica salva com sucesso!")
}

var codigosVerificacao = make(map[string]string)

func checarEmailEEnviarCodigoHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	if err := r.ParseForm(); err != nil {
		http.Error(w, "Erro ao parsear dados do formulário: "+err.Error(), http.StatusBadRequest)
		return
	}
	email := r.FormValue("email")
	if email == "" {
		http.Error(w, "E-mail é obrigatório.", http.StatusBadRequest)
		return
	}

	var pacienteID int
	err := db.QueryRow("SELECT id FROM paciente_infos WHERE email = $1", email).Scan(&pacienteID)

	if err == sql.ErrNoRows {
		http.Error(w, "E-mail não encontrado no cadastro.", http.StatusNotFound)
		return
	} else if err != nil {
		log.Printf("Erro ao consultar email no banco: %v", err)
		http.Error(w, "Erro interno ao verificar o e-mail.", http.StatusInternalServerError)
		return
	}

	codigo := rand.Intn(900000) + 100000
	codigoStr := fmt.Sprintf("%06d", codigo)
	codigosVerificacao[email] = codigoStr

	log.Printf("Código de verificação gerado para %s: %s", email, codigoStr)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Email encontrado e código enviado com sucesso!"))
}

func enviarCodigoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}
	email := r.FormValue("email")
	if email == "" {
		http.Error(w, "E-mail obrigatório", http.StatusBadRequest)
		return
	}
	codigo := rand.Intn(900000) + 100000
	codigoStr := fmt.Sprintf("%06d", codigo)
	codigosVerificacao[email] = codigoStr
	log.Printf("Código de verificação para %s: %s", email, codigoStr)
	w.Write([]byte("Código enviado para seu e-mail!"))
}

func verificarCodigoHandler(w http.ResponseWriter, r *http.Request) {
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
		delete(codigosVerificacao, email)
	} else {
		http.Error(w, "Código incorreto", http.StatusUnauthorized)
	}
}

func redefinirSenhaHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) {
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

	res, err := db.Exec("UPDATE paciente_infos SET senha = $1 WHERE email = $2", string(hash), email)
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

func handleSacForm(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	var form SacForm
	err := json.NewDecoder(r.Body).Decode(&form)
	if err != nil {
		http.Error(w, "Erro ao ler dados: "+err.Error(), http.StatusBadRequest)
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

type LLMPromptRequest struct {
    Prompt string `json:"prompt"`
}

type LLMResponse struct {
    Response string `json:"response"`
}

func llmEmailHelpHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
        return
    }

    var req LLMPromptRequest
    var err error

    if err = json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "Erro ao decodificar JSON da requisição: "+err.Error(), http.StatusBadRequest)
        return
    }

    apiKey := "" 
    apiUrl := "https://generativelanguage.googleapis.com/v1beta/models/gemini-2.0-flash:generateContent?key=" + apiKey

    payload := map[string]interface{}{
        "contents": []map[string]interface{}{
            {
                "role": "user",
                "parts": []map[string]string{
                    {"text": req.Prompt},
                },
            },
        },
    }

    var payloadBytes []byte
    payloadBytes, err = json.Marshal(payload) 
    if err != nil {
        http.Error(w, "Erro ao serializar payload para a API Gemini: "+err.Error(), http.StatusInternalServerError)
        return
    }

    var geminiResp *http.Response
    geminiResp, err = http.Post(apiUrl, "application/json", strings.NewReader(string(payloadBytes))) 
    if err != nil {
        log.Printf("Erro ao chamar a API Gemini: %v", err)
        http.Error(w, "Erro ao se comunicar com o serviço de IA.", http.StatusInternalServerError)
        return
    }
    defer geminiResp.Body.Close()

    var geminiResult map[string]interface{}
    err = json.NewDecoder(geminiResp.Body).Decode(&geminiResult) 
    if err != nil {
        log.Printf("Erro ao decodificar resposta da API Gemini: %v", err)
        http.Error(w, "Erro ao processar resposta do serviço de IA.", http.StatusInternalServerError)
        return
    }

    var llmResponseText string
    if candidates, ok := geminiResult["candidates"].([]interface{}); ok && len(candidates) > 0 {
        if candidate, ok := candidates[0].(map[string]interface{}); ok {
            if content, ok := candidate["content"].(map[string]interface{}); ok {
                if parts, ok := content["parts"].([]interface{}); ok && len(parts) > 0 {
                    if part, ok := parts[0].(map[string]interface{}); ok {
                        if text, ok := part["text"].(string); ok {
                            llmResponseText = text
                        }
                    }
                }
            }
        }
    }

    if llmResponseText == "" {
        log.Println("Resposta vazia ou inválida da API Gemini.")
        http.Error(w, "Não foi possível obter uma resposta útil do serviço de IA.", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(LLMResponse{Response: llmResponseText})
}

func main() {
	db, err := conectar()
	if err != nil {
		log.Fatalf("Erro ao conectar ao banco: %v", err)
	}
	defer db.Close()

	mux := http.NewServeMux()

	mux.HandleFunc("/pacientes", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			listarPacientesAPI(w, r, db)
		case http.MethodPost:
			inserirPacienteHandler(w, r, db)
		default:
			http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/pacientes/", func(w http.ResponseWriter, r *http.Request) {
		idStr := strings.TrimPrefix(r.URL.Path, "/pacientes/")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "ID inválido", http.StatusBadRequest)
			return
		}
		listarPacientePorID(w, r, db, id)
	})

	mux.HandleFunc("/pacientes/busca", func(w http.ResponseWriter, r *http.Request) {
		buscarPacientePorCartaoSUS(w, r, db)
	})

	mux.HandleFunc("/exame_citopatologico", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			inserirExameCitopatologicoAPI(w, r, db)
		default:
			http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/submit_ficha_citopatologica", func(w http.ResponseWriter, r *http.Request) {
		submitFichaCitopatologicaHandler(w, r, db)
	})
	mux.HandleFunc("/submit_ficha_citopatologica/", func(w http.ResponseWriter, r *http.Request) {
		submitFichaCitopatologicaHandler(w, r, db)
	})

	mux.HandleFunc("/checar-email", func(w http.ResponseWriter, r *http.Request) {
		checarEmailEEnviarCodigoHandler(w, r, db)
	})

    mux.HandleFunc("/llm-email-help", llmEmailHelpHandler) 

	mux.HandleFunc("/enviar-codigo", enviarCodigoHandler)
	mux.HandleFunc("/verificar-codigo", verificarCodigoHandler)
	mux.HandleFunc("/redefinir-senha", func(w http.ResponseWriter, r *http.Request) {
		redefinirSenhaHandler(w, r, db)
	})
	mux.HandleFunc("/api/sac", handleSacForm)


	handler := corsMiddleware(mux)

	log.Println("Servidor rodando em http://localhost:3000")
	log.Fatal(http.ListenAndServe(":3000", handler))
}
