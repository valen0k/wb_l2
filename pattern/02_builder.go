package pattern

import "fmt"

/*
Реализовать паттерн «строитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
https://en.wikipedia.org/wiki/Builder_pattern
*/

type IBuilder interface {
	SetFloors(floors uint8)
	SetWalls(walls uint8)
	SetWindows(windows uint8)
	SetDoors(doors uint8)
	SetSquare(square float32)
	SetRooms(rooms uint8)
	Build() House
}

type House struct {
	floors  uint8
	walls   uint8
	windows uint8
	doors   uint8
	square  float32
	rooms   uint8
}

type MyHouse struct {
	floors  uint8
	walls   uint8
	windows uint8
	doors   uint8
	square  float32
	rooms   uint8
}

func (h *MyHouse) SetFloors(floors uint8) {
	h.floors = floors
}

func (h *MyHouse) SetWalls(walls uint8) {
	h.walls = walls
}

func (h *MyHouse) SetWindows(windows uint8) {
	h.windows = windows
}

func (h *MyHouse) SetDoors(doors uint8) {
	h.doors = doors
}

func (h *MyHouse) SetSquare(square float32) {
	h.square = square
}

func (h *MyHouse) SetRooms(rooms uint8) {
	h.rooms = rooms
}

func (h *MyHouse) Build() House {
	return House{
		floors:  h.floors,
		walls:   h.walls,
		windows: h.windows,
		doors:   h.doors,
		square:  h.square,
		rooms:   h.rooms,
	}
}

func main() {
	mHouse := MyHouse{}
	mHouse.SetDoors(2)
	mHouse.SetWindows(2)
	mHouse.SetFloors(1)
	house := mHouse.Build()
	fmt.Println(house)
}

/*
Строитель — это порождающий паттерн проектирования,
который позволяет создавать сложные объекты пошагово.
Строитель даёт возможность использовать один и тот же код строительства
для получения разных представлений объектов.


Преимущества
- Позволяет создавать продукты пошагово.
- Позволяет использовать один и тот же код для создания различных продуктов.
- Изолирует сложный код сборки продукта от его основной бизнес-логики.

Недостатки
- Усложняет код программы из-за введения дополнительных классов.
- Клиент будет привязан к конкретным классам строителей,
так как в интерфейсе директора может не быть метода получения результата.
*/
