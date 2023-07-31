package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
	"os"
	"wb-l0/internal"
	"wb-l0/internal/middleware"
	"wb-l0/internal/usecase"
)

type FiberServer struct {
	server  *fiber.App
	usecase *usecase.Usecase
	Log     *zerolog.Logger
}

func CreateServer(ucase *usecase.Usecase) (*FiberServer, error) {
	server := FiberServer{}
	server.server = fiber.New()
	log := zerolog.New(os.Stderr).With().Timestamp().Caller().Logger()
	server.Log = &log
	server.usecase = ucase
	return &server, nil
}

func (s *FiberServer) StartServer(port string) error {
	s.server.Get("get/:id", s.Get)
	s.server.Post("add", s.Add)
	err := s.server.Listen(port)
	if err != nil {
		s.Log.Error().Err(err).Send()
		return err
	}
	return nil
}

func (s *FiberServer) Add(ctx *fiber.Ctx) error {
	body := ctx.Body()
	id, err := middleware.Validate(body)
	if err != nil {
		s.Log.Error().Err(err).Send()
		return &internal.Error{
			Message: "Validation Error",
			Code:    500,
		}
	}
	err = s.usecase.Add(id, body)
	if err != nil {
		s.Log.Error().Err(err).Send()
		return &internal.Error{
			Message: err.Error(),
			Code:    500,
		}
	}
	return nil

}

func (s *FiberServer) Get(ctx *fiber.Ctx) error {
	id := ctx.Params("id", "")
	if id == "" {
		s.Log.Error().Msg("Empty id")
		return &internal.Error{
			Message: "Empty id",
			Code:    404,
		}
	}
	data, err := s.usecase.GetByIdFromCache(id)
	if err != nil {
		s.Log.Error().Err(err).Send()
		return &internal.Error{
			Message: err.Error(),
			Code:    500,
		}
	}

	return ctx.Send(data)
}
