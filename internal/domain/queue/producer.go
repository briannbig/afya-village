package queue

type Producer interface {
	Produce(any) error
}
