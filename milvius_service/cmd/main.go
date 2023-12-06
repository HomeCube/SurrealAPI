package main

import (
	"log"
	"net/http"

	dbM "github.com/HomeCube/SurrealAPI/milvius_service/internal/db"
	"github.com/HomeCube/SurrealAPI/milvius_service/internal/handler"
	"github.com/go-chi/chi"
)

func main() {
	_, err := dbM.Connect("localhost:19530")
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	r := chi.NewRouter()
	r.Post("/addDB", handler.CreateDBHandler)
	r.Post("/createCollection", handler.CreateCollectionHandler)
	r.Post("/insertData", handler.InsertDataHandler)
	r.Post("/createPartition", handler.CreatePartitionHandler)
	r.Post("/loadPartition", handler.LoadPartitionHandler)
	r.Post("/loadCollection", handler.LoadCollectionHandler)
	r.Post("/createIndex", handler.CreateIndexHandler)
	r.Post("/releaseCollection", handler.ReleaseColelctionHandler)

	log.Println("Starting server on :8082")
	err = http.ListenAndServe(":8082", r)
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
