package sqlite

import (
	"database/sql"

	"github.com/aletomasella/ddd-go/domain/customer"
)

type SqliteRepository struct {
	db *sql.DB
}

type sqliteCustomer struct {
	// This annotation is required for the json package to work
	Id       int    `json:"id"`
	Name     string `json:"name"`
	LastName string `json:"last_name"`
}

func NewSqliteRepository(db *sql.DB) *SqliteRepository {
	return &SqliteRepository{
		db: db,
	}
}

func NewSqlLiteConnection(filePath string) (*SqliteRepository, error) {
	res, err := sql.Open("sqlite3", filePath)

	if err != nil {
		return nil, err
	}

	pingErr := res.Ping()
	if pingErr != nil {
		return nil, pingErr
	}

	return NewSqliteRepository(res), nil
}

func NewFromCustomer(c customer.Customer) sqliteCustomer {
	return sqliteCustomer{
		Id:       c.GetId(),
		Name:     c.GetName(),
		LastName: c.GetLastName(),
	}
}

func (sc sqliteCustomer) ToCustomer() (customer.Customer, error) {
	c, err := customer.NewCustomer(sc.Name, sc.LastName)
	if err != nil {
		return customer.Customer{}, err
	}
	return c, nil
}
