package db

import (
	"fmt"
	"time"

	"github.com/HomeCube/SurrealAPI/surreal_service/model"
	"github.com/surrealdb/surrealdb.go"
)

var DB *surrealdb.DB

// Connect apre una connessione a SurrealDB
func Connect() error {
	var err error
	DB, err = surrealdb.New("ws://localhost:8000/rpc")
	if err != nil {
		return err
	}

	if _, err := DB.Signin(map[string]interface{}{
		"user": "root",
		"pass": "root",
	}); err != nil {
		return err
	}

	if _, err := DB.Use("test", "test"); err != nil {
		return err
	}

	return nil
}

// Close chiude la connessione al database
func Close() {
	if DB != nil {
		// Assumendo che il client SurrealDB abbia un metodo Close
		DB.Close()
	}
}

func Create(table string, record interface{}) error {
	// Logica per inserire il record nel database
	// Ad esempio, utilizzando un metodo del client SurrealDB
	_, err := DB.Create(table, record)
	if err != nil {
		return err
	}
	return nil
}

func GetMessage(id string) ([]model.Message, error) {
	rawData, err := DB.Select(id)
	if err != nil {
		return nil, fmt.Errorf("failed to select message: %w", err)
	}

	// Prova a deserializzare come slice di model.Message
	var messages []model.Message
	errSlice := surrealdb.Unmarshal(rawData, &messages)

	// Se la deserializzazione come slice fallisce, prova come singolo oggetto
	if errSlice != nil {
		var singleMessage model.Message
		errSingle := surrealdb.Unmarshal(rawData, &singleMessage)
		if errSingle != nil {
			// Entrambe le deserializzazioni sono fallite
			return nil, fmt.Errorf("failed to unmarshal messages: %v; %v", errSlice, errSingle)
		}
		// Se la deserializzazione come singolo oggetto ha successo, aggiungilo alla slice
		messages = []model.Message{singleMessage}
	}

	return messages, nil
}

func Update(table, id string, record interface{}) error {
	_, err := DB.Update(id, record)
	if err != nil {
		return err
	}
	return nil
}

// Delete elimina un record dal database
func Delete(table, id string) error {

	_, err := DB.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

func AddPluginNode(node model.PluginNode) error {
	// Costruisci la query per inserire il nuovo nodo
	query := fmt.Sprintf("CREATE plugin:%s SET tipo = %d, content = '%s'", node.ID, node.Tipo, node.Content)
	// Esegui la query utilizzando il client SurrealDB
	_, err := DB.Query(query, nil)
	if err != nil {
		return fmt.Errorf("failed to add plugin node: %w", err)
	}
	return nil
}

func AddRelation(rel model.Relation, table string) error {
	// Costruisci la query per creare la relazione
	var query string
	if rel.Type == "temporale" {
		query = fmt.Sprintf("RELATE %s->temporale->%s", table+rel.From, table+rel.To)
	} else if rel.Type == "dipendenza" {
		query = fmt.Sprintf("RELATE %s->dipende->%s", table+rel.From, table+rel.To)
	} else {
		return fmt.Errorf("unknown relation type: %s", rel.Type)
	}

	// Esegui la query utilizzando il client SurrealDB
	_, err := DB.Query(query, nil)
	if err != nil {
		return fmt.Errorf("failed to add relation: %w", err)
	}
	return nil
}

// GetLastMsgContainer recupera il container di messaggi dal database
func GetLastMsgContainer(id string) ([]model.LastMsgContainer, error) {
	rawData, err := DB.Select("lastmsg_containers:" + id)
	if err != nil {
		return nil, fmt.Errorf("failed to select message: %w", err)
	}

	// Prova a deserializzare come slice di model.Message
	var messages []model.LastMsgContainer
	errSlice := surrealdb.Unmarshal(rawData, &messages)

	// Se la deserializzazione come slice fallisce, prova come singolo oggetto
	if errSlice != nil {
		var singleMessage model.LastMsgContainer
		errSingle := surrealdb.Unmarshal(rawData, &singleMessage)
		if errSingle != nil {
			// Entrambe le deserializzazioni sono fallite
			return nil, fmt.Errorf("failed to unmarshal messages: %v; %v", errSlice, errSingle)
		}
		// Se la deserializzazione come singolo oggetto ha successo, aggiungilo alla slice
		messages = []model.LastMsgContainer{singleMessage}
	}

	fmt.Println(messages)

	return messages, nil
}

// AddMessageToLastMsgContainer aggiunge un messaggio al container
func AddLastMsgContainer(containerID, text string) error {
	containers, err := GetLastMsgContainer(containerID)
	if err != nil {
		return err
	}

	newMessage := model.LastMsg{
		Text:      text,
		Timestamp: time.Now(),
	}

	var container *model.LastMsgContainer

	// Controlla se è stato trovato un container
	if len(containers) > 0 {
		// Utilizza il primo container trovato
		container = &containers[0]
	} else {
		// Nessun container trovato, quindi ne crea uno nuovo
		container = &model.LastMsgContainer{
			ID:       containerID,
			Messages: []model.LastMsg{},
		}
	}

	// Aggiungi il nuovo messaggio all'inizio dell'array
	container.Messages = append([]model.LastMsg{newMessage}, container.Messages...)

	// Mantieni solo gli ultimi 5 messaggi
	if len(container.Messages) > 5 {
		container.Messages = container.Messages[:5]
	}

	// Aggiorna o crea il container nel database
	if len(containers) > 0 {
		// Aggiorna il container esistente
		_, err = DB.Update("lastmsg_containers:"+containerID, container)
	} else {
		// Crea un nuovo container
		_, err = DB.Create("lastmsg_containers", container)
	}
	if err != nil {
		return fmt.Errorf("failed to save last message container: %w", err)
	}

	return nil
}
