package model

import (
	"fmt"
	"log"

	"github.com/milvus-io/milvus-sdk-go/v2/entity"
)

var HotelFields = []*entity.Field{
	{
		Name:       "user_id",
		DataType:   entity.FieldTypeInt64,
		PrimaryKey: true,
		AutoID:     false,
	},
	{
		Name:       "timestamp",
		DataType:   entity.FieldTypeInt16,
		PrimaryKey: false,
		AutoID:     false,
	},
	{
		Name:       "tool_id",
		DataType:   entity.FieldTypeInt8,
		PrimaryKey: false,
		AutoID:     false,
	},
	{
		Name:     "embeddings",
		DataType: entity.FieldTypeFloatVector,
		TypeParams: map[string]string{
			"dim": "1536",
		},
	},
}

type CollectionRequest struct {
	CollectionName string `json:"collection_name"`
	Description    string `json:"description"`
}

type DataRequest struct {
	CollectionName string      `json:"collection_name"`
	PartitionName  string      `json:"partition_name"`
	UserIds        []int64     `json:"user_ids"`
	Timestamps     []int16     `json:"timestamps"`
	ToolIds        []int8      `json:"tool_ids"`
	Embeddings     [][]float32 `json:"embeddings"`
}

type PartitionRequest struct {
	CollectionName string `json:"collection_name"`
	PartitionName  string `json:"partition_name"`
}

type PartitionLoadRequest struct {
	CollectionName string   `json:"collection_name"`
	PartitionNames []string `json:"partition_names"`
}

type IndexRequest struct {
	CollectionName string `json:"collection_name"`
	FieldName      string `json:"field_name"`
}

type CollectionLoadRequest struct {
	CollectionName string `json:"collection_name"`
}

func CreateVDBDataStructure(userIds []int64, timestamps []int16, toolIds []int8, embeddings [][]float32) (entity.ColumnInt64, entity.ColumnInt16, entity.ColumnInt8, entity.ColumnFloatVector, error) {
	// Verifica che le lunghezze delle slice siano coerenti
	var userIdColumn entity.ColumnInt64
	var timestampColumn entity.ColumnInt16
	var toolIdColumn entity.ColumnInt8
	var embeddingColumn entity.ColumnFloatVector

	if len(userIds) != len(timestamps) || len(userIds) != len(toolIds) || len(userIds) != len(embeddings) {
		return userIdColumn, timestampColumn, toolIdColumn, embeddingColumn, fmt.Errorf("le lunghezze delle slice non corrispondono")
	}

	const dim = 1536
	for _, v := range embeddings {
		if len(v) != dim {
			log.Printf("Embedding di lunghezza errata ricevuto: atteso %d, ricevuto %d\n", dim, len(v))
			var embeddingColumn entity.ColumnFloatVector
			return userIdColumn, timestampColumn, toolIdColumn, embeddingColumn, fmt.Errorf("dimensione degli embeddings errata: attesa %d, ricevuta %d", dim, len(v))
		}
	}

	// Creazione delle colonne per l'inserimento nel database
	userIdColumn = *entity.NewColumnInt64("user_id", userIds)
	timestampColumn = *entity.NewColumnInt16("timestamp", timestamps)
	toolIdColumn = *entity.NewColumnInt8("tool_id", toolIds)
	embeddingColumn = *entity.NewColumnFloatVector("embeddings", dim, embeddings)

	return userIdColumn, timestampColumn, toolIdColumn, embeddingColumn, nil
}

func Build_index() (entity.IndexIvfFlat, error) {
	idx, err := entity.NewIndexIvfFlat( // NewIndex func
		entity.L2, // metricType
		1024,      // ConstructParams
	)
	if err != nil {
		log.Fatal("fail to create ivf flat index parameter:", err.Error())
		return *idx, err
	}
	return *idx, nil

}
