package main

import "time"

func main() {
	var dilbert Employee
	dilbert.ID = 4
	dilbert.Position = "Engineer"

	position := &dilbert.Position
	*position = "Senior " + *position

	var EmployeeOfTheMonth *Employee = &dilbert
	EmployeeOfTheMonth.Position += " (proactive team player)"
}

type Employee struct {
	ID        int
	Name      string
	Address   string
	DoB       time.Time
	Position  string
	Salary    int
	ManagerID int
}
