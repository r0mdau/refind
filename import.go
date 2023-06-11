package main

import (
	"context"
	"io/ioutil"
	"log"
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

	// create a class
	classObj := &models.Class{
		Class:      "Document",
		Vectorizer: "text2vec-openai", // If set to "none" you must always provide vectors yourself. Could be any other "text2vec-*" also.
		ModuleConfig: map[string]interface{}{
			"generative-openai": map[string]interface{}{
				"model": "gpt-3.5-turbo",
			},
		},
	}

	if client.Schema().ClassCreator().WithClass(classObj).Do(context.Background()) != nil {
		panic(err)
	}

	// Retrieve the data
	items, err := getData()
	if err != nil {
		panic(err)
	}

	// convert items into a slice of models.Object
	objects := make([]*models.Object, len(items))
	for i := range items {
		objects[i] = &models.Object{
			Class: "Document",
			Properties: map[string]any{
				"content": items[i],
			},
		}
	}

	// batch write items
	batchRes, err := client.Batch().ObjectsBatcher().WithObjects(objects...).Do(context.Background())
	check(err)
	for _, res := range batchRes {
		if res.Result.Errors != nil {
			panic(res.Result.Errors.Error)
		}
	}
}

func getData() ([]string, error) {
	var items []string

	folderPath := "import_docs"

	// Open the folder
	folder, err := os.Open(folderPath)
	check(err)
	defer folder.Close()

	// Read all the file names in the folder
	fileNames, err := folder.Readdirnames(0)
	check(err)

	// Iterate over the file names and print them
	for _, fileName := range fileNames {
		body, err := ioutil.ReadFile(folderPath + "/" + fileName)
		if err != nil {
			log.Fatalf("unable to read file: %v", err)
		}

		items = append(items, string(body))
	}

	return items, nil
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
