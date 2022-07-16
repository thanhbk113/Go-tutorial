package staff

import "log"

var OverPaidLimit = 2000
var underPaidLimit = 2000

type Employee struct {
	FirstName string
	LastName  string
	Salart    int
	FullTime  bool
}

type Office struct {
	Employees []Employee
}

func (e *Office) All() []Employee {
	return e.Employees
}

func (e *Office) Overpaid() []Employee {
	var overpaid []Employee
	for _, employee := range e.Employees {
		if employee.Salart > OverPaidLimit {
			overpaid = append(overpaid, employee)
		}
	}
	return overpaid
}

func (e *Office) Underpaid() []Employee {
	var underpaid []Employee
	for _, employee := range e.Employees {
		if employee.Salart < underPaidLimit {
			underpaid = append(underpaid, employee)
		}
	}
	return underpaid
}

func (e *Office) NotVisible() {
	log.Println("Hello", "Word")
}
