package main

import (
	"flag"
  "fmt"
	"log"
	"net/http"

)

var addr = flag.String("addr", "0.0.0.0:9090", "http service address")

func main() {
	flag.Parse()
  hub := newHub()

  // maybe create a hub for each team?
  go hub.run()
  
	log.SetFlags(0)

  http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
    serveWs(hub, w, r)
  })

  fmt.Println("listening on: ", *addr)
	log.Fatal(http.ListenAndServe(*addr, nil))
}
