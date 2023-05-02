package restaurantRepo

import (
	"github.com/kilianp07/CassandraCRUD/pkg/cassandra"
	envretriever "github.com/kilianp07/CassandraCRUD/utils/envRetriever"
	"github.com/kilianp07/CassandraCRUD/utils/structs"
)

// This function is used to retrieve a restaurant by its ID.
// It is used by the /restaurants/{id} endpoint.
// The restaurant ID is passed in as a parameter named id.
// The function returns a pointer to a Restaurant struct.
// If the restaurant with the given ID is not found, it returns a nil pointer and an error.
func GetById(id string) (*structs.Restaurant, error) {
	vars := envretriever.GetEnvVars()
	restaurant := &structs.Restaurant{}

	c, err := cassandra.NewCassandra(vars.CassandraHost, vars.CassandraUsername, vars.CassandraPassword, vars.CassandraKeyspace)
	if err != nil {
		return nil, err
	}
	defer c.Session.Close()

	query := "SELECT restaurant_id, borough, cuisine, name FROM restaurants WHERE restaurant_id = ?"
	if err := c.Session.Query(query, id).Scan(&restaurant.Id, &restaurant.Borough, &restaurant.Cuisine, &restaurant.Name); err != nil {
		return nil, err
	}
	return restaurant, nil
}

// GetAll fetches all restaurants from the database
// Returns a slice of restaurant structs
// The slice is empty if no restaurants are in the database
func GetAll() ([]*structs.Restaurant, error) {
	vars := envretriever.GetEnvVars()
	restaurants := []*structs.Restaurant{}

	c, err := cassandra.NewCassandra(vars.CassandraHost, vars.CassandraUsername, vars.CassandraPassword, vars.CassandraKeyspace)
	if err != nil {
		return nil, err
	}
	defer c.Session.Close()

	query := "SELECT restaurant_id, borough, cuisine, name FROM restaurants"
	iter := c.Session.Query(query).Iter()

	for {
		restaurant := &structs.Restaurant{}
		if !iter.Scan(&restaurant.Id, &restaurant.Borough, &restaurant.Cuisine, &restaurant.Name) {
			break
		}
		restaurants = append(restaurants, restaurant)
	}
	if err := iter.Close(); err != nil {
		return nil, err
	}

	return restaurants, nil
}

// This function is used to create a new restaurant in the database.
// It takes a pointer to a Restaurant struct, which contains the information about the new restaurant.
// It then uses the cassandra-go-driver to create a connection to the Cassandra database, and then inserts the new restaurant into the restaurants table.
func Create(restaurant *structs.Restaurant) error {
	vars := envretriever.GetEnvVars()

	c, err := cassandra.NewCassandra(vars.CassandraHost, vars.CassandraUsername, vars.CassandraPassword, vars.CassandraKeyspace)
	if err != nil {
		return err
	}
	defer c.Session.Close()

	query := "INSERT INTO restaurants (restaurant_id, borough, cuisine, name) VALUES (?, ?, ?, ?)"
	if err := c.Session.Query(query, restaurant.Id, restaurant.Borough, restaurant.Cuisine, restaurant.Name).Exec(); err != nil {
		return err
	}
	return nil
}

// This function deletes a restaurant from the database.
// The function takes a string parameter, which is the restaurant's ID.
func Delete(id string) error {
	vars := envretriever.GetEnvVars()

	c, err := cassandra.NewCassandra(vars.CassandraHost, vars.CassandraUsername, vars.CassandraPassword, vars.CassandraKeyspace)
	if err != nil {
		return err
	}
	defer c.Session.Close()

	query := "DELETE FROM restaurants WHERE restaurant_id = ?"

	if err := c.Session.Query(query, id).Exec(); err != nil {
		return err
	}
	return nil
}

func Update(restaurant structs.Restaurant, id string) (*structs.Restaurant, error) {
	vars := envretriever.GetEnvVars()

	c, err := cassandra.NewCassandra(vars.CassandraHost, vars.CassandraUsername, vars.CassandraPassword, vars.CassandraKeyspace)
	if err != nil {
		return nil, err
	}
	defer c.Session.Close()

	query := "UPDATE restaurants SET borough = ?, cuisine = ?, name = ? WHERE restaurant_id = ?"

	if err := c.Session.Query(query, restaurant.Borough, restaurant.Cuisine, restaurant.Name, id).Exec(); err != nil {
		return nil, err
	}

	// Get the updated restaurant and return it
	updatedRestaurant, err := GetById(restaurant.Id)
	if err != nil {
		return nil, err
	}

	return updatedRestaurant, nil
}
