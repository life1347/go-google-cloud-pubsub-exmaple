package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"tachingchen.com/googlePubSub/common"
)

// https://github.com/GoogleCloudPlatform/golang-samples/blob/master/appengine_flexible/pubsub/pubsub.go
type pushRequest struct {
	Message struct {
		Attributes map[string]string
		Data       []byte
		ID         string `json:"message_id"`
	}
	Subscription string
}

func handler(w http.ResponseWriter, r *http.Request) {
	msg := &pushRequest{}
	if err := json.NewDecoder(r.Body).Decode(msg); err != nil {
		http.Error(w, fmt.Sprintf("Could not decode body: %v", err), http.StatusBadRequest)
		return
	}
	session := &common.Session{}
	if e := json.Unmarshal(msg.Message.Data, session); e != nil {
		log.Printf("Decode Error: %v. Wrong payload format: %v", e, session)
	} else {
		log.Printf("sessionid: %s timestamp: %d\n", session.SessionID, session.TimeStamp)
	}
}

func main() {
	http.HandleFunc("/", handler)
	// You cann apply for free ssl cert at https://www.sslforfree.com/
	err := http.ListenAndServeTLS("0.0.0.0:443", "/certificate.crt", "/private.key", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
