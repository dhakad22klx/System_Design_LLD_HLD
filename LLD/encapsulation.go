package main

import (
    "errors"
    // "fmt"
)

type BankAccount struct {
    accountHolder string //first letter is lower case, so it is private to a package.
    balance       float64
}

func NewBankAccount(accountHolder string) *BankAccount {
    return &BankAccount{accountHolder: accountHolder, balance: 0.0}
}

func (a *BankAccount) Deposit(amount float64) error {
    if amount <= 0 {
        return errors.New("deposit amount must be positive")
    }
    a.balance += amount
    return nil
}

func (a *BankAccount) Withdraw(amount float64) error {
    if amount <= 0 {
        return errors.New("withdrawal amount must be positive")
    }
    if amount > a.balance {
        return errors.New("insufficient funds")
    }
    a.balance -= amount
    return nil
}

func (a *BankAccount) GetBalance() float64 {
    return a.balance
}

func (a *BankAccount) GetAccountHolder() string {
    return a.accountHolder
}