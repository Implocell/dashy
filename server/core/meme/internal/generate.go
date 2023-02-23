package internal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

// main entry point to generate meme?
// flow will be a post endpoint that people can send text to
// send that text on to OpenAi to make it english and then more poetic
// send that poetic text o Dall-E to generate a picture
// grab the first picture and and save it to firebase
// make an entry in the DB with the link from that saved picture and the text

type OpenAiRequest struct {
	Model            string  `json:"model"`
	Temperature      float64 `json:"temperature"`
	MaxTokens        int     `json:"max_tokens"`
	TopP             int     `json:"top_p"`
	FrequencyPenalty int     `json:"frequency_penalty"`
	PresencePenalty  int     `json:"presence_penalty"`
	Prompt           string  `json:"prompt"`
}

type OpenAiResponse struct {
	ID      string `json:"id"`
	Object  string `json:"object"`
	Created int    `json:"created"`
	Model   string `json:"model"`
	Choices []struct {
		Text         string      `json:"text"`
		Index        int         `json:"index"`
		Logprobs     interface{} `json:"logprobs"`
		FinishReason string      `json:"finish_reason"`
	} `json:"choices"`
	Usage struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
}

func GenerateMemeByText(memeText string) error {
	r := OpenAiRequest{
		Model:            "text-davinci-003",
		Temperature:      0.3,
		MaxTokens:        800,
		TopP:             1,
		FrequencyPenalty: 0,
		PresencePenalty:  0,
		Prompt:           fmt.Sprintf("Translate the following to english from norwegian and create a poem based on the translated text: %s", memeText),
	}

	json, err := json.Marshal(&r)
	if err != nil {
		return err
	}

	buff := bytes.NewBuffer(json)

	url := "https://api.openai.com/v1/completions"
	req, err := http.NewRequest("POST", url, buff)
	if err != nil {
		return err
	}

	bearer := fmt.Sprintf("Bearer %s", os.Getenv("OPEN_AI_API_KEY"))
	req.Header.Add("Authorization", bearer)
	req.Header.Add("content-type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error on response.\n[ERROR] - %s", err)
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	log.Printf("Got text from AI: %s", body)
	return nil
}

// url := "https://api.globalcode.com.br/v1/publico/eventos"

// // Create a Bearer string by appending string access token
// var bearer = "Bearer " + <ACCESS TOKEN HERE>

// // Create a new request using http
// req, err := http.NewRequest("GET", url, nil)

// // add authorization header to the req
// req.Header.Add("Authorization", bearer)

// // Send req using http Client
// client := &http.Client{}
// resp, err := client.Do(req)
// if err != nil {
//     log.Println("Error on response.\n[ERROR] -", err)
// }
// defer resp.Body.Close()
