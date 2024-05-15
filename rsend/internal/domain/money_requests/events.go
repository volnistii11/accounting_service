package money_requests

import (
	"time"

	"github.com/IBM/sarama"
)

type EventType string

const (
	EventMoneyRequestSent EventType = "money-request-sent"
)

type Event struct {
	ID              int64     `json:"id"`
	EventType       EventType `json:"event"`
	RequestID       int64     `json:"request_id"`
	UserID          int64     `json:"user_id"`
	Sum             int64     `json:"sum"`
	OperationMoment time.Time `json:"moment"`
	//IdempotentKey   string    `json:"idempotent_key"`
}

type Msg struct {
	Topic     string `json:"topic"`
	Partition int32  `json:"partition"`
	Offset    int64  `json:"offset"`
	Key       string `json:"key"`
	Payload   string `json:"payload"`
}

func FromKafkaMsg(in *sarama.ConsumerMessage) Msg {
	return Msg{
		Topic:     in.Topic,
		Partition: in.Partition,
		Offset:    in.Offset,
		Key:       string(in.Key),
		Payload:   string(in.Value),
	}
}
