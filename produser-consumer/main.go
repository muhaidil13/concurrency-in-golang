package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/fatih/color"
)

const NumberOfPizzas = 10

var (
	pizzaMade, pissaFailed, total int
)

type Producer struct {
	Data chan PizzaOrder
	Quit chan chan error
}

type PizzaOrder struct {
	pizzaNumber int
	message     string
	success     bool
}

func (p *Producer) Close() error {
	ch := make(chan error)
	p.Quit <- ch
	return nil
}

func makePizza(pizzanumber int) *PizzaOrder {
	pizzanumber++
	if pizzanumber <= NumberOfPizzas {
		delay := rand.Intn(5) + 1
		fmt.Printf("Received number order %d\n", delay)

		rnd := rand.Intn(12) + 1
		msg := ""
		success := false

		if rnd < 5 {
			pissaFailed++
		} else {
			pizzaMade++
		}
		total++

		fmt.Printf("Making Pizza %d it will take %d seccond....\n", pizzanumber, delay)

		// delay for bit
		time.Sleep(time.Duration(delay) * time.Second)

		if rnd <= 2 {
			msg = fmt.Sprintf("*** We ran out of ingredients for pizza %d", pizzanumber)
		} else if rnd <= 4 {
			msg = fmt.Sprintf("*** The cook quit while making  pizza %d", pizzanumber)
		} else {
			success = true
			msg = fmt.Sprintf("*** Pizza order %d is ready", pizzanumber)
		}

		p := PizzaOrder{
			pizzaNumber: pizzanumber,
			message:     msg,
			success:     success,
		}
		return &p
	}

	return &PizzaOrder{
		pizzaNumber: pizzanumber,
	}
}

func pizzaria(pizzaMaker *Producer) {
	// keep track of which we are making
	var i = 0
	// run forever or until we received quit notification

	// try to make pizzas

	for {
		// try to make pizza
		currentpizza := makePizza(i)
		if currentpizza != nil {
			i = currentpizza.pizzaNumber
			select {
			// Send Data using Channel
			case pizzaMaker.Data <- *currentpizza:
			case quitChan := <-pizzaMaker.Quit:
				close(pizzaMaker.Data)
				close(quitChan)
				return

			}
		}
	}
}

func main() {

	// Generate random number generator
	rand.Seed(time.Now().UnixNano())

	// Print out a message
	color.Cyan("the pizza is open for bussines")
	color.Cyan("------------------------------")

	// crate a producer
	pizzajob := &Producer{
		Data: make(chan PizzaOrder),
		Quit: make(chan chan error),
	}

	// run the producer in the background

	go pizzaria(pizzajob)

	// create and run costumer

	for i := range pizzajob.Data {
		if i.pizzaNumber <= NumberOfPizzas {
			if i.success {
				color.Green(i.message)
				color.Green("Order #%d is out for delivery ", i.pizzaNumber)
			} else {
				color.Red(i.message)
				color.Red("The costumer mad")
			}
		} else {
			color.Cyan("done making pizzas")
			err := pizzajob.Close()
			if err != nil {
				color.Red("Closing Channels", err)
			}
		}
	}
	// print out the ending message
	color.Cyan("----------------")
	color.Cyan("Done Making pizza for the day")

	color.Cyan("kita membuat %d pizza, tapi gagal membuat %d, dengan total %d diterima.", pizzaMade, pissaFailed, total)

	switch {
	case pissaFailed > 9:
		color.Red("Hari yang buruk")
	case pissaFailed >= 6:
		color.Red("Hari yang tidak baik")
	case pissaFailed >= 4:
		color.Red("Hari yang buruk tapi lumayan")
	case pissaFailed >= 2:
		color.Yellow("hari ini lumayan baik")
	default:
		color.Green("hari ini adalah hari baik")
	}

}
