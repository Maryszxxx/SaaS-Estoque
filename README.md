# SaaS Estoque API

API REST para gerenciamento e controle de estoque de produtos desenvolvida em Go. O projeto implementa operações completas de CRUD, persistência em banco relacional, validação de dados, suporte a atualizações parciais via PATCH e ambiente containerizado.

---

## Tecnologias e Ferramentas

* **Linguagem:** Go (Golang)
* **Web Framework:** Gin Gonic
* **Banco de Dados:** PostgreSQL & SQL puro
* **Documentação:** Swagger / Swag
* **Containerização:** Docker & Docker Compose

---

## Arquitetura do Projeto

A aplicação adota o padrão de Arquitetura em Camadas (Layered Architecture) com desacoplamento via interfaces (Repository Pattern), garantindo separação de responsabilidades:

```text
  [ HTTP Request ]
         │
         ▼
    [ Handler ]          ---> Validação, binding de payload e resposta HTTP
         │
         ▼
[ Service / Use Case ]   ---> Regras de negócio da aplicação
         │
         ▼
   [ Repository ]        ---> Contrato de acesso a dados (Interface)
         │
         ▼
    [ Database ]         ---> Implementação PostgreSQL

```

### Estrutura de Pastas

```text
saas/
├── database/     # Implementação do repositório PostgreSQL
├── docs/         # Arquivos de documentação gerados pelo Swagger
├── entity/       # Modelos de domínio e structs de dados
├── handler/      # Handlers HTTP e rotas (Gin)
├── repository/   # Interfaces e contratos do repositório
├── usecase/      # Regras de negócio e casos de uso
├── Dockerfile
├── docker-compose.yml
├── go.mod
├── go.sum
└── main.go       # Ponto de entrada da aplicação

```

---

## Endpoints da API

| Método | Endpoint | Descrição |
| --- | --- | --- |
| `POST` | `/products` | Cadastra um novo produto |
| `GET` | `/products` | Lista todos os produtos |
| `GET` | `/products/:id` | Busca um produto específico por ID |
| `PUT` | `/products/:id` | Atualiza todos os campos de um produto |
| `PATCH` | `/products/:id` | Atualiza parcialmente um produto |
| `DELETE` | `/products/:id` | Remove um produto do sistema |

---

## Exemplos de Requisição

### 1. Criar Produto (`POST /products`)

**Payload:**

```json
{
  "name": "Notebook",
  "description": "Notebook para desenvolvimento",
  "price": 4500.00,
  "quantity": 10,
  "category_id": 1
}

```

**Resposta (`201 Created`):**

```json
{
  "message": "Product created successfully"
}

```

---

### 2. Atualização Completa (`PUT /products/:id`)

**Payload:**

```json
{
  "name": "Notebook Gamer",
  "description": "Notebook para desenvolvimento e jogos",
  "price": 5500.00,
  "quantity": 5,
  "category_id": 1
}

```

---

### 3. Atualização Parcial (`PATCH /products/:id`)

Permite alterar apenas os campos desejados.

**Exemplo A (Apenas Preço):**

```json
{
  "price": 4999.90
}

```

**Exemplo B (Apenas Estoque):**

```json
{
  "quantity": 20
}

```

---

## Variáveis de Ambiente

As configurações de banco de dados são gerenciadas por variáveis de ambiente. Crie um arquivo `.env` na raiz do projeto com a seguinte estrutura:

```env
DB_HOST=localhost
DB_PORT=5432
DB_USERNAME=postgres
DB_PASSWORD=your_password
DB_DATABASE=Saas-Estoque

```

Ao executar via Docker Compose, a variável `DB_HOST` deve ser configurada como `postgres`.

---

## Como Executar o Projeto

### Opção 1: Utilizando Docker (Recomendado)

Inicie os containers da API e do banco PostgreSQL com o comando:

```bash
docker compose up --build

```

A API estará acessível em: `http://localhost:8080`

**Comandos de gerenciamento:**

```bash
# Parar os containers
docker compose down

# Parar e remover os volumes (remove os dados persistidos do PostgreSQL)
docker compose down -v

```

---

### Opção 2: Execução Local

1. Configure um servidor PostgreSQL e ajuste as variáveis no arquivo `.env`.
2. Instale as dependências Go:
```bash
go mod download

```


3. Execute a aplicação:
```bash
go run main.go

```



---

## Documentação da API (Swagger)

A API conta com documentação interativa integrada. Após iniciar a aplicação, acesse no seu navegador:

`http://localhost:8080/swagger/index.html`

O Swagger UI permite visualizar os schemas de entrada/saída e testar todos os endpoints diretamente no navegador.

---

## Objetivos do Projeto

Projeto desenvolvido para consolidar os seguintes conceitos de desenvolvimento backend em Go:

* Construção de APIs RESTful utilizando Gin Framework.
* Arquitetura em camadas e padrão Repository para desacoplamento de persistência.
* Manipulação de banco de dados relacional (PostgreSQL) e execução de SQL.
* Validação de payload e suporte a atualizações parciais via PATCH.
* Containerização de aplicações Go e orquestração de ambiente multi-container com Docker Compose.
* Documentação de APIs com Swagger/OpenAPI.