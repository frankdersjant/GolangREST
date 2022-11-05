package Domain

import (
	"database/sql"
	"log"
	"time"
)

//NOTE: PULIC MTHODS RE ALL CAPS
type CustomerRepository interface {
	FindAll() ([]Customer, error)
	FindById(id string) (*Customer, *AppError)
}

//We wrap the sql db stuff in  a struct CustomerRepoDB
type CustomerRepoDB struct {
	db *sql.DB
}

//We implement the interface CustomerRepository
//By using the struct CustomerRepoDB (the real implementation)
func (ch CustomerRepoDB) FindAll() ([]Customer, error) {

	findall_sql := "SELECT * FROM customers"

	rows, err := ch.db.Query(findall_sql)

	if err != nil {
		log.Println("error executing sql")
	}

	customers := make([]Customer, 0)

	for rows.Next() {
		var customer Customer
		err = rows.Scan(&customer.ID, &customer.Name, &customer.Address, &customer.DateofBirth, &customer.Status)
		if err != nil {
			log.Println("Error scanning rows" + err.Error())
		}
		customers = append(customers, customer)
	}

	return customers, nil
}

//We implement the interface CustomerRepository
//By using the struct CustomerRepoDB (the real implementation)
func (ch CustomerRepoDB) FindById(id string) (*Customer, *AppError) {
	var customer Customer
	err := ch.db.QueryRow("select * from customers where id=?", id).Scan(&customer.ID, &customer.Name, &customer.Address, &customer.DateofBirth, &customer.Status)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, NotFoundError("customer not found")
		} else {
			log.Println("Error scanning rows ById" + err.Error())
			return nil, UnexpectedError("Unexpected db error")
		}
	}
	return &customer, nil
}

func NewCustomerRepositoryDB() CustomerRepoDB {

	db, err := sql.Open("mysql", "root:password1@tcp(127.0.0.1:3306)/customers")

	if err != nil {
		panic(err)
	}

	db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(10)
	db.SetConnMaxLifetime(time.Minute * 3)

	return CustomerRepoDB{db}
}
