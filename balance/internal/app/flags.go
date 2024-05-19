package app

import "flag"

type flags struct {
	serverPort  string
	redisServer string
	dbServer    string
}

var cliFlags = flags{}

func init() {
	flag.StringVar(&cliFlags.serverPort, "balance-service-port", "50051", "balance service port")
	flag.StringVar(&cliFlags.redisServer, "redis-host", "balance:6379", "redis host and port")
	flag.StringVar(&cliFlags.redisServer, "db-host", "db:5432", "db host and port")

	flag.Parse()
}
