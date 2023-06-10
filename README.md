# refind

Refind is a project to store my personnal documents in a [Weaviate] vector
 database and use [OpenAI] vectorizer module.

## Getting started

Create and provide your OpenAI API key :

```bash
export OPENAI_APIKEY="..."
```

Let's use it :

```bash
# start the weaviate database
docker-compose up -d

# verify it is up and running
curl http://localhost:8080/v1/meta | jq .

# import data, DO IT ONCE
go run import.go

# query data related to biology
go run query.go | jq .

# when done, gracefully shutdown
docker-compose down
```

In this example, Weaviate is returning biology-related entries.

Weaviate operations

```bash
# get the schema
curl -s http://localhost:8080/v1/schema | jq .

# get one class from the schema
curl -s http://localhost:8080/v1/schema/Question | jq .

# delete a class
curl -s -XDELETE http://localhost:8080/v1/schema/Question | jq .
```

## Goal

[The ChatGPT Retrieval plugin](https://weaviate.io/blog/weaviate-retrieval-plugin).

[Weaviate]: https://weaviate.io
[OpenAI]: https://openai.com/

## Using langchain

[README.md](langchain/README.md)
