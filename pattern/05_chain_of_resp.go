package pattern

import "fmt"

/*
Реализовать паттерн «цепочка вызовов».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
https://en.wikipedia.org/wiki/Chain-of-responsibility_pattern
*/

type department interface {
	setNext(department)
	execute(*patient)
}

type patient struct {
	firstName       string
	lastName        bool
	isRegistration  bool
	isDoctorCheckUp bool
	isMedicine      bool
	isPayment       bool
}

type reception struct {
	next department
}

func (r *reception) setNext(d department) {
	r.next = d
}

func (r *reception) execute(p *patient) {
	if p.isRegistration {
		fmt.Println("Patient registration already done")
	} else {
		fmt.Println("Patient registering patient")
		p.isRegistration = true
	}
	r.next.execute(p)
}

type doctor struct {
	next department
}

func (d *doctor) setNext(next department) {
	d.next = next
}

func (d *doctor) execute(p *patient) {
	if p.isDoctorCheckUp {
		fmt.Println("Doctor checkup already done")
	} else {
		fmt.Println("Doctor checking patient")
		p.isDoctorCheckUp = true
	}
	d.next.execute(p)
}

type medical struct {
	next department
}

func (m *medical) setNext(d department) {
	m.next = d
}

func (m *medical) execute(p *patient) {
	if p.isMedicine {
		fmt.Println("Medicine already given to patient")
	} else {
		fmt.Println("Medical giving medicine to patient")
		p.isMedicine = true
	}
	m.next.execute(p)
}

type cashier struct {
	next department
}

func (c *cashier) setNext(d department) {
	c.next = d
}

func (c *cashier) execute(p *patient) {
	if p.isPayment {
		fmt.Println("Payment Done")
	}
	fmt.Println("Cashier getting money from patient")
}

/*
Цепочка обязанностей — это поведенческий паттерн проектирования,
который позволяет передавать запросы последовательно по цепочке обработчиков.
Каждый последующий обработчик решает,
может ли он обработать запрос сам и стоит ли передавать запрос дальше по цепи.


Преимущества
- Уменьшает зависимость между клиентом и обработчиками.
- Реализует принцип единственной обязанности.
- Реализует принцип открытости/закрытости.

Недостатки
- Запрос может остаться никем не обработанным.
*/
