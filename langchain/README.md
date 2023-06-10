# Langchain

Source [dagster.io](https://dagster.io/blog/chatgpt-langchain).

```bash
python3 -m venv venv
source venv/bin/activate
pip install -r requirements.txt
```

Export OpenAI API key

```bash
export OPENAI_API_KEY="..."
```

Run a search

```bash
python3 -c'from langchain_bot import print_answer; print_answer("What are the main differences between Linux and Windows?")'

python3 -c'from langchain_bot import print_answer; print_answer("Which members of Matchbox 20 play guitar?")'
```
