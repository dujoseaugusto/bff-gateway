### Estrutura de Pastas

```plaintext
.
├── cmd/               # main da aplicação
├── docs/              # documentação do Swagger
├── internal/          # Código fonte da aplicação
|   ├── domain/        # struck de domínio             
|   ├── handlers/      # camada HTTP da aplicação
│       └── modelos    # Modelos Handler                            
|   ├── helpers/       # serviços adicionais, middlewares                  
|   ├── integration/   # Camada de integração                      
|   ├── service/       # Camada de serviço 
│   └── routes/        # Definições de rotas
│       └── routes.go  # Definição das rotas da API
├── postman/           # Arquivos relacionados ao Postman
│   ├── collections/   # Coleções do Postman
│   └── environments/  # Ambientes do Postman
│       ├── local.json # Ambiente para testes locais
│       └── live.json  # Ambiente para testes live
├── docker-compose.yml # Arquivo Docker Compose
├── dockerfile         # Arquivo Docker Compose
└── README.md          # Este arquivo README


### Executar a Aplicação com Docker Compose
Para executar a aplicação com Docker Compose, use o seguinte comando:

```docker-compose up --build```
Isso irá compilar e iniciar a aplicação com as configurações definidas no arquivo docker-compose.yml.

#### Gerar Documentação do Swagger
Para gerar a documentação do Swagger, execute o seguinte comando:

```swag init -g ./internal/routes/routes.go -o .docs/swagger```
Isso irá gerar a documentação do Swagger com base nas rotas definidas no arquivo routes.go e irá salvá-la na pasta .docs/swagger.

### Testar a API com Postman
Na pasta postman/, você encontrará arquivos relacionados ao Postman:

collections/: Contém as coleções do Postman para testar a API.
environments/: Contém os ambientes do Postman para testes locais e em ambiente live.
Para usar o Postman:

Importe as coleções e ambientes do Postman.
Selecione o ambiente apropriado (local ou live).
Execute as requisições para testar a API.
Certifique-se de definir as variáveis de ambiente necessárias de acordo com suas configurações locais e de live.