package main

import "fmt"
type Bitcoin int
type Wallet struct {
    balance Bitcoin
}

func (w Wallet) Balance () int {
    fmt.Println("Balance  ", w.balance)
    return w.balance
}

func (w *Wallet) Deposit (value ) {
    fmt.Println("deposit  ",value)
    w.balance += value
}

