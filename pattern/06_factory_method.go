package pattern

import "errors"

/*
Реализовать паттерн «фабричный метод».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
https://en.wikipedia.org/wiki/Factory_method_pattern
*/

type iGun interface {
	setName(name string)
	setPower(power int)
	getName() string
	getPower() int
}

type gun struct {
	name  string
	power int
}

func (g *gun) setName(name string) {
	g.name = name
}

func (g *gun) setPower(power int) {
	g.power = power
}

func (g *gun) getName() string {
	return g.name
}

func (g *gun) getPower() int {
	return g.power
}

type ak47 struct {
	gun
}

func newAK47() iGun {
	return &ak47{
		gun{
			name:  "AK47",
			power: 5,
		},
	}
}

type musket struct {
	gun
}

func newMusket() iGun {
	return &musket{
		gun{
			name:  "musket",
			power: 1,
		},
	}
}

func getGun(name string) (iGun, error) {
	switch name {
	case "AK47":
		return newAK47(), nil
	case "musket":
		return newMusket(), nil
	default:
		return nil, errors.New("Wrong gun type passed")
	}
}

/*
Фабричный метод — это порождающий паттерн проектирования,
который определяет общий интерфейс для создания объектов в суперклассе,
позволяя подклассам изменять тип создаваемых объектов.


Преимущества
- Избавляет класс от привязки к конкретным классам продуктов.
- Выделяет код производства продуктов в одно место, упрощая поддержку кода.
- Упрощает добавление новых продуктов в программу.
- Реализует принцип открытости/закрытости.

Недостатки
- Может привести к созданию больших параллельных иерархий классов,
так как для каждого класса продукта надо создать свой подкласс создателя.
*/
