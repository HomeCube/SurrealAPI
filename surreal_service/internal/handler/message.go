package handler

import (
	"fmt"

	"github.com/HomeCube/SurrealAPI/surreal_service/internal/db"
	"github.com/HomeCube/SurrealAPI/surreal_service/model"
)

func CreateMessage(m model.Message) error {
	err := db.Create("message2", m)
	if err != nil {
		// gestire l'errore
		return err
	}
	return nil
}

func GetMessage(id string) ([]model.Message, error) {
	// Chiama semplicemente la funzione GetMessage dal package handler
	return db.GetMessage(id)
}

// UpdateMessage aggiorna un messaggio esistente
func UpdateMessage(id string, msg model.Message) error {
	err := db.Update("message", id, msg)
	if err != nil {
		return fmt.Errorf("failed to update message: %w", err)
	}
	return nil
}

// DeleteMessage elimina un messaggio dal database
func DeleteMessage(id string) error {
	err := db.Delete("message", id)
	if err != nil {
		return fmt.Errorf("failed to delete message: %w", err)
	}
	return nil
}

// Handler per aggiungere un nodo
