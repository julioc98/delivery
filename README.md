# Delivery

## Arquitetura e Padrões

### Clean Architecture

Este projeto segue o padrão de arquitetura Clean Architecture, que promove a separação de preocupações e a independência das camadas de software. Isso resulta em um código mais limpo, testável e flexível.

### Padrão de Arquitetura CQRS

Utilizamos o padrão de arquitetura CQRS (Command Query Responsibility Segregation) para separar as operações de leitura (queries) das operações de escrita (commands). Isso melhora a escalabilidade e a manutenção do sistema, permitindo tratamento diferenciado para cada tipo de operação.

# Como Executar o Projeto

Para executar este projeto, você pode usar o Docker Compose. Certifique-se de ter o Docker e o Docker Compose instalados em seu sistema.

1. Clone o repositório para sua máquina:

```sh
   git https://github.com/julioc98/delivery.git
   cd delivery
```

2. Inicie os contêineres Docker com o Docker Compose:

```sh
   docker-compose up
```

Isso iniciará todos os serviços necessários, incluindo PostgreSQL, Redis, NATS, e outros, de acordo com o arquivo docker-compose.yml.

3. Aguarde até que todos os contêineres estejam em execução e que o projeto seja construído e implantado.

4. Após iniciar os contêineres do Docker, você precisa aplicar as migrações do banco de dados. Abra um novo terminal na pasta raiz do projeto e execute o seguinte comando:

```sh
   make migrate-up
```

Isso aplicará todas as migrações do banco de dados e configurará o esquema necessário.

5. Você pode usar a coleção do Postman que esta em `docs/delivery.postman_collection.json` importando para seu Postman ou ferramenta compativel.

Mas caso queira pode fazer todas as chmadas pelo proprio terminal usando `curl`

Salvar nova posição do entregador:
```sh
curl --location 'localhost:3001/drivers/1/locations' \
--header 'Content-Type: application/json' \
--data '{
    "latitude": 154.9988 ,
    "longitude": 89.1622
}'
```

Recuperar a ultima posição de um entregador:
```sh
curl --location 'localhost:3001/drivers/1/locations/current' \
--header 'Content-Type: application/json'
```

Recuperar o historico de posições de um entregador:
```sh
curl --location 'localhost:3001/drivers/1/locations' \
--header 'Content-Type: application/json'
```

Recuperar todos entregadores com a posição atual proxima a um raio de uma localização:
```sh
curl --location 'localhost:3001/drivers/locations/nearby?latitude=154.9987&longitude=89.1621&radius=100' \
--header 'Content-Type: application/json'
```


## Tecnologias Utilizadas

### Golang

A linguagem de programação Go (Golang) é a base deste projeto. Ela é conhecida por sua eficiência, desempenho e suporte a concorrência, tornando-a ideal para aplicativos escaláveis e de alto desempenho.

Saiba mais sobre Golang em [golang.org](https://golang.org/).

### Chi

O Chi é um roteador HTTP minimalista e de alto desempenho para Go. Ele é usado neste projeto para criar endpoints da API REST e gerenciar solicitações HTTP de forma eficiente.

Saiba mais sobre o Chi em [github.com/go-chi/chi](https://github.com/go-chi/chi).

### gRPC

O gRPC é um framework de código aberto que facilita a comunicação entre serviços usando protocol buffers (protobufs). Neste projeto, usamos gRPC para comunicação entre componentes.

Saiba mais sobre o gRPC em [grpc.io](https://grpc.io/).

### Protocol Buffers (protobufs)

Os protobufs são uma linguagem de descrição de interface e um formato de serialização de dados. Eles são usados em conjunto com o gRPC para definir as mensagens e serviços da nossa API.

Saiba mais sobre protobufs em [developers.google.com/protocol-buffers](https://developers.google.com/protocol-buffers).

### golang-migrate

O golang-migrate é uma ferramenta que facilita a migração de bancos de dados em Go. Neste projeto, usamos o golang-migrate para gerenciar nossas migrações de banco de dados.

Saiba mais sobre o golang-migrate em [github.com/golang-migrate/migrate](https://github.com/golang-migrate/migrate).

### Redis

O Redis é um armazenamento de dados em memória de alto desempenho. Neste projeto, usamos o Redis como um sistema de cache para melhorar o desempenho das operações de leitura.

Saiba mais sobre o Redis em [redis.io](https://redis.io/).

### NATS

O NATS é um sistema de mensagens de código aberto altamente escalável e de alto desempenho. Neste projeto, usamos o NATS para comunicação assíncrona entre componentes da aplicação.

Saiba mais sobre o NATS em [nats.io](https://nats.io/).

### PostgreSQL e PostGIS

O PostgreSQL é um poderoso sistema de gerenciamento de banco de dados relacional, e o PostGIS é uma extensão que adiciona suporte a dados geoespaciais. Neste projeto, usamos o PostgreSQL com o PostGIS para armazenar e gerenciar nossos dados.

Saiba mais sobre o PostgreSQL em [postgresql.org](https://www.postgresql.org/) e sobre o PostGIS em [postgis.net](https://postgis.net/).

