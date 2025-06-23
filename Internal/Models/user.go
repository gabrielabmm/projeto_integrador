package models

type Endereco struct {
	CEP        string `json:"cep"`
	Cidade     string `json:"cidade"`
	Logradouro string `json:"logradouro"`
}

type Online struct {
	Email string `json:"email"`
	Login string `json:"login,omitempty"`
	Senha string `json:"senha,omitempty"`
}

type Paciente struct {
	NomeCompleto   string   `json:"nomeCompleto"`
	PrimeiroNome   string   `json:"primeiroNome"`
	DataNascimento string   `json:"dataNascimento"`
	Endereco       Endereco `json:"endereco"`
	CPF            string   `json:"cpf"`
	CartaoSUS      string   `json:"cartaoSUS"`
	Online         Online   `json:"online"`
	RG             string   `json:"rg"`
	Sexo           string   `json:"sexo"`
	Celular        string   `json:"celular"`
	Telefone       string   `json:"telefone"`
	Escolaridade   string   `json:"escolaridade"`
	Observacao     string   `json:"observacao"`
}

type Instituicao struct {
	NomeInstituicao     string `json:"nomeInstituicao"`
	CNES                string `json:"cnes"`
	CRMCorenResponsavel string `json:"crmCorenResponsavel"`
	CNPJ                string `json:"cnpj"`
	Email               string `json:"email"`
	Celular             string `json:"celular"`
	Telefone            string `json:"telefone"`
	CEP                 string `json:"cep"`
	Cidade              string `json:"cidade"`
	Endereco            string `json:"endereco"`
	Numero              string `json:"numero"`
	Complemento         string `json:"complemento"`
	Bairro              string `json:"bairro"`
	Online              Online `json:"online"`
}
