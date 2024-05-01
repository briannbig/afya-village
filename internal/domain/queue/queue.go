package queue

import (
	"github.com/briannbig/afya-village/internal/infra/queue"
	"github.com/nats-io/nats.go"
)

type Producer interface {
	Publish(any) error
}

type EventType string

const (
	EventUserCreated          EventType = "user.create"
	EventPatientCreated       EventType = "patient.create"
	EventMedicalRecordCreated EventType = "medical.record.create"
	EventAppointmentScheduled EventType = "appointment.create"
)

type Queue struct {
	nc *nats.Conn
}

func New(nc *nats.Conn) Queue {
	return Queue{nc: nc}
}

func (q Queue) RegisterProducer(eventType EventType) Producer {
	return queue.NewProducer(q.nc, string(eventType))
}

func (q Queue) RegisterConsumer(eventType EventType, handler func(*nats.Msg)) error {
	return queue.NewConsumer(q.nc, string(eventType), handler)
}
