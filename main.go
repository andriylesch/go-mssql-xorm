package main

import (
	"fmt"
	"go-mssql-xorm/dbmanagers"
	"go-mssql-xorm/models"
)

func main() {

	fmt.Println("\n ------ ORM and MSSQL ------------ \n ")

	// Init DB
	dbmanagers.InitDB(false) // parameter false - will not show Sql code generated by ORM

	// Get Customers
	fmt.Println("------ GetCustomers ------------")
	resultDb := dbmanagers.GetCustomers()
	fmt.Printf("There are customers %d in DB \n\n", len(resultDb))

	// Insert Customers
	fmt.Println("------ InsertCustomer ------------")

	customerObj := models.Customer{
		Email:     "test@gmail.com",
		Nick:      "test1",
		Firstname: "FN_1",
		Lastname:  "LN_1",
		Age:       15}

	customerObj1 := models.Customer{
		Email:     "test2@gmail.com",
		Nick:      "test2",
		Firstname: "FN_2",
		Lastname:  "LN_2",
		Age:       35}

	insertRes := dbmanagers.InsertCustomer(customerObj)
	fmt.Println("Customer was inserted : ", insertRes)

	insertRes1 := dbmanagers.InsertCustomer(customerObj1)
	fmt.Println("Customer was inserted : ", insertRes1)

	resultDb = dbmanagers.GetCustomers()

	for _, item := range resultDb {
		fmt.Println(item.ToString())
	}

	// Update Customer
	fmt.Println("------ Update Customer ------------")

	customerUpd := resultDb[0]

	customerUpd.Nick = "test_upd"
	customerUpd.Email = "test_upd@gmail.com"

	updateRes := dbmanagers.UpdateCustomer(customerUpd)
	fmt.Println("Customer was updated : ", updateRes)

	// Delete Customer
	fmt.Println("------ Delete Customer ------------")
	updateRes = dbmanagers.DeleteCustomer(customerUpd)
	fmt.Println("Customer was deleted : ", updateRes)

	fmt.Println("------ GetCustomer ------------")
	resultDb = dbmanagers.GetCustomers()
	fmt.Printf("There are customers %d in DB \n\n", len(resultDb))

}
