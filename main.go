package main

import (
	"fmt"
	"github.com/dimo/bot"
	"github.com/dimo/database"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

var (
	token string
	games map[string]*bot.Game
)

func main() {
	// Database connection
	env, _ := godotenv.Read(".env")

	dbPort, dbErr := strconv.ParseInt(env["DB_PORT"], 10, 64)

	if dbErr != nil {
		panic(dbErr)
	}

	database.ConnectToDB(env["DB_HOST"], env["DB_NAME"], env["DB_USER"], env["DB_PASSWORD"], dbPort)

	// Start Dimo API
	go func() {
		// creates a new instance of a mux router
		myRouter := mux.NewRouter().StrictSlash(true)

		// replace http.HandleFunc with myRouter.HandleFunc
		myRouter.HandleFunc("/players", bot.FetchAllPlayers)
		myRouter.HandleFunc("/rounds", bot.FetchAllRounds)

		// finally, instead of passing in nil, we want
		// to pass in our newly created router as the second
		// argument
		log.Fatal(http.ListenAndServe(":8081", myRouter))
	}()

	// Start Dimo bot server
	games = make(map[string]*bot.Game)

	var conf = bot.NewConfig
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

	fmt.Println("Dimo API and Bot server is now running...\nPress CTRL + C to exit process.")

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-ch
}

// Player participation in the game
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content == "/play" {
		game := games[m.ChannelID]

		if game == nil {
			var player = bot.NewPlayerFromDiscordAuthor(m.Author)

			message, _ := s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Hello @here :smiley: \n**%v** just created a new Dimo game party :partying_face: \nHis game needs at least **2 players** to be started ! :wink: \n\n**Join him by sending:** :wave: \n**Start the game by sending:** :arrow_forward: \n\nHappy gaming ! :smirk_cat:", player.Name))

			games[m.ChannelID] = bot.NewGame(player, message.ID, func(winner bot.Player) {
				delete(games, m.ChannelID)
			})
		} else {
			s.ChannelMessageSend(m.ChannelID, "There is an active game, you can join it")
		}
	}

	if m.Content == "👋" || m.Content == "/join" {
		game := games[m.ChannelID]

		var resp string
		var status bool
		if game == nil {
			resp = "There is no active game. Create a game"
		}

		player := bot.NewPlayerFromDiscordAuthor(m.Author)
		status, resp = game.AddPlayer(&player)

		if status {
			s.ChannelMessageEdit(m.ChannelID, game.FirstMessageId, fmt.Sprintf(
				"Hello @here :smiley: \n**%v** just created a new Dimo game party :partying_face: \nHis game needs at least **2 players** to be started ! :wink: \n\n**Join him by sending:** :wave: \n**Start the game by sending:** :arrow_forward: \n\n**Playing queue:**\n%v \nHappy gaming ! :smirk_cat:",
				game.StartedBy.Name,
				game.GetPlayersNames(),
			))
		} else {
			s.ChannelMessageSend(m.ChannelID, resp)
		}
	}

	if m.Content == "➡️" || m.Content == "/start" {
		game := games[m.ChannelID]

		if game == nil {
			s.ChannelMessageSend(m.ChannelID, "There is no active game. Create a game")
			return
		}

		if len(game.Players) == 1 {
			s.ChannelMessageSend(m.ChannelID, "You cannot play the game alone. Invite your friends to join you")
			return
		}

		if game.IsActive {
			s.ChannelMessageSend(m.ChannelID, fmt.Sprintln("The game started already"))
			return
		}

		game.Start()

		s.ChannelMessageSend(m.ChannelID, fmt.Sprintln("The game has started"))
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("%s, it is your turn to play", game.GetCurrentPlayer().Name))

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

	game := games[m.ChannelID]
	if game == nil || !game.IsActive {
		return
	}

	player := game.GetPlayerByDiscordId(m.Author.ID)
	if player != nil {
		resp := game.Play(*player, m.Content)
		s.ChannelMessageSend(m.ChannelID, resp)
	}
	// Player not found
}
