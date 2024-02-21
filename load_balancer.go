package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

type RoundRobin struct {
	servers [3]string
	health [3]bool
	i int
}

var rr = RoundRobin{ 
	servers: [3]string{"http://localhost:8081", "http://localhost:8082", "http://localhost:8083"}, 
	health: [3]bool{true, true, true},
	i: 0,
}

func next() {
	for {
		rr.i = (rr.i + 1)%len(rr.servers)
		if(rr.health[rr.i]) {
			break
		}
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	next()
	resp, errServer := http.Get(rr.servers[rr.i])
	if(errServer != nil) {
		rr.health[rr.i] = false
		w.WriteHeader(503)
		fmt.Fprint(w, "Service Unavailable")
		return
	}
	bodyBytes, err := io.ReadAll(resp.Body)
	if(err != nil) {
		log.Fatal(err)
	}
	fmt.Fprint(w, string(bodyBytes))
} 

func main() {
	go func() {
		for range time.Tick(10 * time.Second) {
			for i, server := range rr.servers {
				_, err := http.Get(server)
				if(err == nil) {
					fmt.Printf("%s is OK\n", server)
					rr.health[i] = true
				} else {
					fmt.Printf("%s is DOWN\n", server)
					rr.health[i] = false
				}
			}
		}
	}()
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}