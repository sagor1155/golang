package main

import (
	"errors"
	"fmt"
)

var (
	ErrInsufficientBalance = errors.New("Insufficient Balance!")
)

type BankAccount struct {
	Owner   string
	balance float64
}

// plain object taken as we are not modifying any attribute
func (ba BankAccount) Balance() float64 {
	return ba.balance
}

// pointer object taken as we will modify attribute
func (ba *BankAccount) Deposit(amount float64) {
	ba.balance += amount
}

func (ba *BankAccount) Withdraw(amount float64) (float64, error) {
	if ba.balance < amount {
		return 0, ErrInsufficientBalance
	}
	ba.balance -= amount
	return ba.balance, nil
}

func main() {
	fmt.Println("Example of STRUCT")
	ba := BankAccount{"Joe", 100}
	fmt.Println("Balance:", ba.Balance())
	ba.Deposit(30.0)
	fmt.Println("Balance after deposit:", ba.Balance())

	withdraw_amount, err := ba.Withdraw(200)
	fmt.Println("Error:", err)
	fmt.Println("Withdraw amount:", withdraw_amount)

}
