package objects

import (
	"go-oss/dataServer/locate"
	"go-oss/lib/utils"
	"log"
	"net/url"
	"os"
)

func getFile(hash string) string {
	file := os.Getenv("STORAGE_ROOT") + "/objects/" + hash
	f, _ := os.Open(file)
	d := url.PathEscape(utils.CalculateHash(f))
	f.Close()
	if d != hash {
		log.Println("object hash mismatch", d)
		log.Println("object hash mismatch, remove", file)
		locate.Del(hash)
		os.Remove(file)
		return ""
	}
	return file
}