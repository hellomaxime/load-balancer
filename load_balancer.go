package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

var i = 0;
var servers = [3]int{8081, 8082, 8083}

func handler(w http.ResponseWriter, r *http.Request) {
	resp, _ := http.Get(fmt.Sprintf("http://localhost:%d", servers[i]))
	i = (i + 1)%len(servers)
	bodyBytes, err := io.ReadAll(resp.Body);
	if(err != nil) {
		log.Fatal(err)
	}
	fmt.Fprint(w, string(bodyBytes))
} 

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}