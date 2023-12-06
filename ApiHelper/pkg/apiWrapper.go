package pkg

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type Message struct {
	UserID    string `json:"id,omitempty"`
	SessionID string `json:"session_id"`
	Content   string `json:"content"`
	MsgNum    int    `json:"msg_num"`
}

// CreateIndex crea un indice su una collezione e un campo specificati.
func CreateIndex(collectionName, fieldName string) error {
	// Struct per il corpo della richiesta
	type requestBody struct {
		CollectionName string `json:"collection_name"`
		FieldName      string `json:"field_name"`
	}

	// Crea il corpo della richiesta
	body := requestBody{
		CollectionName: collectionName,
		FieldName:      fieldName,
	}
	bodyBytes, err := json.Marshal(body)
	if err != nil {
		return err
	}

	// Crea la richiesta
	req, err := http.NewRequest("POST", "http://localhost:8082/createIndex", bytes.NewBuffer(bodyBytes))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	// Invia la richiesta
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Gestisce la risposta
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("La richiesta API non è riuscita con il codice di stato: %d", resp.StatusCode)
	}

	return nil
}

// LoadPartition carica partizioni in una collezione specificata.
func LoadPartition(collectionName string, partitionNames []string) error {
	// Struct per il corpo della richiesta
	type requestBody struct {
		CollectionName string   `json:"collection_name"`
		PartitionNames []string `json:"partition_names"`
	}

	// Crea il corpo della richiesta
	body := requestBody{
		CollectionName: collectionName,
		PartitionNames: partitionNames,
	}
	bodyBytes, err := json.Marshal(body)
	if err != nil {
		return err
	}

	// Crea la richiesta
	req, err := http.NewRequest("POST", "http://localhost:8082/loadPartition", bytes.NewBuffer(bodyBytes))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	// Invia la richiesta
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Gestisce la risposta
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("La richiesta API non è riuscita con il codice di stato: %d", resp.StatusCode)
	}

	return nil
}

// CreatePartition crea una partizione in una collezione specificata.
func CreatePartition(collectionName, partitionName string) error {
	// Struct per il corpo della richiesta
	type requestBody struct {
		CollectionName string `json:"collection_name"`
		PartitionName  string `json:"partition_name"`
	}

	// Crea il corpo della richiesta
	body := requestBody{
		CollectionName: collectionName,
		PartitionName:  partitionName,
	}
	bodyBytes, err := json.Marshal(body)
	if err != nil {
		return err
	}

	// Crea la richiesta
	req, err := http.NewRequest("POST", "http://localhost:8082/createPartition", bytes.NewBuffer(bodyBytes))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	// Invia la richiesta
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Gestisce la risposta
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("La richiesta API non è riuscita con il codice di stato: %d", resp.StatusCode)
	}

	return nil
}

// CreateCollection crea una nuova collezione.
func CreateCollection(collectionName, description string) error {
	// Struct per il corpo della richiesta
	type requestBody struct {
		CollectionName string `json:"collection_name"`
		Description    string `json:"description"`
	}

	// Crea il corpo della richiesta
	body := requestBody{
		CollectionName: collectionName,
		Description:    description,
	}
	bodyBytes, err := json.Marshal(body)
	if err != nil {
		return err
	}

	// Crea la richiesta
	req, err := http.NewRequest("POST", "http://localhost:8082/createCollection", bytes.NewBuffer(bodyBytes))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	// Invia la richiesta
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Gestisce la risposta
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("La richiesta API non è riuscita con il codice di stato: %d", resp.StatusCode)
	}

	return nil
}

// AddDB aggiunge un database.
func AddDB(dbName string) error {
	// Crea la richiesta
	req, err := http.NewRequest("POST", "http://localhost:8080/addDB", strings.NewReader(`"`+dbName+`"`))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	// Invia la richiesta
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Gestisce la risposta
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("La richiesta API non è riuscita con il codice di stato: %d", resp.StatusCode)
	}

	return nil
}

// AddLastMessage aggiunge un messaggio all'ultimo messaggio in un container specificato.
func AddLastMessage(containerID, text string) error {
	// Struct per il corpo della richiesta
	type requestBody struct {
		ContainerID string `json:"containerID"`
		Text        string `json:"text"`
	}

	// Crea il corpo della richiesta
	body := requestBody{
		ContainerID: containerID,
		Text:        text,
	}
	bodyBytes, err := json.Marshal(body)
	if err != nil {
		return err
	}

	// Crea la richiesta
	req, err := http.NewRequest("POST", "http://localhost:8080/lastmessage/add", bytes.NewBuffer(bodyBytes))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	// Invia la richiesta
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Gestisce la risposta
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("La richiesta API non è riuscita con il codice di stato: %d", resp.StatusCode)
	}

	return nil
}

// AddNode aggiunge un nodo a un plugin.
func AddNode(id string, tipo int, content string) error {
	// Struct per il corpo della richiesta
	type requestBody struct {
		ID      string `json:"id"`
		Tipo    int    `json:"tipo"`
		Content string `json:"content"`
	}

	// Crea il corpo della richiesta
	body := requestBody{
		ID:      id,
		Tipo:    tipo,
		Content: content,
	}
	bodyBytes, err := json.Marshal(body)
	if err != nil {
		return err
	}

	// Crea la richiesta
	req, err := http.NewRequest("POST", "http://localhost:8080/plugins/nodes", bytes.NewBuffer(bodyBytes))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	// Invia la richiesta
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Gestisce la risposta
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("La richiesta API non è riuscita con il codice di stato: %d", resp.StatusCode)
	}

	return nil
}

// AddRelation crea una relazione tra nodi in un plugin.
func AddRelation(from, to, relationType string) error {
	// Struct per il corpo della richiesta
	type requestBody struct {
		From string `json:"from"`
		To   string `json:"to"`
		Type string `json:"type"`
	}

	// Crea il corpo della richiesta
	body := requestBody{
		From: from,
		To:   to,
		Type: relationType,
	}
	bodyBytes, err := json.Marshal(body)
	if err != nil {
		return err
	}

	// Crea la richiesta
	req, err := http.NewRequest("POST", "http://localhost:8080/plugins/relations", bytes.NewBuffer(bodyBytes))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	// Invia la richiesta
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Gestisce la risposta
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("La richiesta API non è riuscita con il codice di stato: %d", resp.StatusCode)
	}

	return nil
}

// DeleteMessage elimina un messaggio specifico.
func DeleteMessage(messageID string) error {
	// Crea la richiesta
	req, err := http.NewRequest("DELETE", "http://localhost:8080/messages/"+messageID, nil)
	if err != nil {
		return err
	}

	// Invia la richiesta
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Gestisce la risposta
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("La richiesta API non è riuscita con il codice di stato: %d", resp.StatusCode)
	}

	return nil
}

func GetMessage(messageID string) (*Message, error) {
	// Crea la richiesta
	req, err := http.NewRequest("GET", "http://localhost:8080/messages/"+messageID, nil)
	if err != nil {
		return nil, err
	}

	// Invia la richiesta
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Verifica lo stato della risposta
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("La richiesta API non è riuscita con il codice di stato: %d", resp.StatusCode)
	}

	// Decodifica il corpo della risposta
	var message Message
	if err := json.NewDecoder(resp.Body).Decode(&message); err != nil {
		return nil, err
	}

	return &message, nil
}

// UpdateMessage aggiorna un messaggio specifico.
func UpdateMessage(messageID, userID, sessionID, content string) error {
	// Struct per il corpo della richiesta
	type requestBody struct {
		UserID    string `json:"user_id"`
		SessionID string `json:"session_id"`
		Content   string `json:"content"`
	}

	// Crea il corpo della richiesta
	body := requestBody{
		UserID:    userID,
		SessionID: sessionID,
		Content:   content,
	}
	bodyBytes, err := json.Marshal(body)
	if err != nil {
		return err
	}

	// Crea la richiesta
	req, err := http.NewRequest("PUT", "http://localhost:8080/messages/"+messageID, bytes.NewBuffer(bodyBytes))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	// Invia la richiesta
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Gestisce la risposta
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("La richiesta API non è riuscita con il codice di stato: %d", resp.StatusCode)
	}

	return nil
}
