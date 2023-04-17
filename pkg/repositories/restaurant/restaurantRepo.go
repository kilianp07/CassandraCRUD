package restaurantRepo

import (
	"github.com/kilianp07/CassandraCRUD/pkg/cassandra"
	envretriever "github.com/kilianp07/CassandraCRUD/utils/envRetriever"
	"github.com/kilianp07/CassandraCRUD/utils/structs"
)

// Get restaurant from cassandra DB using go-cql-driver
func GetRestaurantById(id int) (restaurant *structs.Restaurant, err error) {
	vars := envretriever.GetEnvVars()

	c, err := cassandra.NewCassandra(vars.CassandraHost, vars.CassandraUsername, vars.CassandraPassword, vars.CassandraKeyspace)
	defer c.Session.Close()

	// Query
	query := "SELECT * FROM restaurant WHERE id = ?"
	if err := c.Session.Query(query, id).Scan(&restaurant); err != nil {
		return nil, err
	}
	return restaurant, nil
}
