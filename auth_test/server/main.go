package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	var c Config
	flag.StringVar(&(c.Address), "address", ":8752", "address:port")
	flag.Parse()

	hc := NewHandleCore()

	http.Handle("/", &rootHandler{hc})

	http.Handle("/users/sign_up", &signUpHandler{hc})
	http.Handle("/users/sign_down", &signDownHandler{hc})
	http.Handle("/users/sign_in", &signInHandler{hc})
	http.Handle("/users/sign_out", &signOutHandler{hc})

	http.Handle("/work", &workHandler{hc})

	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	log.Printf("listening address %s", c.Address)

	err := http.ListenAndServe(c.Address, nil)
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

type Config struct {
	Address string
}

type HandleCore struct {
}

func NewHandleCore() *HandleCore {
	return &HandleCore{}
}
