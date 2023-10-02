package main
//Program to print salary earned from each source in a month and print final bankBalance
import (
	"fmt"
	"sync"
)

type SourceIncome struct {
	source string
	income int
}

var wg sync.WaitGroup // maintain go routine concurrency
var balance sync.Mutex  // to get  lock on bankBalance 

func main() {

	var bankBalance int = 0
	incomes := []SourceIncome{
		{source: "A",
			income: 5,
		},
		{
			source: "B",
			income: 1,
		},
		{
			source: "C",
			income: 2,
		},
		{
			source: "D",
			income: 3,
		},
	}
	wg.Add(len(incomes))
	for i, value := range incomes {

		fmt.Println(i, value)
		go func(i int, total SourceIncome) {
			defer wg.Done()

			//fmt.Printf("inside go")
			//bankBalance = 0
			for wk := 1; wk <= 4; wk++ {
				balance.Lock()
				temp := bankBalance
				temp = temp + total.income
				bankBalance = temp
				balance.Unlock()
				fmt.Printf("on week %d you earned %d from source %s \n", wk, total.income, total.source)

			}
		}(i, value)

	}
	wg.Wait()
	fmt.Printf("final bank balance %d\n", bankBalance)
	// var bankbalance
	//source income
	// loop through 52 weeks
	//print bank balance
}
