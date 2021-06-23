package domain

type Employees []*Employee

type Employee struct {
	id       EmployeeID
	name     Name
	gender   Gender
	birthday Birthday
	salary   Salary
}

func (e Employee) ID() EmployeeID {
	return e.id
}

func (e Employee) Name() Name {
	return e.name
}

func (e Employee) Gender() Gender {
	return e.gender
}

func (e Employee) Birthday() Birthday {
	return e.birthday
}

func (e Employee) Salary() Salary {
	return e.salary
}

func NewEmployee(id EmployeeID, name Name, gender Gender, birthday Birthday, salary Salary) *Employee {
	return &Employee{id: id, name: name, gender: gender, birthday: birthday, salary: salary}
}
