package pkg

import (
	"context"
	"io"
	"log"
	"os/exec"
	"time"

	"github.com/HomeCube/SurrealAPI/openApI/config"
	"github.com/haguro/elevenlabs-go"
)

func TtsSetup() io.WriteCloser {
	elevenlabs.SetAPIKey(config.ELEVEN_API_KEY)
	elevenlabs.SetTimeout(1 * time.Minute)
	cmd := exec.CommandContext(context.Background(), "mpv", "--no-cache", "--no-terminal", "--", "fd://0")
	pipe, err := cmd.StdinPipe()
	if err != nil {
		log.Fatal(err)
	}

	// Attempt to run the command in a separate process
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}
	return pipe
}

func Speaker(pipe io.WriteCloser, message string) error {
	err := elevenlabs.TextToSpeechStream(
		pipe,
		"pNInz6obpgDQGcFmaJgB",
		elevenlabs.TextToSpeechRequest{
			Text:    message,
			ModelID: "eleven_multilingual_v2",
		})
	if err != nil {
		log.Printf("Got %T error: %q\n", err, err)
		return err
	}
	return nil
}
