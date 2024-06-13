Traductions:

* [Anglais](README.md)
* [Portugais (Brésil)](README_pt_br.md)

---

# 🔨 Système de Vente aux Enchères Basique (auction-go)

![Project Logo](assets/auction-logo.png)

Bienvenue dans le système de vente aux enchères basique développé en Go ! Ce projet permet de créer et consulter des enchères, des utilisateurs et des offres (bids), ainsi que de déterminer le gagnant d'une enchère.

## 📑&nbsp;Table des Matières

- [📖 Introduction](#introduction)
- [🛠 Prérequis](#prérequis)
- [⚙️ Installation](#installation)
- [🚀 Utilisation](#utilisation)
- [🔍 Exemples](#exemples)
- [🤝 Contribution](#contribution)
- [📜 Licence](#licence)

## 📖&nbsp;Introduction

Ce système de vente aux enchères basique est un projet développé en Go qui permet la création et la consultation des enchères, des utilisateurs et des offres. Il fournit un moyen simple et efficace de gérer les enchères et de déterminer le gagnant en fonction des offres reçues.

## 🛠&nbsp;Prérequis

Assurez-vous d'avoir les éléments suivants installés avant de continuer :

- [Go](https://golang.org/doc/install)
- [Docker](https://www.docker.com/get-started)
- [Docker Compose](https://docs.docker.com/compose/install/)

## ⚙️&nbsp;Installation

1. Clonez ce dépôt :

    ```sh
    git clone git@github.com:rodrigoachilles/auction-go.git
    cd auction-go
    ```

2. Exécutez Docker Compose :

    ```sh
    docker-compose up -d
    ```

## 🚀&nbsp;Utilisation

Après avoir démarré Docker Compose, vous pouvez utiliser l'API pour créer et consulter des enchères, des utilisateurs et des offres.

### 🔧&nbsp;Exécution des Services

1. Naviguez jusqu'au dossier principal du projet :

    ```sh
    cd auction-go
    ```

2. Exécutez le serveur Go :

    ```sh
    go run cmd/auction/main.go
    ```

### 📚&nbsp;Points de Terminaison Disponibles

#### Créer un Utilisateur

```sh
POST /users
```

- Corps (JSON) :

```json
{
   "name": "John Doe"
}
```

#### Consulter un Utilisateur

```sh
GET /users/{id}
```

#### Lister les Utilisateurs

```sh
GET /users
```

#### Créer une Enchère

```sh
POST /auctions
```

- Corps (JSON) :

```json
{
   "product_name": "Product Name",
   "category": "Category",
   "description": "Product Description",
   "condition": 0
}
```

#### Consulter une Enchère

```sh
GET /auctions/{id}
```

#### Lister les Enchères

```sh
GET /auctions?status=0&category=""&productName=""
```

#### Déterminer le Gagnant de l'Enchère

```sh
GET /auctions/{id}/winner
```

#### Consulter les Offres d'une Enchère

```sh
GET /auctions/{id}/bids
```

#### Créer une Offre

```sh
POST /bids
```

- Corps (JSON) :

```json
{
  "user_id": 1,
  "auction_id": 1,
  "amount": 100.0
}
```

## 🔍&nbsp;Exemples

Voici quelques exemples d'utilisation des points de terminaison du système de vente aux enchères :

- Créer un nouvel utilisateur et consulter ses informations.
- Créer une nouvelle enchère et ajouter des offres.
- Consulter toutes les offres d'une enchère spécifique et déterminer le gagnant.

## 🤝&nbsp;Contribution

N'hésitez pas à ouvrir des issues ou à soumettre des "pull requests" pour des améliorations et des corrections de bugs.

## 📜&nbsp;Licence

Ce projet est sous licence MIT.
