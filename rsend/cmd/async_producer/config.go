package main

import (
	"github.com/volnistii11/accounting_service/rsend/internal/infra/kafka"
)

type appConfig struct {
	repeatCnt int
	startID   int
	count     int
}

type producerConfig struct {
	topic string
}

type config struct {
	app      appConfig
	kafka    kafka.Config
	producer producerConfig
}

func newConfig(f flags) config {
	return config{
		app: appConfig{
			repeatCnt: f.repeatCnt,
			startID:   f.startID,
			count:     f.count,
		},
		kafka: kafka.Config{
			Brokers: []string{
				f.bootstrapServer,
			},
		},
		producer: producerConfig{
			topic: f.topic,
		},
	}
}
