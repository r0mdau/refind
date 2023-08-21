/*
Copyright © 2023 Romain Dauby (https://github.com/r0mdau)
*/
package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/weaviate/weaviate-go-client/v4/weaviate"
	"github.com/weaviate/weaviate-go-client/v4/weaviate/graphql"
)

// queryCmd represents the query command
var queryCmd = &cobra.Command{
	Use:   "query",
	Short: "Query Weaviate and generate results",
	Long: `This command queries Weaviate and generates results.
The generative search is done using the OpenAI API.`,
	Run: func(cmd *cobra.Command, args []string) {
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
	},
}

func init() {
	rootCmd.AddCommand(queryCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// queryCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// queryCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
