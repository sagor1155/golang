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

type Depositable interface {
	Deposit(float64)
}

type Withdrawable interface {
	Withdraw(float64) (float64, error)
}

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
type OverdraftableBankAccount struct {
	BankAccount
	Fee float64
}

func (oba *OverdraftableBankAccount) Withdraw(amount float64) (float64, error) {
	var err error
	if oba.balance < amount {
		oba.balance -= oba.Fee
		err = ErrOverdraftIncurred
	}
	oba.balance -= amount
	return oba.balance, err
}

func Transfer(debtor Withdrawable, creditor Depositable, amount float64) error {
	balance, err := debtor.Withdraw(amount)

	switch err {
	case ErrInsufficientBalance:
		return err
	case ErrOverdraftIncurred:
		log.Printf("debtor incurred overdraft: new balance is %.2f", balance)
	}

	creditor.Deposit(amount)
	return nil
}

func main() {
	fmt.Println("Example of STRUCT")
	ba := &BankAccount{"Joe", 100}
	fmt.Println("Balance:", ba.Balance())
	ba.Deposit(30.0)
	fmt.Println("Balance after deposit:", ba.Balance())

	balance, err := ba.Withdraw(200)
	fmt.Println("Error:", err)
	fmt.Println("Withdraw amount:", balance)

	oba := &OverdraftableBankAccount{BankAccount{"Billy", 50}, 20}
	fmt.Println(oba.Owner, "has:", oba.balance, "bucks")
	balance, err = oba.Withdraw(20)
	if err == nil {
		fmt.Println("After withdraw,", oba.Owner, "has:", oba.balance, "bucks")
	}

	_, err = oba.Withdraw(150)
	if err != nil {
		log.Println(err)
		log.Printf("Overdraft incurred: balance is now %.2f\n", oba.Balance())
	}

	balance, err = ba.Withdraw(150)
	if err != nil {
		fmt.Println(err)
	}

	// ba.Deposit(50)
	fmt.Println("Acount 1 balance:", ba.Balance(), ", Account 2 balance:", oba.Balance())

	err = Transfer(ba, oba, 150)
	if err != nil {
		log.Printf("Transfer Unsuccessfull: %v", err)
	} else {
		log.Printf("Transfer Successfull")
		fmt.Println("Acount 1 balance:", ba.Balance(), ", Account 2 balance:", oba.Balance())
	}

	err = Transfer(oba, ba, 150)
	if err != nil {
		log.Printf("Transfer Unsuccessfull: %v", err)
	} else {
		log.Printf("Transfer Successfull")
		fmt.Println("Acount 1 balance:", ba.Balance(), ", Account 2 balance:", oba.Balance())
	}
}
