package consumer_group

import (
	"context"
	"encoding/json"
	"log"

	"github.com/IBM/sarama"
	"github.com/volnistii11/accounting_service/transfer/internal/app/models"
)

type UseCaseProvider interface {
	CreateMoneyTransfer(ctx context.Context, info models.MoneyTransfer) error
}

type ConsumerGroupHandler struct {
	useCase UseCaseProvider
	ready   chan bool
}

func NewConsumerGroupHandler(useCase UseCaseProvider) *ConsumerGroupHandler {
	return &ConsumerGroupHandler{
		useCase: useCase,
		ready:   make(chan bool),
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
				log.Printf("Error converting message: %s", err.Error())
				continue
			}

			log.Printf("Message payload: %+v", msg.Payload)

			if err = h.useCase.CreateMoneyTransfer(session.Context(), msg.Payload); err != nil {
				log.Printf("Error creating money transfer: %s", err.Error())
				continue
			}

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
