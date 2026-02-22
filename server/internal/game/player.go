package game

import "fmt"

type PortfolioCard struct {
	quantity int
}

type Player struct {
	Id        string
	portfolio map[string]*PortfolioCard
	money     float64
}

func NewPlayer(id string, money float64) *Player {

	return &Player{
		Id:        id,
		portfolio: make(map[string]*PortfolioCard),
		money:     money,
	}
}

func (p *Player) Credit(amount float64) {
	p.money += amount
}

func (p *Player) Debit(amount float64) bool {
	if amount > p.money {
		return false
	}

	p.money = p.money - amount

	return true
}

func (p *Player) AddCard(cardId string, qty int) {
	card := p.portfolio[cardId]

	if card == nil {
		p.portfolio[cardId] = &PortfolioCard{quantity: 0}
		card = p.portfolio[cardId]
	}

	card.quantity += qty
}

func (p *Player) RemoveCard(cardId string, qty int) bool {
	card := p.portfolio[cardId]

	if card == nil {
		return false
	}

	if qty > card.quantity {
		return false
	}

	card.quantity -= qty

	if card.quantity == 0 {
		delete(p.portfolio, cardId)
	}

	return true
}

func (p *Player) getCard(cardId string) (*PortfolioCard, error) {
	if portfolioCard, ok := p.portfolio[cardId]; !ok {
		return portfolioCard, fmt.Errorf("No card found")
	} else {
		return portfolioCard, nil
	}
}
