# rAPIquette

raquette + API = rAPIquette

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
[GraphQL](https://graphql.org) is a Query Language for APIs, and aims to replace
RESTful APIs.
[Redis](https://redis.io/) is an in-memory key-value store.

# Getting started

## Run the API

``` sh
docker-compose up
```

## DEBUG

Get and build the dependencies:

``` sh
# Install go generate
go get -u github.com/jteeuwen/go-bindata/...
go get -u github.com/golang/dep/cmd/dep
# Generate the schema builder
cd schema && go generate && cd ..
# Install the dependencies
dep ensure
# If using cmd.exe, omit the './'
./build.sh
```

Start the databases:

``` sh
docker-compose up tennis-redis raquette-redis
```

Start the API:

``` sh
# Unix
RAQUETTE_HOST=localhost:6380
TENNIS_HOST=localhost:6379
./rapiquette
```

``` powershell
# Powershell
$env:RAQUETTE_HOST="localhost:6380"
$env:TENNIS_HOST="localhost:6379"
.\rapiquette
```

``` batch
:: cmd.exe
SET RAQUETTE_HOST=localhost:6380
SET TENNIS_HOST=localhost:6379
rapiquette
```

# Deploy

Set the `GIN_MODE` environment variable to `release`.
Set the `RAQUETTE_HOST` and `TENNIS_HOST` environment variables to your Redis
location (can be the same instance or 2 different databases).

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

## Examples

Create a new player, with the following query:

``` graphql
mutation CreateNewPlayer($player: CreatePlayerInput!) {
  createPlayer(player: $player) {
    id
    name
  }
}
```

and the Query Variables:

``` graphql
{
    "player": {
        "name": "Thomasauvajon"
    }
}
```

You can then query the players:

``` graphql
query {
  players {
    name
    id
  }
}
```

or a particular player:

``` graphql
query GetPlayerByID($id: ID!) {
  player(id: $id) {
    id
    name
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
  createMatch(date: $date, homePlayers: $homePlayers, awayPlayers: $awayPlayers, referee: $referee, stadium: $stadium) {
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
    }
    awayPlayers {
      id
      name
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
