package domain

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"

	"github.com/neoyewchuan/RestDevGo/banking/errs"
	"github.com/neoyewchuan/RestDevGo/banking/logger"
)

type CustomerRepositoryDB struct {
	client *sqlx.DB
}

func (crdb CustomerRepositoryDB) FindAll(status string) ([]Customer, *errs.AppError) {
	//var rows *sql.Rows
	var err error
	customers := make([]Customer, 0)
	findAllSql := `select customer_id, name, city, zipcode, date_of_birth, status from customers`
	if status == "" {
		err = crdb.client.Select(&customers, findAllSql)
	} else {
		findAllSql += ` where status = ?`
		err = crdb.client.Select(&customers, findAllSql, status)
	}

	if err != nil {
		logger.Error("Error while querying customer table: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}
	//if len(customers) == 0 {
	if err == sql.ErrNoRows {
		logger.Info("Customer table is empty")
		return nil, errs.NewNotFoundError("Customer table is empty")
	}
	// using of sqlx StructScan
	// err = sqlx.StructScan(rows, &customers)
	// if err != nil {
	// 	logger.Error("Error while querying customer table: " + err.Error())
	// 	return nil, errs.NewUnexpectedError("Unexpected database error")
	// }

	// standard sql linear scanning of returned result rows
	// for rows.Next() {
	// 	var c Customer
	// 	err := rows.Scan(&c.ID, &c.Name, &c.City, &c.ZipCode, &c.DateOfBirth, &c.Status)
	// 	if err != nil {
	// 		logger.Error("Error while querying customer table: " + err.Error())
	// 		return nil, errs.NewUnexpectedError("Unexpected database error")
	// 	}
	// 	customers = append(customers, c)
	// }
	return customers, nil
}

func (crdb CustomerRepositoryDB) ByID(id string) (*Customer, *errs.AppError) {
	findOneSql := `select customer_id, name, city, zipcode, date_of_birth, status from customers 
			where customer_id = ?`
	//row := crdb.client.QueryRow(findOneSql, id)
	var customer Customer
	err := crdb.client.Get(&customer, findOneSql, id)

	//err := row.Scan(&c.ID, &c.Name, &c.City, &c.ZipCode, &c.DateOfBirth, &c.Status)
	if err != nil {
		//if err == sql.ErrNoRows {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("Customer not found")
		}
		logger.Error("Error while scanning customer: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}
	return &customer, nil
}

func NewCustomerRepositoryDB() CustomerRepositoryDB {
	client, err := sqlx.Open("mysql", "root:toorroot@tcp(localhost:3306)/banking")
	if err != nil {
		panic(err)
	}
	// See important settings section
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	return CustomerRepositoryDB{client}
}
