package bot

import (
	"github.com/dimo/database"
	"time"
)

type Round struct {
	GameId      int
	PlayerId    string
	Reason      string
	Word        string
	SubmittedAt int64
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
