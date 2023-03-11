package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) { // we add r and we can use it
		log.Println("hello world")
		// we read everything from the body
		// d, _ := io.ReadAll(r.Body)     //d -> data // _ -> err in the future
		d, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(rw, "Ooops", http.StatusBadRequest) //rw, message, status code
			return
		}

		log.Printf("hi hi %s", d)      // this write info to logs
		fmt.Fprintf(rw, "hi hi %s", d) // this send response to user
	})
	http.HandleFunc("/goodbye", func(http.ResponseWriter, *http.Request) {
		log.Println("bye")
	})

	http.ListenAndServe(":9090", nil)
}
