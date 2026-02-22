package game

type Market struct {
	Players map[string]*Player
	Cards   map[string]*Card
}

func NewMarket() *Market {
	return &Market{
		Players: make(map[string]*Player),
		Cards:   make(map[string]*Card),
	}
}

func (m *Market) purchase(playerId, cardId string, quantity int) bool {
	if quantity < 0 {
		return false
	}

	card := m.Cards[cardId]
	player := m.Players[playerId]

	if card == nil || player == nil {
		return false
	}

	cost := float64(quantity) * card.Price()

	if !card.DecreaseAvailableQuantity(quantity) {
		return false
	}

	if !player.Debit(cost) {
		card.IncreaseAvailableQuantity(quantity)

		return false
	}

	player.AddCard(cardId, quantity)

	card.UpdatePrice(card.Price() + priceDeltaFromTrade(card.Price(), quantity, 100))

	return true
}

func (m *Market) sell(playerId, cardId string, quantity int) bool {
	card := m.Cards[cardId]
	player := m.Players[playerId]

	if card == nil || player == nil {
		return false
	}

	cost := float64(quantity) * card.Price()

	if !card.IncreaseAvailableQuantity(quantity) {
		return false
	}

	if !player.RemoveCard(cardId, quantity) {
		return false
	}

	player.Credit(cost)

	card.UpdatePrice(card.Price() - priceDeltaFromTrade(card.Price(), quantity, 100))

	return true
}

func (m *Market) Card(cardId string) *Card {
	return m.Cards[cardId]
}

func (m *Market) AddCard(card *Card) bool {

	if _, ok := m.Cards[card.Id]; ok {
		return false
	}

	m.Cards[card.Id] = card

	return true

}

func (m *Market) AddPlayer(player *Player) bool {
	if _, ok := m.Players[player.Id]; ok {
		return false
	}

	m.Players[player.Id] = player

	return true
}

func priceDeltaFromTrade(price float64, qty, divider int) float64 {
	return price * (float64(qty) / float64(divider))
}
