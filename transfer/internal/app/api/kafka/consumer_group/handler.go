package consumer_group

import (
	"encoding/json"
	"log"

	"github.com/IBM/sarama"
	"github.com/volnistii11/accounting_service/transfer/internal/app/models"
)

type ConsumerGroupHandler struct {
	ready chan bool
}

func NewConsumerGroupHandler() *ConsumerGroupHandler {
	return &ConsumerGroupHandler{
		ready: make(chan bool),
	}
}

func (h *ConsumerGroupHandler) Setup(_ sarama.ConsumerGroupSession) error {
	return nil
}

func (h *ConsumerGroupHandler) Cleanup(_ sarama.ConsumerGroupSession) error {
	return nil
}

func (h *ConsumerGroupHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for {
		select {
		case message, ok := <-claim.Messages():
			if !ok {
				return nil
			}

			msg, err := convertMsg(message)
			if err != nil {
				log.Printf("Err: %s", err.Error())
			}
			log.Printf("Message payload: %+v", msg.Payload)

			data, err := json.Marshal(msg)
			if err != nil {
				log.Printf("Err: %s", err.Error())
			}
			log.Printf("Message claimed: %s", data)

			session.MarkMessage(message, "")
		case <-session.Context().Done():
			log.Print("ctx closed")
			return nil
		}
	}
}

type Msg struct {
	Topic     string               `json:"topic"`
	Partition int32                `json:"partition"`
	Offset    int64                `json:"offset"`
	Key       string               `json:"key"`
	Payload   models.MoneyTransfer `json:"payload"`
}

func convertMsg(in *sarama.ConsumerMessage) (*Msg, error) {
	var moneyTransfer models.MoneyTransfer

	err := json.Unmarshal(in.Value, &moneyTransfer)
	if err != nil {
		return nil, err
	}

	return &Msg{
		Topic:     in.Topic,
		Partition: in.Partition,
		Offset:    in.Offset,
		Key:       string(in.Key),
		Payload:   moneyTransfer,
	}, err
}
