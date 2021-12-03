package bot

import (
	"fmt"
	"github.com/dimo/database"
	"strings"
	"time"

	"github.com/scylladb/go-set/strset"
)

const maxPlayersInParty int = 5

type Game struct {
	Id                 int
	Players            []Player
	CurrentPlayerIndex int
	Words              *strset.Set
	LastWord           string
	OnGameEnd          func(winner Player)
	IsActive           bool
	StartedAt          int64
	FinishedAt         int64
	FirstMessageId     string
	StartedBy          Player
}

func insertGame(game *Game) (*Game, error) {
	sql := `
		INSERT INTO game (started_at, finished_at)
		VALUES ($1, $2)
		RETURNING game_id;
	`

	row := database.Database.QueryRow(sql, game.StartedAt, game.FinishedAt)
	err := row.Scan(&game.Id)

	if err != nil {
		return game, err
	}

	return game, nil
}

func updateGame(game *Game) {
	sql := `
		UPDATE game
		SET
		    started_at = $1,
		    finished_at = $2,
		WHERE game_id = $3;
	`

	database.Database.Exec(sql, game.StartedAt, game.FinishedAt, game.Id)
}

func NewGame(gameAdmin Player, firstMessageId string, onGameEnd func(winner Player)) *Game {
	game := Game{
		Players:            make([]Player, 0),
		CurrentPlayerIndex: 0,
		Words:              strset.New(),
		LastWord:           "",
		OnGameEnd:          onGameEnd,
		IsActive:           false,
		StartedAt:          0,
		FinishedAt:         0,
		FirstMessageId:     firstMessageId,
		StartedBy:          gameAdmin,
	}

	game.AddPlayer(&gameAdmin)

	_, err := insertGame(&game)
	if err != nil {
		fmt.Println("An error occurred when persisting game data", err)
	}

	return &game
}

func (g *Game) Start() {
	g.IsActive = true
	g.StartedAt = time.Now().Unix()

	updateGame(g)
}

func (g *Game) Finish() {
	g.IsActive = false
	g.FinishedAt = time.Now().Unix()

	updateGame(g)
}

func (g *Game) GetPlayersNames() string {
	var allPlayersNames = ""

	for index, player := range g.Players {
		allPlayersNames += fmt.Sprintf("%v- %v\n", index+1, player.Name)
	}

	return allPlayersNames
}

func (g *Game) GetCurrentPlayer() Player {
	return g.Players[g.CurrentPlayerIndex]
}

func (g *Game) GetPlayerByDiscordId(playerDiscordId string) *Player {
	for _, player := range g.Players {
		if player.DiscordId == playerDiscordId {
			return &player
		}
	}
	return nil
}

func (g *Game) GetNextPlayerIndex() int {
	nextPlayerIndex := g.CurrentPlayerIndex + 1

	if nextPlayerIndex >= len(g.Players) {
		nextPlayerIndex = 0
	}

	return nextPlayerIndex
}

func (g *Game) AddPlayer(candidatePlayer *Player) (bool, string) {
	if len(g.Players) < maxPlayersInParty {
		isInGameAlready := false

		for _, player := range g.Players {
			if player.DiscordId == candidatePlayer.DiscordId {
				isInGameAlready = true
				break
			}
		}

		if isInGameAlready {
			return false, fmt.Sprintln("You are already in the game")
		} else {
			g.Players = append(g.Players, *candidatePlayer)

			_, err := insertPlayer(candidatePlayer)
			if err != nil {
				fmt.Println("An error occurred when persisting player data", err)
			}
		}
	} else {
		return false, fmt.Sprintln("Game is filled up, sorry :(")
	}

	return true, candidatePlayer.Name
}

func (g *Game) removePlayer(playerToRemove Player) Player {
	index := -1
	for i, player := range g.Players {
		if player.DiscordId == playerToRemove.DiscordId {
			index = i
		}
	}

	g.Players[index] = g.Players[len(g.Players)-1]
	g.Players = g.Players[:len(g.Players)-1]

	return playerToRemove
}

func (g *Game) removeCurrentPlayer() Player {
	playerToRemove := g.Players[g.CurrentPlayerIndex]
	return g.removePlayer(playerToRemove)
}

func isValidPlay(g *Game, word string) bool {
	if g.Words.Size() == 0 {
		return true
	} else if true /*Is Valid Word*/ {
		if !g.Words.Has(word) && word[0] == g.LastWord[len(g.LastWord)-1] {
			return true
		}
	}
	return false
}

func (g *Game) Play(player Player, word string) string {
	word = strings.ToLower(word)
	var removedPlayer Player

	if player == g.GetCurrentPlayer() {
		if isValidPlay(g, word) {
			g.Words.Add(word)
			g.LastWord = word
			g.CurrentPlayerIndex = g.GetNextPlayerIndex()

			round := NewRound(*g, player, word)
			fmt.Println(round)
			insertRound(&round)

			return fmt.Sprintf("@%s, it is your turn. Suggest a word", g.GetCurrentPlayer().Name)
		} else {
			removedPlayer = g.removeCurrentPlayer()
			failed(g, player, word)
		}
	} else {
		removedPlayer = g.removePlayer(player)
		failed(g, player, word)
	}

	resp := fmt.Sprintf("@%s has been eliminated.\n", removedPlayer.Name)

	if len(g.Players) == 1 {
		g.OnGameEnd(g.GetCurrentPlayer())
		resp += fmt.Sprintf("Congrat @%s, You have won!!!", g.GetCurrentPlayer().Name)

		round := NewRoundWithReason(*g, player, word, resp)
		insertRound(&round)
		g.Finish()
	} else {
		resp += fmt.Sprintf("@%s, it is your turn. Suggest a word", g.GetCurrentPlayer().Name)
	}

	return resp
}

func failed(g *Game, player Player, word string) {
	round := NewRoundWithReason(*g, player, word, fmt.Sprintf("@%s has been eliminated.", player.Name))
	insertRound(&round)
}
