package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"chat-gpt/internal/descriptions/dtos"
)

const apiKey = "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"

func GetDescriptions(w http.ResponseWriter, r *http.Request) {
	err, req := buildRequest(r)
	if err != nil {
		fmt.Println("error on building request: ", err.Error())
		return
	}

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("error on sending request: ", err.Error())
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("error on reading response body: ", err.Error())
		return
	}

	var descriptionResponse dtos.DescriptionResponse
	if err := json.Unmarshal(body, &descriptionResponse); err != nil {
		fmt.Println("error on unmarshalling response body: ", err.Error())
		return
	}

	json.NewEncoder(w).Encode(descriptionResponse)
}

func buildRequest(r *http.Request) (error, *http.Request) {
	productName := strings.TrimPrefix(r.URL.Path, "/descriptions/")

	descriptionRequest := dtos.DescriptionRequest{
		MaxTokens:   3000,
		Model:       "text-davinci-003",
		Prompt:      "Gere para mim uma descrição para um anúncio de um violão " + productName,
		Temperature: 1.0,
	}

	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(descriptionRequest)
	if err != nil {
		log.Fatal(err)
	}

	req, err := http.NewRequest("POST", "https://api.openai.com/v1/completions", &buf)
	if err != nil {
		fmt.Println("error on creating request: ", err.Error())
		return nil, nil
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	return err, req
}
