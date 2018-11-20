package main

type TPlayer struct {
  Name string
  Team int
}

func newPlayer(name string, team int) *TPlayer {
  return &TPlayer{
    Name: name,
    Team: team,
  }
}

