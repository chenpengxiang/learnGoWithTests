package main

import (
	"errors"
	"fmt"
)

type Bitcoin int

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

func (w *Wallet) Balance() Bitcoin {
	fmt.Printf("address of balance in test is %v \n", &w.balance)
	return w.balance
}

func (w *Wallet) Withdraw(amount Bitcoin) (err error) {
	if w.balance >= amount {
		w.balance -= amount
	} else {
		err = errors.New("cannot withdraw, insufficient funds")
	}
	return
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}
