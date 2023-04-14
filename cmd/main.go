package main

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/kilianp07/CassandraCRUD/pkg/cassandra"
)

func main() {

	// Read .env file
	err := godotenv.Load("../.env")
	if err != nil {
		panic(err)
	}

	// Initialize cassandra
	cassandra, err := cassandra.NewCassandra(os.Getenv("CASSANDRA_HOST"), os.Getenv("CASSANDRA_USERNAME"), os.Getenv("CASSANDRA_PASSWORD"))
	if err != nil {
		panic(err)
	}
	cassandra.Initialize()
}
