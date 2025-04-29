package accounts

// Account struct
type Account struct {
	owner   string
	balance int
}

// NewAccount creates Account
// 객체를 복사하지 않고 원본 객체를 반환
// main에서 원본 객체에 접근하지 못하게 하면서 값을 설정 가능하게함 -> constructor
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}

// Deposit x amount on your account
func (a *Account) Deposit(amount int) {
	a.balance += amount
}

// Balance of your account
func (a Account) Balance() int {
	return a.balance
}

// Withdarw x
func (a *Account) Withdraw(amount int) {
	a.balance -= amount
}
