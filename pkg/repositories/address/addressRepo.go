package addressRepo

import (
	"github.com/kilianp07/CassandraCRUD/pkg/cassandra"
	envretriever "github.com/kilianp07/CassandraCRUD/utils/envRetriever"
	"github.com/kilianp07/CassandraCRUD/utils/structs"
)

func Create(address *structs.Address) error {
	vars := envretriever.GetEnvVars()

	c, err := cassandra.NewCassandra(vars.CassandraHost, vars.CassandraUsername, vars.CassandraPassword, vars.CassandraKeyspace)
	if err != nil {
		return err
	}
	defer c.Session.Close()

	query := "INSERT INTO addresses (address_id, building, street, zipcode, restaurant_id) VALUES (?, ?, ?, ?, ?)"
	if err := c.Session.Query(query, address.Id, address.Building, address.Street, address.Zipcode, address.RestaurantId).Exec(); err != nil {
		return err
	}
	return nil
}
