package main

import (
	"encoding/json"
	"net/http"
	"time"
)

type CalendarDay struct {
	Day        int  `json:"day"`
	IsToday    bool `json:"is_today"`
	IsSelected bool `json:"is_selected"`
	IsEmpty    bool `json:"is_empty"`
}

type Appointment struct {
	Date string `json:"date"`
	Ubs  string `json:"ubs"`
	Time string `json:"time"`
}

type ExamCard struct {
	Protocolo   string `json:"protocolo"`
	Status      string `json:"status"` // "confirmed" ou "cancelled"
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

func apiCalendario(w http.ResponseWriter, r *http.Request) {
	now := time.Now()
	year, month, day := now.Date()

	firstDay := time.Date(year, month, 1, 0, 0, 0, 0, time.Local)
	firstWeekday := int(firstDay.Weekday())
	lastDay := time.Date(year, month+1, 0, 0, 0, 0, 0, time.Local).Day()

	var dias []CalendarDay
	for i := 0; i < firstWeekday; i++ {
		dias = append(dias, CalendarDay{IsEmpty: true})
	}

	for d := 1; d <= lastDay; d++ {
		dias = append(dias, CalendarDay{
			Day:        d,
			IsToday:    d == day,
			IsSelected: d == day,
		})
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(dias)
}

func apiConfirmar(w http.ResponseWriter, r *http.Request) {
	var appt Appointment
	if err := json.NewDecoder(r.Body).Decode(&appt); err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	resp := map[string]string{
		"mensagem": "Agendamento confirmado com sucesso!",
		"data":     appt.Date,
		"ubs":      appt.Ubs,
		"horario":  appt.Time,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func apiExames(w http.ResponseWriter, r *http.Request) {
	exames := []ExamCard{
		{
			Protocolo:   "091207",
			Status:      "confirmed",
			Medico:      "Ana Lopes",
			Exame:       "Preventivo",
			Laboratorio: "LabHealth",
			Data:        "12/06/2024",
			Telefone:    "(62) 3212-3456",
			Email:       "contato@labhealth.com",
			Endereco:    "Rua da Amostra, 123 - Centro, Goiânia - GO",
			Horario:     "10:00",
			Local:       "UBS Centro",
			Preparo:     "Jejum de 8 horas e evitar consumo de álcool 24h antes.",
			Observacoes: "Chegar com 15 minutos de antecedência. Trazer documento de identificação com foto e pedido médico.",
		},
		{
			Protocolo:   "090909",
			Status:      "confirmed",
			Medico:      "Joana Sousa",
			Exame:       "Ultrassom",
			Laboratorio: "SaúdeIdeal",
			Data:        "14/05/2024",
			Telefone:    "(62) 3456-7890",
			Email:       "atendimento@saudeideal.com",
			Endereco:    "Avenida da Saúde, 456 - Setor Bueno, Goiânia - GO",
			Horario:     "14:00",
			Local:       "Clínica Particular XYZ",
			Preparo:     "Não é necessário preparo específico.",
			Observacoes: "Chegar no horário. Apresentar encaminhamento.",
		},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(exames)
}

func RegistrarRotasAgendar() {
	http.HandleFunc("/api/calendario", apiCalendario)
	http.HandleFunc("/api/confirmar", apiConfirmar)
	http.HandleFunc("/api/exames", apiExames)
}
