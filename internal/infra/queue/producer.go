package queue

import (
	"github.com/briannbig/afya-village/internal/util"
	"github.com/nats-io/nats.go"
)

type Producer struct {
	con     *nats.Conn
	subject string
}

func NewProducer(con *nats.Conn, subject string) *Producer {
	return &Producer{
		con:     con,
		subject: subject,
	}
}

func (p *Producer) Publish(payload any) error {
	data, err := util.GetBytes(payload)
	if err != nil {
		return err
	}

	if err := p.con.Publish(p.subject, data); err != nil {
		return err
	}

	return nil
}
