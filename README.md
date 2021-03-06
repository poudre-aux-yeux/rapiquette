# rAPIquette

[![Go Report Card](https://goreportcard.com/badge/github.com/poudre-aux-yeux/rapiquette)](https://goreportcard.com/report/github.com/poudre-aux-yeux/rapiquette)
[![GoDoc](https://camo.githubusercontent.com/ac30242392a5470effdd2b008d7be055b6f6f8d6/68747470733a2f2f696d672e736869656c64732e696f2f62616467652f676f646f632d7265666572656e63652d626c75652e7376673f7374796c653d666c61742d737175617265)](https://godoc.org/github.com/poudre-aux-yeux/rapiquette)

raquette + API = rAPIquette (Raquette is the French for tennis racket).

Get the Docker image:

```sh
docker pull poudreauxyeux/rapiquette
```

# Ecosystem

**Main repo:**
[Matte Ma Raquette](https://github.com/poudre-aux-yeux/matte-ma-raquette/)

**Back-end:**

[rAPIquette](https://github.com/poudre-aux-yeux/rapiquette): GraphQL API

**Front-ends:**

[RefeGreen](https://github.com/poudre-aux-yeux/refegreen):  
Allow referees to live update scores. Used by the referees.

[MatteMaRaquette](https://github.com/poudre-aux-yeux/ATP_LIVE):  
Web application to watch live scores and results. Used by tennis fans.

[MonAdminRaquette](https://github.com/poudre-aux-yeux/mon-admin-raquette):  
Back-office Web application to manage accounts, players, matches, stadiums... Used by the ATP (Association of Tennis Professionals) staff.

# Languages and frameworks

The back-end is developped in [Go](https://golang.org/), an open-source
programming language created at Google.
It uses the standard library, the
[poudre-aux-yeux/graphql-go](https://github.com/poudre-aux-yeux/graphql-go) GraphQL framework
and the [Redigo](https://github.com/garyburd/redigo) Redis client.
[GraphQL](https://graphql.org) is a Query Language for APIs, and aims to
replace RESTful APIs.
[Redis](https://redis.io/) is an in-memory key-value store.

# Getting started

## Requirements

You need to have [Docker](https://docs.docker.com/install/) (> 17.05) installed
and running to use Docker Compose with this repository.

You also need to have a [Go](https://golang.org/doc/install) installation
if you want to build the repository.

The repository needs to be located at
`$GOPATH/src/github.com/poudre-aux-yeux/rapiquette` or else it won't compile.

## Run the API

``` sh
docker-compose up
```

## Debugging / Development

Get the repository ready for building:

``` sh
# Install mage
go get -u github.com/magefile/mage
# Initialize the repo
mage setup

# Note: Check all available 'mage' commands
mage -l
```

Debug:

``` sh
# This will run the 2 database containers and set the necessary env variables
mage databases

# Start the API
# Unix
./rapiquette
# Windows
rapiquette.exe
```

# Deploy

Set the `GIN_MODE` environment variable to `release`.
Set the `RAQUETTE_HOST` and `TENNIS_HOST` environment variables to your Redis
location (can be the same instance or 2 distinct databases).

Or just run `docker-compose up`.

# API Schema

## GraphQL Schema:

To read the GraphQL schema documentation,
run `docker-compose up` and browse `http://localhost:3333`.

# Consume the API

## GraphiQL

[GraphiQL](https://github.com/graphql/graphiql) is a graphical interactive
in-browser GraphQL IDE.

![GraphiQL screenshot](https://raw.githubusercontent.com/graphql/graphiql/master/resources/graphiql.png)
To use it run `docker-compose up` and browse `http://localhost:3333`.

## Apollo

[Apollo](https://www.apollographql.com) provides GraphQL clients for React,
Vue.js, Angular, Android, iOS and other frontend platforms.
It is a good choice to query this API with an Apollo Client, but other solutions exist.

## Other clients

Check the "official" list at https://graphql.org/code/#graphql-clients

## Custom client

But what is a GraphQL client and how does it work? Check the [GraphQL clients documentation](https://graphql.org/graphql-js/graphql-clients/).

## Examples

You can run these queries in GraphiQL or with any GraphQL client. Some of them use [Query Variables](https://graphql.org/learn/queries/#variables).  


Create a new player, with the following query:

``` graphql
mutation CreateNewPlayer($player: CreatePlayerInput!) {
  createPlayer(player: $player) {
    id
    name
    image
    birth
    nationality
    weight
    ranking
    titles
    height
  }
}
```

and the Query Variables:

``` graphql
{
  "player": {
    "name": "Somas Thauvajon",
    "image": "https://media-exp2.licdn.com/mpr/mpr/shrinknp_400_400/AAMABADGAAwAAQAAAAAAAAwuAAAAJDI0MTJkMmNiLThkYTQtNDhkMC1iNzM4LTdkNjcxYzc1Y2RlZA.jpg",
    "birth": "1993-04-10T00:00:00Z",
    "nationality": "FRA",
    "weight": 72,
    "ranking": 999,
    "titles": 1,
    "height": 175
  }
}
```

You can then query all the players:

``` graphql
query {
  players {
    id
    name
    image
    birth
    nationality
    weight
    ranking
    titles
    height
  }
}
```

or a particular player:

``` graphql
query GetPlayerByID($id: ID!) {
  player(id: $id) {
    id
    name
    image
    birth
    nationality
    weight
    ranking
    titles
    height
  }
}
```

with the Query Variables:

``` graphql
{
  	"id": "10Abbwj2Zg8sETDPUq99I52FiuA"
}
```

You can also update an existing player:

``` graphql
mutation UpdateExistingPlayer($player: UpdatePlayerInput!) {
  updatePlayer(player: $player) {
    id
    name
    image
    birth
    nationality
    weight
    ranking
    titles
    height
  }
}

# Query variables
{
  "player": {
    "id": "10wl445Jm4nsiFsvZ7JRGIRRA2T",
    "name": "Buillaume GAECHLER",
    "image": "https://media.licdn.com/mpr/mpr/shrinknp_200_200/AAEAAQAAAAAAAASvAAAAJDQyOTA5MjYwLWQ1OTEtNDEwNS1hYjc4LWJjZmJjYmY5MTM1Yw.jpg",
    "birth": "1996-06-28T00:00:00Z",
    "nationality": "FRA",
    "weight": 60,
    "ranking": 988,
    "titles": 1,
    "height": 190
  }
}
```

Create a stadium:

``` graphql
mutation CreateNewStadium($stadium: CreateStadiumInput!) {
  createStadium(stadium: $stadium) {
    id
    name
    city
    surface
  }
}

# Query Variables
{
  "stadium": {
    "GroundType": "clay",
    "Name": "Center Court",
    "City": "Wimbledon, London"
  }
}
```

Create a referee:

``` graphql
mutation CreateNewReferee($ref: CreateTennisRefereeInput!) {
  createTennisReferee(referee: $ref) {
    id
    name
  }
}

# Query Variables
{
  "ref": {
    "name": "James Keothavong"
  }
}
```

Create a match:

``` graphql
mutation CreateNewMatch(
  $match: CreateMatchInput!
) {
  createMatch(
    match: $match
  ) {
    id
    date
    stadium {
      id
      name
      city
      surface
    }
    referee {
      id
      name
    }
    homePlayers {
      id
      name
      image
      birth
      nationality
      weight
      ranking
      titles
      height
    }
    awayPlayers {
      id
      name
      image
      birth
      nationality
      weight
      ranking
      titles
      height
    }
  }
}

# Query Variables

{
  "match": {
    "date": "2018-11-10T23:00:00Z",
    "homePlayersLinks": ["10JEIUJTiophMWzEGqaAb3fvMY3"],
    "awayPlayersLinks": ["10JEOE9xEjIb5LpQzKIGfgvIsJo"],
    "refLink": "10JG02R4lHisO0T3I5r2FclC10s",
    "stdLink": "10JDckheW3F2AcXNNdkfsOXz4zM"
  }
}
```

Subscribe to points scored:

``` graphql
subscription {
  pointScored  {
    match {
      id
      date
    }
    team
  }
}
```

Score a point:

``` graphql
mutation ScorePoint(
  $pt: ScorePointInput!
) {
  scorePoint(
    point: $pt
  ) {
    match {
      id
    }
    team
  }
}

# Query variables

{
  "pt": {
    "matchID": "160COqwzM98vhRshJ4vxt2Y5ELF",
    "team": true 
  }
}

# note: team: true when team A scores, team: false when team B scores
```
