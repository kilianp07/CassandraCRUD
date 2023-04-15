package envretriever

import (
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

type envVars struct {
	CassandraHost     string
	CassandraUsername string
	CassandraPassword string
	CassandraKeyspace string
}

func GetEnvVars() envVars {
	// Get the path of the currently executing binary
	executablePath, err := os.Executable()
	if err != nil {
		panic(err)
	}
	// Derive the path of the .env file relative to the binary path
	envFilePath := filepath.Join(filepath.Dir(executablePath), ".env")

	// Read .env file
	err = godotenv.Load(envFilePath)
	if err != nil {
		panic(err)
	}
	return envVars{
		CassandraHost:     os.Getenv("CASSANDRA_HOST"),
		CassandraUsername: os.Getenv("CASSANDRA_USERNAME"),
		CassandraPassword: os.Getenv("CASSANDRA_PASSWORD"),
		CassandraKeyspace: os.Getenv("CASSANDRA_KEYSPACE"),
	}
}
