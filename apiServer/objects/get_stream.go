package objects

import (
	locate "go-oss/apiServer/locate"
	"fmt"
	"io"
	objectstream "go-oss/lib/objectstream"
)

func getStream(object string) (io.Reader, error) {
	server := locate.Locate(object)
	if server == "" {
		return nil, fmt.Errorf("object %s locate fail", object)
	}
	return objectstream.NewGetStream(server, object)
}
