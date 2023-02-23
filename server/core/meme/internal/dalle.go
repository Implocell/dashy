package internal

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

type DallERequest struct {
	Prompt string `json:"prompt"`
	N      int    `json:"n"`
	Size   string `json:"size"`
}

func NewDallERequest(prompt string) *DallERequest {
	return &DallERequest{
		Size:   "256x256",
		N:      1,
		Prompt: prompt,
	}
}

type DallEResponse struct {
	Created int `json:"created"`
	Data    []struct {
		URL string `json:"url"`
	} `json:"data"`
}

func GetImageFromPoem(poemText string) (string, error) {
	r := NewDallERequest(poemText)
	res, err := sendDallEOpenAiRequest(r)
	if err != nil {
		return "", err
	}
	return res.Data[0].URL, nil
}

func mapRequestToDallEResponse(body []byte) (*DallEResponse, error) {
	var res DallEResponse
	err := json.Unmarshal(body, &res)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &res, nil
}

func sendDallEOpenAiRequest(openAiRequest *DallERequest) (*DallEResponse, error) {
	url := "https://api.openai.com/v1/images/generations"
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

	DallEResponse, err := mapRequestToDallEResponse(openAiResponse)
	if err != nil {
		return nil, err
	}

	return DallEResponse, nil
}
