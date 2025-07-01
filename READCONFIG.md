# Sistema de Vendas Base

Versão básica e open-source do código usado na criação de um sistema de uma loja real.  
A versão completa com os dados específicos do projeto será criada e vendida de forma privada.

Este documento descreve como rodar o sistema completo, incluindo backend (Go), banco de dados (Postgres via Docker) e frontend (Flutter).

---

## Pré-requisitos

- [Go](https://golang.org/dl/) 1.20 ou superior
- [Docker](https://docs.docker.com/get-docker/) e [Docker Compose](https://docs.docker.com/compose/)
- [Flutter](https://docs.flutter.dev/get-started/install) (para rodar o app)

---

## Configuração do Banco de Dados

1. **Clone o repositório:**
   ```sh
   git clone https://github.com/Chris-cez/BaseShopSystem.git
   cd BaseShopSystem
   ```

2. **Copie o arquivo de exemplo de variáveis de ambiente:**
   ```sh
   cp .env.example .env
   ```
   Edite o arquivo `.env` se desejar alterar usuário, senha ou nome do banco.

3. **Suba o banco de dados com Docker Compose:**
   ```sh
   docker-compose up -d
   ```

---

## Instalação das Dependências do Backend

1. **Instale as dependências Go:**
   ```sh
   go mod tidy
   ```

---

## Rodando o Backend

1. **Execute as migrações e rode o servidor:**
   ```sh
   go run .
   ```
   ou
   ```sh
   go run main.go seed.go
   ```
   O backend irá aplicar as migrações e popular dados de teste automaticamente.

---

## Documentação da API (Swagger)

Após rodar o backend, acesse a documentação interativa da API pelo navegador:

```
http://localhost:8080/swagger/index.html
```

Nessa página você pode visualizar, testar e explorar todos os endpoints disponíveis.

---

## Rodando os Testes

Execute todos os testes automatizados com:
```sh
go test ./...
```

---

## Rodando o Frontend (Flutter)

1. **Acesse a pasta do frontend:**
   ```sh
   cd ui
   ```

2. **Instale as dependências do Flutter:**
   ```sh
   flutter pub get
   ```

3. **Rode o app:**
   ```sh
   flutter run
   ```
   > **Obs:** Certifique-se de que o endpoint da API está correto no código Flutter, caso rode em ambiente diferente.

---

## Exemplo de Variáveis de Ambiente (`.env`)

```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=root
DB_NAME=postgres
DB_SSLMODE=disable
```

---

## Exemplo de Uso da API

- **Criar Produto:**
  ```sh
  curl -X POST http://localhost:8080/api/products \
    -H "Authorization: Bearer <seu_token_jwt>" \
    -H "Content-Type: application/json" \
    -d '{"code":"001","price":10.0,"name":"Produto Teste","ncm":"12345678","gtin":"7891234567890","um":"UN","description":"Descrição teste","class_id":1,"stock":100,"valtrib":0.5}'
  ```

- **Listar Produtos:**
  ```sh
  curl -H "Authorization: Bearer <seu_token_jwt>" http://localhost:8080/api/products
  ```

---

## Problemas Comuns

- **Erro de conexão com o banco:**  
  Verifique se o Docker está rodando e se as variáveis de ambiente estão corretas.

- **Erro 401 nas rotas:**  
  Certifique-se de enviar um token JWT válido no header `Authorization`.

- **Erro 422 ao criar produto:**  
  Verifique se todos os campos obrigatórios estão sendo enviados no JSON.

- **Erro ao acessar a documentação Swagger:**  
  Certifique-se de que a aplicação está rodando e acesse `http://localhost:8080/swagger/index.html`.  
  Se aparecer erro "Failed to load API definition", confira se a pasta `docs/` foi gerada com `swag init` e está no mesmo diretório do `main.go`.

---

## Licença

Este projeto está licenciado sob os termos da [Apache License 2.0](LICENSE).

---