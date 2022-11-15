package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

type Income struct {
	Source string
	Amount int
}

func main() {
	// variable for bank balance
	var bankBalance int
	var balace sync.Mutex

	// printout starting values
	fmt.Printf("Initial acount Balace : %d.00", bankBalance)
	fmt.Println()

	//  Define weekly revenue
	incomes := []Income{
		{
			Source: "Main job",
			Amount: 2000,
		},
		{
			Source: "Yotuber job",
			Amount: 3400,
		},
		{
			Source: "Army job",
			Amount: 1900,
		},
		{
			Source: "Inverstor",
			Amount: 2900,
		},
	}

	wg.Add(len(incomes))

	// loop throuh 52 weeks and print out how much is made; keep a running total

	for i, income := range incomes {
		go func(i int, income Income) {
			defer wg.Done()
			for week := 1; week <= 52; week++ {

				balace.Lock()
				temp := bankBalance
				temp += income.Amount
				bankBalance = temp
				balace.Unlock()

				fmt.Printf("On Week %d you Earn	%d Form %s", week, income.Amount, income.Source)
				fmt.Println()
			}
			// value yang dimasukkan disini
		}(i, income)
	}

	wg.Wait()
	// print a final value

	fmt.Printf("Final Balaces :%d.00", bankBalance)
	fmt.Println()
}
