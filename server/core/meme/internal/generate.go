package internal

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func sendRequest(req *http.Request) ([]byte, error) {
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func setOpenAiRequestHeaders(req *http.Request) {
	bearer := fmt.Sprintf("Bearer %s", os.Getenv("OPEN_AI_API_KEY"))
	req.Header.Add("Authorization", bearer)
	req.Header.Add("content-type", "application/json")
}
