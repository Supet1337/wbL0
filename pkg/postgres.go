package pkg

import (
	"fmt"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	"wb-l0/config"
	"wb-l0/internal"
)

func GetConn(c *config.Config) (*sqlx.DB, error) {
	connectionUrl := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		c.Postgres.Host,
		c.Postgres.Port,
		c.Postgres.User,
		c.Postgres.Password,
		c.Postgres.DbName)
	fmt.Println(connectionUrl)
	database, err := sqlx.Connect("pgx", connectionUrl)
	if err != nil {
		return nil, &internal.Error{
			Message: err.Error(),
			Code:    502,
		}
	}
	return database, nil
}
