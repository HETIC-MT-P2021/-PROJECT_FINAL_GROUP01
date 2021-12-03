package Dimo

type Player struct {
	PlayerId  int `json:"player_id"`
	Name      string `json:"name"`
	Avatar    string `json:"avatar"`
	DiscordId int `json:"discord_id"`
}