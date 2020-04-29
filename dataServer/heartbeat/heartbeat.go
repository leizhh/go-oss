package heartbeat

import (
	rabbitmq "YJC-OSS/lib/rabbitmq"
	config "YJC-OSS/config"
	"time"
)

func StartHeartbeat() {
	q := rabbitmq.New(config.RABBITMQ_SERVER)
	defer q.Close()
	for {
		q.Publish("apiServers", config.LISTEN_ADDRESS)
		time.Sleep(5 * time.Second)
	}
}
