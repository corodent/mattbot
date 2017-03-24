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

type GroupMeAttachment struct {
     Type         string
     Url          string
}

type GroupMeBotPost struct {
     Attachments  []GroupMeAttachment
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

// picture of jack in the image server
// {"payload":{"url":"https://i.groupme.com/450x450.jpeg.cd107edf1c27497d8a45a239a4746154","picture_url":"https://i.groupme.com/450x450.jpeg.cd107edf1c27497d8a45a239a4746154"}}

const PICTURE_URL = "https://i.groupme.com/450x450.jpeg.cd107edf1c27497d8a45a239a4746154"

func postMessage( w http.ResponseWriter, post *GroupMeBotPost ) {
//     body := fmt.Sprintf( `{"bot_id":"%s","text":"%s"}`, BOT_ID, "Hello from MattBot" )
     body := fmt.Sprintf( `{"bot_id":"%s","text":"%s","attachments":[{"type":"image","url":"%s"}]}`, BOT_ID, "Jack", PICTURE_URL )

     resp, err := http.Post( "https://api.groupme.com/v3/bots/post", "application/json", bytes.NewBufferString(body) )
     if err !=nil {
     	log.Fatal(err)
     }
     _, err = ioutil.ReadAll(resp.Body)
     resp.Body.Close()
     if err !=nil {
     	log.Fatal(err)
     }
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
//     var m interface{}
     e := json.Unmarshal( b, &m )
     if e!=nil {
     	fmt.Printf( "error: Unable to parse body JSON: %s\n", e );
	fmt.Printf( "%s\n", b )
        http.Error( w, "Unable to parse request", 500 )
	return
     }

     if m.Sender_type==SenderTypeUser {
        postMessage( w, &m )
     }
}

func handler(w http.ResponseWriter, req *http.Request) {
     if req.RequestURI=="/img/jack.jpg" {
     	http.ServeFile( w, req, "img/jack.jpg" )
	return
     }
     
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
