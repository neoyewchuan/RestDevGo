package domain

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"

	"github.com/neoyewchuan/RestDevGo/banking/banking/errs"
)

type CustomerRepositoryDB struct {
	client *sql.DB
}

func (crdb CustomerRepositoryDB) FindAll() ([]Customer, *errs.AppError) {

	findAllSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers"

	rows, err := crdb.client.Query(findAllSql)
	if err != nil {
		log.Println("Error while querying customer table: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	customers := make([]Customer, 0)
	for rows.Next() {
		var c Customer
		err := rows.Scan(&c.ID, &c.Name, &c.City, &c.ZipCode, &c.DateOfBirth, &c.Status)
		if err != nil {
			log.Println("Error while querying customer table: " + err.Error())
			return nil, errs.NewUnexpectedError("Unexpected database error")
		}
		customers = append(customers, c)
	}
	return customers, nil
}

func (crdb CustomerRepositoryDB) ByID(id string) (*Customer, *errs.AppError) {
	findOneSql := `select customer_id, name, city, zipcode, date_of_birth, status from customers 
			where customer_id = ?`
	row := crdb.client.QueryRow(findOneSql, id)
	var c Customer
	err := row.Scan(&c.ID, &c.Name, &c.City, &c.ZipCode, &c.DateOfBirth, &c.Status)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("Customer not found")
		}
		log.Println("Error while scanning customer: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}
	return &c, nil
}

func NewCustomerRepositoryDB() CustomerRepositoryDB {
	client, err := sql.Open("mysql", "root:toorroot@tcp(localhost:3306)/banking")
	if err != nil {
		panic(err)
	}
	// See important settings section
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	return CustomerRepositoryDB{client}
}
