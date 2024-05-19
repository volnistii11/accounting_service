package app

type config struct {
	serverPort  string
	redisServer string
	dbServer    string
}

func newConfig(f flags) config {
	return config{
		serverPort:  f.serverPort,
		redisServer: f.redisServer,
		dbServer:    f.dbServer,
	}
}
