package domain

import (
	"database/sql"
	"github.com/bankingApp/errs"
	"github.com/bankingApp/logger"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type CustomerRepositoryDb struct {
	client *sqlx.DB
}

func (d CustomerRepositoryDb) FindAll(status string) ([]Customer, *errs.AppError) {
	var err error
	customers := make([]Customer, 0)

	if status == "" {
		findAllSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers"
		err = d.client.Select(&customers, findAllSql)

	} else {

		findAllSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers where status = ?"
		err = d.client.Select(&customers, findAllSql, status)

	}
	if err != nil {
		logger.Error("Error while querying customer table" + err.Error())
		return nil, errs.NewUnexpectedError("unexpected database error")

	}

	return customers, nil
}

func (d CustomerRepositoryDb) FindAllByStatus(status string) ([]Customer, *errs.AppError) {

	statusId := ""
	if status == "active" {
		statusId = "1"
	}
	if status == "inactive" {
		statusId = "0"
	}

	findAllSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers where status = ?"

	rows, err := d.client.Query(findAllSql, statusId)
	if err != nil {
		logger.Error("Error while querying customer table" + err.Error())
		return nil, errs.NewUnexpectedError("unexpected database error")

	}
	customers := make([]Customer, 0)
	err = sqlx.StructScan(rows, &customers)

	if err != nil {
		logger.Error("Error while scanning customers" + err.Error())
		return nil, errs.NewUnexpectedError("unexpected database error")
	}

	return customers, nil
}

func (d CustomerRepositoryDb) FindCustomerById(id string) (*Customer, *errs.AppError) {
	customerSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers where customer_id = ?"

	var c Customer
	err := d.client.Get(&c, customerSql, id)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("customer not found")
		} else {
			logger.Error("Error while scanning customer " + err.Error())
			return nil, errs.NewUnexpectedError("unexpected database error")
		}
	}
	return &c, nil
}

func NewCustomerRepositoryDb(dbClient *sqlx.DB) CustomerRepositoryDb {
	return CustomerRepositoryDb{dbClient}
}
