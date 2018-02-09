rAPIquette

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
It uses the standard library and the
[neelance/graphql-go](https://github.com/neelance/graphql-go) GraphQL framework.
[GraphQL](https://graphql.org) is a Query Language for APIs, and aims to replace
RESTful APIs.

# Getting started

## Run the API

``` sh
docker-compose up
```

## DEBUG

Get and build the dependencies:

``` sh
# Install the dependencies
go get
# Install go generate
go get -u github.com/jteeuwen/go-bindata/...
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
location (can be the same instance of 2 different databases).

Or just run `docker-compose up`.

# API Schema

GraphQL Schema:

``` graphql
scalar Time

schema {
    query: Query
    mutation: Mutation
}

type Query {
    matches(): [Match!]!
    
    admin(id: ID!): Admin
    refereeRaquette(id: ID!): RefereeRaquette
    
    match(id: ID!): Match
    player(id: ID!): Player
    stadium(id: ID!): Stadium
    tennisReferee(id: ID!): TennisReferee
    set(id: ID!): Set
    game(id: ID!): Game
}

type Mutation {
    createMatch(date: Time!, players: [ID!]!, referee: ID!): Match
    startMatch(id: ID!): Match
}

interface User {
    id: ID!
    hash: String!
    username: String!
    email: String!
}

type Admin implements User {
    id: ID!
    hash: String!
    username: String!
    email: String!
}

type RefereeRaquette implements User {
    id: ID!
    hash: String!
    username: String!
    email: String!
}

type Match {
    id: ID!
    date: Time!
    stadium: Stadium!
    referee: TennisReferee!
    players: [Player!]!
    sets: [Set!]!
    service: Boolean
}

type Player {
    id: ID!
    name: String!
}

type Stadium {
    id: ID!
    name: String!
    city: String!
}

type TennisReferee {
    id: ID!
    name: String!
}

type Set {
    id: ID!
    games: [Game!]!
}

type Game {
    id: ID!
    homePoints: Int
    awayPoints: Int
}
```

# Consume the API

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
