Translations:

* [French](README_fr.md)
* [Portuguese (Brazil)](README_pt_br.md)

# 🔨 Basic Auction System (auction-go)

![Project Logo](assets/auction-logo.png)

Welcome to the basic auction system developed in Go! This project allows you to create and consult auctions, users, and bids, as well as determine the winner of an auction.

## 📑&nbsp;Table of Contents

- [📖 Introduction](#introduction)
- [🛠 Prerequisites](#prerequisites)
- [⚙️ Installation](#installation)
- [🚀 Usage](#usage)
- [🔍 Examples](#examples)
- [🤝 Contribution](#contribution)
- [📜 License](#license)

## 📖&nbsp;Introduction

This basic auction system is a project developed in Go that allows the creation and consultation of auctions, users, and bids. It provides a simple and efficient way to manage auctions and determine the winner based on received bids.

## 🛠&nbsp;Prerequisites

Make sure you have the following items installed before continuing:

- [Go](https://golang.org/doc/install)
- [Docker](https://www.docker.com/get-started)
- [Docker Compose](https://docs.docker.com/compose/install/)

## ⚙️&nbsp;Installation

1. Clone this repository:

    ```sh
    git clone git@github.com:rodrigoachilles/auction-go.git
    cd auction-go
    ```

2. Run Docker Compose:

    ```sh
    docker-compose up -d
    ```

## 🚀&nbsp;Usage

After starting Docker Compose, you can use the API to create and consult auctions, users, and bids.

### 🔧&nbsp;Running Services

1. Navigate to the main project folder:

    ```sh
    cd auction-go
    ```

2. Run the Go server:

    ```sh
    go run cmd/auction/main.go
    ```

### 📚&nbsp;Available Endpoints

#### Create User

```sh
POST /users
```

- Body (JSON):

```json
{
   "name": "John Doe"
}
```

#### Get User

```sh
GET /users/{id}
```

#### List Users

```sh
GET /users
```

#### Create Auction

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

#### Get Auction

```sh
GET /auctions/{id}
```

#### List Auctions

```sh
GET /auctions?status=0&category=""&productName=""
```

#### Determine Auction Winner

```sh
GET /auctions/{id}/winner
```

#### Get Auction Bids

```sh
GET /auctions/{id}/bids
```

#### Create Bid

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

## 🔍&nbsp;Examples

Here are some usage examples of the auction system endpoints:

- Create a new user and consult their information.
- Create a new auction and add bids.
- Consult all bids of a specific auction and determine the winner.

## 🤝&nbsp;Contribution

Feel free to open issues or submit pull requests for improvements and bug fixes.

## 📜&nbsp;License

This project is licensed under the MIT License.
