package models

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
