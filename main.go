package main

import (
	"github.com/kilianp07/CassandraCRUD/pkg/api"
	"github.com/kilianp07/CassandraCRUD/pkg/cassandra"
	envretriever "github.com/kilianp07/CassandraCRUD/utils/envRetriever"
)

func main() {

	// Read .env file
	vars := envretriever.GetEnvVars()

	// Test cassandra connection
	cassandra, err := cassandra.NewCassandra(vars.CassandraHost, vars.CassandraUsername, vars.CassandraPassword, "contacts")
	if err != nil {
		panic(err)
	}
	cassandra.Close()

	api.Start()
}
