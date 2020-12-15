# Desafio SmartMEI 

Candidato: Vinicius da Silva Cardozo

LinkedIn: https://www.linkedin.com/in/vinicius-cardozo-669a15136/

## Requirementos

* [Docker](https://docs.docker.com/engine/install/ubuntu/)
* [Go](https://golang.org/dl/)

## Instalação

Após clonar o repositório, é necessário ir no caminho raiz e rodar o seguinte comando 

    docker-compose up

Isso subirá o banco de dados na porta 27017.

## Iniciando o projeto

Após subir o banco de dados, para iniciar o projeto basta rodar o seguinte comando no caminho raiz.

    go run main.go run

O servidor HTTP está rodando localmente na porta `8082`

## Endpoints

### Pegar dados de um usuário

METÓDO `GET`v1/

    /v1/getUser/:userid
    

**Parametros**: 

* userid: Id do usuário do qual os dados devem ser retornados

#### Exemplo

No terminal execute

    curl --location --request GET 'http://localhost:8082/v1/getUser/5421'

Saída esperada

_Status http_: `200 OK`
```json
{
  "status": "success",
  "data": {
    "id": "5fd8d7273722a1d48c225244",
    "name": "teste01",
    "email": "teste01@email.com.br",
    "pages": "2020-12-15T15:32:55.909Z",
    "collection": [
      {
        "id": "5e6e21ba-6fef-4a1c-b1c8-bc208992e6a1",
        "title": "Livro 02",
        "pages": "156",
        "createdAt": "2020-12-15T15:33:10.848Z"
      }
    ],
    "lentBooks": [
      {
        "book": {
          "id": "5e6e21ba-6fef-4a1c-b1c8-bc208992e6a1",
          "title": "Livro 02",
          "pages": "156",
          "createdAt": "2020-12-15T15:33:10.848Z"
        },
        "fromUser": "5fd8d7273722a1d48c225244",
        "toUser": "5fd8d7563722a1d48c225245",
        "lentAt": "2020-12-15T15:34:29.006Z",
        "returnedAt": "2020-12-15T15:35:44.898Z",
        "returned": true
      },
      {
        "book": {
          "id": "5e6e21ba-6fef-4a1c-b1c8-bc208992e6a1",
          "title": "Livro 02",
          "pages": "156",
          "createdAt": "2020-12-15T15:33:10.848Z"
        },
        "fromUser": "5fd8d7273722a1d48c225244",
        "toUser": "5fd8d7563722a1d48c225245",
        "lentAt": "2020-12-15T15:36:17.056Z",
        "returnedAt": "0001-01-01T00:00:00Z",
        "returned": false
      }
    ],
    "borrowedBooks": []
  }
}
```

### Criar um novo usuário 

METÓDO `POST`

    /v1/createUser

**Body** (application/json): 

* createUser: Nome e email do usuário a ser criado

#### Exemplo

No terminal execute

    curl --location --request POST 'http://localhost:8082/v1/createUser' \
    --header 'Content-Type: application/json' \
    --data-raw '{
        "name": "user", 
	      "email": "user@email.com.br"
      }'

Ou em alguma outra interface tipo Postman ou Insomnia

    {
	    "name": "user", 
	    "email": "user@email.com.br"
    }

Saída esperada

_Status http_: `200 OK`
```json
{
  "status": "success",
  "data": {
    "id": "5fd8d7563722a1d48c225245",
    "name": "user",
    "email": "user@email.com.br",
    "createdAt": "2020-12-15T15:33:42.626Z",
    "collection": [],
    "lentBooks": [],
    "borrowedBooks": []
  }
}
```

### Adicionar um novo livro 

METÓDO `POST`

    /v1/addBook/:userid

**Parametros**: 

* userid: Id do usuário onde vai ser adicionado o livro   

**Body** (application/json): 

* addBook: Adiciona um novo livro ao usuario

#### Exemplo

No terminal execute

    curl --location --request POST 'http://localhost:8082/v1/addBook/33' \
    --header 'Content-Type: application/json' \
    --data-raw '{
	      "title": "Livro 02",
	      "pages": "156"
      }'

Ou em alguma outra interface tipo Postman ou Insomnia

      {
	      "title": "Livro 02",
	      "pages": "156"
      }

Saída esperada

_Status http_: `200 OK`
```json
{
  "status": "success",
  "data": {
    "id": "5e6e21ba-6fef-4a1c-b1c8-bc208992e6a1",
    "title": "Livro 02",
    "pages": "156",
    "createdAt": "2020-12-15T12:33:10.848192737-03:00"
  }
}
```

### Emprestar um livro

METÓDO `POST`

    /v1/lendBook/:userid

**Parametros**: 

* userid: Id do usuário que vai emprestar o livro    

**Body** (application/json): 

* lendBook: Empresta um livro de sua coleção para um usuário

#### Exemplo

No terminal execute

    curl --location --request POST 'http://localhost:8082/v1/lendBook/33' \
    --header 'Content-Type: application/json' \
    --data-raw '{
	      "bookId": "5e6e21ba-6fef-4a1c-b1c8-bc208992e6a1",
	      "toUserId": "5fd8d7563722a1d48c225245"
      }'

Ou em aluma outra interface tipo Postman ou Insomnia

    {
	    "bookId": "5e6e21ba-6fef-4a1c-b1c8-bc208992e6a1",
	    "toUserId": "5fd8d7563722a1d48c225245"
    }

Saída esperada

_Status http_: `200 OK`
```json
{
  "status": "success",
  "data": {
    "book": {
      "id": "5e6e21ba-6fef-4a1c-b1c8-bc208992e6a1",
      "title": "Livro 02",
      "pages": "156",
      "createdAt": "2020-12-15T15:33:10.848Z"
    },
    "fromUser": "5fd8d7273722a1d48c225244",
    "toUser": "5fd8d7563722a1d48c225245",
    "lentAt": "2020-12-15T15:34:29.006Z",
    "returnedAt": "0-0-0T0:0:0.0-03:00",
    "returned": false
  }
}
```

### Devolver um livro

METÓDO `POST`

    /v1/returnBook/:userid/:bookid

**Parametros**: 

* userid: Id do usuário que vai devolver o livro    
* bookid: Id do livro a ser devolvido

**Body** (application/json): 

* lendBook: Empresta um livro de sua coleção para um usuário

#### Exemplo

No terminal execute

    curl --location --request POST 'http://localhost:8082/v1/returnBook/34/3823' \

Ou em aluma outra interface tipo Postman ou Insomnia

Saída esperada

_Status http_: `200 OK`
```json
{
  "status": "success",
  "data": {
    "book": {
      "id": "5e6e21ba-6fef-4a1c-b1c8-bc208992e6a1",
      "title": "Livro 02",
      "pages": "156",
      "createdAt": "2020-12-15T15:33:10.848Z"
    },
    "fromUser": "5fd8d7273722a1d48c225244",
    "toUser": "5fd8d7563722a1d48c225245",
    "lentAt": "2020-12-15T15:34:29.006Z",
    "returnedAt": "2020-12-15T12:35:44.898687163-03:00",
    "returned": true
  }
}
```

### Requerementos

* GoLang 1:15

# Contato

E-mail: viniiciuscardozo@gmail.com
LinkedIn: https://www.linkedin.com/in/vinicius-cardozo-669a15136/