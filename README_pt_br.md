Traduções:

* [Inglês](README.md)
* [Francês](README_fr.md)

---

# 🔨 Sistema Básico de Leilão (auction-go)

![Project Logo](assets/auction-logo.png)

Bem-vindo ao sistema básico de leilão desenvolvido em Go! Este projeto permite criar e consultar leilões, usuários e lances (bids), além de determinar o vencedor de um leilão.

## 📑&nbsp;Sumário

- [📖 Introdução](#introdução)
- [🛠 Pré-requisitos](#pré-requisitos)
- [⚙️ Instalação](#instalação)
- [🚀 Uso](#uso)
- [🔍 Exemplos](#exemplos)
- [🤝 Contribuição](#contribuição)
- [📜 Licença](#licença)

## 📖&nbsp;Introdução

Este sistema básico de leilão é um projeto desenvolvido em Go que permite a criação e consulta de leilões, usuários e lances. Ele fornece uma maneira simples e eficiente de gerenciar leilões e determinar o vencedor com base nos lances recebidos.

## 🛠&nbsp;Pré-requisitos

Certifique-se de ter os seguintes itens instalados antes de continuar:

- [Go](https://golang.org/doc/install)
- [Docker](https://www.docker.com/get-started)
- [Docker Compose](https://docs.docker.com/compose/install/)

## ⚙️&nbsp;Instalação

1. Clone este repositório:

    ```sh
    git clone git@github.com:rodrigoachilles/auction-go.git
    cd auction-go
    ```

2. Execute o Docker Compose:

    ```sh
    docker-compose up -d
    ```

## 🚀&nbsp;Uso

Após iniciar o Docker Compose, você pode usar a API para criar e consultar leilões, usuários e lances.

### 🔧&nbsp;Executando Serviços

1. Navegue até a pasta principal do projeto:

    ```sh
    cd auction-go
    ```

2. Execute o servidor Go:

    ```sh
    go run cmd/auction/main.go
    ```

### 📚&nbsp;Endpoints Disponíveis

#### Criar Usuário

```sh
POST /users
```

- Body (JSON):

```json
{
   "name": "John Doe"
}
```

#### Consultar Usuário

```sh
GET /users/{id}
```

#### Consultar Usuários

```sh
GET /users
```

#### Criar Leilão

```sh
POST /auctions
```

- Body (JSON):

```json
{
   "product_name": "Product Name",
   "category": "Category",
   "description": "Product Description",
   "condition": 0
}
```

#### Consultar Leilão

```sh
GET /auctions/{id}
```

#### Consultar Leilões

```sh
GET /auctions?status=0&category=""&productName=""
```

#### Determinar Vencedor do Leilão

```sh
GET /auctions/{id}/winner
```

#### Consultar Lances de um Leilão

```sh
GET /auctions/{id}/bids
```

#### Criar Lance

```sh
POST /bids
```

- Body (JSON):

```json
{
  "user_id": 1,
  "auction_id": 1,
  "amount": 100.0
}
```

## 🔍&nbsp;Exemplos

Aqui estão alguns exemplos de uso dos endpoints do sistema de leilão:

- Criar um novo usuário e consultar suas informações.
- Criar um novo leilão e adicionar lances.
- Consultar todos os lances de um leilão específico e determinar o vencedor.

## 🤝&nbsp;Contribuição

Sinta-se à vontade para abrir issues ou enviar pull requests para melhorias e correções de bugs.

## 📜&nbsp;Licença

Este projeto está licenciado sob a Licença MIT.
