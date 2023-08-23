# Langchain

Prepare the virtual environment

```bash
python3 -m venv venv
source venv/bin/activate
pip install -r requirements.txt
```

Export OpenAI API key

```bash
export OPENAI_APIKEY="..."
```

Import pdf from the `../import_docs` folder

```bash
python3 import.py
```

Ask questions

```bash
python3 query.py -in "How google handles an incident that spans multiple services ?"
```
