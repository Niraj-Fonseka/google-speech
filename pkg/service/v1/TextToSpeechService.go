package v1

import (
	"bytes"
	"fmt"
	v1 "google-speech/pkg/api/v1"
	"io/ioutil"
	"log"
	"os/exec"

	texttospeech "cloud.google.com/go/texttospeech/apiv1"
	"golang.org/x/net/context"
	texttospeechpb "google.golang.org/genproto/googleapis/cloud/texttospeech/v1"
)

// Server represents the gRPC server
type TextToSpeechServiceServer struct {
}

func NewTextToSpeechServiceServer() v1.TextToSpeechServer {
	return &TextToSpeechServiceServer{}
}

func (s *TextToSpeechServiceServer) GenerateSpeech(ctx context.Context, in *v1.TextToSpeechMessage) (*v1.TextToSpeechMessage, error) {

	// Instantiates a client.
	client, err := texttospeech.NewClient(ctx)
	if err != nil {
		log.Fatal(err)
	}

	req := texttospeechpb.SynthesizeSpeechRequest{
		Input: &texttospeechpb.SynthesisInput{
			InputSource: &texttospeechpb.SynthesisInput_Text{Text: in.Data},
		},
		// Build the voice request, select the language code ("en-US") and the SSML
		// voice gender ("neutral").
		// need to set the gender
		Voice: &texttospeechpb.VoiceSelectionParams{
			LanguageCode: "en-US",
			SsmlGender:   texttospeechpb.SsmlVoiceGender_NEUTRAL,
		},
		// Select the type of audio file you want returned.
		AudioConfig: &texttospeechpb.AudioConfig{
			AudioEncoding: texttospeechpb.AudioEncoding_MP3,
		},
	}

	resp, err := client.SynthesizeSpeech(ctx, &req)
	if err != nil {
		log.Fatal(err)
	}

	// The resp's AudioContent is binary.
	filename := "output.mp3"
	err = ioutil.WriteFile(filename, resp.AudioContent, 0644)
	if err != nil {
		log.Fatal(err)
	}

	if in.Play {
		fmt.Println("Playing")
		cmd := exec.Command("play", filename)

		var out bytes.Buffer
		cmd.Stdout = &out
		err := cmd.Run()
		if err != nil {
			log.Fatal(err)
		} else {
			fmt.Println(out.String())
		}
	}

	fmt.Printf("Audio content written to file: %v\n", filename)

	return &v1.TextToSpeechMessage{Response: "Audio Generated"}, nil
}
