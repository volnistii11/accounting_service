package app

import "flag"

type flags struct {
	topic             string
	bootstrapServer   string
	consumerGroupName string

	balanceServiceServer string
}

var cliFlags = flags{}

func init() {
	flag.StringVar(&cliFlags.topic, "topic", "accounting-events", "topic to produce")
	flag.StringVar(&cliFlags.bootstrapServer, "bootstrap-server", "kafka:9092", "kafka broker host and port")
	flag.StringVar(&cliFlags.consumerGroupName, "cg-name", "accounting-events-consumer-group", "topic to produce")

	flag.StringVar(&cliFlags.balanceServiceServer, "balance-service-host", "balance:50051", "balance service host and port")

	flag.Parse()
}
