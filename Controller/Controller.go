package controller

import (
	"fmt"

	"github.com/gocql/gocql"
)

type DBconnection struct {
	cluster *gocql.ClusterConfig
	session *gocql.Session
}

var connection DBconnection

func SetupDB() *gocql.Session {
	connection.cluster = gocql.NewCluster("127.0.0.1:9042")
	// connection.cluster.Port = 9042
	// connection.cluster.Consistency = gocql.Quorum
	connection.cluster.Keyspace = "first_app"
	// connection.cluster.ConnectTimeout = time.Second * 30 // Example: 30 seconds timeout
	// connection.cluster.Timeout = time.Second * 60        // Example: 60 seconds timeout

	var err error
	connection.session, err = connection.cluster.CreateSession()
	if err != nil {
		fmt.Println("db connection failed", err)
		return nil
	}

	fmt.Println("db connection successful")
	return connection.session
}
