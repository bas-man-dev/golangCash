package main

import "fmt"

func main() {
	var change float64
	for {
		fmt.Println("change owed: ")
		fmt.Scan(&change)
		if change > 0 {
			break
		}
	}

	amount := int(change * 100)

	fmt.Println(countCoins(1, amount))

}

func countCoins(n, amount int) int {
	if amount == 0 {
		return n - 1
	}
	if amount >= 25 {
		amount = amount - 25
		n++
		return countCoins(n, amount)
	}
	if amount >= 10 {
		amount = amount - 10
		n++
		return countCoins(n, amount)
	}
	if amount >= 5 {
		amount = amount - 5
		n++
		return countCoins(n, amount)
	}
	if amount >= 1 {
		amount = amount - 1
		n++
		return countCoins(n, amount)
	}
	return 0
}
