package main

import (
	"time"

	"github.com/fatih/color"
)

type BarberShop struct {
	ShopCapacity    int
	HairCutDuration time.Duration
	NumberOfBarber  int
	BarberDoneChan  chan bool
	ClientChan      chan string
	Open            bool
}

func (shop *BarberShop) AddBarber(barber string) {
	shop.NumberOfBarber++
	go func() {
		for {
			issleeping := false
			if len(shop.ClientChan) == 0 {
				color.Yellow("There is no Client %s take a nap", barber)
				issleeping = true
			}
			client, shopopen := <-shop.ClientChan
			if shopopen {
				if issleeping {
					color.Red("%s wake %s up", client, barber)
					issleeping = false
				}
				// cut hair
				shop.Cuthair(barber, client)
			} else {
				// close barber
				shop.SendBarberToHome(barber)
				return
			}

		}
	}()
}

func (shop *BarberShop) Cuthair(barber, client string) {
	color.Green("%s is cutting hairs %s", barber, client)
	time.Sleep(shop.HairCutDuration)
	color.Green("%s is finish cutting %s hairs ", barber, client)
}
func (shop *BarberShop) SendBarberToHome(barber string) {
	color.Green("%s is go home", barber)
	shop.BarberDoneChan <- true
}

func (shop *BarberShop) CloseBarberForToday() {
	color.Cyan("Closing barber for day")
	close(shop.ClientChan)
	shop.Open = false
	for a := 1; a <= shop.NumberOfBarber; a++ {
		<-shop.BarberDoneChan
	}
	close(shop.BarberDoneChan)
	color.Green("====================================")
	color.Green("Barber is now Close comeback tomorow")
}

func (shop *BarberShop) addClient(client string) {
	color.Green("**** %s Arives", client)

	if shop.Open {
		select {
		case shop.ClientChan <- client:
			color.Green("%s waiting in room", client)
		default:
			color.Yellow("Rooms is full client %s leave", client)
		}
	} else {
		color.Red("The shoop is closed, so %s leaves", client)
	}
}
