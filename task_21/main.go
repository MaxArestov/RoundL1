// Реализовать паттерн «адаптер» на любом примере.

package main

import (
	"fmt"
)

// Gun - интерфейс для оружия
type Gun interface {
	Shoot()
}

// Pistol - структура, реализующая интерфейс Gun
type Pistol struct{}

func (p *Pistol) Shoot() {
	fmt.Println("Pistol shot!")
}

// Rifle - структура, реализующая интерфейс Gun
type Rifle struct{}

func (r *Rifle) Shoot() {
	fmt.Println("Rifle shot!")
}

// GunAdapter - адаптер для оружия
type GunAdapter struct {
	gun Gun
}

func (ga *GunAdapter) Shoot() {
	ga.gun.Shoot()
}

func main() {
	pistol := &Pistol{}
	rifle := &Rifle{}

	pistolAdapter := &GunAdapter{gun: pistol}
	rifleAdapter := &GunAdapter{gun: rifle}

	// Тестирование
	pistolAdapter.Shoot()
	rifleAdapter.Shoot()
}
