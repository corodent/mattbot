package main

import (
  "bytes"
  "fmt"
  "io/ioutil"
  "log"
  "net/http"
  "encoding/json"
)

const SenderTypeBot = "bot"
const SenderTypeUser = "user"

type GroupMeBotPost struct {
     Attachments  []string
     Avatar_url   string
     Created_at   int
     Group_id     string
     Id           string
     Name         string
     Sender_id    string
     Sender_type  string
     User         string
     Source_guid  string
     System       bool
     Text         string
     User_id      string
}

const BOT_ID = "53e068b257ee15680e15338acf"

func handleGet( w http.ResponseWriter, req *http.Request ) {
     fmt.Printf( "handler: %s %s\n", req.Method, req.Proto )
     w.Header().Set("Content-Type", "text/plain")
     responseString := "mattbot: " + req.Method
     w.Write([]byte(responseString))
}

func postMessage( w http.ResponseWriter, post *GroupMeBotPost ) {
     body := fmt.Sprintf( `{"bot_id":"%s","text":"%s"}`, BOT_ID, "Hello from MattBot" )

     resp, err := http.Post( "https://api.groupme.com/v3/bots/post", "application/json", bytes.NewBufferString(body) )
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

func handlePost( w http.ResponseWriter, req *http.Request ) {
     fmt.Printf( "handlePost\n" );

     b, err := ioutil.ReadAll( req.Body )
     if err!=nil {
     	fmt.Printf( "error: %s\n", err )
        http.Error( w, "Unable to read request", 500 )
	return
     }

     var m GroupMeBotPost
     e := json.Unmarshal( b, &m )
     if e!=nil {
     	fmt.Printf( "error: Unable to parse body JSON: %s\n", e );
        http.Error( w, "Unable to parse request", 500 )
	return
     }

     if m.Sender_type==SenderTypeUser {
        postMessage( w, &m )
     }
}

func handler(w http.ResponseWriter, req *http.Request) {
     if req.Method == http.MethodGet {
     	handleGet( w, req );
     } else if req.Method == http.MethodPost {
       handlePost( w, req );
     } else {
        log.Printf( "ERROR: Invalid method %s\n", req.Method );
        http.Error( w, "Method not supported", 405 )
     }
}

func main() {
     fmt.Printf( "Starting mattbot\n" )
     http.HandleFunc("/", handler)
     log.Fatal(http.ListenAndServe(":5000", nil))
}     
