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

func (m *Market) card(cardId string) *Card {
	return m.Cards[cardId]
}

func priceDeltaFromTrade(price float64, qty, divider int) float64 {
	return price * (float64(qty) / float64(divider))
}
