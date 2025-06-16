package main

import (
	"encoding/json"
	"net/http"
)

var paciente = Paciente{
	NomeCompleto:   "Maria Linda da Silva",
	PrimeiroNome:   "Maria",
	DataNascimento: "01/01/1985",
	Endereco: Endereco{
		CEP:        "12345-678",
		Cidade:     "Exemplópolis",
		Logradouro: "Rua dos bobos, n°0",
	},
	CPF:          "012.345.678-90",
	CartaoSUS:    "123.4567.7890.1234",
	Online:       Online{Email: "maria88@linda.com.br", Login: "maria88@linda.com.br", Senha: "Pi2025"},
	RG:           "00.000.000-0",
	Sexo:         "Feminino",
	Celular:      "(11) 98765-4321",
	Telefone:     "(11) 1234-5678",
	Escolaridade: "Superior Completo",
	Observacao:   "Tem alzheimer e uma cirurgia no estômago.",
}

var instituicao = Instituicao{
	NomeInstituicao:     "Joana",
	CNES:                "1234567",
	CRMCorenResponsavel: "CRM/SP 123456",
	CNPJ:                "12.345.678/0001-99",
	Email:               "contato@clinicajoana.com",
	Celular:             "(11) 91234-5678",
	Telefone:            "(11) 5555-1234",
	CEP:                 "01000-001",
	Cidade:              "São Paulo",
	Endereco:            "Rua da Clínica",
	Numero:              "100",
	Complemento:         "Andar 5, Sala 502",
	Bairro:              "Saúde",
	Online:              Online{Senha: "Pi2025"},
}

func RegistrarRotasAPI() {
	http.HandleFunc("/api/paciente", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(paciente)
	})

	http.HandleFunc("/api/instituicao", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(instituicao)
	})
}
