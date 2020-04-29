package objects

import (
	heartbeat "YJC-OSS/apiServer/heartbeat"
	"fmt"
	objectstream "YJC-OSS/lib/objectstream"
)

func putStream(object string) (*objectstream.PutStream, error) {
	server := heartbeat.ChooseRandomDataServer()
	if server == "" {
		return nil, fmt.Errorf("cannot find any dataServer")
	}

	return objectstream.NewPutStream(server, object), nil
}
