# rAPIquette

raquette + API = rAPIquette

Get the Docker image:

```sh
docker pull poudreauxyeux/rapiquette
```

# Ecosystem

**Back-end:**

[rAPIquette](https://github.com/poudre-aux-yeux/rapiquette): GraphQL API

**Front-ends:**

[MonArbitreRaquette](https://github.com/poudre-aux-yeux/mon-arbitre-raquette):
Web application for the referee

[MatteMaRaquette](https://github.com/poudre-aux-yeux/ATP_LIVE):
Web application to watch the live scores and results

GereMaRaquette: Back-office Web application to administrate accounts, players,
matches, stadiums ...

# Languages and frameworks

The back-end is developped in [Go](https://golang.org/), an open-source
programming language created at Google.
It uses the standard library, the
[neelance/graphql-go](https://github.com/neelance/graphql-go) GraphQL framework
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

## Debugging

Get the repository ready for building:

``` sh
# Install mage
go get -u github.com/magefile/mage
# Initialize the repo
mage setup

# Note : Check all available 'mage' commands
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

[![GraphiQL screenshot](https://raw.githubusercontent.com/graphql/graphiql/master/resources/graphiql.png)](http://graphql.org/swapi-graphql)

To use it run `docker-compose up` and browse `http://localhost:3333`.

## Apollo

[Apollo](https://www.apollographql.com) provides GraphQL clients for React,
Vue.js, Angular, Android, iOS and other frontend platforms.
It is a good choice to query this API with an Apollo Client.

## Other clients

Check the "official" list at http://graphql.org/code/#graphql-clients

## Custom client

[GraphQL clients documentation](http://graphql.org/graphql-js/graphql-clients/)

## Examples

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
    "name": "Thomas Sauvajon",
    "image": "https://media-exp2.licdn.com/mpr/mpr/shrinknp_400_400/AAMABADGAAwAAQAAAAAAAAwuAAAAJDI0MTJkMmNiLThkYTQtNDhkMC1iNzM4LTdkNjcxYzc1Y2RlZA.jpg",
    "birth": "1993-04-10",
    "nationality": "FRA",
    "weight": 72,
    "ranking": 999,
    "titles": 1,
    "height": 175
  }
}
```

You can then query the players:

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
  $date: Time!,
  $homePlayers: [ID!]!,
  $awayPlayers: [ID!]!,
  $referee: ID!,
  $stadium: ID!
) {
  createMatch(
    date: $date,
    homePlayers: $homePlayers,
    awayPlayers: $awayPlayers,
    referee: $referee,
    stadium: $stadium
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
  "date": "2018-11-10T23:00:00Z",
  "homePlayers": ["10JEIUJTiophMWzEGqaAb3fvMY3"],
  "awayPlayers": ["10JEOE9xEjIb5LpQzKIGfgvIsJo"],
  "referee": "10JG02R4lHisO0T3I5r2FclC10s",
  "stadium": "10JDckheW3F2AcXNNdkfsOXz4zM"
}
```
