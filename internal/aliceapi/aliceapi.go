package aliceapi

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type Client struct {
	model       string
	temperature float64
	stream      bool
	endpoint    string
	client      *http.Client
}

type Request struct {
	Model       string  `json:"model"`
	Prompt      string  `json:"prompt"`
	Temperature float64 `json:"temperature"`
	Stream      bool    `json:"stream"`
}

type Response struct {
	Model    string `json:"model"`
	Response string `json:"response"`
	Done     bool   `json:"done"`
}

func New(endpoint string, model string) *Client {
	return &Client{
		model:       model,
		temperature: 0.7,
		stream:      false,
		endpoint:    endpoint,
		client:      http.DefaultClient,
	}
}

func (ac *Client) SendRequest(prompt string) string {
	request := Request{
		Model:       ac.model,
		Prompt:      prompt,
		Temperature: ac.temperature,
		Stream:      ac.stream,
	}

	jsonReq, err := json.Marshal(request)
	if err != nil {
		log.Fatal(err)
	}

	response, err := ac.client.Post(ac.endpoint, "application/json", bytes.NewReader(jsonReq))
	if err != nil {
		log.Fatal(err)
	}

	responseBytes, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var result Response
	err = json.Unmarshal(responseBytes, &result)
	if err != nil {
		log.Fatal(err)
	}

	return result.Response
}

func (ac *Client) ChangeTemperature(temperature float64) {
	ac.temperature = temperature
}

func (ac *Client) ChangeStream(stream bool) {
	ac.stream = stream
}
