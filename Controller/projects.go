package controller

import "github.com/gocql/gocql"

type Project struct {
	ProjectID gocql.UUID
	Name      string
	Type      string
	Users     []gocql.UUID
	IsCurrent bool
}
