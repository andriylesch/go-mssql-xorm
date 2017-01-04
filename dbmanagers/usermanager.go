package dbmanagers

import (
	"fmt"
	"log"

	"go-mssql-xorm/models"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/go-xorm/xorm"
)

const connectionString string = "driver={SQL Server};server=.\\GO;database=LocalDb;user id=testUser;password=123;"

func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}

func getOrm() *xorm.Engine {

	orm, err := xorm.NewEngine("mssql", connectionString)
	checkErr(err, "create ORM object was failed")

	return orm
}

func InitDB(isShowSql bool) {

	orm := getOrm()
	defer orm.Close()
	orm.ShowSQL(isShowSql)

	err := orm.Sync2(new(models.Customer))
	checkErr(err, "Sync structures with DB was failed")
}

func GetCustomers() []models.Customer {
	var customers []models.Customer

	orm := getOrm()
	defer orm.Close()

	err := orm.Find(&customers)

	// rows, err := orm.Rows(&models.Customer{})
	if err != nil {
		fmt.Println(err)
	}

	return customers
}

func UpdateCustomer(customer models.Customer) bool {

	orm := getOrm()
	defer orm.Close()

	affected, err := orm.Id(customer.Id).Update(&customer)

	if err != nil {
		fmt.Println(err)
	}

	return affected > 0
}

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

// func XormLogic() {

// 	engine, err := xorm.NewEngine("mssql", connectionString)
// 	checkErr(err, "sql.Open failed")

// 	defer engine.Close()

// 	//engine.ShowSQL(true)
// 	//engine.SetMaxOpenConns(5)

// 	err = engine.Sync2(new(Customer))
// 	checkErr(err, "Sync2 failed")

// 	results, err := engine.Query("select * from Customer")
// 	checkErr(err, "select request")

// 	fmt.Println(results)

// 	engine.ShowSQL(true)
// 	customer := Customer{Email: "test1@gmail.com", Age: 20, Name: "test1"}

// 	affected, err := engine.Insert(&customer)
// 	checkErr(err, "insert request")

// 	fmt.Println(affected)

// 	//var customer1 Customer

// 	// results1, err1 := engine.Get(&customer1)
// 	// checkErr(err1, "Get request")
// 	// fmt.Println(results1)
// 	// fmt.Println(customer1)

// 	//results, err := engine.Query("select * from userdata")

// }

// func initDb() *gorp.DbMap {
// 	// connect to db using standard Go database/sql API
// 	// use whatever database/sql driver you wish
// 	db, err := sql.Open("mssql", connectionString)
// 	checkErr(err, "sql.Open failed")

// 	// construct a gorp DbMap
// 	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.SqlServerDialect{}}

// 	// add a table, setting the table name to 'posts' and
// 	// specifying that the Id property is an auto incrementing PK
// 	dbmap.AddTableWithName(Post{}, "posts").SetKeys(true, "Id")

// 	// create the table. in a production system you'd generally
// 	// use a migration tool, or create the tables via scripts
// 	err = dbmap.CreateTablesIfNotExists()
// 	checkErr(err, "Create tables failed")

// 	return dbmap
// }

// func TestData() {

// 	// initialize the DbMap
// 	dbmap := initDb()
// 	defer dbmap.Db.Close()

// 	// delete any existing rows
// 	err := dbmap.TruncateTables()
// 	checkErr(err, "TruncateTables failed")

// 	// create two posts
// 	p1 := Post{Title: "Go 1.1 released!", Body: "Lorem ipsum lorem ipsum"}
// 	p2 := Post{Title: "Go 1.2 released!", Body: "Lorem ipsum lorem ipsum"}

// 	// insert rows - auto increment PKs will be set properly after the insert
// 	err = dbmap.Insert(&p1, &p2)
// 	checkErr(err, "Insert failed")

// }

// func GetUsers() models.Foo {

// 	// var resultsDb []models.User
// 	// var resultDb models.User

// 	db := getSqlDB()
// 	defer db.Close()

// 	// if db.Where("Id = ?", "15000010").Find(&resultsDb).Error == nil {

// 	// }

// 	// db.DropTable(&models.Foo{})

// 	// if ok := db.HasTable("foos"); ok {
// 	// 	fmt.Println("Table should not exist, but does")
// 	// }

// 	// if ok := db.HasTable(&models.Foo{}); ok {
// 	// 	fmt.Println("Table should not exist, but does")
// 	// }

// 	// // We create the table
// 	// if err := db.CreateTable(&models.Foo{}).Error; err != nil {
// 	// 	fmt.Println("Table should be created")
// 	// }

// 	// u := models.Foo{Desc: "test4"}

// 	// if err := db.Save(&u).Error; err != nil {
// 	// 	fmt.Println("No error should raise :", err)
// 	// }

// 	// var fooObj models.Foo
// 	// db.Find

// 	// fmt.Println(fooObj)

// 	// // And now it should exits, and HasTable should return true
// 	// if ok := db.HasTable("foos"); !ok {
// 	// 	fmt.Println("Table should exist, but HasTable informs it does not")
// 	// }

// 	// if ok := db.HasTable(&models.Foo{}); !ok {
// 	// 	fmt.Println("Table should exist, but HasTable informs it does not")
// 	// }

// 	// fmt.Println(resultDb.ToString())

// 	return fooObj
// }

// var (
// DB                 *gorm.DB)

// func getSqlDB() *gorm.DB {

// 	db, err := gorm.Open("mssql", connectionString)
// 	if err != nil {
// 		fmt.Println("Fatal error : ", err)
// 		return nil
// 	}

// 	return db
// }

// func InsertUser(model models.User) bool {

// 	// var userDb []models.User

// 	// db := getSqlDB()
// 	// defer db.Close()

// 	//db.Where("",)

// 	// if (models.User{}) == model {
// 	// 	return false
// 	// }

// 	// db := getSqlDB()
// 	// if db == nil {
// 	// 	return false
// 	// }

// 	// defer db.Close()

// 	// _, err := db.Exec(
// 	// 	"INSERT INTO [dbo].[Users] ([Nick],[EMail],[Nom],[Prenom],[Adr1],[Ville]) VALUES (? ,? ,? ,? ,? ,?)	",
// 	// 	model.Nick,
// 	// 	model.Email,
// 	// 	model.Nom,
// 	// 	model.Prenom,
// 	// 	model.Adr1,
// 	// 	model.Ville)

// 	// if err != nil {
// 	// 	fmt.Println(err)
// 	// 	return false
// 	// }

// 	return false
// }
