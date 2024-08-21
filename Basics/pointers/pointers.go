package pointers

import (
	"fmt"
	"errors"
)

// Bitcoin is a type that represents a number of Bitcoins
// We're making a new type and we can declare methods on them
// This can be very useful when you want to add some domain specific functionality on top of existing types.
type Bitcoin int

// Stringer is an interface that has a single method String
// This is a way to define a string representation of a type
type Stringer interface {
	String() string
}

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

func (w *Wallet) Withdraw(amount Bitcoin) error{
	if amount > w.balance {
		return ErrInsufficientFunds
	}

	w.balance -= amount
	return nil
}