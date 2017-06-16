package graphql

var Schema = `
	schema {
		query: Query
	}
	# The query type, represents all of the entry points into our object graph
	type Query {
		CanIKick(It: String): String
	}
`
