package main

import (
  "fmt"
  "io/ioutil"
  "log"
  "net/http"
)

func getPage() {
     resp, err := http.Get("http://www.mmkendall.com/")
     if err !=nil {
     	log.Fatal(err)
     }
     responseText, err := ioutil.ReadAll(resp.Body)
     resp.Body.Close()
     if err !=nil {
     	log.Fatal(err)
     }
     fmt.Printf( "%s\n", responseText )
}

func handler(w http.ResponseWriter, req *http.Request) {
     w.Header().Set("Content-Type", "text/plain")
     w.Write([]byte("mattbot\n"))
}

func main() {
     http.HandleFunc("/", handler)
     log.Fatal(http.ListenAndServe(":5000", nil))
}     
