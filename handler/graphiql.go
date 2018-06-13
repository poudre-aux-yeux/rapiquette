package handler

import (
	"bytes"
	"fmt"
	"net/http"
)

// GraphiQL is an in-browser IDE for exploring GraphiQL APIs.
// This handler returns GraphiQL when requested.
//
// For more information, see https://github.com/graphql/graphiql.
type GraphiQL struct{}

func (h GraphiQL) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		respond(w, errorJSON("only GET requests are supported"), http.StatusMethodNotAllowed)
		return
	}

	w.Write(graphiql)
}

func respond(w http.ResponseWriter, body []byte, code int) {
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(code)
	_, _ = w.Write(body)
}

func errorJSON(msg string) []byte {
	buf := bytes.Buffer{}
	fmt.Fprintf(&buf, `{"error": "%s"}`, msg)
	return buf.Bytes()
}

var graphiql = []byte(`
<!DOCTYPE html>
<html>
	<head>
		<link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/graphiql/0.11.10/graphiql.css" />
		<script src="https://cdnjs.cloudflare.com/ajax/libs/fetch/1.1.0/fetch.min.js"></script>
		<script src="https://cdnjs.cloudflare.com/ajax/libs/react/15.5.4/react.min.js"></script>
		<script src="https://cdnjs.cloudflare.com/ajax/libs/react/15.5.4/react-dom.min.js"></script>
		<script src="https://cdnjs.cloudflare.com/ajax/libs/graphiql/0.11.10/graphiql.js"></script>
		<script src="//unpkg.com/subscriptions-transport-ws@0.8.3/browser/client.js"></script>
		<script src="//unpkg.com/graphiql-subscriptions-fetcher@0.0.2/browser/client.js"></script>
	</head>
	<body style="width: 100%; height: 100%; margin: 0; overflow: hidden;">
		<div id="graphiql" style="height: 100vh;">Loading...</div>
		<script>
			function graphQLFetcher(graphQLParams) {
				return fetch("/graphql", {
					method: "post",
					body: JSON.stringify(graphQLParams),
					credentials: "include",
				}).then(function (response) {
					return response.text();
				}).then(function (responseBody) {
					try {
						return JSON.parse(responseBody);
					} catch (error) {
						return responseBody;
					}
				});
			}
			var subscriptionsClient = new window.SubscriptionsTransportWs.SubscriptionClient('ws://localhost:3333/graphql', { reconnect: true });
			var subscriptionsFetcher = window.GraphiQLSubscriptionsFetcher.graphQLFetcher(subscriptionsClient, graphQLFetcher);
			ReactDOM.render(
				React.createElement(GraphiQL, {fetcher: subscriptionsFetcher}),
				document.getElementById("graphiql")
			);
		</script>
	</body>
</html>
`)
