package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/_ah/push-handlers/first-topic", pushHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s\n", port)
	}

	log.Printf("Listening on port %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	fmt.Fprint(w, "Hello, World!!")
}

func pushHandler(w http.ResponseWriter, r *http.Request) {
	type pushRequest struct {
		Message struct {
			Attributes map[string]string
			Data       []byte
			ID         string `json:"message_id"`
		}
		Subscription string
	}

	message := &pushRequest{}
	if err := json.NewDecoder(r.Body).Decode(message); err != nil {
		log.Printf("Could not decode body: %v\n", err)
		http.Error(w, fmt.Sprintf("Could not decode body: %v", err), http.StatusBadRequest)
		return
	}

	log.Printf("Data = %v\n", string(message.Message.Data))
	fmt.Fprint(w, "ok")
}
