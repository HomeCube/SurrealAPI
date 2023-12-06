package main

import (
	"log"
	"net/http"

	"github.com/HomeCube/SurrealAPI/surreal_service/internal/db"
	"github.com/HomeCube/SurrealAPI/surreal_service/internal/handler"
	"github.com/go-chi/chi/v5"
)

func main() {

	err := db.Connect()
	if err != nil {

	}
	r := chi.NewRouter()

	// Definizione delle route
	r.Post("/messages", handler.CreateMessageHandler)        // Crea un nuovo messaggio
	r.Get("/messages/{id}", handler.GetMessageHandler)       // Legge un messaggio specifico
	r.Put("/messages/{id}", handler.UpdateMessageHandler)    // Aggiorna un messaggio specifico
	r.Delete("/messages/{id}", handler.DeleteMessageHandler) // Elimina un messaggio specifico
	r.Post("/plugins/nodes", handler.AddPluginNodeHandler)
	r.Post("/plugins/relations", handler.AddRelationHandler)
	r.Post("/lastmessages/add", handler.AddLastMessageHandler)
	// Avvio del server
	log.Println("Starting server on :8080")
	err = http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
	// Aggiungi altre route qui...

}
