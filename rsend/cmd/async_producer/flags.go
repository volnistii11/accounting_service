package main

import "flag"

type flags struct {
	repeatCnt       int
	startID         int
	count           int
	bootstrapServer string
	topic           string
}

var cliFlags = flags{}

func init() {
	flag.IntVar(&cliFlags.repeatCnt, "repeat-count", 3, "count times all messages sent")
	flag.IntVar(&cliFlags.startID, "start-id", 1, "start order-id of all messages")
	flag.IntVar(&cliFlags.count, "count", 10, "count of orders to emit events")
	flag.StringVar(&cliFlags.bootstrapServer, "bootstrap-server", "kafka0:9092", "kafka broker host and port")
	flag.StringVar(&cliFlags.topic, "topic", "accounting-events", "topic to produce")

	flag.Parse()
}
