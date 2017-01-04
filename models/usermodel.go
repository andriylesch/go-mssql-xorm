package models

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
)

type Customer struct {
	Id        int    `xorm:"ID PK int autoincr"`
	Nick      string `xorm:"Nick varchar(100) notnull"`
	Email     string `xorm:"Email varchar(100) notnull"`
	Firstname string `xorm:"Firstname varchar(100)"`
	Lastname  string `xorm:"Lastname varchar(100)"`
	Age       int    `xorm:"Age int default 0"`
}

func (customer *Customer) ToString() string {

	var buffer bytes.Buffer
	buffer.WriteString("------- User --------")
	buffer.WriteString("\n Id : " + strconv.Itoa(customer.Id))
	buffer.WriteString("\n Nick : " + customer.Nick)
	buffer.WriteString("\n Email : " + customer.Email)
	buffer.WriteString("\n Firstname : " + customer.Firstname)
	buffer.WriteString("\n Lastname : " + customer.Lastname)
	buffer.WriteString("\n Age : " + strconv.Itoa(customer.Age))
	buffer.WriteString("\n---------------------\n")

	return buffer.String()
}

func (customer *Customer) ToJson() string {
	result, err := json.Marshal(customer)
	if err != nil {
		fmt.Println(err)
		return err.Error()
	}

	return string(result)
}
