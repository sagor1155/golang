package main

import (
	"errors"
	"fmt"
	"log"
)

var (
	ErrInsufficientBalance = errors.New("Insufficient Balance!")
	ErrOverdraftIncurred   = errors.New("Overdraft Incurred!")
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

// struct embedded with another struct
type OverdrafttableBankAccount struct {
	account BankAccount
	Fee     float64
}

func (oba *OverdrafttableBankAccount) Withdraw(amount float64) (float64, error) {
	var err error
	if oba.account.balance < amount {
		oba.account.balance -= oba.Fee
		err = ErrOverdraftIncurred
	}
	oba.account.balance -= amount
	return oba.account.balance, err
}

func main() {
	fmt.Println("Example of STRUCT")
	ba := BankAccount{"Joe", 100}
	fmt.Println("Balance:", ba.Balance())
	ba.Deposit(30.0)
	fmt.Println("Balance after deposit:", ba.Balance())

	balance, err := ba.Withdraw(200)
	fmt.Println("Error:", err)
	fmt.Println("Withdraw amount:", balance)

	oba := OverdrafttableBankAccount{BankAccount{"Billy", 50}, 20}
	fmt.Println(oba.account.Owner, "has:", oba.account.balance, "bucks")
	balance, err = oba.account.Withdraw(20)
	if err == nil {
		fmt.Println("After withdraw,", oba.account.Owner, "has:", oba.account.balance, "bucks")
	}

	_, err = oba.Withdraw(150)
	if err != nil {
		log.Println(err)
		log.Printf("Overdraft incurred: balance is now %.2f\n", oba.account.Balance())
	}

	balance, err = ba.Withdraw(150)
	if err != nil {
		fmt.Println(err)
	}

}
