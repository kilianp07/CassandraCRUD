package cassandra

import (
	"fmt"

	"github.com/gocql/gocql"
)

type Cassandra struct {
	cluster *gocql.ClusterConfig
	Session *gocql.Session

	host     string
	username string
	password string
	keyspace string
}

// NewCassandra creates a new Cassandra object with the given host, username, password and keyspace.
func NewCassandra(host string, username string, password string, keyspace string) (*Cassandra, error) {
	c := &Cassandra{
		host:     host,
		username: username,
		password: password,
		keyspace: keyspace,
	}

	if err := c.initialize(); err != nil {
		return nil, err
	}

	return c, nil

}

func (c *Cassandra) initialize() error {
	var (
		err error
	)

	c.cluster = gocql.NewCluster(c.host)
	c.cluster.Authenticator = gocql.PasswordAuthenticator{
		Username: c.username,
		Password: c.password,
	}

	c.cluster.Keyspace = c.keyspace

	c.Session, err = c.cluster.CreateSession()
	if err != nil {
		fmt.Printf("Error creating session: %v \n", err)
		return err
	}
	fmt.Println("cassandra well initialized")
	return nil
}

func (c *Cassandra) Close() {
	c.Session.Close()
}
