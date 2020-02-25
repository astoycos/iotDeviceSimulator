package main


import (
	
	"time"
	"os"
	"os/exec"
	"log"
	"strings"
)

func main() { 

	urlpage := os.Getenv("STREAMURL")
	url,err := exec.Command("youtube-dl","-g",urlpage).Output()
	if err != nil {
        log.Fatal(err)
    }	
	

	for{
	
	getFrame(strings.TrimSpace(string(url)))

	pushFrame("iot-http-adapter-enmasse-infra.apps.astoycos-ocp.shiftstack.com", "sensor1@myapp.iot", "hono-secret")

	
	<-time.After(20* time.Second)
	}
}	