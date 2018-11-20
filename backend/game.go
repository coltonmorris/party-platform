package main

type TGame struct {
  // each hub is a team
  Teams []*Hub
  
  Players []*TPlayer
  CurrentRoundNumber int
  Round *TRound
  MaxScore int
  // TODO number of lives? maybe put it in the team type
}

type GameConfig struct {
  NumTeams int
  MaxScore int
}


func newGame(config *GameConfig) *TGame {
  var newTeams []*Hub
  for i := 0; i < config.NumTeams; i++ {
    newTeams = append(newTeams, newHub())
  }


  return &TGame{
    Teams: newTeams,
    Players: nil,
    CurrentRoundNumber: 0,
    Round: nil,
    MaxScore: config.MaxScore,
  }
}

func (g *TGame) addPlayer(player *TPlayer) {
  g.Players = append(g.Players, player)
}
