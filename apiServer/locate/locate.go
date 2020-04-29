package locate

import (
	rabbitmq "YJC-OSS/lib/rabbitmq"
	config "YJC-OSS/config"
	"strconv"
	"time"
)

func Locate(name string) string {
	q := rabbitmq.New(config.RABBITMQ_SERVER)
	q.Publish("dataServers", name)
	c := q.Consume()
	go func() {
		time.Sleep(time.Second)
		q.Close()
	}()
	msg := <-c
	s, _ := strconv.Unquote(string(msg.Body))
	return s
}

func Exist(name string) bool {
	return Locate(name) != ""
}
