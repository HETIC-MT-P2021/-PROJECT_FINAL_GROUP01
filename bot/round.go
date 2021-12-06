package bot

import (
	"encoding/json"
	"fmt"
	"github.com/dimo/database"
	"net/http"
	"time"
)

type Round struct {
	GameId      int    `json:"game_id"`
	PlayerId    string `json:"player_id"`
	Reason      string `json:"reason"`
	Word        string `json:"word"`
	SubmittedAt int64  `json:"submitted_at"`
}

func NewRound(game Game, player Player, word string) Round {
	return Round{
		GameId:      game.Id,
		PlayerId:    player.DiscordId,
		Reason:      "-",
		Word:        word,
		SubmittedAt: time.Now().Unix(),
	}
}

func NewRoundWithReason(game Game, player Player, word string, reason string) Round {
	return Round{
		GameId:      game.Id,
		PlayerId:    player.DiscordId,
		Reason:      reason,
		Word:        word,
		SubmittedAt: time.Now().Unix(),
	}
}

func insertRound(round *Round) (*Round, error) {
	sql := `
		INSERT INTO round (game_id, player_id, reason, word, submitted_at)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING game_id;
	`

	row := database.Database.QueryRow(sql, round.GameId, round.PlayerId, round.Reason, round.Word, round.SubmittedAt)
	err := row.Scan(&round.GameId)

	if err != nil {
		return round, err
	}

	return round, nil
}

func FetchAllRounds(w http.ResponseWriter, r *http.Request) {
	rows, err := database.Database.Query("SELECT * FROM round")
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var rounds []Round

	for rows.Next() {
		var rd Round

		err := rows.Scan(&rd.GameId, &rd.PlayerId, &rd.Reason, &rd.Word, &rd.SubmittedAt)
		if err != nil {
			panic(err)
		}

		rounds = append(rounds, rd)
	}
	if err = rows.Err(); err != nil {
		panic(err)
	}

	fmt.Println("Endpoint Hit: returnAllRounds")
	json.NewEncoder(w).Encode(rounds)
}
