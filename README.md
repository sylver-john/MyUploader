# MyUploader
a Golang image uploader using Kraken.io 

Fill your Kraken's Api Key and Secret here :
```go
    kr, err := kraken.New("api key", "api secret")
```

First we send the image to Kraken using HTTP and Kraken.io's golang package with :
```go
      params := map[string]interface {} {
        "wait": true,
    }
    imgPath := "../upload.png"
    data, err := kr.Upload(params, imgPath)
```

And then we retrieve the otpimized image using HTTP GET with :
```go
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
```

Finally we send the file to a other web server using HTTP POST with a Pipe :
```go

r, w := io.Pipe()
	go func() {
		defer w.Close()
    	newFile, err := ioutil.ReadFile("./upload.png")
		_, err = w.Write(newFile)
	    if err != nil {
	        log.Fatal(err)
	    }
	}()

	resp, err := http.Post("http://localhost:4321", "application/image", r)
  ```

