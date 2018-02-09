rAPIquette

# Ecosystem

Front-ends :

[MonArbitreRaquette](https://github.com/poudre-aux-yeux/mon-arbitre-raquette) : Application Web pour l'arbitre

MatteMaRaquette : Application Web pour consulter en direct les résultats

GereMaRaquette : Application Web de gestion (back-office) des joueurs, terrains, matchs ...


# Getting started

``` sh
# Install the dependencies
go get
# Install go generate
go get -u github.com/jteeuwen/go-bindata/...
# If using cmd.exe, omit the './'
./build.sh
# Run the executable
./rapiquette
```

# Deploy

Set the `GIN_MODE` environment variable to `release`
Set the `REDIS_HOST` environment variable to your Redis location

Or just run docker compose up

# API

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

## Mutations

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
