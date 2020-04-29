package objects

import (
	locate "YJC-OSS/apiServer/locate"
	"fmt"
	"io"
	objectstream "YJC-OSS/lib/objectstream"
)

func getStream(object string) (io.Reader, error) {
	server := locate.Locate(object)
	if server == "" {
		return nil, fmt.Errorf("object %s locate fail", object)
	}
	return objectstream.NewGetStream(server, object)
}
