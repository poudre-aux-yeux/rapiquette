package handler

import "net/http"

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

var graphiql = []byte(`
	<!DOCTYPE html>
	<head>
	  <style>body {height: 100vh; margin: 0; width: 100%; overflow: hidden;}</style>
	  <link rel="stylesheet" href="//cdn.jsdelivr.net/sweetalert/1.1.3/sweetalert.css" />
	  <!--<link rel="stylesheet" href="//cdn.jsdelivr.net/graphiql/0.6.3/graphiql.css" />-->
	  <link rel="stylesheet" href="//cdnjs.cloudflare.com/ajax/libs/graphiql/0.11.11/graphiql.css" />
	  <script src="//cdn.jsdelivr.net/sweetalert/1.1.3/sweetalert.min.js"></script>
	  <script src="//cdn.jsdelivr.net/fetch/0.9.0/fetch.min.js"></script>
	  <script src="//cdn.jsdelivr.net/react/0.14.7/react.min.js"></script>
	  <script src="//cdn.jsdelivr.net/react/0.14.7/react-dom.min.js"></script>
	  <!-- <script src="//cdn.jsdelivr.net/graphiql/0.6.3/graphiql.min.js"></script> -->
	  <script src="//cdnjs.cloudflare.com/ajax/libs/graphiql/0.11.11/graphiql.min.js"></script>
	  <script>
		(function () {
		  var PROMPT_OPTIONS = {
			title: "GraphQL Endpoint",
			text: "Please give the GraphQL HTTP Endpoint",
			type: "input",
			showCancelButton: false,
			inputPlaceholder: window.location.origin + '/graphql',
		  };
		  document.addEventListener('DOMContentLoaded', function () {
			swal(PROMPT_OPTIONS, function(endpoint){
			  if (!endpoint) {
				endpoint = window.location.origin + '/graphql';
			  }
			  function fetcher(params) {
				var options = {
				  method: 'post',
				  headers: {'Accept': 'application/json', 'Content-Type': 'application/json'},
				  body: JSON.stringify(params),
				  credentials: 'include',
				};
				return fetch(endpoint, options)
				  .then(function (res) { return res.json() });
			  }
			  var body = React.createElement(GraphiQL, {fetcher: fetcher, query: '', variables: ''});
			  ReactDOM.render(body, document.body);
			});
		  });
		}());
	  </script>
	</head>
	<body>
	</body>	
`)
