package queue

import "github.com/nats-io/nats.go"

type Consumer struct {
	con      *nats.Conn
	subject  string
	callback func(*nats.Msg)
}

func NewConsumer(con *nats.Conn, subject string, handler func(*nats.Msg)) error {
	var c = &Consumer{
		con:      con,
		subject:  subject,
		callback: handler,
	}
	_, err := c.con.Subscribe(c.subject, c.callback)
	return err
}
