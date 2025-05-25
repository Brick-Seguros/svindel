# Svindel

**AI-powered fraud analysis, document understanding, and tool orchestration engine.**  
Svindel helps detect fraud, analyze documents (CPF, CNPJ, Name, Plate, Email, Phone, Address), retrieve relevant reports, and generate context-aware completions using Large Language Models (LLMs).

---

## 🚀 What Is Svindel?

Svindel is a **completion engine for fraud detection workflows.** It transforms raw user prompts or queries into intelligent, structured AI responses and evaluates their quality with AI-powered judges.

---

## ✅ Svindel Capabilities

- 🔍 **Document & Entity Extraction:** Detects CPF, CNPJ, Name, Plate, Email, Phone, or Address inside natural language prompts.
- 🔗 **Contextual Retrieval:** Queries internal systems for relevant **Reports** and **Resources** based on detected documents or input types.
- 🧠 **Prompt Augmentation:** Automatically injects retrieved data into AI prompts for enriched completions.
- 💬 **AI Completions:** Generates **structured multi-message responses** using LLMs (OpenAI-powered, backend-agnostic).
- 🔧 **Tool Execution:** Suggests validators, enrichments, and APIs relevant to phones, emails, addresses, and documents.
- 🎯 **AI Response Evaluation:** Automated evaluation of AI responses for:
  - ✅ **Entity Recognition Correctness**
  - ✅ **Relevance to the user’s input**
  - ✅ **Hallucination Risk** (invented data)
  - 🧠 (Planned) Accuracy, Completeness, Reasoning Quality, and Toxicity detection.
- 🌐 **APIs:** Exposes completion services via **REST** and **WebSocket** interfaces for both synchronous and streaming responses.

---

## 💡 How Svindel Works

### 🔥 Pipeline per User Prompt:

1. **Document Extraction**
   - Detects:
     - `CPF`
     - `CNPJ`
     - `Plate` (vehicle)
     - `Name`
     - `Email`
     - `Phone`
     - `Address`
   - If none are found → classifies as a **QUESTION**.

2. **Contextual Retrieval (if document exists)**
   - Retrieves:
     - Historical **Reports** (for CPF, CNPJ, Plate).
     - Available **Resources** (e.g., credit checks, KYC, validators, public registries).

3. **Prompt Augmentation**
   - Builds an enriched AI prompt containing:
     - User’s original input.
     - Extracted document metadata.
     - Retrieved reports (if applicable).
     - Relevant resources and tools.

4. **Completion Generation**
   - Sends the augmented prompt to an LLM (OpenAI GPT models).
   - Returns a **structured array of messages**, including:
     - `TEXT` → Textual responses or context.
     - `REPORT_SHORTCUT` → Shortcut to internal reports.
     - `RESOURCE_SELECTOR` → Suggested validators, APIs, and checks.
     - `AGENT_TRIGGER` → Trigger specialized AI agents (planned).

5. **Automated Evaluation (Eval Engine)**
   - AI responses are automatically evaluated in the background based on:
     - 🔎 **Entity Recognition:** Did the AI detect the correct document type (CPF, CNPJ, Plate, etc.)?
     - 🎯 **Relevance:** Is the response relevant to the user's input and context?
     - 🚫 **Hallucination Risk:** Did the AI invent any data not present in the context (e.g., fake reports or resources)?
   - ✔️ Generates a structured `EvaluationResult` containing:
     - Scorecards
     - Risk Levels
     - Tags (e.g., `correct`, `minor_hallucination`, `irrelevant`)
     - Comments for traceability.

---

## 🧠 Resource Catalog

Svindel provides dynamically suggested resources per document type:

| Document Type | Resources Example                                   |
|----------------|-----------------------------------------------------|
| CPF            | Processos, KYC, Análise de Crédito                 |
| CNPJ           | KYB, Balanço Patrimonial, Análise de Crédito       |
| Plate          | Gravame, Débitos, Sinistros, Donos                 |
| Email          | Validador de E-mail                                |
| Phone          | Validador de Telefone, Consulta de IMEI            |
| Address        | Validador de Endereço, Área de Risco, Google Search|
| Name           | Google Search                                       |

---

## 🧩 AI Response Structure Example

```json
{
  "messages": [
    {
      "type": "TEXT",
      "text": "Encontramos algumas análises já realizadas nesse CPF"
    },
    {
      "type": "REPORT_SHORTCUT",
      "shortcut": { "id": "report-123", "title": "Análise CPF", ... }
    },
    {
      "type": "RESOURCE_SELECTOR",
      "resources": [
        { "id": "ANALYSIS_PF_KYC", "title": "KYC", ... }
      ]
    },
    {
      "type": "TEXT",
      "text": "Estes são os relatórios e recursos disponíveis para este CPF."
    }
  ]
}
```

---

## 🧠 Evaluation Result Example

```json
{
  "ID": "e7335cb1-6ba7-46b1-8258-91c577ea9a63",
  "Input": {
    "UserInput": "Guilherme Ninov de Meira\n",
    "Context": "",
    "AIResponse": {
      "Document": "",
      "DocumentType": "NONE",
      "IsQuestion": true
    }
  },
  "Results": [
    {
      "CriteriaType": "field",
      "CriteriaName": "extracted_document_type",
      "Value": "NAME"
    },
    {
      "CriteriaType": "field",
      "CriteriaName": "extracted_document_value",
      "Value": ""
    },
    {
      "CriteriaType": "field",
      "CriteriaName": "expected_document_type",
      "Value": "NAME"
    },
    {
      "CriteriaType": "field",
      "CriteriaName": "expected_document_value",
      "Value": "Guilherme Ninov de Meira"
    },
    {
      "CriteriaType": "tag",
      "CriteriaName": "tags",
      "Value": "missing, wrong_type"
    }
  ],
  "Comments": "The AI failed to identify the input as a NAME type.",
  "Rating": "bad",
  "Strategy": "entity_recognition",
  "EvaluatedAt": "2025-05-25T16:02:38.512602-03:00"
}
```

## 🔥 Next Steps

### 📄 Talk to One Report
Enable the AI chat to be **contextualized to a single report.**

- The AI will respond considering **only the data, resources, and history linked to that specific report.**
- 🔍 Useful for deep investigations of a specific case or analysis.


### 🔄 Talk to One Flow Execution
Enable the AI chat to be **contextualized to a specific workflow execution.**

- The AI will have access **only to the variables, history, and results of that workflow execution.**
- 🔄 Essential for auditing, reviewing, or expanding analyses in progress.
- ✔️ Supports interactions fully scoped to an automation or operational flow.


### 🕸️ Text to Graph
Transform free-form text into a **graph-based structure of entities and relationships.**

- Extract names, documents, companies, addresses, and their connections directly from textual descriptions.
- 🔗 Forms the foundation for fraud investigations, network detection, and data enrichment.
- ✔️ Allows automatic creation of **relationship graphs** for analysis and visualization.


### 🧠 Abstract the OpenAI Vendor
Implement an abstraction layer that allows **easily switching the LLM provider**, such as:

- OpenAI
- Anthropic
- Google Gemini
- Open-source models like Llama, Mistral, Mixtral, etc.

---

## 🚀 How to Run the Project

### 🔑 Environment Variables

Create a `.env` file in the root folder:

```env
PORT=8080
OPENAI_API_KEY=your-openai-key
BRICK_API_URL=https://api.brickseguros.com.br
BRICK_API_TOKEN=your-brick-api-token
```


### ▶️ Run 

### 1. Install dependencies:

```bash
make deps
```

### 2. Run the server:

```bash
make run
```

### 3. Websocket Request:

- Connect to ws://localhost:8080/ws/chat

- Send any message