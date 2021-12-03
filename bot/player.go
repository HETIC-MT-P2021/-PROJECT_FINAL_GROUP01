package bot

import "github.com/bwmarrin/discordgo"

type Player struct {
	Id        string
	Name      string
	DiscordId string
	Avatar    string
}

func NewPlayerFromDiscordAuthor(author *discordgo.User) Player {
	return Player{
		Id:        author.ID,
		Name:      author.Username,
		DiscordId: author.Token,
		Avatar:    author.Avatar,
	}
}
