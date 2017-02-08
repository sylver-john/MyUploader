package main

import (
  "log"
  "github.com/kraken-io/kraken-go"
  "net/http"
  "os"
  "fmt"
  "io"
)
func main() {
    kr, err := kraken.New("key", "secret")
    if err != nil {
        log.Fatal(err)
    }
    params := map[string]interface {} {
        "wait": true,
    }
    imgPath := "../upload.png"
    data, err := kr.Upload(params, imgPath)
    if err != nil {
        log.Fatal("err ", err)
    }
    if data["success"] != true {
        log.Println("Failed, error message ", data["message"])
    } else {
        log.Println("Success, Optimized image URL: ", data["kraked_url"])
    }

    url := data["kraked_url"].(string)
    response, e := http.Get(url)
    if e != nil {
        log.Fatal(e)
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
    fmt.Println("Success!")
}