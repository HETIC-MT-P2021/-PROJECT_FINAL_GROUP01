package Dimo

type Round struct {
	GameId      int `json:"game_id"`
	PlayerId    int `json:"player_id"`
	Reason      string `json:"reason"`
	Word        string `json:"word"`
	SubmittedAt string `json:"submitted_at"`
}

