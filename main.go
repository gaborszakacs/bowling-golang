package main

import (
	"fmt"
	"os"

	"github.com/gaborszakacs/bowling-golang/bowling"
	"github.com/gaborszakacs/bowling-golang/giphy"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
		os.Exit(1)
	}

	g := bowling.Game{Out: os.Stdout}
	for i := 0; i < 9; i++ {
		g.Roll(2)
		g.Roll(3)
	}
	g.Roll(10)
	g.PrintRolls()

	key := os.Getenv("GIPHY_API_KEY")
	url, err := g.Celebrate(giphy.Client{APIkey: key})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(url)
}
