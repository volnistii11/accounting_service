package app

import (
	"github.com/volnistii11/accounting_service/transfer/internal/app/api/kafka"
	"github.com/volnistii11/accounting_service/transfer/internal/app/api/kafka/consumer_group"
)

type config struct {
	kafka         kafka.Config
	consumerGroup consumer_group.Config

	balanceServiceServer string
}

func newConfig(f flags) config {
	return config{
		kafka: kafka.Config{
			Brokers: []string{
				f.bootstrapServer,
			},
		},
		consumerGroup: consumer_group.Config{
			GroupName: f.consumerGroupName,
			Topics:    []string{f.topic},
		},
		balanceServiceServer: f.balanceServiceServer,
	}
}
