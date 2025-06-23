package models

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
