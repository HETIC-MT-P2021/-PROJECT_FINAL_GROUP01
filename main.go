package PROJECT_FINAL_GROUP01


import (
	"fmt"
	"github.com/joho/godotenv"
	"strconv"

	_ "github.com/lib/pq"

	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP01/bot"
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP01/config"
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP01/model"
)



func main() {
	fmt.Println("Starting Bot")

	env, _ := godotenv.Read(".env")

	dbPort, dbErr := strconv.ParseInt(env["DB_PORT"], 10, 64)

	if dbErr != nil {
		panic(dbErr)
	}

	model.ConnectToDB(env["DB_HOST"], env["DB_NAME"], env["DB_USER"], env["DB_PASSWORD"], dbPort)

	botCredsErr := config.ReadConfig()

	if botCredsErr != nil {
		fmt.Println(botCredsErr.Error())
		return
	}

	bot.Start()

	<-make(chan struct{})

	return
}
