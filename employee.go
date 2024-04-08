package main 

import (
	"fmt"
)

type Employee struct {
	id uint32
	firstName, lastName string
	salary int64
}

func NewEmployee(id uint32, firstName, lastName string, salary int64) *Employee {
	return &Employee{id, firstName, lastName, salary}
}

func (s *Employee) GetAnnualSalary() int64 {
	return s.salary * 12
}

func (s *Employee) SetSalary(pecent int64) {
	s.salary = pecent
}

func main() {
	emp := NewEmployee(01, "Peter", "Tan", 1000)
	fmt.Printf(" Id: %v\n First Name: %s\n Last Name: %s\n Salary: %v\n Annual Salary: %v\n", emp.id, emp.firstName, emp.lastName, emp.salary, emp.GetAnnualSalary())
	emp.SetSalary(1500)
	fmt.Printf(" Salary after raise: %v\n -------------------------------- \n", emp.salary)
	fmt.Printf(" Id: %v\n First Name: %s\n Last Name: %s\n Salary: %v\n Annual Salary: %v\n", emp.id, emp.firstName, emp.lastName, emp.salary, emp.GetAnnualSalary())
}