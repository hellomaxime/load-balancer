package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

var port string

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w,`<!DOCTYPE html>
		<html lang="en">
			<head>
				<meta charset="utf-8">
				<title>Index Page</title>
			</head>
			<body>
				Hello from the web server running on port %s.
			</body>
		</html>`, port)
} 

func main() {
	if(len(os.Args) != 2) {
		fmt.Println("args error")
		return;
	}
	port = os.Args[1]
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}