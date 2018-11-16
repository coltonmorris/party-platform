## Ideas for Backend

golang wrapped grpc-web

database



### Datastructures:

player: {
  name: string
  team: id
}

team: {
  id: int
  score: int
}

game_session: {
  players: []player
  teams: []team
  gamePaused: boolean // the state between rounds
  currentRoundNumber: int
  round: Round
}

round: {
  category: string
  word: string
  startTime: timestamp
  timerLength: int
}
  

lobby: {
  players: []player
  host: int // the players index
}
