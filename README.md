# refind

Refind is a project to store my personnal documents in a [Weaviate] vector
 database and use [OpenAI] vectorizer module and generative search.

Here is a shell recording with [asciinema](https://asciinema.org/) showing the
different answers with the same input question from :

1. the current ChatGPt version (using
[sgpt](https://github.com/TheR1D/shell_gpt) cli)
1. the generative search with Weaviate containing a sample document from the
[Enteprise Roadmap to SRE](https://sre.google/resources/practices-and-processes/enterprise-roadmap-to-sre/)
book

![shell](demo.svg)

## Getting started

Create and provide your OpenAI API key :

```bash
export OPENAI_APIKEY="..."
```

Put full-text documents with less than 4097 [tokens](https://help.openai.com/en/articles/4936856-what-are-tokens-and-how-to-count-them).
If your prompt is 4000 tokens, your completion can be 97 tokens at most.

Let's use it :

```bash
# start the weaviate database
docker-compose up -d

# verify it is up and running
curl http://localhost:8080/v1/meta | jq .
docker-compose logs

# import data, DO IT ONCE
go run main.go import

# query data related to software
go run main.go query | jq .

# when done, gracefully shutdown
docker-compose down
```

In this example, Weaviate is returning software-related entries.

Weaviate operations :

```bash
# get the schema
curl -s http://localhost:8080/v1/schema | jq .

# get objects
curl -s http://localhost:8080/v1/objects | jq .

# get one class from the schema
curl -s http://localhost:8080/v1/schema/Chatbot | jq .

# delete a class
curl -s -XDELETE http://localhost:8080/v1/schema/Chatbot | jq .
```

## Goal

[The ChatGPT Retrieval plugin](https://weaviate.io/blog/weaviate-retrieval-plugin).

[Weaviate]: https://weaviate.io
[OpenAI]: https://openai.com/

### TODO

1. automatically add more metadata in weaviate when ingesting documents
1. automatically split documents bigger than the OpenAI token limit
1. add query string as a flag when running the `query` command
1. unit tests, let Copilot write them \o/

## Trying langchain

[README.md](langchain/README.md)
