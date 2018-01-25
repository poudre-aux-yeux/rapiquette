rAPIquette

# Ecosystem

Front-ends :

[MonArbitreRaquette](https://github.com/poudre-aux-yeux/mon-arbitre-raquette) : Application Web pour l'arbitre
MatteMaRaquette : Application Web pour consulter en direct les résultats
GereMaRaquette : Application Web de gestion (back-office) des joueurs, terrains, matchs ...


# Getting started

``` sh
go get
go build
./rapiquette
```

# Deploy

Set the `GIN_MODE` environment variable to `release`

# API spec

```
/api
    /mails
        /organisation : POST
        /public : POST
        /presse : POST
    /users : CRUD
    /terrains : CRUD
    /arbitres : CRUD
    /lieu : CRUD
    /matches : GET
        /:id : GET (socket)
    /arbitrage
        /scores : CRUD
        /appel médecin : POST
        /enregistre faute : POST
```
