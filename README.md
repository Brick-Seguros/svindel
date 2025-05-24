# Svindel

**AI-powered fraud analysis, document understanding, and tool orchestration engine.**  
Svindel helps detect fraud, analyze documents (CPF, CNPJ, Name, Plate, Email, Phone, Address), retrieve relevant reports, and generate context-aware completions using Large Language Models (LLMs).

---

## 🚀 What Is Svindel?

Svindel is a **completion engine for fraud detection workflows.** It transforms raw user prompts or queries into intelligent, structured AI responses.

### ✅ Svindel Capabilities:
- 🔍 **Document & Entity Extraction:** Detects CPF, CNPJ, Name, Plate, Email, Phone, or Address inside natural language prompts.
- 🔗 **Contextual Retrieval:** Queries internal systems for relevant **Reports** and **Resources** based on detected documents or input types.
- 🧠 **Prompt Augmentation:** Automatically injects retrieved data into AI prompts for enriched completions.
- 💬 **AI Completions:** Generates **structured multi-message responses** using LLMs (OpenAI-powered, backend-agnostic).
- 🔧 **Tool Execution:** Suggests validators, enrichments, and APIs relevant to phones, emails, or addresses.
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

---

## 🧠 Resource Catalog

Svindel provides dynamically suggested resources per document type:

| Document Type | Resources Example                                   |
|----------------|-----------------------------------------------------|
| CPF            | Processos, KYC, Análise de Crédito                 |
| CNPJ           | KYB, Balanço Patrimonial, Análise de Crédito       |
| Plate          | Gravame, Débitos, Sinistros, Donos                 |
| Email          | Validador de E-mail                                |
| Phone          | Validador de Telefone, Consulta de IMEI           |
| Address        | Validador de Endereço, Área de Risco, Google Search |
| Name           | Google Search                                       |

---

## 🧩 Response Structure Example

```json
{
  "messages": [
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

## 🔥 Próximos Passos

### 📄 Talk to One Report
Permitir que o chat de IA seja **contextualizado a um único report**.

- O AI responderá considerando exclusivamente os dados, recursos e histórico vinculados a um report específico.
- 🔍 Útil para investigações aprofundadas sobre um único caso ou análise.

### 🔄 Talk to One Flow Execution
Permitir que o chat de IA seja **contextualizado a uma execução específica de workflow**.

- O AI terá acesso apenas às variáveis, histórico e resultados daquele fluxo.
- 🔄 Essencial para auditar, revisar ou expandir análises em andamento.
- ✔️ Permite interações no contexto de uma automação ou fluxo operacional específico.

---

### 🕸️ Text to Graph
Transformar texto livre em uma estrutura de dados baseada em **grafo de entidades e relações**.

- Extrair nomes, documentos, empresas, endereços e suas conexões diretamente a partir de descrições textuais.
- 🔗 Base para investigações antifraude, detecção de redes e enriquecimento de dados.
- ✔️ Permite criar visualizações e análises de redes de relacionamentos automaticamente.

---

### 🧠 Abstract the OpenAI Vendor
Implementar uma camada de abstração que permita **trocar facilmente o provedor de LLM**, como:

- OpenAI
- Anthropic
- Google Gemini
- Modelos open-source (Llama, Mistral, Mixtral, etc.)
