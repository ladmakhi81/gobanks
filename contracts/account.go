package contracts

import "github.com/ladmakhi81/gobanks/entities"

type AccountRepository interface {
	CreateAccount(account *entities.Account) error
	UpdateAccount(account *entities.Account) error
	DeleteAccount(id int) error
	GetAccounts() ([]*entities.Account, error)
	GetAccountByID(id int) (*entities.Account, error)
	GetAccountByNumber(number int) (*entities.Account, error)
	WithDrawCredit(number int, amount float64) error
	DepositCredit(number int, amount float64) error
}
