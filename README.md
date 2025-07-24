# ⚔️ Stress-test (CLI em Go)

## 📄 Descrição

O **Stress-test** é uma ferramenta de linha de comando (CLI) desenvolvida em **Go** para realizar testes de carga (_stress test_) em serviços web HTTP.

💻 Ideal para desenvolvedores, equipes de QA e operadores que desejam avaliar:

- Robustez 🛡️  
- Desempenho ⚙️  
- Estabilidade 🌐  

da API ou aplicação sob diferentes níveis de carga.

A ferramenta permite definir:

- A **URL** do serviço a ser testado  
- O número total de **requisições**  
- O nível de **concorrência** (requisições simultâneas)

✅ Ao final do teste, gera um **relatório detalhado** com estatísticas que facilitam a identificação de gargalos e falhas.

---

## 🧠 Regras de Negócio

- O usuário deve informar obrigatoriamente:
  - `--url`: URL do serviço (ex: `http://localhost:8080/health`)
  - `--requests`: Total de requisições (ex: `1000`)
  - `--concurrency`: Nível de concorrência (ex: `10`)
  
- As requisições são distribuídas com **goroutines** para máxima performance.
- Todos os testes são realizados com o método `GET`.
- O sistema **garante** que todas as requisições serão feitas, mesmo em caso de erros.
- O relatório final contém:
  - ⏱️ Tempo total da execução  
  - 📥 Total de requisições realizadas  
  - ✅ Quantidade de `HTTP 200`  
  - 📊 Distribuição de códigos de status HTTP  
  - ❌ Quantidade de erros ou falhas de rede  

---

## 🚀 Como usar

### ⚙️ Pré-requisitos

- [Go 1.22.4+](https://go.dev/dl/) instalado *(para rodar localmente)*
- [Docker](https://www.docker.com/) instalado *(para rodar com container)*

---

### 🧪 Rodando **sem Docker**

1. Clone o repositório:

```bash
git clone https://github.com/Eliezer2000/Stress-test.git
cd Stress-test
```

2. Compile o projeto:

```bash
go build -o stress-test ./cmd
```

3. Execute a ferramenta com os parâmetros desejados:

```bash
./stress-test --url=http://google.com --requests=100 --concurrency=10
```

Parâmetros disponíveis:

- `--url`: URL do serviço a ser testado  
- `--requests`: Número total de requisições  
- `--concurrency`: Número de requisições simultâneas  

---

### 🐳 Rodando **com Docker**

1. Construa a imagem Docker:

```bash
docker build -t stress-test .
```

2. Execute o container:

```bash
docker run --rm stress-test --url=http://google.com --requests=100 --concurrency=10
```

ℹ️ Para acessar um serviço na sua máquina local usando Docker, utilize `host.docker.internal` no lugar de `localhost` (válido para **Windows** e **Mac**).

---

## 📊 Exemplo de saída

```text
====== Stress Test Report ======
Total time: 2.3456789s
Total requests: 100
HTTP 200: 98
Errors: 2
Status code distribution:
  200: 98
  Errors (no response): 2
================================
```

---

## 🗂️ Estrutura do Projeto

```
Stress-test/
  ├── cmd/
  │   └── main.go           # Ponto de entrada da aplicação
  ├── internal/
  │   ├── cli.go            # Parsing e validação dos parâmetros CLI
  │   ├── report.go         # Geração e exibição do relatório
  │   ├── runner.go         # Execução dos requests concorrentes
  │   └── types.go          # Definição de structs e tipos auxiliares
  ├── Dockerfile            # Dockerfile para build e execução via container
  ├── go.mod                # Gerenciamento de dependências Go
  └── README.md             # Este arquivo
```

---

## ⚙️ Detalhes Técnicos

- 🔁 **Concorrência**: uso de goroutines e canais para controlar requisições simultâneas com alta performance
- 🧩 **Coleta de resultados**: cada requisição retorna status HTTP e erro (se houver), permitindo análise completa
- 📈 **Relatório final**: exibe estatísticas úteis para avaliar performance
- 🛠️ **Extensibilidade**: arquitetura modular, fácil de estender para:
  - Suporte a outros métodos HTTP
  - Autenticação
  - Payloads personalizados
  - Exportação dos resultados

---

## 📄 Licença

Este projeto está licenciado sob a licença **MIT**.  
Consulte o arquivo [LICENSE](LICENSE) para mais detalhes.

---
