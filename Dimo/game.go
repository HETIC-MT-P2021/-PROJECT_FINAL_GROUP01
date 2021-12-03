package Dimo

import (
	"container/ring"
	"fmt"
	"log"
	"strings"

	"github.com/scylladb/go-set/strset"
)

type Game struct {
	activePlayer *ring.Ring
	words        *strset.Set
	lastWord     string
	onGameEnd    func(winner string)
	isActive     bool
}

func NewGame(user string, onGameEnd func(winner string)) *Game {
	player := ring.New(1)
	player.Value = user
	game := Game{
		words:        strset.New(),
		lastWord:     "",
		activePlayer: player,
		onGameEnd:    onGameEnd,
		isActive:     false,
	}

	return &game
}

func (g *Game) GetCurrentPlayer() string {
	return g.activePlayer.Value.(string)
}

func (g *Game) IsActive() bool {
	return g.isActive
}

func (g *Game) Start() {
	g.isActive = true
}

func (g *Game) AddPlayer(user string) string {
	if g.activePlayer.Len() < 5 {
		isInGameAlready := false
		g.activePlayer.Do(func(player interface{}) {
			if player == user {
				isInGameAlready = true
			}
		})
		if isInGameAlready {
			return fmt.Sprintln("You are aleady in the game")
		} else {
			newPlayer := ring.New(1)
			newPlayer.Value = user
			g.activePlayer.Link(newPlayer)
		}
	} else {
		return fmt.Sprintln("Game is filled up, sorry :(")
	}
	log.Println(g.activePlayer)
	return fmt.Sprintf("%s joined the game", user)
}

func removePlayer(g *Game, player string) string {
	p := g.activePlayer
	for p.Next().Value != player {
		p = p.Next()
	}
	return p.Unlink(1).Value.(string)
}

func removeCurrentPlayer(g *Game) string {
	g.activePlayer = g.activePlayer.Prev()
	player := g.activePlayer.Unlink(1).Value
	g.activePlayer = g.activePlayer.Next()
	return player.(string)
}

func isValidPlay(g *Game, word string) bool {
	if g.words.Size() == 0 {
		return true
	} else if true /*Is Valid Word*/ {
		if !g.words.Has(word) && word[0] == g.lastWord[len(g.lastWord)-1] {
			return true
		}
	}
	return false
}

func (g *Game) Play(player string, word string) string {
	word = strings.ToLower(word)
	var removedPlayer string
	if player == g.activePlayer.Value {
		if isValidPlay(g, word) {
			g.words.Add(word)
			g.lastWord = word
			g.activePlayer = g.activePlayer.Next()
			return fmt.Sprintf("@%s, it is your turn. Suggest a word", g.activePlayer.Value)
		} else {
			removedPlayer = removeCurrentPlayer(g)
		}
	} else {
		removedPlayer = removePlayer(g, player)
	}

	resp := fmt.Sprintf("@%s has been eliminated.\n", removedPlayer)
	if g.activePlayer.Len() == 1 {
		g.onGameEnd(g.activePlayer.Value.(string))
		resp += fmt.Sprintf("Congrat @%s, You have won!!!", g.activePlayer.Value)
	} else {
		resp += fmt.Sprintf("@%s, it is your turn. Suggest a word", g.activePlayer.Value)
	}
	return resp
}
