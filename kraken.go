package main

import (
  "log"
  "github.com/kraken-io/kraken-go"
  "net/http"
  "os"
  "io"
)
func main() {
	/*================SEND THE IMAGE================*/
    kr, err := kraken.New("api key", "api secret")
    if err != nil {
        log.Fatal(err)
    }
    params := map[string]interface {} {
        "wait": true,
    }
    imgPath := os.Args[1]
    data, err := kr.Upload(params, imgPath)
    if err != nil {
        log.Fatal("err ", err)
    }
    if data["success"] != true {
        log.Println("Failed, error message ", data["message"])
    } else {
        log.Println("Success, Optimized image URL: ", data["kraked_url"])
    }
	/*================AND RETRIEVE THE OPTIMIZED IMAGE================*/
    url := data["kraked_url"].(string)
    response, err := http.Get(url)
    if err != nil {
        log.Fatal(err)
    }
    defer response.Body.Close()
    file, err := os.Create("./upload.png")
    if err != nil {
        log.Fatal(err)
    }
    _, err = io.Copy(file, response.Body)
    if err != nil {
        log.Fatal(err)
    }
    file.Close()
	log.Print("optimized image received")
}