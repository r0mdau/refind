package main

import (
	"context"
	"encoding/json"
	"net/http"
	"os"

	"github.com/weaviate/weaviate-go-client/v4/weaviate"
	"github.com/weaviate/weaviate/entities/models"
)

func main() {
	cfg := weaviate.Config{
		Host:   "localhost:8080",
		Scheme: "http",
		Headers: map[string]string{
			"X-OpenAI-Api-Key": os.Getenv("OPENAI_APIKEY"),
		},
	}
	client, err := weaviate.NewClient(cfg)
	if err != nil {
		panic(err)
	}

	// add the schema
	classObj := &models.Class{
		Class:      "Question",
		Vectorizer: "text2vec-openai",
	}

	if client.Schema().ClassCreator().WithClass(classObj).Do(context.Background()) != nil {
		panic(err)
	}

	// Retrieve the data
	items, err := getJSONdata()
	if err != nil {
		panic(err)
	}

	// convert items into a slice of models.Object
	objects := make([]*models.Object, len(items))
	for i := range items {
		objects[i] = &models.Object{
			Class: "Question",
			Properties: map[string]any{
				"category": items[i]["Category"],
				"question": items[i]["Question"],
				"answer":   items[i]["Answer"],
			},
		}
	}

	// batch write items
	batchRes, err := client.Batch().ObjectsBatcher().WithObjects(objects...).Do(context.Background())
	if err != nil {
		panic(err)
	}
	for _, res := range batchRes {
		if res.Result.Errors != nil {
			panic(res.Result.Errors.Error)
		}
	}
}

func getJSONdata() ([]map[string]string, error) {
	// Retrieve the data
	data, err := http.DefaultClient.Get("https://raw.githubusercontent.com/weaviate-tutorials/quickstart/main/data/jeopardy_tiny.json")
	if err != nil {
		return nil, err
	}
	defer data.Body.Close()

	// Decode the data
	var items []map[string]string
	if err := json.NewDecoder(data.Body).Decode(&items); err != nil {
		return nil, err
	}
	return items, nil
}
