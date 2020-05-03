package temp

import (
	"go-oss/dataServer/locate"
	"os/exec"
	"os"
	"log"
)

func commitTempObject(datFile string, tempinfo *tempInfo) {
	log.Println("rename",datFile, " to ", os.Getenv("STORAGE_ROOT")+"/objects/"+tempinfo.Name)
	//e := os.Rename(datFile, os.Getenv("STORAGE_ROOT")+"/objects/"+tempinfo.Name)
	var cmd *exec.Cmd
	cmd = exec.Command("mv",datFile,os.Getenv("STORAGE_ROOT")+"/objects/"+tempinfo.Name)
	if _, err := cmd.Output();err != nil {
		log.Println("rename:",err)
	}	
	
	locate.Add(tempinfo.Name)
}
