package usecase

import (
	"github.com/rs/zerolog"
	"os"
	"wb-l0/config"
	"wb-l0/internal"
	"wb-l0/internal/repository"
)

type Usecase struct {
	Cache      map[string][]byte
	repository *repository.DataBase
	Log        *zerolog.Logger
}

// Конструткор usecase
func NewUsecase(conf *config.Config) (*Usecase, error) {
	usecase := Usecase{}
	var err error
	usecase.Cache = make(map[string][]byte)
	log := zerolog.New(os.Stderr).With().Timestamp().Caller().Logger()
	usecase.Log = &log
	usecase.repository, err = repository.CreatePostgres(conf)
	if err != nil {
		usecase.Log.Error().Err(err).Send()
		return nil, &internal.Error{
			Message: err.Error(),
			Code:    500,
		}
	}
	err = usecase.CacheRecovery()
	if err != nil {
		usecase.Log.Error().Err(err)
		return nil, &internal.Error{
			Message: err.Error(),
			Code:    500,
		}
	}
	return &usecase, nil
}

func (u *Usecase) Add(id string, data []byte) error {
	if _, ok := u.Cache[id]; ok {
		u.Log.Error().Msg("Key already exist")
		return &internal.Error{
			Message: "Key already exist",
			Code:    412,
		}
	}
	u.Cache[id] = data
	err := u.repository.Add(id, data)

	if err != nil {
		u.Log.Error().Err(err).Send()
		return &internal.Error{
			Message: err.Error(),
			Code:    500,
		}
	}
	return nil
}

func (u *Usecase) CacheRecovery() error {
	cache, err := u.repository.GetAll()
	if err != nil {
		u.Log.Error().Err(err).Send()
		return &internal.Error{
			Message: err.Error(),
			Code:    500,
		}
	}
	u.Cache = cache
	return nil
}

func (u *Usecase) GetByIdFromCache(id string) ([]byte, error) {
	if _, ok := u.Cache[id]; !ok {
		u.Log.Error().Msg("Key doesnt exist")
		return nil, &internal.Error{
			Message: "Key doesnt exist",
			Code:    404,
		}
	}

	return u.Cache[id], nil
}
