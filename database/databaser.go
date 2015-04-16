package database

import (
	"database/sql"
)

type DataBaser interface {
	Open(string, string) (*sql.DB, error)
	Create(string) error
	Drop(string) error
	Using(string) error
}
