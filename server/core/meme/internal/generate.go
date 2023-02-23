package internal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

// main entry point to generate meme?
// flow will be a post endpoint that people can send text to
// send that text on to OpenAi to make it english and then more poetic
// send that poetic text o Dall-E to generate a picture
// grab the first picture and and save it to firebase
// make an entry in the DB with the link from that saved picture and the text

type DavinciRequest struct {
	Model            string  `json:"model"`
	Temperature      float64 `json:"temperature"`
	MaxTokens        int     `json:"max_tokens"`
	TopP             int     `json:"top_p"`
	FrequencyPenalty int     `json:"frequency_penalty"`
	PresencePenalty  int     `json:"presence_penalty"`
	Prompt           string  `json:"prompt"`
}

type DavinciResponse struct {
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

type DaliRequest struct {
	Prompt string `json:"promt"`
	N      int    `json:"n"`
	Size   string `json:"size"`
}

type DaliAiResponse struct {
	Created int `json:"created"`
	Data    []struct {
		URL string `json:"url"`
	} `json:"data"`
}

func GetPoem(memeText string) (string, error) {
	r := DavinciRequest{
		Model:            "text-davinci-003",
		Temperature:      0.3,
		MaxTokens:        800,
		TopP:             1,
		FrequencyPenalty: 0,
		PresencePenalty:  0,
		Prompt:           fmt.Sprintf("Translate the following to english from norwegian and create a poem based on the translated text: %s", memeText),
	}
	res, err := sendDavinciOpenAiRequest(&r, "https://api.openai.com/v1/completions")
	if err != nil {
		return "", err
	}
	fmt.Printf("\nres: %v", res)
	return res.Choices[0].Text, nil
}

func GetImageFromPoem(poemText string) (string, error) {
	r := DaliRequest{
		Size:   "256x256",
		N:      1,
		Prompt: poemText,
	}
	res, err := sendDaliOpenAiRequest(&r, "https://api.openai.com/v1/images/generations")
	if err != nil {
		return "", err
	}
	return res.Data[0].URL, nil
}

func sendDavinciOpenAiRequest(openAiRequest *DavinciRequest, url string) (*DavinciResponse, error) {
	openAiJson, err := json.Marshal(openAiRequest)
	if err != nil {
		return nil, err
	}
	buff := bytes.NewBuffer(openAiJson)

	req, err := http.NewRequest("POST", url, buff)
	if err != nil {
		return nil, err
	}
	setOpenAiRequestHeaders(req)
	openAiResponse, err := sendRequest(req)
	if err != nil {
		return nil, err
	}

	davinciResponse, err := mapRequestToDavinciResponse(openAiResponse)
	if err != nil {
		return nil, err
	}

	return davinciResponse, nil
}

func sendDaliOpenAiRequest(openAiRequest *DaliRequest, url string) (*DaliAiResponse, error) {
	openAiJson, err := json.Marshal(openAiRequest)
	if err != nil {
		return nil, err
	}
	buff := bytes.NewBuffer(openAiJson)

	req, err := http.NewRequest("POST", url, buff)
	if err != nil {
		return nil, err
	}
	setOpenAiRequestHeaders(req)
	openAiResponse, err := sendRequest(req)
	if err != nil {
		return nil, err
	}

	daliResponse, err := mapRequestToDaliResponse(openAiResponse)
	if err != nil {
		return nil, err
	}

	return daliResponse, nil
}

func setOpenAiRequestHeaders(req *http.Request) error {
	bearer := fmt.Sprintf("Bearer %s", os.Getenv("OPEN_AI_API_KEY"))
	req.Header.Add("Authorization", bearer)
	req.Header.Add("content-type", "application/json")

	return nil
}

func sendRequest(req *http.Request) ([]byte, error) {
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error on response.\n[ERROR] - %s", err)
		return nil, err
	}
	defer resp.Body.Close()

	fmt.Printf("\nresp: %v\n", resp)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	fmt.Printf("\nbody: %v\n", body)
	return body, nil
}

func mapRequestToDavinciResponse(body []byte) (*DavinciResponse, error) {
	var res DavinciResponse
	err := json.Unmarshal(body, &res)
	if err != nil {
		return nil, err
	}
	return &res, nil

}
func mapRequestToDaliResponse(body []byte) (*DaliAiResponse, error) {
	var res DaliAiResponse
	err := json.Unmarshal(body, &res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}
