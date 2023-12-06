package handler

import (
	"testing"

	"reflect"

	"github.com/HomeCube/SurrealAPI/surreal_service/internal/db"
	"github.com/HomeCube/SurrealAPI/surreal_service/model"
)

func TestCRUDMessage(t *testing.T) {
	// Crea un messaggio di test
	msg := model.Message{
		UserID:    "user123",
		SessionID: "session456",
		Content:   "Hello, World!",
		MsgNum:    1,
	}
	err := db.Connect()
	if err != nil {
		t.Fatalf("Failed to connect to the database: %v", err)
	}
	// Create
	CreateMessage(msg)
	if err != nil {
		t.Fatalf("CreateMessage failed: %v", err)
	}

	retrievedMsg, err := GetMessage("message2:" + msg.UserID)
	if err != nil {
		t.Fatalf("GetMessage failed: %v", err)
	}

	msg.UserID = "message2:" + msg.UserID

	if !reflect.DeepEqual(msg, retrievedMsg[0]) {
		t.Errorf("Expected retrieved message to be %+v, got %+v", msg, retrievedMsg)
	}

	// Update
	updatedContent := "Updated Content"
	msg.Content = updatedContent
	err = UpdateMessage(msg.UserID, msg)
	if err != nil {
		t.Fatalf("UpdateMessage failed: %v", err)
	}

	// Re-Read per verificare l'update
	updatedMsg, err := GetMessage(msg.UserID)
	if err != nil {
		t.Fatalf("GetMessage (post-update) failed: %v", err)
	}
	if updatedMsg[0].Content != updatedContent {
		t.Errorf("Expected updated message content to be %v, got %v", updatedContent, updatedMsg[0].Content)
	}

	// Delete
	err = DeleteMessage(msg.UserID)
	if err != nil {
		t.Fatalf("DeleteMessage failed: %v", err)
	}

	// Verifica che il messaggio sia stato effettivamente eliminato
	_, err = GetMessage(msg.UserID)
	if err == nil {
		t.Errorf("Expected GetMessage to fail after delete, but it succeeded")
	}

}
