package main

import "fmt"
import "go-mssql-xorm/dbmanagers"

func main() {

	fmt.Println("\n ------ ORM and MSSQL ------------ \n ")

	// Init DB
	dbmanagers.InitDB(false)

	fmt.Println("------ GetUsers ------------")
	resultDb := dbmanagers.GetCustomers()
	fmt.Printf("There are users %d in DB \n\n", len(resultDb))

	fmt.Println("------ InsertUser ------------")

	// customerObj := models.Customer{
	// 	Email:     "test@gmail.com",
	// 	Nick:      "test1",
	// 	Firstname: "FN_1",
	// 	Lastname:  "LN_1",
	// 	Age:       15}

	// customerObj1 := models.Customer{
	// 	Email:     "test2@gmail.com",
	// 	Nick:      "test2",
	// 	Firstname: "FN_2",
	// 	Lastname:  "LN_2",
	// 	Age:       35}

	// insertRes := dbmanagers.InsertCustomer(customerObj)
	// fmt.Println("Customer was inserted : ", insertRes)

	// insertRes1 := dbmanagers.InsertCustomer(customerObj1)
	// fmt.Println("Customer was inserted : ", insertRes1)

	resultDb = dbmanagers.GetCustomers()

	for _, item := range resultDb {
		fmt.Println(item.ToString())
	}

	fmt.Println("------ Update User ------------")

	customerUpd := resultDb[0]

	customerUpd.Nick = "test_upd"
	customerUpd.Email = "test_upd@gmail.com"

	updateRes := dbmanagers.UpdateCustomer(customerUpd)
	fmt.Println("Customer was updated : ", updateRes)

	fmt.Println("------ Delete User ------------")

	updateRes = dbmanagers.DeleteCustomer(customerUpd)
	fmt.Println("Customer was deleted : ", updateRes)

	fmt.Println("------ GetUsers ------------")
	resultDb = dbmanagers.GetCustomers()
	fmt.Printf("There are users %d in DB \n\n", len(resultDb))

}
