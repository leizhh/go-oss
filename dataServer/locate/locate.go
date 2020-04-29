package locate

import (
	rabbitmq "YJC-OSS/lib/rabbitmq"
	config "YJC-OSS/config"
	"strconv"
	"os"
)

func Locate(name string) bool {
	_, err := os.Stat(name)
	return !os.IsNotExist(err)
}

func StartLocate() {
	q := rabbitmq.New(config.RABBITMQ_SERVER)
	defer q.Close()
	q.Bind("dataServers")
	c := q.Consume()
	for msg := range c {
		object, e := strconv.Unquote(string(msg.Body))
		if e != nil {
			panic(e)
		}
		if Locate(config.STORAGE_ROOT + "/objects/" + object) {
			q.Send(msg.ReplyTo, config.LISTEN_ADDRESS)
		}
	}
}
