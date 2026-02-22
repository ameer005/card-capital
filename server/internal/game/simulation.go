package game

import (
	"fmt"
	"math/rand"
)

type Simulation struct {
	market *Market
	tick   int
}

func NewSimulation(m *Market, t int) *Simulation {
	return &Simulation{market: m, tick: t}
}

func (s *Simulation) PlayerAction(p *Player) {

	for cardId := range s.market.Cards {
		action := rand.Intn(3)

		switch action {
		// buy
		case 0:
			qty := rand.Intn(8) + 1
			s.market.purchase(p.Id, cardId, qty)
		//sell
		case 1:
			qty := rand.Intn(8) + 1
			s.market.sell(p.Id, cardId, qty)
		default:
		}

	}

}

func (s *Simulation) Run() {
	s.printStats(-1)
	for i := 0; i < s.tick; i++ {
		for _, player := range s.market.Players {
			s.PlayerAction(player)
		}
		if (i+1)%100 == 0 {
			s.printStats(i + 1)
		}
	}

	if s.tick%100 != 0 {
		s.printStats(s.tick)
	}
}

func (s *Simulation) printStats(tick int) {
	fmt.Printf("\nStats after %d ticks\n", tick)
	fmt.Printf("Card Price: %f Qty left: %d\n",
		s.market.Card("1").Price(),
		s.market.Card("1").AvailaleQuantity(),
	)

	for _, player := range s.market.Players {
		playerCard, err := player.getCard("1")
		if err != nil {
			continue
		}

		fmt.Printf("player: %s, money: %f, holding: %d\n", player.Id, player.money, playerCard.quantity)
	}
}
