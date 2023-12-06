package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/HomeCube/SurrealAPI/surreal_service/internal/db"
	"github.com/HomeCube/SurrealAPI/surreal_service/model"
	"github.com/go-chi/chi/v5"
	// altri import necessari
)

// CreateMessageHandler gestisce la creazione di nuovi messaggi
func CreateMessageHandler(w http.ResponseWriter, r *http.Request) {
	var msg model.Message

	// Decodifica il corpo della richiesta in un oggetto Message
	err := json.NewDecoder(r.Body).Decode(&msg)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Qui puoi impostare MsgNum o eseguire qualsiasi altra logica necessaria

	// Creazione del messaggio nel database
	err = CreateMessage(msg)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Imposta lo stato di risposta e, se necessario, invia dati di risposta
	w.WriteHeader(http.StatusCreated)
	// Opzionale: invia indietro il messaggio creato o un messaggio di conferma
	json.NewEncoder(w).Encode(msg)
}

func GetMessageHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id") // Estrae l'ID dal percorso della richiesta
	message, err := GetMessage(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(message)
}

// UpdateMessageHandler gestisce l'aggiornamento di un messaggio specifico
func UpdateMessageHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	var msg model.Message
	err := json.NewDecoder(r.Body).Decode(&msg)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = UpdateMessage(id, msg)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent) // Risposta senza contenuto per una PUT riuscita
}

// DeleteMessageHandler gestisce la richiesta DELETE per eliminare un messaggio
func DeleteMessageHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	err := DeleteMessage(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent) // Risposta senza contenuto per una DELETE riuscita
}

func AddPluginNodeHandler(w http.ResponseWriter, r *http.Request) {
	var node model.PluginNode
	if err := json.NewDecoder(r.Body).Decode(&node); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := db.AddPluginNode(node); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println(node)
	// Resto dell'implementazione
}

func AddRelationHandler(w http.ResponseWriter, r *http.Request) {
	var rel model.Relation
	if err := json.NewDecoder(r.Body).Decode(&rel); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := db.AddRelation(rel, "plugin:"); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Resto dell'implementazione
}

func AddLastMessageHandler(w http.ResponseWriter, r *http.Request) {
	var req model.AddMessageRequest

	// Decodifica il corpo della richiesta
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Aggiungi il messaggio al container
	err = db.AddLastMsgContainer(req.ContainerID, req.Text)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Invia una risposta di successo
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"status": "success"})
}
