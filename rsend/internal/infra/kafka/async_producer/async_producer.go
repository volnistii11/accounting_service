package async_producer

import (
	"fmt"

	"github.com/IBM/sarama"
	"github.com/volnistii11/accounting_service/rsend/internal/infra/kafka"
)

func NewAsyncProducer(conf kafka.Config, opts ...Option) (sarama.AsyncProducer, error) {
	config := PrepareConfig(opts...)

	asyncProducer, err := sarama.NewAsyncProducer(conf.Brokers, config)
	if err != nil {
		return nil, fmt.Errorf("NewAsyncProducer failed: %w", err)
	}

	return asyncProducer, nil
}
