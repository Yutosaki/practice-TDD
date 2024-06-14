package wallet

import (
	"errors"
	"fmt"
)

type Bitcoin int
type Wallet struct{
	balance Bitcoin
}

func (w *Wallet)Deposit(amount Bitcoin){
	// fmt.Printf("address of balance in Deposit is %v \n", &w.balance)
	(*w).balance += amount//w.balanceでも自動的に(*w)のように逆参照してくれる
}

func (w *Wallet)Balance()Bitcoin{
	return w.balance
}

type Stringer interface{
	String() string
}

func (b Bitcoin)String()string{
	return fmt.Sprintf("%d BTC", b)
}

func (w *Wallet)Withdraw(amount Bitcoin)error{
	if w.balance < amount {
		return errors.New("oh no ")
	}
	w.balance -= amount
	return nil
}