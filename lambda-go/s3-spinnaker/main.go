package main

import (
	"bytes"
	"fmt"
	"context"
	"os"
	"strings"
	"github.com/aws/aws-lambda-go/lambda"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"crypto/tls"
)

type record struct {
	S3 struct {
		Bucket struct {
			Name string `json:"name"`
			ARN string `json:"arn"`
		} `json:"bucket"`
		Object struct {
			Key string `json:"key"`
		} `json:"object"`
	} `json:"s3"`
}

type s3Event struct {
	Records []record `json:"Records"`
}

type webhookResponse struct {
	EventProcessed bool `json:"eventProcessed"`
	EventID string `json:"eventId"`
}

type parameterBody struct {
	Parameters struct {
		Object string `json:"object"`
	} `json:"parameters"`
}

func handleRequest(ctx context.Context, event s3Event) (string, error) {
	url := os.Getenv("TRIGGER_URL")
	insecure := os.Getenv("INSECURESKIPVERIFY")

	// Put together the client
	var client *http.Client
	if strings.ToUpper(insecure) == "TRUE" {
		tr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
		client = &http.Client{Transport: tr}
	} else {
		client = &http.Client{}
	}

	// Make the body
	object := "s3://" + event.Records[0].S3.Bucket.Name + "/" + event.Records[0].S3.Object.Key
	parameterBody := &parameterBody{}
	parameterBody.Parameters.Object = object
	var jsonData []byte
	jsonData, err := json.Marshal(parameterBody)
	if err != nil {
		panic(err)
	}
	
	// Make the request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	// Read the response body
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var webhookResponse webhookResponse
	err = json.Unmarshal(body, &webhookResponse)
	if err != nil {
		panic(err)
	}

	return fmt.Sprintf("Triggered event %s with object %s (JSON body %s)", webhookResponse.EventID, object, jsonData), nil
}

func main() {
	lambda.Start(handleRequest)
}