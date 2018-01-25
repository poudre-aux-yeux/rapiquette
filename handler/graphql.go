package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"sync"

	graphql "github.com/neelance/graphql-go"
)

// A query represents a single GraphQL query.
type query struct {
	OpName    string                 `json:"operationName"`
	Query     string                 `json:"query"`
	Variables map[string]interface{} `json:"variables"`
}

// A request respresents an HTTP request to the GraphQL endpoint.
// A request can have a single query or a batch of requests with one or more queries.
// It is important to distinguish between a single query request and a batch request with a single query.
// The shape of the response will differ in both cases.
type request struct {
	queries []query
	isBatch bool
}

// The GraphQL handler handles GraphQL API requests over HTTP.
// It can handle batched requests as sent by the apollo-client.
type GraphQL struct {
	Schema *graphql.Schema
}

// ServeHTTP handles the GraphQL requets (queries and mutations)
func (g GraphQL) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	req, err := parse(r)
	if err != nil {
		respond(w, errorJSON(err.Error()), http.StatusBadRequest)
		return
	}

	n := len(req.queries)
	if n == 0 {
		respond(w, errorJSON("no queries to execute"), http.StatusBadRequest)
	}

	// TODO: add authentication

	var (
		ctx       = context.Background()
		responses = make([]*graphql.Response, n)
		wg        sync.WaitGroup
	)

	wg.Add(n)

	for i, q := range req.queries {
		// Queries are executed in separate goroutines so they process in parallel
		go func(i int, q query) {
			res := g.Schema.Exec(ctx, q.Query, q.OpName, q.Variables)

			// TODO Expand errors
			responses[i] = res
			wg.Done()
		}(i, q)
	}

	wg.Wait()

	// TODO: Error handling

	var resp []byte
	if req.isBatch {
		resp, err = json.Marshal(responses)
	} else if len(responses) > 0 {
		resp, err = json.Marshal(responses[0])
	}

	if err != nil {
		respond(w, errorJSON("server error"), http.StatusInternalServerError)
		return
	}

	respond(w, resp, http.StatusOK)
}
