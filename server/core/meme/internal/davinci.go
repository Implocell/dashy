package internal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

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

func GetPoem(memeText string) (string, error) {
	r := DavinciRequest{
		Model:            "text-davinci-003",
		Temperature:      0.3,
		MaxTokens:        200,
		TopP:             1,
		FrequencyPenalty: 0,
		PresencePenalty:  0,
		Prompt:           fmt.Sprintf("Translate the following to english from norwegian and create a poem based on the translated text: %s", memeText),
	}
	res, err := sendDavinciOpenAiRequest(&r, "https://api.openai.com/v1/completions")
	if err != nil {
		log.Println(err)
		return "", err
	}
	return res.Choices[0].Text, nil
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

func mapRequestToDavinciResponse(body []byte) (*DavinciResponse, error) {
	var res DavinciResponse
	err := json.Unmarshal(body, &res)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &res, nil

}
