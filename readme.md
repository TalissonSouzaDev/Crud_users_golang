
# Documentação da API

## Retorna todos os usuarios

```http
  GET http://localhost:8000/userlist
```
## Retorna um usuario

```http
  GET http://localhost:8000/showuser/{id}
```

| Parâmetro   | Tipo       | Descrição                                   |
| :---------- | :--------- | :------------------------------------------ |
| `id`      | `string` | **Obrigatório**. ID do usuario|

## Criação do usuario
```http
  POST http://localhost:8000/createuser
```
##### Parâmetro Body  

```http
  {
  "name": "teste",
  "email": "teste@example.com"
  }
```

## Edição do usuario
```http
  PUT http://localhost:8000/updateuser{id}
```
| Parâmetro   | Tipo       | Descrição                                   |
| :---------- | :--------- | :------------------------------------------ |
| `id`      | `string` | **Obrigatório**. ID do usuario|

##### Parâmetro Body  

```http
  {
  "name": "teste",
  "email": "teste@example.com"
  }
```

## Deleta um usuario

```http
  DELETE http://localhost:8000/deleteuser/{id}
```

| Parâmetro   | Tipo       | Descrição                                   |
| :---------- | :--------- | :------------------------------------------ |
| `id`      | `string` | **Obrigatório**. ID do usuario|
