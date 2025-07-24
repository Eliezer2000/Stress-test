# Stress-test

## Descrição

O **Stress-test** é uma ferramenta de linha de comando (CLI) desenvolvida em Go para realizar testes de carga (stress test) em serviços web HTTP. Ela foi criada para desenvolvedores, equipes de QA e operadores que desejam avaliar a robustez, desempenho e estabilidade de APIs e aplicações web sob diferentes níveis de carga.

A ferramenta permite que o usuário defina a URL do serviço a ser testado, o número total de requisições e o nível de concorrência (quantidade de chamadas simultâneas). Ao final do teste, a ferramenta gera um relatório detalhado com estatísticas sobre o desempenho do serviço, facilitando a identificação de gargalos e falhas.

---

## Regras de Negócio

- O usuário deve informar obrigatoriamente:
  - `--url`: URL do serviço a ser testado (ex: `http://localhost:8080/health`).
  - `--requests`: Número total de requisições a serem realizadas (ex: `1000`).
  - `--concurrency`: Número de requisições simultâneas (concorrência, ex: `10`).
- O sistema distribui as requisições de acordo com o nível de concorrência informado, utilizando goroutines para maximizar a performance.
- Todas as requisições são feitas via método HTTP GET.
- O sistema garante que o número total de requisições será cumprido, mesmo em caso de falhas ou erros de rede.
- Ao final, é gerado um relatório contendo:
  - Tempo total gasto na execução do teste.
  - Quantidade total de requisições realizadas.
  - Quantidade de requisições com status HTTP 200 (sucesso).
  - Distribuição dos demais códigos de status HTTP (ex: 404, 500, etc.).
  - Quantidade de erros (requisições que não obtiveram resposta ou retornaram erro de rede).

---

## Como usar

### Pré-requisitos

- [Go 1.22.4+](https://go.dev/dl/) instalado **(para rodar sem Docker)**
- [Docker](https://www.docker.com/) instalado **(para rodar com Docker)**

---

### Rodando **sem Docker**

1. Clone o repositório:
   ```sh
   git clone https://github.com/Eliezer2000/Stress-test.git
   cd Stress-test
   ```

2. Compile o projeto:
   ```sh
   go build -o stress-test ./cmd
   ```

3. Execute a ferramenta, passando os parâmetros desejados:
   ```sh
   ./stress-test --url=http://google.com --requests=100 --concurrency=10
   ```

   - **--url**: URL do serviço a ser testado.
   - **--requests**: Número total de requisições a serem feitas.
   - **--concurrency**: Número de requisições simultâneas.

---

### Rodando **com Docker**

1. Construa a imagem Docker:
   ```sh
   docker build -t stress-test .
   ```

2. Execute o container, passando os parâmetros desejados:
   ```sh
   docker run --rm stress-test --url=http://google.com --requests=100 --concurrency=10
   ```

   - O parâmetro `--rm` remove o container após a execução.
   - Para acessar um serviço rodando na sua máquina local, use `host.docker.internal` no lugar de `localhost` (válido para Windows e Mac).

---

## Exemplo de saída

```
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

## Estrutura do Projeto

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

## Detalhes Técnicos

- **Concorrência:** Utiliza goroutines e canais para controlar o número de requisições simultâneas, garantindo performance e respeito ao limite de concorrência definido pelo usuário.
- **Coleta de resultados:** Cada requisição retorna um struct com o status HTTP e possíveis erros, permitindo análise detalhada dos resultados.
- **Relatório:** O relatório final apresenta estatísticas essenciais para análise de performance e estabilidade do serviço testado.
- **Extensibilidade:** O projeto está organizado em múltiplos arquivos e módulos, facilitando futuras melhorias, como suporte a outros métodos HTTP, autenticação, payloads customizados, etc.


---

## Licença

Este projeto está licenciado sob a licença MIT. Consulte o arquivo [LICENSE](LICENSE) para mais detalhes.