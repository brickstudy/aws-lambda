package mysql

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/robert-min/aws-lambda/data-catalog/adapter/config"
)

type DB struct {
	*sql.DB
}

// Create New MySQL Client DB instance.
func New(config *config.DB) (*DB, error) {
	url := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		config.Username,
		config.Password,
		config.Hostname,
		config.Port,
		config.Name,
	)
	db, err := sql.Open("mysql", url)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return &DB{
		db,
	}, nil
}
