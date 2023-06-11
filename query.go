package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/weaviate/weaviate-go-client/v4/weaviate"
	"github.com/weaviate/weaviate-go-client/v4/weaviate/graphql"
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

	fields := []graphql.Field{
		{Name: "content"},
	}

	nearText := client.GraphQL().
		NearTextArgBuilder().
		WithConcepts([]string{"software", "reliability"})

	gs := graphql.NewGenerativeSearch().SingleResult("Summarize where to start in adopting SRE in enterprise {content}")
	//gs := graphql.NewGenerativeSearch().GroupedResult("Explain why these documents are about engineering levels")

	result, err := client.GraphQL().Get().
		WithClassName("Document").
		WithFields(fields...).
		WithGenerativeSearch(gs).
		WithNearText(nearText).
		WithLimit(2).
		Do(context.Background())
	if err != nil {
		panic(err)
	}

	out, err := json.Marshal(result)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(out))
}
