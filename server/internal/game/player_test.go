package game

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddCard(t *testing.T) {

	t.Run("Add card", func(t *testing.T) {
		p := NewPlayer("1", 500)
		p.AddCard("1", 10)
		pCard, err := p.getCard("1")
		assert.NoError(t, err)

		assert.Equal(t, pCard.quantity, 10)
	})

	t.Run("No Card Found", func(t *testing.T) {
		p := NewPlayer("1", 500)
		_, err := p.getCard("1")

		assert.Error(t, err)
	})

	t.Run("Remove card", func(t *testing.T) {
		p := NewPlayer("1", 500)
		p.AddCard("1", 10)

		ok := p.RemoveCard("1", 10)
		assert.True(t, ok)

		_, err := p.getCard("1")

		assert.Error(t, err)

	})

	t.Run("Remove card Quantity", func(t *testing.T) {
		p := NewPlayer("1", 500)
		p.AddCard("1", 10)

		ok := p.RemoveCard("1", 5)
		assert.True(t, ok)

		pCard, err := p.getCard("1")
		assert.NoError(t, err)

		assert.Equal(t, pCard.quantity, 5)

		ok = p.RemoveCard("1", 5)
		assert.True(t, ok)

		assert.Equal(t, pCard.quantity, 0)

		_, err = p.getCard("1")
		assert.Error(t, err)
	})

}
