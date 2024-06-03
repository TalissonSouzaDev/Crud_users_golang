
# Crud com Golang

É um Crud Basico para um grande inicio na jornada com Golang, lista usuario, ler pelo ID, Criar, Edita e Deleta.

# Documentação da API

## Retorna todos os usuarios
##### Metodo: GET

```http
http://localhost:8000/userlist
```
## Retorna um usuario
##### Metodo: GET

```http
http://localhost:8000/showuser/{id}
```

| Parâmetro   | Tipo       | Descrição                                   |
| :---------- | :--------- | :------------------------------------------ |
| `id`      | `string` | **Obrigatório**. ID do usuario|

## Criação do usuario
##### Metodo: POST
```http
http://localhost:8000/createuser
```
##### Parâmetro Body  

```http
  {
  "name": "teste",
  "email": "teste@example.com"
  }
```

## Edição do usuario
##### Metodo: PUT
```http
http://localhost:8000/updateuser{id}
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
##### Metodo: DELETE

```http
http://localhost:8000/deleteuser/{id}
```

| Parâmetro   | Tipo       | Descrição                                   |
| :---------- | :--------- | :------------------------------------------ |
| `id`      | `string` | **Obrigatório**. ID do usuario|
