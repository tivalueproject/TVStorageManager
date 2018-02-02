package logic

import (
    "os"
    "bytes"
    "io"
    "mime/multipart"
    "net/http"
    //"strconv"
  )

//upload file
func Upload(filename string) (response string, stat_code int, err error) {
	buf := &bytes.Buffer{}
    w := multipart.NewWriter(buf)

    fileWriter, err := w.CreateFormFile("uploadfile", filename)
    if err != nil {
        panic(err)
    }

    fh, err := os.Open(filename)
    if err != nil {
        panic(err)
    }
    defer fh.Close()

    //iocopy
    _, err = io.Copy(fileWriter, fh)
    if err != nil {
        panic(err)
    }
    w.Close()

    request, err := http.NewRequest("POST","http://localhost:5001/api/v0/add", buf)
    if err !=nil {
        panic(err)
    }

    request.Header.Set("Content-Type", w.FormDataContentType())
    var client http.Client
    res, err := client.Do(request)
    if err != nil {
        panic(err)
    }

    resbuf := new(bytes.Buffer) 
	resbuf.ReadFrom(res.Body)
    response = resbuf.String()

    stat_code = res.StatusCode
    return
}

//pin file 
func PinAdd(hash string)(response string, stat_code int, err error) {
    buf := new(bytes.Buffer) 
    r := multipart.NewWriter(buf)
    defer r.Close()

    request, err := http.NewRequest("POST", "http://localhost:5001/api/v0/pin/add?arg=" + hash, buf)
    if err != nil {
        panic(err)
    }
    var client http.Client
    res, err := client.Do(request)
    if err != nil {
        panic(err)
    }
    resbuf := new(bytes.Buffer) 
	resbuf.ReadFrom(res.Body)
    response = resbuf.String()
    stat_code = res.StatusCode
    return
}


//download file
func Download(hash string) (response string, stat_code int, err error) {
    buf := new(bytes.Buffer) 
    r := multipart.NewWriter(buf)
    defer r.Close()

    request, err := http.NewRequest("POST", "http://localhost:5001/api/v0/cat?arg=" + hash, buf)
    if err !=nil {
        panic(err)
    }

    var client http.Client
    res, err := client.Do(request)
    if err != nil {
        panic(err)
    }
    resbuf := new(bytes.Buffer) 
    resbuf.ReadFrom(res.Body)
    response = resbuf.String()
    stat_code = res.StatusCode
    return
}




