package pkg

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"strings"

	"github.com/HomeCube/SurrealAPI/openApI/config"

	openai "github.com/sashabaranov/go-openai"
	"github.com/sashabaranov/go-openai/jsonschema"
)

type ToolCreator struct {
	Tools []openai.Tool
}

type Asker struct {
	Client openai.Client
	Model  string
	Req    openai.ChatCompletionRequest
	Stream openai.ChatCompletionStream
}

type Embedder struct {
	Client openai.Client
	Model  string
	req    openai.EmbeddingRequest
}

func (e *Embedder) CreateEmbedder() {
	e.Client = *openai.NewClient(config.API_KEY)

}

func (e *Embedder) CreateEmbeddingRequest(prompt string) {
	e.req = openai.EmbeddingRequest{
		Input: []string{prompt},
		Model: openai.AdaEmbeddingV2,
	}
}

func (t *ToolCreator) CreateTool(params jsonschema.Definition, name string, description string) {

	f := openai.FunctionDefinition{
		Name:        name,
		Description: description,
		Parameters:  params,
	}

	tool := openai.Tool{
		Type:     openai.ToolTypeFunction,
		Function: f,
	}

	t.Tools = append(t.Tools, tool)

}

func (e *Embedder) GetEmbedding() openai.EmbeddingResponse {
	targetResponse, err := e.Client.CreateEmbeddings(context.Background(), e.req)
	if err != nil {
		log.Fatal("Error creating target embedding:", err)
	}
	return targetResponse
}

func (a *Asker) CreateClient() {
	a.Client = *openai.NewClient(config.API_KEY)

}

func (a *Asker) CreateCompletionRequest(prompt string, tools []openai.Tool) openai.ChatCompletionRequest {
	a.Req = openai.ChatCompletionRequest{
		Model:     openai.GPT3Dot5Turbo,
		MaxTokens: 800,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleUser,
				Content: prompt,
			},
		},
		Tools:  tools,
		Stream: true,
	}

	return a.Req
}

func (a *Asker) GetChatCompletitionStream() {
	stream, err := a.Client.CreateChatCompletionStream(context.Background(), a.Req)
	if err != nil {
		fmt.Printf("ChatCompletionStream error: %v\n", err)
		return
	}
	if stream != nil {
		a.Stream = *stream
	}
}

const (
	WordCountThreshold = 20 // Numero di parole da leggere prima di inviarle a Speaker
)

// Delimitatori di fine frase
var sentenceDelimiters = []rune{'.', ';', '?', '!', '\n'}

func isSentenceEnd(buffer string) bool {
	for _, delimiter := range sentenceDelimiters {
		if strings.ContainsRune(buffer, delimiter) {
			return true
		}
	}
	return false
}

func (a *Asker) PrintResults() {
	pipe := TtsSetup()
	fmt.Printf("Stream response: ")
	defer pipe.Close() // Chiude il pipe alla fine

	var buffer string

	for {
		response, err := a.Stream.Recv()
		if errors.Is(err, io.EOF) {
			if buffer != "" {
				Speaker(pipe, buffer) // Gestisce le parole rimanenti nel buffer
			}
			fmt.Println("\nStream finished")
			return
		}

		if err != nil {
			fmt.Printf("\nStream error: %v\n", err)
			return
		}

		buffer += response.Choices[0].Delta.Content

		// Processa il buffer se termina con un delimitatore di fine frase o ha abbastanza parole
		if isSentenceEnd(buffer) || len(strings.Fields(buffer)) >= WordCountThreshold {
			Speaker(pipe, buffer)
			buffer = "" // Resetta il buffer dopo aver inviato il testo a Speaker
		}

		fmt.Printf(response.Choices[0].Delta.Content)
	}
}
