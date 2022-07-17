package pattern

import "fmt"

/*
Реализовать паттерн «команда».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
https://en.wikipedia.org/wiki/Command_pattern
*/

type command interface {
	execute()
}

type alarm struct {
	model string
}

func NewAlarm(model string) *alarm {
	return &alarm{model: model}
}

func (a *alarm) off() {
	fmt.Printf("alarm (%s) off\n", a.model)
}

func (a *alarm) on() {
	fmt.Printf("alarm (%s) on\n", a.model)
}

type offCommand struct {
	alarm *alarm
}

func (o *offCommand) execute() {
	o.alarm.off()
}

type onCommand struct {
	alarm *alarm
}

func (o *onCommand) execute() {
	o.alarm.on()
}

/*
Команда — это поведенческий паттерн проектирования,
который превращает запросы в объекты,
позволяя передавать их как аргументы при вызове методов,
ставить запросы в очередь, логировать их,
а также поддерживать отмену операций.


Преимущества
- Убирает прямую зависимость между объектами, вызывающими операции, и объектами, которые их непосредственно выполняют.
- Позволяет реализовать простую отмену и повтор операций.
- Позволяет реализовать отложенный запуск операций.
- Позволяет собирать сложные команды из простых.
- Реализует принцип открытости/закрытости.

Недостатки
- Усложняет код программы из-за введения множества дополнительных классов.
*/
