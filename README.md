# Desafio go-expert #3 - Clean Architecture

## Pré-requisitos
- Subir o docker-compose local: `docker-compose up -d`;
- Localizar o mysql: `docker ps`;
- Entrar no container do mysql. Ex: `docker exec -it {id do container} bash`
- Executar o mysql: `mysql -uroot -p` (senha padrão é root);
- Criar a base de dados: `create database orders; use orders`
- Criar a tabela: `CREATE TABLE orders (id varchar(255) NOT NULL, price float NOT NULL, tax float NOT NULL, final_price float NOT NULL, PRIMARY KEY (id));`

## Execução
Para executar o aplicativo, navegue até a pasta `cmd/ordersystem` e execute o comando:
`go run main.go wire_gen.go`.

## Subindo gRPC
Para subir o gRPC, execute `evans -r repl`.

Em seguida, selecione o serviço:
- `package pb`;
- `service OrderService`.