package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/fatih/color"
)

// variable
var seeatingCapacity = 10
var arival = 100
var cutDuration = 1000 * time.Millisecond
var timeopen = 10 * time.Second

func main() {
	// seed random number generator
	rand.Seed(time.Now().UnixNano())

	// print random message
	color.Yellow("Sleeping Barber Problem")
	color.Yellow("=======================")

	// create channel if we need any

	clientChan := make(chan string, seeatingCapacity)
	doneChan := make(chan bool)
	// create barbershop
	shoop := BarberShop{
		ShopCapacity:    seeatingCapacity,
		HairCutDuration: cutDuration,
		NumberOfBarber:  0,
		ClientChan:      clientChan,
		BarberDoneChan:  doneChan,
		Open:            true,
	}
	color.Green("The Shoop Open for the day")

	// add barber
	shoop.AddBarber("aidil")
	shoop.AddBarber("farid")
	shoop.AddBarber("alif")
	shoop.AddBarber("zulk")
	shoop.AddBarber("kaka")
	shoop.AddBarber("fikar")
	shoop.AddBarber("maksi")

	// start barber as goroutine
	shopClosing := make(chan bool)
	Closed := make(chan bool)
	go func() {
		<-time.After(timeopen)
		shopClosing <- true
		shoop.CloseBarberForToday()
		Closed <- true
	}()

	// add client
	i := 1
	go func() {
		for {
			randomMillisecond := rand.Int() % (2 * arival)
			select {
			case <-shopClosing:
				return
			case <-time.After(time.Millisecond * time.Duration(randomMillisecond)):
				shoop.addClient(fmt.Sprintf("Client #%d", i))
				i++
			}
		}
	}()
	// block until barber is closed
	<-Closed
}
