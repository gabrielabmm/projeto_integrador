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
	"time" // Importação necessária para time.Time e time.Parse

	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

// --- MODELOS ---

// Paciente representa a estrutura de dados de um paciente no sistema.
// Inclui campos do primeiro e segundo código, mantendo a consistência do banco de dados.
type Paciente struct {
	ID               int    `json:"id"` // Adicionado do segundo código
	CartaoSUS        string `json:"cartao_sus"`
	CPFPaciente      string `json:"cpf_paciente"`    // Nome original do primeiro código
	NomeCompleto     string `json:"nome_completo"`   // Nome original do primeiro código
	DataNascimento   string `json:"data_nascimento"` // Mantido como string para input do HTML, será formatado para output
	CEP              string `json:"cep"`
	DDD              string `json:"ddd"`
	Telefone         string `json:"telefone"`
	Fixo             string `json:"fixo"`           // Do primeiro código
	EmailPaciente    string `json:"email_paciente"` // Do primeiro código
	Nacionalidade    string `json:"nacionalidade"`
	UF               string `json:"uf"`
	RacaCor          string `json:"raca_cor"` // Nome original do primeiro código
	Escolaridade     string `json:"escolaridade"`
	NomeMae          string `json:"nome_mae"`
	NomeSocial       string `json:"nome_social"`
	Logradouro       string `json:"logradouro"`
	NumeroResidencia string `json:"numero_residencia"`
	Complemento      string `json:"complemento"`
	Setor            string `json:"setor"`
	Cidade           string `json:"cidade"` // Do primeiro código
	CodMunicipio     string `json:"cod_municipio"`
	PontoReferencia  string `json:"ponto_referencia"`
}

// ExameCitopatologico representa a estrutura de dados para o exame citopatológico.
type ExameCitopatologico struct {
	PacienteID             int          `json:"paciente_id"`
	MotivoExame            string       `json:"motivo_exame"`
	PrimeiraVezExame       string       `json:"primeira_vez_exame"`
	UsaDIU                 string       `json:"usa_diu"`
	UsaAnticoncepcional    string       `json:"usa_anticoncepcional"`
	EstaGestante           string       `json:"esta_gestante"`
	UsaHormonio            string       `json:"usa_hormonio"`
	JaFezRadioterapia      string       `json:"ja_fez_radioterapia"`
	DataUltimaMenstruacao  sql.NullTime `json:"data_ultima_menstruacao"` // Pode ser nula
	EstaMenopausa          string       `json:"esta_menopausa"`
	TeveCorrimento         string       `json:"teve_corrimento"`
	TeveSangramentoAnormal string       `json:"teve_sangramento_anormal"`
	// Campos relacionados a profissional e data de coleta não estão diretamente na tabela exame_citopatologico
	// do SQL, mas sim em 'resultado_exame' ou precisariam de uma FK para 'profissional_saude'.
	// Para inserção aqui, ignoraremos esses campos por enquanto ou os trataremos separadamente.
	// ProfissionalResponsavel string `json:"profissional_responsavel"` // Não é um campo direto em exame_citopatologico
	// DataColeta              string `json:"data_coleta"`              // Não é um campo direto em exame_citopatologico
}

// UBSInfoRequest representa os dados da UBS vindos do formulário HTML.
type UBSInfoRequest struct {
	UF            string `json:"uf"`
	Protocolo     string `json:"protocolo"`
	Cnes          string `json:"cnes"`
	Unidade       string `json:"unidade"`
	MunicipiosUbs string `json:"municipios_ubs"`
	Prontuario    string `json:"prontuario"` // Adicionado para capturar o campo prontuário da UBS
}

// FichaCitopatologicaPayload é o payload completo recebido do frontend.
type FichaCitopatologicaPayload struct {
	UBSInfoData UBSInfoRequest      `json:"ubs_info_data"`
	ExameData   ExameCitopatologico `json:"exame_data"`
}

// SacForm representa a estrutura para o formulário de SAC.
// Do primeiro código.
type SacForm struct {
	Email    string `json:"email"`
	Mensagem string `json:"mensagem"`
}

// --- UTILITÁRIOS ---

// limparMascara remove caracteres não numéricos de uma string.
func limparMascara(str string) string {
	return strings.Map(func(r rune) rune {
		if r >= '0' && r <= '9' {
			return r
		}
		return -1
	}, str)
}

// isValidEmail verifica se uma string é um formato de e-mail válido.
func isValidEmail(email string) bool {
	emailRegex := regexp.MustCompile(`^[^\s@]+@[^\s@]+\.[^\s@]+$`)
	return emailRegex.MatchString(email)
}

// --- CONEXÃO COM O BANCO ---

// conectar estabelece uma conexão com o banco de dados PostgreSQL.
func conectar() (*sql.DB, error) {
	connStr := "host=localhost port=5432 user=postgres password=postgres dbname=ProjetoIntegrador sslmode=disable"
	return sql.Open("postgres", connStr)
}

// --- MIDDLEWARE CORS ---

// corsMiddleware é um middleware que adiciona os cabeçalhos CORS a todas as respostas.
// Ele também lida com as requisições OPTIONS (preflight) respondendo com 200 OK.
func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Define os cabeçalhos que permitem todas as origens
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, X-Requested-With, Accept, Origin, Authorization")

		// Se a requisição for do tipo OPTIONS (preflight), responde com 200 OK e encerra.
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		// Chama o próximo handler na cadeia
		next.ServeHTTP(w, r)
	})
}

// --- HANDLERS DE PACIENTES ---

// inserirPacienteHandler lida com o cadastro de um novo paciente.
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

	// Limpeza de máscaras e truncamento
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

	// --- Tratamento da Data de Nascimento para Inserção ---
	var dataNascimentoTime time.Time
	if paciente.DataNascimento != "" {
		// Assume o formato DDMMYYYY vindo do frontend char-input
		dataNascimentoTime, err = time.Parse("02012006", paciente.DataNascimento)
		if err != nil {
			http.Error(w, "Formato de data de nascimento inválido (esperado DDMMYYYY): "+err.Error(), http.StatusBadRequest)
			return
		}
	} else {
		http.Error(w, "Data de nascimento é obrigatória", http.StatusBadRequest)
		return
	}
	// --- Fim Tratamento da Data de Nascimento ---

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
		paciente.CartaoSUS, paciente.CPFPaciente, paciente.NomeCompleto, dataNascimentoTime, // Usa a data parseada
		paciente.CEP, paciente.DDD, paciente.Telefone, paciente.Fixo, paciente.EmailPaciente, paciente.Nacionalidade, paciente.UF, paciente.RacaCor, paciente.Escolaridade, paciente.NomeMae, paciente.NomeSocial, paciente.Logradouro, paciente.NumeroResidencia, paciente.Complemento, paciente.Setor, paciente.Cidade, codMunicipio, paciente.PontoReferencia,
	)
	if err != nil {
		http.Error(w, "Erro ao inserir no banco: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Paciente cadastrado com sucesso!"))
}

// listarPacientesAPI lida com a listagem de pacientes (ID, Nome Completo, CPF).
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
		// dataNascimentoDB time.Time está comentada porque esta função lista apenas ID, nome e CPF.
		err := rows.Scan(&p.ID, &p.NomeCompleto, &p.CPFPaciente)
		if err == nil {
			pacientes = append(pacientes, p)
		} else {
			log.Printf("Erro ao escanear paciente: %v", err)
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pacientes)
}

// listarPacientePorID lida com a busca de um paciente por ID.
func listarPacientePorID(w http.ResponseWriter, r *http.Request, db *sql.DB, id int) {
	var p Paciente
	var codMunicipio sql.NullString
	var dataNascimentoDB time.Time // Variável para escanear a data do banco

	err := db.QueryRow(`SELECT id, cartao_sus, cpf_paciente, nome_completo, data_nascimento, cep, ddd, telefone, nacionalidade, uf, raca_cor, escolaridade, nome_mae, nome_social, logradouro, numero_residencia, complemento, setor, cidade, cod_municipio, ponto_referencia, fixo, email
		FROM paciente_infos WHERE id = $1`, id).Scan(
		&p.ID, &p.CartaoSUS, &p.CPFPaciente, &p.NomeCompleto, &dataNascimentoDB, // Escaneia para time.Time
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

	// Formata a data de nascimento para YYYY-MM-DD para o frontend
	p.DataNascimento = dataNascimentoDB.Format("2006-01-02")

	if codMunicipio.Valid {
		p.CodMunicipio = codMunicipio.String
	} else {
		p.CodMunicipio = ""
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(p)
}

// buscarPacientePorCartaoSUS lida com a busca de um paciente pelo cartão SUS.
func buscarPacientePorCartaoSUS(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	cartao := r.URL.Query().Get("cartao_sus")
	if cartao == "" {
		http.Error(w, "Informe o cartao_sus", http.StatusBadRequest)
		return
	}

	var p Paciente
	var codMunicipio sql.NullString
	var dataNascimentoDB time.Time // Variável para escanear a data do banco

	err := db.QueryRow(`SELECT id, cartao_sus, cpf_paciente, nome_completo, data_nascimento, cep, ddd, telefone, nacionalidade, uf, raca_cor, escolaridade, nome_mae, nome_social, logradouro, numero_residencia, complemento, setor, cidade, cod_municipio, ponto_referencia, fixo, email
		FROM paciente_infos WHERE cartao_sus = $1`, cartao).Scan(
		&p.ID, &p.CartaoSUS, &p.CPFPaciente, &p.NomeCompleto, &dataNascimentoDB, // Escaneia para time.Time
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

	// Formata a data de nascimento para YYYY-MM-DD para o frontend
	p.DataNascimento = dataNascimentoDB.Format("2006-01-02")

	if codMunicipio.Valid {
		p.CodMunicipio = codMunicipio.String
	} else {
		p.CodMunicipio = ""
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(p)
}

// --- HANDLERS DE EXAMES CITOPATOLÓGICOS ---

// inserirExameCitopatologicoAPI lida com a inserção de um novo exame citopatológico.
func inserirExameCitopatologicoAPI(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	var e ExameCitopatologico
	if err := json.NewDecoder(r.Body).Decode(&e); err != nil {
		http.Error(w, "Dados inválidos: "+err.Error(), http.StatusBadRequest)
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

// submitFichaCitopatologicaHandler lida com a submissão completa da ficha citopatológica.
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

	// --- 1. Inserir/Atualizar informações da UBS ---
	// Para simplificar, faremos um INSERT OR UPDATE (UPSERT) na tabela ubs_infos baseado no CNES e Protocolo.
	// Se a UBS já existe, atualiza. Caso contrário, insere.
	// Nota: `municipios_ubs` é TEXT no SQL, mas estamos recebendo um único `municipio` no HTML.
	// 'estado' no SQL é tipo_uf. Vou usar 'uf' do request como 'estado'.

	// Limpar máscaras para CNES e Protocolo
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

	// --- 2. Inserir informações do Exame Citopatológico ---
	// Mapear strings do HTML para ENUMs do Go e tipos do PostgreSQL.
	// id_profissional: A tabela SQL espera um ID. Como o HTML envia um nome (Profissional Responsável),
	// você precisaria de uma lógica para buscar o ID desse profissional na tabela profissional_saude,
	// ou modificar o frontend para enviar o ID do profissional.
	// Por simplicidade para este exemplo, vamos usar sql.NullInt32 para id_profissional.
	var idProfissional sql.NullInt32
	// Exemplo de como você buscaria o ID do profissional pelo nome, se houvesse uma tabela de profissionais
	// var profissionalID int
	// err = db.QueryRow("SELECT id FROM profissional_saude WHERE nome = $1", payload.ExameData.ProfissionalResponsavel).Scan(&profissionalID)
	// if err == nil {
	// 	idProfissional = sql.NullInt32{Int32: int32(profissionalID), Valid: true}
	// } else {
	// 	log.Printf("Profissional '%s' não encontrado, id_profissional será NULL. Erro: %v", payload.ExameData.ProfissionalResponsavel, err)
	// }

	// Tratamento da Data da Última Menstruação
	var dataUltimaMenstruacao sql.NullTime
	if payload.ExameData.DataUltimaMenstruacao.Valid { // Use .Valid para verificar se há valor em sql.NullTime
		// Frontend envia DD/MM/YYYY, precisamos parsear
		parsedTime, err := time.Parse("02/01/2006", payload.ExameData.DataUltimaMenstruacao.Time.Format("02/01/2006")) // Formato DD/MM/YYYY
		if err != nil {
			log.Printf("Erro ao parsear data da última menstruação '%v': %v", payload.ExameData.DataUltimaMenstruacao, err)
			// Se houver erro no parse, a data permanecerá nula. Você pode escolher retornar um erro HTTP também.
		} else {
			dataUltimaMenstruacao = sql.NullTime{Time: parsedTime, Valid: true}
		}
	}

	_, err = db.Exec(`
		INSERT INTO exame_citopatologico (
			id_profissional, paciente_id, motivo_exame, primeira_vez_exame, usa_diu,
			usa_anticoncepcional, esta_gestante, usa_hormonio, ja_fez_radioterapia,
			data_ultima_menstruacao, esta_menopausa, teve_corrimento, teve_sangramento_anormal
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)
	`,
		idProfissional, // Usando sql.NullInt32 (será NULL por padrão se não for preenchido)
		payload.ExameData.PacienteID,
		payload.ExameData.MotivoExame,
		payload.ExameData.PrimeiraVezExame,
		payload.ExameData.UsaDIU,
		payload.ExameData.UsaAnticoncepcional,
		payload.ExameData.EstaGestante,
		payload.ExameData.UsaHormonio,
		payload.ExameData.JaFezRadioterapia,
		dataUltimaMenstruacao, // Usa sql.NullTime
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

// --- HANDLERS DE CÓDIGO POR EMAIL (primeiro código) ---

var codigosVerificacao = make(map[string]string) // Mapa para armazenar códigos em memória

// checarEmailEEnviarCodigoHandler verifica se o email existe no banco de dados
// e, se existir, gera e "envia" um código de verificação.
func checarEmailEEnviarCodigoHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	email := r.FormValue("email")
	if email == "" {
		http.Error(w, "E-mail obrigatório", http.StatusBadRequest)
		return
	}

	// 1. Verificar se o email existe no banco de dados (tabela paciente_infos)
	var pacienteID int
	err := db.QueryRow("SELECT id FROM paciente_infos WHERE email = $1", email).Scan(&pacienteID)

	if err == sql.ErrNoRows {
		// Email não encontrado
		http.Error(w, "E-mail não encontrado no cadastro.", http.StatusNotFound)
		return
	} else if err != nil {
		// Erro ao consultar o banco de dados
		log.Printf("Erro ao consultar email no banco: %v", err)
		http.Error(w, "Erro interno ao verificar o e-mail.", http.StatusInternalServerError)
		return
	}

	// 2. Se o email existe, gerar e armazenar o código de verificação
	// (Usando a mesma lógica que você já tinha no enviarCodigoHandler)
	rand.Seed(time.Now().UnixNano())     // Sempre inicialize o gerador de números aleatórios
	codigo := rand.Intn(900000) + 100000 // Gera um código de 6 dígitos entre 100000 e 999999
	codigoStr := fmt.Sprintf("%06d", codigo)
	codigosVerificacao[email] = codigoStr

	// Em um ambiente real, aqui você enviaria o email de verdade.
	// Por enquanto, vamos logar para fins de depuração.
	log.Printf("Código de verificação gerado para %s: %s", email, codigoStr)

	// 3. Responder ao frontend que o email foi encontrado e o código foi "enviado"
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Email encontrado e código enviado com sucesso!"))
}

// enviarCodigoHandler envia um código de verificação para o email.
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
	rand.Seed(time.Now().UnixNano())
	codigo := rand.Intn(900000) + 100000 // Gera um código de 6 dígitos
	codigoStr := fmt.Sprintf("%06d", codigo)
	codigosVerificacao[email] = codigoStr
	log.Printf("Código de verificação para %s: %s", email, codigoStr)
	w.Write([]byte("Código enviado para seu e-mail!"))
}

// verificarCodigoHandler verifica se o código recebido é válido para o email.
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
		delete(codigosVerificacao, email) // Remove o código após o uso
	} else {
		http.Error(w, "Código incorreto", http.StatusUnauthorized)
	}
}

// redefinirSenhaHandler lida com a redefinição de senha para um email.
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

	res, err := db.Exec("UPDATE paciente_infos SET senha = $1 WHERE email = $2", string(hash), email) // Convertendo hash para string
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

// --- HANDLER DE SAC (primeiro código) ---

// handleSacForm lida com o envio do formulário de SAC, validando o email e a mensagem.
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

// --- MAIN ---

func main() {
	// Conecta ao banco de dados uma vez na inicialização
	db, err := conectar()
	if err != nil {
		log.Fatalf("Erro ao conectar ao banco: %v", err)
	}
	defer db.Close() // Garante que a conexão seja fechada ao finalizar o programa

	// Cria um novo ServeMux para registrar as rotas
	mux := http.NewServeMux()

	// Roteamento de API para pacientes
	mux.HandleFunc("/pacientes", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			listarPacientesAPI(w, r, db) // Lista todos os pacientes
		case http.MethodPost:
			inserirPacienteHandler(w, r, db) // Cadastra um novo paciente
		default: // OPTIONS é tratado pelo middleware, outros métodos são inválidos aqui
			http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		}
	})

	// Rota para buscar paciente por ID (usa uma sub-rota dinâmica)
	mux.HandleFunc("/pacientes/", func(w http.ResponseWriter, r *http.Request) {
		// A rota deve ser do tipo /pacientes/{id}
		idStr := strings.TrimPrefix(r.URL.Path, "/pacientes/")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "ID inválido", http.StatusBadRequest)
			return
		}
		listarPacientePorID(w, r, db, id) // Busca paciente por ID
	})

	// Rota para buscar paciente por Cartão SUS (usa query parameter)
	mux.HandleFunc("/pacientes/busca", func(w http.ResponseWriter, r *http.Request) {
		buscarPacientePorCartaoSUS(w, r, db)
	})

	// Roteamento de API para exames citopatológicos
	mux.HandleFunc("/exame_citopatologico", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			inserirExameCitopatologicoAPI(w, r, db) // Insere um novo exame
		default: // OPTIONS é tratado pelo middleware, outros métodos são inválidos aqui
			http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		}
	})

	// Nova rota para submissão completa da ficha citopatológica
	mux.HandleFunc("/submit_ficha_citopatologica", func(w http.ResponseWriter, r *http.Request) {
		submitFichaCitopatologicaHandler(w, r, db)
	})

	// Roteamento para funcionalidades de redefinição de senha e SAC
	mux.HandleFunc("/enviar-codigo", enviarCodigoHandler)
	mux.HandleFunc("/verificar-codigo", verificarCodigoHandler)
	mux.HandleFunc("/redefinir-senha", func(w http.ResponseWriter, r *http.Request) {
		redefinirSenhaHandler(w, r, db) // Passa db para o handler
	})
	mux.HandleFunc("/api/sac", handleSacForm)

	// Aplica o middleware CORS a todas as rotas do mux
	handler := corsMiddleware(mux)

	log.Println("Servidor rodando em http://localhost:3000")
	log.Fatal(http.ListenAndServe(":3000", handler)) // Usa o handler com o middleware aplicado
}
