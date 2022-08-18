package pattern

import (
	"errors"
	"log"
)

/*
Реализовать паттерн «фасад».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
https://en.wikipedia.org/wiki/Facade_pattern
*/

type WalletFacade struct {
	wallet   *Wallet
	security *Security
}

func NewWalletFacade() *WalletFacade {
	return &WalletFacade{
		wallet:   NewWallet(),
		security: NewSecurity(),
	}
}

func (f *WalletFacade) Buy(amount float32, code uint) error {
	if f.security.Check(code) {
		err := f.wallet.DebitBalance(amount)
		if err != nil {
			return err
		}
		return nil
	}
	return errors.New("некорректный код")
}

type Security struct {
	code uint
}

func NewSecurity() *Security {
	return &Security{code: 123}
}

func (s *Security) Check(code uint) bool {
	return s.code == code
}

type Wallet struct {
	balance float32
}

func NewWallet() *Wallet {
	return &Wallet{balance: 0}
}

func (w *Wallet) CreditBalance(amount float32) {
	w.balance += amount
	log.Println("баланс пополнен")
}

func (w *Wallet) DebitBalance(amount float32) error {
	if w.balance < amount {
		return errors.New("на балансе недостаточно средств")
	}
	w.balance -= amount
	log.Println("снятие с баланса пополнено")
	return nil
}

func main() {
	facade := NewWalletFacade()
	log.Println(facade.Buy(99, 12))
	log.Println(facade.Buy(99, 123))
}

/*
Фасад — это структурный паттерн проектирования,
который предоставляет простой интерфейс к сложной системе классов,
библиотеке или фреймворку.


Преимущества
- Изолирует клиентов от компонентов сложной подсистемы.

Недостатки
- Фасад рискует стать божественным объектом, привязанным ко всем классам программы.
*/
