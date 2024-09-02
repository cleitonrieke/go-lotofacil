package handlers

import (
	"backlotofacil/models"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

func (h Handler) ResumoNumerosLotoFacil(w http.ResponseWriter, r *http.Request) {
	var participantes []models.Participante
	h.DB.Find(&participantes)

	tabNumeros := make(map[string]int, 25)
	for i := 0; i < 25; i++ {
		tabNumeros[strconv.Itoa(i+1)] = 0
	}

	for _, participante := range participantes {
		numeros := strings.Split(participante.NumerosSelecionados, ",")
		for i := 0; i < len(numeros); i++ {
			tabNumeros[numeros[i]]++
		}
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tabNumeros)
}
