package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	dbM "github.com/HomeCube/SurrealAPI/milvius_service/internal/db"
	"github.com/HomeCube/SurrealAPI/milvius_service/model"
)

func CreateDBHandler(w http.ResponseWriter, r *http.Request) {
	var msg string

	// Decodifica il corpo della richiesta in un oggetto Message
	err := json.NewDecoder(r.Body).Decode(&msg)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

	}

	err = dbM.CreateDb(msg)

	if err != nil {

	}

}

func CreateCollectionHandler(w http.ResponseWriter, r *http.Request) {
	var req model.CollectionRequest

	// Decodifica il corpo della richiesta in un oggetto CollectionRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Chiama la funzione CreateCollection con i dati ricevuti
	dbM.CreateCollection(req.CollectionName, req.Description)
	fmt.Println("Created Collection with name", req.CollectionName)
	// Imposta una risposta di successo
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Collection created successfully"))
}

func InsertDataHandler(w http.ResponseWriter, r *http.Request) {
	var req model.DataRequest

	// Decodifica il corpo della richiesta
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Assicurati che ogni embedding abbia la dimensione corretta
	for _, emb := range req.Embeddings {
		if len(emb) != 1536 {
			fmt.Println(len(emb))
			http.Error(w, "Dimensione degli embeddings errata", http.StatusBadRequest)
			return
		}
	}

	// Chiama la funzione InsertData
	dbM.InsertData(req.CollectionName, req.PartitionName, req.UserIds, req.Timestamps, req.ToolIds, req.Embeddings)

	// Risposta di successo
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Data inserted successfully"))
}

func CreatePartitionHandler(w http.ResponseWriter, r *http.Request) {
	var req model.PartitionRequest

	// Decodifica il corpo della richiesta in un oggetto PartitionRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Chiama la funzione CreatePartition con i dati ricevuti
	dbM.CreatePartition(req.CollectionName, req.PartitionName)
	if err != nil {
		log.Printf("failed to create partition: %v", err)
		http.Error(w, "failed to create partition", http.StatusInternalServerError)
		return
	}

	// Imposta una risposta di successo
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Partition created successfully"))
}

func LoadPartitionHandler(w http.ResponseWriter, r *http.Request) {
	var req model.PartitionLoadRequest

	// Decodifica il corpo della richiesta in un oggetto PartitionLoadRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Chiama la funzione LoadPartition con i dati ricevuti
	dbM.LoadPartition(req.CollectionName, req.PartitionNames)

	// Imposta una risposta di successo
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Partitions loaded successfully"))
}

func CreateIndexHandler(w http.ResponseWriter, r *http.Request) {
	var req model.IndexRequest

	// Decodifica il corpo della richiesta in un oggetto IndexRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Chiama la funzione CreateIndex con i dati ricevuti
	dbM.CreateIndex(req.CollectionName, req.FieldName)

	// Imposta una risposta di successo
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Index created successfully"))
}

func LoadCollectionHandler(w http.ResponseWriter, r *http.Request) {
	var req model.CollectionLoadRequest

	// Decodifica il corpo della richiesta in un oggetto CollectionRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Chiama la funzione LoadCollection con i dati ricevuti
	dbM.LoadCollection(req.CollectionName)

	// Imposta una risposta di successo
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Collection loaded successfully"))
}

func ReleaseColelctionHandler(w http.ResponseWriter, r *http.Request) {
	var req model.CollectionLoadRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	dbM.ReleaseCollection(req.CollectionName)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Collection released successfully"))
}
