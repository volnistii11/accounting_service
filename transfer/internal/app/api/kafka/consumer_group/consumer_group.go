package consumer_group

import (
	"context"
	"github.com/volnistii11/accounting_service/transfer/internal/app/api/kafka"
	"log"
	"sync"
	"time"

	"github.com/IBM/sarama"
)

type Config struct {
	GroupName string
	Topics    []string
}

type consumerGroup struct {
	sarama.ConsumerGroup
	handler sarama.ConsumerGroupHandler
	topics  []string
}

func NewConsumerGroup(kafkaConfig kafka.Config, groupConfig Config, consumerGroupHandler sarama.ConsumerGroupHandler, opts ...Option) (*consumerGroup, error) {
	config := sarama.NewConfig()
	config.Version = sarama.MaxVersion
	config.Consumer.Offsets.Initial = sarama.OffsetOldest
	config.Consumer.Group.ResetInvalidOffsets = true
	config.Consumer.Group.Heartbeat.Interval = 3 * time.Second
	config.Consumer.Group.Session.Timeout = 60 * time.Second
	config.Consumer.Group.Rebalance.Timeout = 60 * time.Second
	config.Consumer.Return.Errors = true

	config.Consumer.Offsets.AutoCommit.Enable = true
	config.Consumer.Offsets.AutoCommit.Interval = 5 * time.Second

	// Применяем свои конфигурации
	for _, opt := range opts {
		opt.Apply(config)
	}

	cg, err := sarama.NewConsumerGroup(kafkaConfig.Brokers, groupConfig.GroupName, config)
	if err != nil {
		return nil, err
	}

	return &consumerGroup{
		ConsumerGroup: cg,
		handler:       consumerGroupHandler,
		topics:        groupConfig.Topics,
	}, nil
}

func (c *consumerGroup) Run(ctx context.Context, wg *sync.WaitGroup) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		log.Println("[consumer-group] run")

		for {
			// `Consume` should be called inside an infinite loop, when a
			// server-side rebalance happens, the consumer session will need to be
			// recreated to get the new claims
			if err := c.ConsumerGroup.Consume(ctx, c.topics, c.handler); err != nil {
				log.Printf("Error from consume: %v\n", err)
			}
			// check if context was cancelled, signaling that the consumer should stop
			if ctx.Err() != nil {
				log.Printf("[consumer-group]: ctx closed: %s\n", ctx.Err().Error())
				return
			}
		}
	}()
}
