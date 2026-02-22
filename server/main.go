package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/ameer005/card-capital/internal/game"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	card := game.NewCard("Banana", 50, 100, "1")
	market := game.NewMarket()

	market.AddCard(card)

	for i := 0; i < 3; i++ {
		player := game.NewPlayer(fmt.Sprintf("%d", i), 500.00)

		market.AddPlayer(player)
	}

	sim := game.NewSimulation(market, 500)

	sim.Run()

}
