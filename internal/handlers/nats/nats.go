package nats

import (
	"github.com/nats-io/nats.go"
	"github.com/nats-io/stan.go"
	"github.com/rs/zerolog"
	"os"
	"time"
	"wb-l0/internal"
	"wb-l0/internal/middleware"
	"wb-l0/internal/usecase"
)

type Nats struct {
	Conn    stan.Conn
	usecase *usecase.Usecase
	Log     *zerolog.Logger
}

func NewNats(usecase *usecase.Usecase) (*Nats, error) {
	nts := Nats{}
	log := zerolog.New(os.Stderr).With().Timestamp().Caller().Logger()
	nts.Log = &log

	nts.usecase = usecase

	nc, err := nats.Connect("nats://127.0.0.1:4222")
	if err != nil {
		nts.Log.Fatal().Err(err).Msg("Can not connect to NATS")
		return nil, &internal.Error{
			Message: err.Error(),
			Code:    500,
		}
	}
	defer nc.Close()
	time.Sleep(3 * time.Second)

	sc, err := stan.Connect("test-cluster", "client",
		stan.SetConnectionLostHandler(func(_ stan.Conn, reason error) {
			nts.Log.Fatal().Err(reason).Msg("Can not connect to NATS STREAMING")
		}))

	if err != nil {
		nts.Log.Fatal().Err(err).Msg("Can not connect to NATS STREAMING")
		return nil, &internal.Error{
			Message: err.Error(),
			Code:    500,
		}
	}
	nts.Conn = sc

	return &nts, nil
}

func (n *Nats) Subscribe(topic string) error {
	_, err := n.Conn.Subscribe(topic, func(msg *stan.Msg) {
		id, err := middleware.Validate(msg.Data)
		if err != nil {
			n.Log.Error().Err(err).Send()
			return
		}
		err = n.usecase.Add(id, msg.Data)
		if err != nil {
			n.Log.Error().Err(err).Send()
			return
		}
		n.Log.Info().Msg("Message received")
	}, stan.DurableName("supet"))
	if err != nil {
		_ = n.Conn.Close()
		n.Log.Error().Err(err)
		return &internal.Error{
			Message: err.Error(),
			Code:    500,
		}
	}
	return nil
}
