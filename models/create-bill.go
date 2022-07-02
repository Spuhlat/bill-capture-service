package models

import (
	"encoding/json"
	"fmt"
	"main/utils"

	"github.com/google/uuid"
)

type User struct {
	Id        uuid.UUID
	FirstName string
	LastName  string
}

type Bill struct {
	Name         string
	Participants []User
}

// temp solution without db
func CreateBill(name string, members []string) {
	file := utils.ReadFromFile("bills.json")
	var bills []Bill
	json.Unmarshal(file, &bills)

	memberList := utils.Map(members, func(el string, i int) User {
		return User{Id: uuid.New(), FirstName: el, LastName: el}
	})

	newBill := Bill{Name: name, Participants: memberList}
	newBills := append(bills, newBill)
	bill, err := json.Marshal(newBills)
	if err != nil {
		fmt.Println("err")
		return
	}
	utils.WriteToFile("bills.json", string(bill))
}
