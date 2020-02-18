package main 

import(
	"net/http"
	"io"
	"log"
	"crypto/tls"
	"os"
)

func getFrame(url string){

	resp, err := http.Get(url)
	if err != nil {
        log.Fatalln(err)
	}
	log.Println("Got File")
	defer resp.Body.Close()
	
	// Create the file
    out, err := os.Create("index.m3u8")
    if err != nil {
        log.Fatalln(err)
	}
	
	_, err = io.Copy(out, resp.Body)
    if err != nil {
        log.Fatalln(err)
    }
	log.Println("Wrote file to index.m3u8")
}

func pushFrame(url string, user string, pass string){
	log.Println("Opening file index.m3u8")
	data2, err := os.Open("index.m3u8")
    if err != nil {
        log.Fatal(err)
    }
    defer data2.Close()
	
	// TODO: This is insecure; use only in dev environments.
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	
	req, err := http.NewRequest("POST", "https://" + url + "/telemetry", data2)
	if err != nil {
		log.Fatalln(err)
	}
	req.SetBasicAuth(user, pass)
	req.Header.Set("Content-Type", "application/vnd.apple.mpegurl")
	
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Sent file to http adapter")
	log.Println(resp)

	defer resp.Body.Close()
}