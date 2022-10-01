// ./platform/database/open_db_connection.go

package database

import "github.com/snickers31/go-with-fiber/app/queries"

type Queries struct {
	*queries.DBQueries
}

func OpenDBConnection() (*Queries, error) {
	db, err := PostgreSQLConnection()
	if err != nil {
		return nil, err
	}

	return &Queries{DBQueries: &queries.DBQueries{DB: db}}, nil

}
