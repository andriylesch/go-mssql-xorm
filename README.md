# go-mssql-xorm
how to make friends GO &amp; XORM ORM &amp; MSSQL

# Installation

    go get github.com/andriylesch/go-mssql-xorm
    
# Documents

* [Manual] (http://xorm.io/docs)
* [GoDoc](http://godoc.org/github.com/go-xorm/xorm)

# How to start

* Create engine/ORM object

```Go
const connectionString string = "driver={SQL Server};server=.\\GO;database=LocalDb;user id=####;password=####;"

func getOrm() *xorm.Engine {
	orm, err := xorm.NewEngine("mssql", connectionString)
	checkErr(err, "Creating ORM object was failed")
	return orm
}
```
* Define your custom struct and made Sync2 *table* struct to DB

```Go
// "go-mssql-xorm/models"
type Customer struct {
	Id        int    `xorm:"ID PK int autoincr"`
	Nick      string `xorm:"Nick varchar(100) notnull"`
	Email     string `xorm:"Email varchar(100) notnull"`
	Firstname string `xorm:"Firstname varchar(100)"`
	Lastname  string `xorm:"Lastname varchar(100)"`
	Age       int    `xorm:"Age int default 0"`
}

//"go-mssql-xorm/dbmanagers"
func InitDB(isShowSql bool) {

	orm := getOrm()
	defer orm.Close()
	orm.ShowSQL(isShowSql)

	err := orm.Sync2(new(models.Customer))
	checkErr(err, "Sync structures with DB was failed")
}

```
*Detail information how you can configure some columns in your custom struct follow link [Column definition (https://lunny.gitbooks.io/xorm-manual-en-us/content/chapter-02/4.columns.html)*

* Insert record to database

```Go
  customerObj := models.Customer{
		Email:     "test@gmail.com",
		Nick:      "test1",
		Firstname: "FN_1",
		Lastname:  "LN_1",
		Age:       15}
  
  // init orm obj ...  
  
  affected, err := orm.Insert(&customerObj)
	checkErr(err, "Insert was failed")
```

* Update record in database
```Go
  customerObj := models.Customer{
    ID: 1
		Email:     "test@gmail.com",
		Nick:      "test1",
		Firstname: "FN_1_1",
		Lastname:  "LN_1_1",
		Age:       15}
    
  // init orm obj ...
   
  affected, err := orm.Id(customerObj.Id).Update(&customerObj)
	checkErr(err, "Insert was failed")
```

* Get records from DB
```Go
  var customers []models.Customer
  
  // init orm obj ...
  
	err := orm.Find(&customers)
	checkErr(err, "Get Data was failed")
```

* Delete record from DB
```Go
   customerObj := models.Customer{
    ID: 1
		Email:     "test@gmail.com",
		Nick:      "test1",
		Firstname: "FN_1_1",
		Lastname:  "LN_1_1",
		Age:       15}
  
  // init orm obj ...
  
	affected, err := orm.Id(customerObj.Id).Delete(&customerObj)
	checkErr(err, "Delete was failed")
```

# Reference

Source code
- https://github.com/go-xorm/xorm

Manual
- http://xorm.io/docs/ 








