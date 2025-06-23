package models

type Appointment struct {
    Date string `json:"date"`
    Ubs  string `json:"ubs"`
    Time string `json:"time"`
}

type ExamCard struct {
    Protocolo   string `json:"protocolo"`
    Status      string `json:"status"` 
    Medico      string `json:"medico"`
    Exame       string `json:"exame"`
    Laboratorio string `json:"laboratorio"`
    Data        string `json:"data"`
    Telefone    string `json:"telefone"`
    Email       string `json:"email"`
    Endereco    string `json:"endereco"`
    Horario     string `json:"horario"`
    Local       string `json:"local"`
    Preparo     string `json:"preparo"`
    Observacoes string `json:"observacoes"`
}
