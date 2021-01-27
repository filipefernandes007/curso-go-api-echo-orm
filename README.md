### Requisitos

[Docker](https://docs.docker.com/desktop/)

### Instalar

```shell
git clone git@github.com:filipefernandes007/curso-go-api-echo-orm.git
cd curso-go-api-echo-orm
```

Ou 

```shell
git clone https://github.com/filipefernandes007/curso-go-api-echo-orm.git
cd curso-go-api-echo-orm
```

#### Instalar dependências:
```shell
go mod tidy
```

### Correr o docker
```shell
docker-compose up 
```

### Aceder ao MySQL

```shell
Porta: 33060
User: user1
Password: 123456
```

Ver configurações no `docker-compose.yml`

### Correr a API

```shell
go run api/main.go
```
