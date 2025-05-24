# Svindel

**AI-powered fraud analysis and document understanding engine.**  
Svindel helps detect fraud, analyze documents (CPF, CNPJ, Name, Plate), retrieve relevant reports, and generate context-aware completions using Large Language Models (LLMs).

---

## 🚀 What Is Svindel?

Svindel is a **completion engine for fraud detection workflows.** It transforms raw user prompts or queries into intelligent, structured AI responses.

### ✅ Svindel Capabilities:
- 🔍 **Document Extraction:** Detects CPF, CNPJ, Name, or Vehicle Plate inside natural language prompts.
- 🔗 **Contextual Retrieval:** Queries internal systems for relevant reports and resources based on detected documents.
- 🧠 **Prompt Augmentation:** Injects retrieved data into AI prompts for enhanced completions.
- 💬 **AI Completions:** Generates structured responses using LLMs (OpenAI-powered, with swappable backends).
- 🌐 **APIs:** Exposes completion services via REST and WebSocket interfaces.

---

## 💡 How Svindel Works

### 🔥 Pipeline per User Prompt:

1. **Document Extraction**
   - Parses input for CPF, CNPJ, Name, or Plate.
   - If none → classifies it as a **QUESTION**.

2. **Contextual Retrieval (if document exists)**
   - Retrieves:
     - Historical **Reports**.
     - Available **Resources** (external APIs, validators, or checks relevant to that document type).

3. **Prompt Augmentation**
   - Combines the user’s input with the retrieved context.

4. **Completion Generation**
   - Sends the augmented prompt to the AI (OpenAI GPT models).
   - Returns **structured outputs**, including:
     - `TEXT`
     - `REPORT_SHORTCUT`
     - `AGENT_TRIGGER`
     - `RESOURCE_SELECTOR`

---

## Next Steps

- Connect the Retrieval package to the database and Brick API.
- Implement "Talk to One Report".
- Implement "Talk to One Flow Execution".
- Implement "Text to Graph".
- Abstract the OpenAI vendor from the completion package.

