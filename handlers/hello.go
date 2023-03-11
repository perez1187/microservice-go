package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type Hello struct {
	l *log.Logger //we want to have control od logger message
}

func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

func (h *Hello) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	h.l.Println("hello world") // we replaced log.Println

	d, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(rw, "Ooops", http.StatusBadRequest) //rw, message, status code
		return
	}

	log.Printf("hi hi %s", d)      // this write info to logs
	fmt.Fprintf(rw, "hi hi %s", d) // this send response to user
}
