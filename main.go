package main


import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP01/Dimo"
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP01/model"

	"github.com/bwmarrin/discordgo"
)

var (
	token string
	games map[string]*Dimo.Game
)

func main() {
	//Database connection
	env, _ := godotenv.Read(".env")

	dbPort, dbErr := strconv.ParseInt(env["DB_PORT"], 10, 64)

	if dbErr != nil {
		panic(dbErr)
	}

	model.ConnectToDB(env["DB_HOST"], env["DB_NAME"], env["DB_USER"], env["DB_PASSWORD"], dbPort)

	//Bot connection

	games = make(map[string]*Dimo.Game)

	var conf = Dimo.NewConfig
	token = conf().Token
	sess, err := discordgo.New("Bot " + token)
	if err != nil {
		log.Println("Unable to create a session on the server for the bot dimo")
		return
	}

	sess.AddHandler(messageCreate)

	err = sess.Open()
	if err != nil {
		fmt.Println("Error opening connection", err)
		return
	}

	defer sess.Close()

	fmt.Println("Dimo Now running, Press CTRL-C to exit")
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-ch
}
//player participation in the game

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	log.Println(m.Content)

	if m.Content == "/play" {
		game := games[m.ChannelID]
		if game == nil {
			s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("%v just created a game press /join to join the game", m.Author.Username))
			games[m.ChannelID] = Dimo.NewGame(m.Author.Username, func(winner string) {
				delete(games, m.ChannelID)
			})
		} else {
			s.ChannelMessageSend(m.ChannelID, "There is an active game, you can join it")
		}
	}

	if m.Content == "👋" {
		game := games[m.ChannelID]
		var resp string
		if game == nil {
			resp = "There is no active game. Create a game"
		}
		resp = game.AddPlayer(m.Author.Username)
		s.ChannelMessageSend(m.ChannelID, resp)
	}

	if m.Content == "➡️" {
		game := games[m.ChannelID]
		if game == nil {
			s.ChannelMessageSend(m.ChannelID, "There is no active game. Create a game")
			return
		}
		if game.IsActive() {
			s.ChannelMessageSend(m.ChannelID, fmt.Sprintln("The game started already"))
			return
		}
		game.Start()
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintln("The game has started"))
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("%s, it is your turn to play", game.GetCurrentPlayer()))
		return
	}

	if m.Content[0:1] != "/" {
		handleGame(s, m)
		return
	}
}

func handleGame(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	log.Println("")
	game := games[m.ChannelID]
	if game == nil || !game.IsActive() {
		return
	}

	resp := game.Play(m.Author.Username, m.Content)
	log.Println(resp)
	s.ChannelMessageSend(m.ChannelID, resp)
}
