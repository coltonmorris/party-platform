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
  name: string
  score: int
}

game_session: {
  players: []player
  teams: []team
  currentRoundNumber: int
  round: Round
}

round: {
  category: string
  word: string
  startTime: timestamp
  timerLength: int
  started: bool
}
  

lobby: {
  players: []player
  host: int // the players index
}


// not sure if it's better to split words into categories, or bundle all words together
categories: []category
words: []word

category: {
  name: string
  words: []string
  id: int
}

word: {
  name: string
  category: string
  id: int
}
  
