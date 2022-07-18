package pattern

import "fmt"

/*
Реализовать паттерн «состояние».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
https://en.wikipedia.org/wiki/State_pattern
*/

type state interface {
	getName() string
	freeze(*stateContext)
	heat(*stateContext)
}

type stateContext struct {
	st state
}

func newStateContext() *stateContext {
	return &stateContext{st: newSolidState()}
}

func (s *stateContext) freeze() {
	fmt.Println("Freezing", s.st.getName(), "substance...")
	s.st.freeze(s)
}

func (s *stateContext) heat() {
	fmt.Println("Heating", s.st.getName(), "substance...")
	s.st.heat(s)
}

func (s *stateContext) setState(st state) {
	s.st = st
}

func (s *stateContext) getState() state {
	return s.st
}

type solidState struct {
	name string
}

func newSolidState() *solidState {
	return &solidState{name: "solid"}
}

func (s *solidState) getName() string {
	return s.name
}

func (s *solidState) freeze(st *stateContext) {
	fmt.Println("Nothing happens.")
}

func (s *solidState) heat(st *stateContext) {
	st.setState(newLiquidState())
}

type liquidState struct {
	name string
}

func newLiquidState() *liquidState {
	return &liquidState{name: "liquid"}
}

func (l *liquidState) getName() string {
	return l.name
}

func (l *liquidState) freeze(st *stateContext) {
	st.setState(newSolidState())
}

func (l *liquidState) heat(st *stateContext) {
	st.setState(newGaseousState())
}

type gaseousState struct {
	name string
}

func newGaseousState() *gaseousState {
	return &gaseousState{name: "gaseous"}
}

func (g *gaseousState) getName() string {
	return g.name
}

func (g *gaseousState) freeze(st *stateContext) {
	st.setState(newLiquidState())
}

func (g *gaseousState) heat(*stateContext) {
	fmt.Println("Nothing happens.")
}

/*


 */

/*
Состояние — это поведенческий паттерн проектирования,
который позволяет объектам менять поведение в зависимости от своего состояния.
Извне создаётся впечатление, что изменился класс объекта.


Преимущества
- Избавляет от множества больших условных операторов машины состояний.
- Концентрирует в одном месте код, связанный с определённым состоянием.
- Упрощает код контекста.

Недостатки
- Может неоправданно усложнить код, если состояний мало и они редко меняются.
*/
