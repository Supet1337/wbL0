package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog"
	"os"
	"wb-l0/config"
	"wb-l0/internal"
	"wb-l0/pkg"
)

type DataBase struct {
	DB  *sqlx.DB
	Log *zerolog.Logger
}

func CreatePostgres(config *config.Config) (*DataBase, error) {
	database := DataBase{}
	log := zerolog.New(os.Stderr).With().Timestamp().Caller().Logger()
	database.Log = &log
	db, err := pkg.GetConn(config)
	if err != nil {
		log.Fatal().Err(err).Send()
		return nil, err
	}
	database.DB = db

	_, err = database.DB.Exec(`CREATE TABLE IF NOT EXISTS orders
	(
	   order_uid text primary key,
	   model json  not null
	)`)
	if err != nil {
		log.Fatal().Err(err).Send()
		return nil, err
	}

	return &database, nil
}

func (p *DataBase) Add(id string, data []byte) error {
	_, err := p.DB.Exec(`INSERT INTO orders VALUES ($1, $2)`, id, data)
	if err != nil {
		p.Log.Error().Err(err).Send()
		return err
	}
	return nil
}

func (p *DataBase) GetAll() (map[string][]byte, error) {

	data, err := p.DB.Queryx("SELECT * FROM orders")
	if err != nil {
		p.Log.Error().Err(err).Send()
		return nil, err
	}

	result := make(map[string][]byte)

	for data.Next() {
		id := ""
		byt := []byte{}
		err := data.Scan(&id, &byt)
		if err != nil {
			p.Log.Err(err).Timestamp().Send()
			return nil, &internal.Error{
				Message: err.Error(),
				Code:    500,
			}
		}
		result[id] = byt
	}

	return result, nil
}
