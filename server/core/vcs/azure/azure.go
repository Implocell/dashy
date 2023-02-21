package azure

import (
	b64 "encoding/base64"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func GetActivePullRequests() ([]byte, error) {
	organization := os.Getenv("AZURE_ORGANIZATION")
	project := os.Getenv("AZURE_PROJECT")
	url := fmt.Sprintf("https://dev.azure.com/{%s}/{%s}/_apis/git/pullrequests?api-version=7.1-preview.1&searchCriteria.status=active", organization, project)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	addAuthorizationToRequest(req)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return body, nil
}

func addAuthorizationToRequest(req *http.Request) {
	authorizationToken := ""
	bearer := fmt.Sprintf("Bearer %s", b64.StdEncoding.EncodeToString([]byte(authorizationToken)))
	req.Header.Add("Authorization", bearer)
}
