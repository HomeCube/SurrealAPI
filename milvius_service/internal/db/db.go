package dbM

import (
	"context"
	"fmt"
	"log"

	"github.com/HomeCube/SurrealAPI/milvius_service/model"
	"github.com/milvus-io/milvus-sdk-go/v2/client"
	"github.com/milvus-io/milvus-sdk-go/v2/entity"
)

var DB client.Client

func Connect(address string) (client.Client, error) {
	var err error
	DB, err = client.NewGrpcClient(context.Background(), address)
	if err != nil {
		log.Fatal("failed to connect to Milvus:", err.Error())
	} else {
		log.Println("Connected To VDB")
	}
	return DB, err
}

func CreateDb(name string) error {
	fmt.Println(name)
	err := DB.CreateDatabase(context.Background(), name)
	if err != nil {
		log.Println("Could not create DB")
		return err
	}
	return nil
}

func CreateSchema(collection_name string, description string) entity.Schema {
	schema := &entity.Schema{
		CollectionName: collection_name,
		Description:    description,
		Fields:         model.HotelFields,
	}
	return *schema
}

func CreateCollection(collection_name string, description string) {
	schema := CreateSchema(collection_name, description)
	err := DB.CreateCollection(context.Background(), &schema, 1)
	if err != nil {
		log.Fatal("failed to create collection:", err.Error())
	}
}

func CreatePartition(collection_name string, partition_name string) {
	err := DB.CreatePartition(
		context.Background(), // ctx
		collection_name,      // CollectionName
		partition_name,       // partitionName
	)
	if err != nil {
		log.Fatal("failed to create partition:", err.Error())
	}
}

func InsertData(coll_name string, part_name string, userIds []int64, timestamps []int16, toolIds []int8, embeddings [][]float32) {
	userIdColumn, timestampColumn, toolIdColumn, embeddingColumn, _ := model.CreateVDBDataStructure(userIds, timestamps, toolIds, embeddings)

	_, err := DB.Insert(
		context.Background(), // ctx
		coll_name,            // CollectionName
		part_name,            // partitionName
		&userIdColumn,        // columnarData
		&timestampColumn,     // columnarData
		&toolIdColumn,
		&embeddingColumn, // columnarData
	)
	if err != nil {
		log.Fatal("failed to insert data:", err.Error())
	}
}

func LoadPartition(collection_name string, partition_names []string) {

	// Assicurati che 'partition_names' non sia vuoto
	if len(partition_names) == 0 {
		log.Println("Nessun nome di partizione fornito")
		return
	}

	err := DB.LoadPartitions(
		context.Background(), // ctx
		collection_name,      // CollectionName
		partition_names,      // partitionNames
		false,
	)
	if err != nil {
		log.Fatalf("failed to load partitions: %v", err)
	}
}

func LoadCollection(collection_name string) {
	err := DB.LoadCollection(
		context.Background(),
		collection_name,
		false,
	)

	if err != nil {
		log.Fatalf("failed to load colelction: %v", err)
	}
}

func CreateIndex(collection_name string, field_name string) {
	idx, _ := model.Build_index()
	DB.CreateIndex(
		context.Background(), // ctx
		collection_name,      // CollectionName
		field_name,           // fieldName
		&idx,                 // entity.Index
		false,                // async
	)
}

func ReleaseCollection(collection_name string) {
	err := DB.ReleaseCollection(context.Background(), collection_name)
	if err != nil {
		log.Fatalf("failed to release colelction: %v", err)
	}
}
