package account
import (
	"sync"
)

// Create a global variable as a mutex, to limit read/write access to the account object.

var mu = &sync.Mutex{}

type Account struct {
	balance int64
	status bool
}

func Open(initialDeposit int64) *Account {

	if (initialDeposit < 0) {
		return nil
	}

	account := Account{initialDeposit, true}
	return &account

}

func (account *Account) Close() (payout int64, ok bool) {

	mu.Lock()

	if (!account.status) {
		mu.Unlock()

		return 0, false
	}

	ok = true
	payout = account.balance
	account.status = false

	mu.Unlock()

	return
}

func (account *Account) Balance() (balance int64, ok bool) {

	mu.Lock()

	if (!account.status) {

		mu.Unlock()

		return 0, false
	}

	balance = account.balance
	ok = true

	mu.Unlock()

	return

}

func (account *Account) Deposit(amount int64) (newBalance int64, ok bool) {

	mu.Lock()

	if (!account.status) {

		mu.Unlock()

		return 0, false
	}

	if (account.balance + amount < 0) {
		mu.Unlock()

		return 0, false
	}

	account.balance += amount
	newBalance = account.balance
	ok = true

	mu.Unlock()

	return
}
