package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

//código para manipulação de pacientes

type Paciente struct {
	ID            int
	CartaoSUS     string
	CPF           string
	Nome          string
	DataNasc      string
	CEP           string
	DDD           string
	Telefone      string
	Nacionalidade string
	UF            string
}

func conectar() (*sql.DB, error) {
	connStr := "host=localhost port=5432 user=postgres password=postgres dbname=ProjetoIntegrador sslmode=disable"
	return sql.Open("postgres", connStr)
}

func inserirPacientes(db *sql.DB) {
	pacientes := []Paciente{
		{CartaoSUS: "123456789012345", CPF: "11122233344", Nome: "Ana Souza", DataNasc: "1990-05-10", CEP: "74000000", DDD: "62", Telefone: "912345678", Nacionalidade: "Brasileira", UF: "GO"},
		{CartaoSUS: "987654321098765", CPF: "55566677788", Nome: "Carlos Lima", DataNasc: "1985-12-22", CEP: "74000001", DDD: "62", Telefone: "998765432", Nacionalidade: "Brasileira", UF: "GO"},
		{CartaoSUS: "111122223333444", CPF: "99988877766", Nome: "Maria Oliveira", DataNasc: "2000-07-15", CEP: "74000002", DDD: "62", Telefone: "996547321", Nacionalidade: "Brasileira", UF: "GO"},
		{CartaoSUS: "000000000000000", CPF: "00000000000", Nome: "Emilly Linda", DataNasc: "2005-06-25", CEP: "74693090", DDD: "62", Telefone: "993857260", Nacionalidade: "Brasileira", UF: "GO"},
	}

	for _, p := range pacientes {
		_, err := db.Exec(`
			INSERT INTO paciente_infos (
				cartao_sus, cpf_paciente, nome_completo, data_nascimento,
				cep, ddd, telefone, nacionalidade, uf
			) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`,
			p.CartaoSUS, p.CPF, p.Nome, p.DataNasc,
			p.CEP, p.DDD, p.Telefone, p.Nacionalidade, p.UF,
		)
		if err != nil {
			log.Println("Erro ao inserir paciente:", err)
		}
	}
	fmt.Println("Pacientes inseridos com sucesso.")
}

func listarPacientes(db *sql.DB) {
	rows, err := db.Query("SELECT id, nome_completo, cpf_paciente FROM paciente_infos")
	if err != nil {
		log.Fatal("Erro ao listar:", err)
	}
	defer rows.Close()

	fmt.Println("Pacientes no banco:")
	for rows.Next() {
		var id int
		var nome, cpf string
		err := rows.Scan(&id, &nome, &cpf)
		if err != nil {
			log.Println(err)
			continue
		}
		fmt.Printf("ID: %d | Nome: %s | CPF: %s\n", id, nome, cpf)
	}
}

func atualizarPaciente(db *sql.DB, id int, novoNome string) {
	_, err := db.Exec("UPDATE paciente_infos SET nome_completo = $1 WHERE id = $2", novoNome, id)
	if err != nil {
		log.Println("Erro ao atualizar paciente:", err)
	} else {
		fmt.Println("Paciente atualizado com sucesso.")
	}
}

func deletarPaciente(db *sql.DB, id int) {
	_, err := db.Exec("DELETE FROM paciente_infos WHERE id = $1", id)
	if err != nil {
		log.Println("Erro ao deletar paciente:", err)
	} else {
		fmt.Println("Paciente deletado com sucesso.")
	}
}

func main() {
	db, err := conectar()
	if err != nil {
		log.Fatal("Erro ao conectar:", err)
	}
	defer db.Close()

	// ✅ Escolha o que testar comentando/descomentando abaixo:

	inserirPacientes(db)
	listarPacientes(db)
	//deletarPaciente(db, 3) // Exemplo: deleta paciente com ID 2
	//atualizarPaciente(db, 1, "Ana Paula Souza") // Exemplo: atualiza nome do ID 1

}
