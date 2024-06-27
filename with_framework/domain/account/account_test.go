package account

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAccount(t *testing.T) {
	t.Run("New 테스트", func(t *testing.T) {
		t.Run("userId를 받아서 새로운 Account 객체를 만든다", func(t *testing.T) {
			account := New("123")
			assert.Equal(t, account.UserId, "123")
			assert.Equal(t, account.Balance, 0)
		})
	})
	t.Run("Deposit 테스트", func(t *testing.T) {
		t.Run("금액을 입금할 수 있다.", func(*testing.T) {
			account := Account{Balance: 0}
			account.Deposit(100)
			assert.Equal(t, account.Balance, 100)
		})
		t.Run("입금할 금액이 0보다 작으면 에러를 반환한다.", func(*testing.T) {
			account := Account{Balance: 0}
			err := account.Deposit(-100)
			assert.Error(t, err)
			assert.Equal(t, err.Error(), "amount must be positive")
		})
	})
	t.Run("WithDraw 테스트", func(t *testing.T) {
		t.Run("금액을 출금할 수 있다.", func(*testing.T) {
			account := Account{Balance: 100}
			account.Withdraw(50)
			assert.Equal(t, account.Balance, 50)
		})

		t.Run("출금할 금액이 0보다 작으면 에러를 반환한다.", func(*testing.T) {
			account := Account{Balance: 100}
			err := account.Withdraw(-100)
			assert.Error(t, err)
			assert.Equal(t, err.Error(), "amount must be positive")
		})
	})

}
