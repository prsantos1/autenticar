package server

import (
	"encoding/json"
	"net/http"
)

// User representa a estrutura dos dados de login
type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// loginHandler lida com requisições de login
func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Erro ao decodificar JSON", http.StatusBadRequest)
		return
	}

	// Verifica as credenciais (exemplo simples)
	if user.Username == "admin" && user.Password == "password" {
		response := map[string]bool{"success": true}
		json.NewEncoder(w).Encode(response)
	} else {
		response := map[string]bool{"success": false}
		json.NewEncoder(w).Encode(response)
	}
}
