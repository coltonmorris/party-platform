package main

import (
	"flag"
  "fmt"
	"log"
	"net/http"
  "time"

)

var addr = flag.String("addr", "0.0.0.0:9090", "http service address")

type TRound struct {
  Category string
  Word string
  Timer *time.Timer
  TimerLength int
  Started bool
}

func main() {
	flag.Parse()
  hub := newHub()

  game := newGame(&GameConfig{NumTeams: 2, MaxScore: 7})

  // TODO replace hub functionality into team? Or maybe into game?
  // maybe create a hub for each team?
  go hub.run()
  
	log.SetFlags(0)

  http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
    serveWs(hub, w, r)
    game.addPlayer(newPlayer("colton", 2))

    fmt.Println("Players: ", game.Players[0].Name)
  })

  fmt.Println("listening on: ", *addr)
	log.Fatal(http.ListenAndServe(*addr, nil))
}
