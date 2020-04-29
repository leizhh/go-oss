package objects

import (
	"io"
	"log"
	"net/http"
	config "YJC-OSS/config"
	"strings"
	"os"
)

func get(w http.ResponseWriter, r *http.Request) {
	f, e := os.Open(config.STORAGE_ROOT + "/objects/" +
		strings.Split(r.URL.EscapedPath(), "/")[2])
	if e != nil {
		log.Println(e)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	defer f.Close()
	io.Copy(w, f)
}
