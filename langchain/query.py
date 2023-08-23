import os
from langchain.document_loaders import DirectoryLoader
from langchain.text_splitter import RecursiveCharacterTextSplitter
from langchain.embeddings.openai import OpenAIEmbeddings
import weaviate
from langchain.vectorstores import Weaviate
from langchain.chains.question_answering import load_qa_chain
from langchain.llms import OpenAI
import argparse

parser = argparse.ArgumentParser()

parser.add_argument("-in", "--input", dest="input", help="Input text (required)")

client = weaviate.Client(
    url="http://localhost:8080",
    additional_headers={"X-OpenAI-Api-Key": os.environ["OPENAI_APIKEY"]}
)

vectorstore = Weaviate(client, "Chatbot", "content", attributes=["source"])

args = parser.parse_args()
query = args.input

# retrieve text related to the query
docs = vectorstore.similarity_search(query, top_k=20)

# define chain
chain = load_qa_chain(
    OpenAI(openai_api_key = os.environ["OPENAI_APIKEY"],temperature=0.4),
    chain_type="stuff")

# create answer
print(chain.run(input_documents=docs, question=query, return_only_outputs=True))