package domain

import (
	"database/sql"
	"strconv"

	"github.com/jmoiron/sqlx"

	"github.com/neoyewchuan/RestDevGo/banking/errs"
	"github.com/neoyewchuan/RestDevGo/banking/logger"
)

type TransactionRepositoryDB struct {
	client *sqlx.DB
}

func (trdb TransactionRepositoryDB) Deposit(txn Transaction) (*Transaction, *errs.AppError) {

	var resAccount Account

	sqlSelect := `SELECT account_id,customer_id,account_type,amount,status FROM accounts WHERE account_id = ?`
	err := trdb.client.Get(&resAccount, sqlSelect, txn.AccountID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("Account not found")
		}
		logger.Error("Error while scanning account: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	tx, err := trdb.client.Begin()
	if err != nil {
		logger.Error("Error while beginning sql transaction: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}
	sqlInsert := `INSERT INTO transactions (account_id, transaction_type, transaction_date, amount) VALUES (?, ?, ?, ?)`
	resInsert, err := tx.Exec(sqlInsert, txn.AccountID, txn.TransactionType, txn.TransactionDate, txn.Amount)
	if err != nil {
		tx.Rollback()
		logger.Error("Error while creating new transaction: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error from database")
	}
	id, err := resInsert.LastInsertId()
	if err != nil {
		logger.Error("Error while getting last insert id for new account: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error from database")
	}
	txn.TransactionID = strconv.FormatInt(id, 10)

	newBalance := resAccount.Amount + txn.Amount
	txn.Amount = newBalance
	// update account Balance
	sqlUpdate := `UPDATE accounts SET amount = ? WHERE account_id = ?`
	_, err = tx.Exec(sqlUpdate, newBalance, txn.AccountID)
	if err != nil {
		tx.Rollback()
		logger.Error("Error while updating account balance: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error from database")
	} else {
		tx.Commit()
	}
	return &txn, nil
}

func (trdb TransactionRepositoryDB) Withdrawal(txn Transaction) (*Transaction, *errs.AppError) {

	var resAccount Account

	// Need to verify that the account do have enough amount in account
	sqlSelect := `SELECT account_id,customer_id,account_type,amount,status FROM accounts WHERE account_id = ?`
	err := trdb.client.Get(&resAccount, sqlSelect, txn.AccountID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("Account not found")
		}
		logger.Error("Error while scanning account: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}
	// need to maintain at least $1.00 in account
	if resAccount.Amount-1 < txn.Amount {
		return nil, errs.NewValidationError("Insufficient balance in the account.")
	}
	tx, err := trdb.client.Begin()
	if err != nil {
		logger.Error("Error while beginning sql transaction: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}
	sqlInsert := `INSERT INTO transactions (account_id, transaction_type, transaction_date, amount) VALUES (?, ?, ?, ?)`
	resInsert, err := tx.Exec(sqlInsert, txn.AccountID, txn.TransactionType, txn.TransactionDate, txn.Amount)
	if err != nil {
		tx.Rollback()
		logger.Error("Error while creating new transaction: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error from database")
	}
	id, err := resInsert.LastInsertId()
	if err != nil {
		logger.Error("Error while getting last insert id for new account: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error from database")
	}
	txn.TransactionID = strconv.FormatInt(id, 10)

	newBalance := resAccount.Amount - txn.Amount
	txn.Amount = newBalance
	// update account Balance
	sqlUpdate := `UPDATE accounts SET amount = ? WHERE account_id = ?`
	_, err = tx.Exec(sqlUpdate, newBalance, txn.AccountID)
	if err != nil {
		tx.Rollback()
		logger.Error("Error while updating account balance: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error from database")
	} else {
		tx.Commit()
	}

	return &txn, nil
}

func NewTransactionRepositoryDB(dbClient *sqlx.DB) TransactionRepositoryDB {
	return TransactionRepositoryDB{dbClient}
}
