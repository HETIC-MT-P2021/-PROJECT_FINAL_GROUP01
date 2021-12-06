package bot

import (
	"encoding/json"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/dimo/database"
	"net/http"
)

type Player struct {
	Id        int    `json:"player_id"`
	Name      string `json:"name"`
	DiscordId string `json:"discord_id"`
	Avatar    string `json:"avatar"`
}

func NewPlayerFromDiscordAuthor(author *discordgo.User) Player {
	return Player{
		Id:        0,
		Name:      author.Username,
		DiscordId: author.ID,
		Avatar:    author.Avatar,
	}
}

func insertPlayer(player *Player) (*Player, error) {
	sql := `
		INSERT INTO player (name, avatar, discord_id)
		VALUES ($1, $2, $3)
		RETURNING player_id;
	`

	row := database.Database.QueryRow(sql, player.Name, player.Avatar, player.DiscordId)
	err := row.Scan(&player.Id)

	if err != nil {
		return player, err
	}

	return player, nil
}

func FetchAllPlayers(w http.ResponseWriter, r *http.Request) {
	rows, err := database.Database.Query("SELECT * FROM player")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var players []Player

	for rows.Next() {
		var player Player

		err := rows.Scan(&player.Id, &player.Name, &player.Avatar, &player.DiscordId)
		if err != nil {
			panic(err)
		}

		players = append(players, player)
	}
	if err = rows.Err(); err != nil {
		panic(err)
	}

	fmt.Println("Endpoint Hit: returnAllPlayers")
	json.NewEncoder(w).Encode(players)
}
