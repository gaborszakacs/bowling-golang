package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/gaborszakacs/bowling-golang/bowling"
)

func main() {
	rollsInput := flag.String("rolls", "", "rolls separated by comma")
	flag.Parse()
	rolls := strings.Split(*rollsInput, ",")
	g := bowling.Game{}
	for _, roll := range rolls {
		n, err := strconv.Atoi(roll)
		if err != nil {
			os.Exit(1)
		}
		g.Roll(n)
	}
	fmt.Printf("Score: %d", g.Score())
}
