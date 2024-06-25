package accountRepository

import (
	"database/sql"

	"without.framework/domain/account"
)

type repository interface {
	Save(account *account.Account) error
}

type accountRepository struct {
	db *sql.DB
}

func New(db *sql.DB) *accountRepository {
	return &accountRepository{db: db}
}

func (r *accountRepository) Save(account *account.Account) error {
	_, err := r.db.Exec("INSERT INTO account (id, balance, userId) VALUES (?, ?, ?)", account.Id, account.Balance, account.UserId)
	return err
}
