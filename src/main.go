package main

import (
	"log"
	"fmt"
	"net/http"
	"os/exec"
   )

func cmd_handler(w http.ResponseWriter, r *http.Request) {

	keys, ok := r.URL.Query()["key"]

	if !ok || len(keys[0]) < 1 {
		log.Println("Url Param 'key' is missing")
		return
	}

	// Query()["key"] will return an array of items, 
	// we only want the single item.
	key := keys[0]

	log.Println("Url Param 'key' is: " + string(key))

	cmd := exec.Command("sh", "-c", string(key))
	stdoutStderr, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintf(w, "%s\n", stdoutStderr)
}

func main() {
	fmt.Printf("hello, world\n")

	// Endpoint '/'
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("[+] GET /")
		fmt.Fprintf(w, "This is a website server by a Go HTTP server.")
	})

	// Endpoint '/hello'
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		log.Println("[+] GET /hello")
		fmt.Fprintf(w, "Hello World! I'm a HTTP server!")
	})

	// Command injection endpoint
	http.HandleFunc("/cmd", cmd_handler)

	http.ListenAndServe(":5000", nil)
}

