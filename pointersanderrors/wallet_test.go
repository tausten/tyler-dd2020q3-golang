package pointersanderrors

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWallet(t *testing.T) {

	t.Run("Deposit", func(t *testing.T) {
		// Given
		wallet := Wallet{}

		// When
		wallet.Deposit(Bitcoin(10.0))
		result := wallet.Balance()

		// Then
		assert.Equal(t, Bitcoin(10.0), result)
	})

	t.Run("Withdraw", func(t *testing.T) {
		// Given
		wallet := Wallet{balance: Bitcoin(20)}

		// When
		err := wallet.Withdraw(Bitcoin(10.0))
		result := wallet.Balance()

		// Then
		assert.NoError(t, err)
		assert.Equal(t, Bitcoin(10), result)
	})

	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		// Given
		startingBalance := Bitcoin(20)
		wallet := Wallet{balance: startingBalance}

		// When
		err := wallet.Withdraw(Bitcoin(100.0))

		// Then
		assert.Equal(t, ErrInsufficientFunds, err)
	})
}

func TestBitcoinStringer(t *testing.T) {
	// Given
	b := Bitcoin(12.34567890)

	// When
	result := fmt.Sprintf("%s", b)

	// Then
	assert.Equal(t, "12.345679 BTC", result)
}
