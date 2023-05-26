package envretriever

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

type envVars struct {
	CassandraHost     string
	CassandraUsername string
	CassandraPassword string
}

func GetEnvVars() (envVars, error) {
	// Get the path of the currently executing binary
	executablePath, err := os.Executable()
	if err != nil {
		panic(err)
	}
	// Derive the path of the .env file relative to the binary path
	envFilePath := filepath.Join(filepath.Dir(executablePath), ".env")

	// Read .env file
	err = godotenv.Load(envFilePath)
	if err == nil {
		return envVars{
			CassandraHost:     os.Getenv("CASSANDRA_HOST"),
			CassandraUsername: os.Getenv("CASSANDRA_USERNAME"),
			CassandraPassword: os.Getenv("CASSANDRA_PASSWORD"),
		}, nil
	}

	var requiredEnv = []string{
		"CASSANDRA_HOST",
		"CASSANDRA_USERNAME",
		"CASSANDRA_PASSWORD",
	}

	for _, key := range requiredEnv {
		_, found := os.LookupEnv(key)
		if !found {
			fmt.Printf("Missing %s key", key)
			return envVars{}, fmt.Errorf("Missing %s key", key)
		}
	}

	return envVars{
		CassandraHost:     os.Getenv("CASSANDRA_HOST"),
		CassandraUsername: os.Getenv("CASSANDRA_USERNAME"),
		CassandraPassword: os.Getenv("CASSANDRA_PASSWORD"),
	}, nil
}
