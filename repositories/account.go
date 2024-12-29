package repositories

import (
	"github.com/ladmakhi81/gobanks/database"
	"github.com/ladmakhi81/gobanks/entities"
)

type AccountRepository struct {
	DatabaseServer *database.DatabaseServer
}

func (repo AccountRepository) CreateAccount(account *entities.Account) error {
	sql := `
		INSERT INTO "_accounts" 
		("first_name", "last_name", "balance", "created_at") 
		VALUES ($1, $2, $3, $4)
		RETURNING "id", "first_name", "last_name", "balance", "created_at";
	`
	result := repo.DatabaseServer.DB.QueryRow(sql, account.FirstName, account.LastName, account.Balance, account.CreatedAt)
	err := result.Scan(
		&account.ID,
		&account.FirstName,
		&account.LastName,
		&account.Balance,
		&account.CreatedAt,
	)
	if err != nil {
		return err
	}
	return nil
}

func (repo AccountRepository) DeleteAccount(id int) error {
	sql := `
		DELETE FROM "_accounts" WHERE id = $1;
	`
	statement, pErr := repo.DatabaseServer.DB.Prepare(sql)
	if pErr != nil {
		return pErr
	}
	_, eErr := statement.Exec(id)
	if eErr != nil {
		return eErr
	}
	return nil
}

func (repo AccountRepository) UpdateAccount(account *entities.Account) error {
	sql := `
		UPDATE "_accounts" 
		SET "first_name"=$1, 
		SET "last_name"=$2, 
		SET "number"=$3, 
		SET "balance"=$4 
		WHERE "id"=$5;
	`
	statement, pErr := repo.DatabaseServer.DB.Prepare(sql)
	if pErr != nil {
		return pErr
	}
	_, eErr := statement.Exec(
		account.FirstName,
		account.LastName,
		account.Number,
		account.Balance,
		account.ID,
	)
	if eErr != nil {
		return eErr
	}
	return nil
}

func (repo AccountRepository) GetAccounts() ([]*entities.Account, error) {
	sql := `
		SELECT * FROM "_accounts" ORDER BY "id" DESC;
	`
	rows, err := repo.DatabaseServer.DB.Query(sql)
	if err != nil {
		return nil, err
	}

	accounts := []*entities.Account{}

	for rows.Next() {
		account := new(entities.Account)
		sErr := rows.Scan(
			&account.ID,
			&account.FirstName,
			&account.LastName,
			&account.Number,
			&account.Balance,
			&account.CreatedAt,
		)
		if sErr != nil {
			return nil, sErr
		}
		accounts = append(accounts, account)
	}
	return accounts, nil
}

func (repo AccountRepository) GetAccountByID(id int) (*entities.Account, error) {
	sql := `
		SELECT * FROM "_accounts" WHERE "id"=$1
	`
	row := repo.DatabaseServer.DB.QueryRow(sql, id)
	account := new(entities.Account)
	sErr := row.Scan(
		&account.ID,
		&account.FirstName,
		&account.LastName,
		&account.Number,
		&account.Balance,
		&account.CreatedAt,
	)
	if sErr != nil {
		return nil, sErr
	}
	return account, nil
}

func (repo AccountRepository) GetAccountByNumber(number int) (*entities.Account, error) {
	sql := `
		SELECT * FROM "_accounts" WHERE "number"=$1
	`
	row := repo.DatabaseServer.DB.QueryRow(sql, number)
	account := new(entities.Account)
	sErr := row.Scan(
		&account.ID,
		&account.FirstName,
		&account.LastName,
		&account.Number,
		&account.Balance,
		&account.CreatedAt,
	)
	if sErr != nil {
		return nil, sErr
	}
	return account, nil
}

func (repo AccountRepository) WithDrawCredit(number int, amount float64) error {
	sql := `
		UPDATE "_accounts" SET "balance"=balance - $1 WHERE "number"=$2;
	`
	statement, pErr := repo.DatabaseServer.DB.Prepare(sql)

	if pErr != nil {
		return pErr
	}

	_, eErr := statement.Exec(amount, number)

	if eErr != nil {
		return eErr
	}

	return nil
}
func (repo AccountRepository) DepositCredit(number int, amount float64) error {
	sql := `
		UPDATE "_accounts" SET "balance"=balance + $1 WHERE "number"=$2;
	`
	statement, pErr := repo.DatabaseServer.DB.Prepare(sql)

	if pErr != nil {
		return pErr
	}

	_, eErr := statement.Exec(amount, number)

	if eErr != nil {
		return eErr
	}

	return nil
}
