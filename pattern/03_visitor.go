package pattern

import "fmt"

/*
Реализовать паттерн «посетитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
https://en.wikipedia.org/wiki/Visitor_pattern
*/

type visitor interface {
	visitForSquare(*square)
	visitForCircle(*circle)
	visitForTriangle(*triangle)
}

type square struct {
	side int
}

func (s *square) getType() string {
	return "Square"
}

func (s *square) accept(v visitor) {
	v.visitForSquare(s)
}

type circle struct {
	radius int
}

func (c *circle) getType() string {
	return "Circle"
}

func (c *circle) accept(v visitor) {
	v.visitForCircle(c)
}

type triangle struct {
	radius int
}

func (t *triangle) getType() string {
	return "Triangle"
}

func (t *triangle) accept(v visitor) {
	v.visitForTriangle(t)
}

type shape interface {
	getType() string
	accept(visitor)
}

type areaCalculator struct {
	area int
}

func (a *areaCalculator) visitForSquare(s *square) {
	fmt.Println("Calculating area for", s.getType())
}

func (a *areaCalculator) visitForCircle(c *circle) {
	fmt.Println("Calculating area for", c.getType())
}

func (a *areaCalculator) visitForTriangle(t *triangle) {
	fmt.Println("Calculating area for", t.getType())
}

/*
Посетитель — это поведенческий паттерн проектирования,
который позволяет добавлять в программу новые операции,
не изменяя классы объектов,
над которыми эти операции могут выполняться.


Преимущества
- Упрощает добавление операций, работающих со сложными структурами объектов.
- Объединяет родственные операции в одном классе.
- Посетитель может накапливать состояние при обходе структуры элементов.

Недостатки
- Паттерн не оправдан, если иерархия элементов часто меняется.
- Может привести к нарушению инкапсуляции элементов.
*/
