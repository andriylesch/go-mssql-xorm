package dbmanagers

/*
	mandatory moment
	package "github.com/denisenkom/go-mssqldb" should be exists in import section
	without it. it will not be possible to do request to DB
*/

import (
	"log"

	"go-mssql-xorm/models"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/go-xorm/xorm"
)

var (
	isShowSQLCode    bool   = false
	connectionString string = "driver={SQL Server};server=.\\GO;database=LocalDb;user id=testUser;password=123;"
)

func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}

// Create ORM object
func getOrm() *xorm.Engine {

	orm, err := xorm.NewEngine("mssql", connectionString)
	checkErr(err, "Creating ORM object was failed")

	orm.ShowSQL(isShowSQLCode)
	return orm
}

// InitDB method
// parameter isShowSql - true - show sql script in debug console/ false - will skip it
func InitDB(isShowSQL bool) {

	orm := getOrm()
	defer orm.Close()
	isShowSQLCode = isShowSQL
	err := orm.Sync2(new(models.Customer))
	checkErr(err, "Sync structures with DB was failed")
}

// Get list of customers
func GetCustomers() []models.Customer {
	var customers []models.Customer

	orm := getOrm()
	defer orm.Close()

	err := orm.Find(&customers)
	checkErr(err, "Get Data was failed")

	return customers
}

// Update Customer
func UpdateCustomer(customer models.Customer) bool {

	orm := getOrm()
	defer orm.Close()

	affected, err := orm.Id(customer.Id).Update(&customer)
	checkErr(err, "Update was failed")

	return affected > 0
}

// Insert Customer
func InsertCustomer(customer models.Customer) bool {

	if (models.Customer{}) == customer {
		return false
	}

	orm := getOrm()
	defer orm.Close()

	affected, err := orm.Insert(&customer)
	checkErr(err, "Insert was failed")

	return affected > 0
}

// Delete Customer
func DeleteCustomer(customer models.Customer) bool {

	if (models.Customer{}) == customer {
		return false
	}

	orm := getOrm()
	defer orm.Close()

	affected, err := orm.Id(customer.Id).Delete(&customer)
	checkErr(err, "Delete was failed")

	return affected > 0
}
