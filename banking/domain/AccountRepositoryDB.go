package domain

import (
	"database/sql"
	"strconv"

	"github.com/jmoiron/sqlx"

	"github.com/neoyewchuan/RestDevGo/banking/errs"
	"github.com/neoyewchuan/RestDevGo/banking/logger"
)

type AccountRepositoryDB struct {
	client *sqlx.DB
}

func (ardb AccountRepositoryDB) Save(a Account) (*Account, *errs.AppError) {
	// need to first check if account (type) already exist for the given customer
	// to avoid creating duplicate accounts
	account := make([]Account, 0)
	acctExistQuery := "SELECT customer_id,account_type FROM accounts WHERE customer_id = ? AND account_type = ?"
	err := ardb.client.Select(&account, acctExistQuery, a.CustomerID, a.AccountType)
	if err != nil {
		//if err == sql.ErrNoRows {
		if err != sql.ErrNoRows {
			//return nil, errs.NewNotFoundError("Account already exist")
			return nil, errs.NewNotFoundError(err.Error())
		}
	}
	if len(account) > 0 {
		return nil, errs.NewNotFoundError("Account already exist")
	}

	sqlInsert := `INSERT INTO accounts (customer_id, opening_date, account_type, amount, status) values (?, ?, ?, ?, ?) `
	//sqlInsert += ` WHERE ? NOT IN (SELECT customer_id+account_type FROM accounts)`
	result, err := ardb.client.Exec(sqlInsert, a.CustomerID, a.OpeningDate, a.AccountType, a.Amount, a.Status)
	if err != nil {
		logger.Error("Error while creating new account: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error from database")
	}

	id, err := result.LastInsertId()
	if err != nil {
		logger.Error("Error while getting last insert id for new account: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error from database")
	}
	a.AccountID = strconv.FormatInt(id, 10)
	return &a, nil

}

func NewAccountRepositoryDB(dbClient *sqlx.DB) AccountRepositoryDB {
	return AccountRepositoryDB{dbClient}
}
