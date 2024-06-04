package server

import (
	"github.com/T4jgat/cobalt+/config"
	"github.com/streadway/amqp"
	"log"
)

type Server struct {
	conn *amqp.Connection
	ch   *amqp.Channel
}

func New(cfg config.RabbitMQConfig) (*Server, error) {
	conn, err := amqp.Dial(cfg.URL)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}
	defer ch.Close()

	return &Server{conn: conn, ch: ch}, nil
}

func (s *Server) Close() {
	if err := s.ch.Close(); err != nil {
		log.Println("failed to close channel", err)
	}
	if err := s.conn.Close(); err != nil {
		log.Println("failed to close connection", err)
	}
}
