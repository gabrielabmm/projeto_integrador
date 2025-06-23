package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"projeto_integrador_versao_final/internal/models"
)

func InserirExameCitopatologicoAPI(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	EnableCORS(w)
	if r.Method == http.MethodOptions {
		return
	}

	var e models.ExameCitopatologico
	if err := json.NewDecoder(r.Body).Decode(&e); err != nil {
		http.Error(w, "Dados inválidos", http.StatusBadRequest)
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
