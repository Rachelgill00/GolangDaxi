package main

import "fmt"

type account struct {
	AccountNo string
	Pwd       string
	Balance   float64
}

func NewAccount(a string, p string, b float64) *account {
	return &account{
		AccountNo: a,
		Pwd:       p,
		Balance:   b,
	}
}

// Method
// 1. Deposite
func (account *account) Deposite(money float64, pwd string) {
	if pwd != account.Pwd {
		fmt.Println("Password is wrong!")
	}

	if money <= 0 {
		fmt.Println("The input balance is wrong!")
	}

	account.Balance += money
	fmt.Println("Successfully deposit.")
	//fmt.Println(account.Balance)

}

func (account *account) String() string {
	return fmt.Sprintf("[AccountNo: %v, Balance: %v]", account.AccountNo, account.Balance)
}

func main() {
	// var account1 = Account{
	// 	AccountNo: "s111111",
	// 	Pwd:       "aaa111",
	// 	Balance:   82827,
	// }
	account1 := NewAccount("s111111", "aaa111", 82827)

	account1.Deposite(100, "aaa111")

	fmt.Println(account1)
}
