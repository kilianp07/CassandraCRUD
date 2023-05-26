package contactRepo

import (
	"github.com/kilianp07/CassandraCRUD/pkg/cassandra"
	envretriever "github.com/kilianp07/CassandraCRUD/utils/envRetriever"
	"github.com/kilianp07/CassandraCRUD/utils/structs"
)

func GetById(id string) (*structs.Contact, error) {
	vars, _ := envretriever.GetEnvVars()
	contact := &structs.Contact{}

	c, err := cassandra.NewCassandra(vars.CassandraHost, vars.CassandraUsername, vars.CassandraPassword, "contacts")
	if err != nil {
		return nil, err
	}
	defer c.Session.Close()

	query := "SELECT id, title, name, address, realaddress, department, country, tel, email FROM my_table WHERE id = ?"
	if err := c.Session.Query(query, id).Scan(&contact.Id, &contact.Title, &contact.Name, &contact.Address, &contact.RealAddress, &contact.Departement, &contact.Country, &contact.Tel, &contact.Email); err != nil {
		return nil, err
	}
	return contact, nil

}

func GetAll() ([]*structs.Contact, error) {
	vars, _ := envretriever.GetEnvVars()
	contacts := []*structs.Contact{}

	c, err := cassandra.NewCassandra(vars.CassandraHost, vars.CassandraUsername, vars.CassandraPassword, "contacts")
	if err != nil {
		return nil, err
	}
	defer c.Session.Close()

	query := "SELECT id, title, name, address, realaddress, department, country, tel, email FROM my_table"
	iter := c.Session.Query(query).Iter()
	for {
		contact := &structs.Contact{}
		if !iter.Scan(&contact.Id, &contact.Title, &contact.Name, &contact.Address, &contact.RealAddress, &contact.Departement, &contact.Country, &contact.Tel, &contact.Email) {
			break
		}
		contacts = append(contacts, contact)
	}

	return contacts, nil

}

func Create(contact *structs.Contact) error {
	vars, _ := envretriever.GetEnvVars()

	c, err := cassandra.NewCassandra(vars.CassandraHost, vars.CassandraUsername, vars.CassandraPassword, "contacts")
	if err != nil {
		return err
	}
	defer c.Session.Close()

	query := "INSERT INTO my_table (id, title, name, address, realaddress, department, country, tel, email) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)"
	if err := c.Session.Query(query, contact.Id, contact.Title, contact.Name, contact.Address, contact.RealAddress, contact.Departement, contact.Country, contact.Tel, contact.Email).Exec(); err != nil {
		return err
	}
	return nil
}

func Update(contact *structs.Contact) error {
	vars, _ := envretriever.GetEnvVars()

	c, err := cassandra.NewCassandra(vars.CassandraHost, vars.CassandraUsername, vars.CassandraPassword, "contacts")
	if err != nil {
		return err
	}
	defer c.Session.Close()

	query := "UPDATE my_table SET title = ?, name = ?, address = ?, realaddress = ?, department = ?, country = ?, tel = ?, email = ? WHERE id = ?"
	if err := c.Session.Query(query, contact.Title, contact.Name, contact.Address, contact.RealAddress, contact.Departement, contact.Country, contact.Tel, contact.Email, contact.Id).Exec(); err != nil {
		return err
	}
	return nil
}

func Delete(id string) error {
	vars, _ := envretriever.GetEnvVars()

	c, err := cassandra.NewCassandra(vars.CassandraHost, vars.CassandraUsername, vars.CassandraPassword, "contacts")
	if err != nil {
		return err
	}
	defer c.Session.Close()

	query := "DELETE FROM my_table WHERE id = ?"
	if err := c.Session.Query(query, id).Exec(); err != nil {
		return err
	}
	return nil
}
