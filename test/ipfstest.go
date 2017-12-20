package test

import (
  "TVStorageManager/json"
  "fmt"
  "os"
  "bytes"
  "io"
  "mime/multipart"
  "net/http"
)

func IpfsTest() {

	response, err := Upload("Ti_Value.exe")
    if err != nil {
    	panic(err)
    }
	const jsonStream = `{"jsonrpc":"2.0","id":1,"method":"login","params":["username","password"]}`
	result, _ := json.JsonParser([]byte(response))
	if err != nil {
    	panic(err)
    }
	if result != nil {
		name, _ := result.Get("Name").String()
		hash, _ := result.Get("Hash").String()
		size, _ := result.Get("Size").String()
		fmt.Println("Name, Hash, Size: ", name, hash, size)
		Download(hash, name+".exe")
	}
}



//客户端上传文件代码：
 func Upload(filepath string) (response string, err error) {
    
    response = ""
    // Create buffer
    buf := new(bytes.Buffer) // caveat IMO dont use this for large files, \
    // create a tmpfile and assemble your multipart from there (not tested)
    w := multipart.NewWriter(buf)
    // Create file field
    fw, err := w.CreateFormFile("file", filepath) //这里的file很重要，必须和服务器端的FormFile一致
    if err !=nil {
        panic(err)
    }
    fd, err := os.Open(filepath)
    if err !=nil {
        panic(err)
    }
    defer fd.Close()
    // Write file field from file to upload
    _, err = io.Copy(fw, fd)
    if err !=nil {
        panic(err)
    }
    // Important if you do not close the multipart writer you will not have a
    // terminating boundry
    w.Close()
    req, err := http.NewRequest("POST","http://localhost:5001/api/v0/add", buf)
    if err !=nil {
        panic(err)
    }
    req.Header.Set("Content-Type", w.FormDataContentType())
      var client http.Client
    res, err := client.Do(req)
    if err !=nil {
        panic(err)
    }

    resbuf := new(bytes.Buffer) 
	resbuf.ReadFrom(res.Body)
	response = resbuf.String()
    io.Copy(os.Stdout, res.Body) // Replace this with Status.Code check
    fmt.Println("h")
    return
}

//客户端下载文件代码：
 func Download(hash string, filepath string) (err error) {
    // Create buffer
    buf := new(bytes.Buffer) // caveat IMO dont use this for large files, \
    // create a tmpfile and assemble your multipart from there (not tested)
    r := multipart.NewWriter(buf)


    defer r.Close()
    req, err := http.NewRequest("POST","http://localhost:5001/api/v0/cat?arg=" + hash, buf)
    if err !=nil {
        panic(err)
    }

    var client http.Client
    res, err := client.Do(req)
    if err !=nil {
        panic(err)
    }

    file, err := os.Create(filepath)
    if err !=nil {
        panic(err)
    }
    defer file.Close()
    io.Copy(file, res.Body) // Replace this with Status.Code check
    return err
}
