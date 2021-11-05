package Dimo

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Token string
}

func NewConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal()
	}
	token := os.Getenv("DISCORD_TOKEN")
	return &Config{
		Token: token,
	}
}

func (c *Config) GetToken() string {
	return c.Token
}
