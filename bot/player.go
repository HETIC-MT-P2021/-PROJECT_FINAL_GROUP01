package bot

import (
	"github.com/bwmarrin/discordgo"
	"github.com/dimo/database"
)

type Player struct {
	Id        int
	Name      string
	DiscordId string
	Avatar    string
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
