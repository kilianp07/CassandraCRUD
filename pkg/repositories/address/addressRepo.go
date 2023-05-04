package addressRepo

import (
	"github.com/gocql/gocql"
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

func GetAll() ([]*structs.Address, error) {
	vars := envretriever.GetEnvVars()
	addresses := []*structs.Address{}

	c, err := cassandra.NewCassandra(vars.CassandraHost, vars.CassandraUsername, vars.CassandraPassword, vars.CassandraKeyspace)
	if err != nil {
		return nil, err
	}
	defer c.Session.Close()

	query := "SELECT address_id, building, street, zipcode, restaurant_id FROM addresses"
	iter := c.Session.Query(query).Iter()

	for {
		address := &structs.Address{}
		if !iter.Scan(&address.Id, &address.Building, &address.Street, &address.Zipcode, &address.RestaurantId) {
			break
		}
		addresses = append(addresses, address)
	}
	if err := iter.Close(); err != nil {
		return nil, err
	}

	return addresses, nil
}

func GetById(id string) (*structs.Address, error) {
	vars := envretriever.GetEnvVars()
	address := &structs.Address{}

	c, err := cassandra.NewCassandra(vars.CassandraHost, vars.CassandraUsername, vars.CassandraPassword, vars.CassandraKeyspace)
	if err != nil {
		return nil, err
	}
	defer c.Session.Close()

	query := "SELECT address_id, building, street, zipcode, restaurant_id FROM addresses WHERE address_id = ? LIMIT 1"
	if err := c.Session.Query(query, id).Scan(&address.Id, &address.Building, &address.Street, &address.Zipcode, &address.RestaurantId); err != nil {
		return nil, err
	}
	return address, nil
}

func GetByRestaurantId(restaurantId string, c *cassandra.Cassandra) (*structs.Address, error) {
	address := &structs.Address{}

	query := "SELECT address_id, building, street, zipcode, restaurant_id FROM addresses WHERE restaurant_id = ? LIMIT 1 ALLOW FILTERING"
	if err := c.Session.Query(query, restaurantId).Scan(&address.Id, &address.Building, &address.Street, &address.Zipcode, &address.RestaurantId); err != nil {
		if err == gocql.ErrNotFound {
			return nil, nil
		}
		return nil, err
	}
	return address, nil
}

func Update(address *structs.Address) (*structs.Address, error) {
	vars := envretriever.GetEnvVars()

	c, err := cassandra.NewCassandra(vars.CassandraHost, vars.CassandraUsername, vars.CassandraPassword, vars.CassandraKeyspace)
	if err != nil {
		return nil, err
	}
	defer c.Session.Close()

	query := "UPDATE addresses SET building = ?, street = ?, zipcode = ?, restaurant_id = ? WHERE address_id = ?"
	if err := c.Session.Query(query, address.Building, address.Street, address.Zipcode, address.RestaurantId, address.Id).Exec(); err != nil {
		return nil, err
	}
	return address, nil
}

func Delete(id string) error {
	vars := envretriever.GetEnvVars()

	c, err := cassandra.NewCassandra(vars.CassandraHost, vars.CassandraUsername, vars.CassandraPassword, vars.CassandraKeyspace)
	if err != nil {
		return err
	}
	defer c.Session.Close()

	query := "DELETE FROM addresses WHERE address_id = ?"
	if err := c.Session.Query(query, id).Exec(); err != nil {
		return err
	}
	return nil
}
