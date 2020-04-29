package objects

import (
	"io"
	"log"
	"net/http"
	config "YJC-OSS/config"
	"strings"
	"os"
)

func put(w http.ResponseWriter, r *http.Request) {
	f, e := os.Create(config.STORAGE_ROOT + "/objects/" +
		strings.Split(r.URL.EscapedPath(), "/")[2])
	if e != nil {
		log.Println(e)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer f.Close()
	io.Copy(f, r.Body)
}
