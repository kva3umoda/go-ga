package main

import "fmt"

// Вещь
type Item struct {
	Name   string  // наименование
	Weight float64 // вес
	Value  float64 // ценность
}

var (
	items = []Item{
		{"map", 9, 150},
		{"compass", 13, 35},
		{"water", 153, 200},
		{"sandwich", 50, 160},
		{"glucose", 15, 60},
		{"tin", 68, 45},
		{"banana", 27, 60},
		{"apple", 39, 40},
		{"cheese", 23, 30},
		{"beer", 52, 10},
		{"suntan cream", 11, 70},
		{"camera", 32, 30},
		{"t-shirt", 24, 15},
		{"trousers", 48, 10},
		{"umbrella", 73, 40},
		{"waterproof trousers", 42, 70},
		{"waterproof overclothes", 43, 75},
		{"note-case", 22, 80},
		{"sunglasses", 7, 20},
		{"towel", 18, 12},
		{"socks", 4, 50},
		{"book", 30, 10},
	}

	maxCapacity = 400.0 // Максимальнай допустимый вес
)

// Расчет веса 
func calcValue(genome []float64) float64 {
	totalWeight := 0.0
	totalValue := 0.0
	for i := range genome {
		item := items[i]
		if totalWeight+item.Weight <= maxCapacity {
			totalWeight += genome[i] * item.Weight
			totalValue += genome[i] * item.Value
		}
	}
	return totalValue
}

func printItems(genome []float64) {
	totalWeight := 0.0
	totalValue := 0.0
	for i := range genome {
		item := items[i]
		if totalWeight+item.Weight <= maxCapacity {
			totalWeight += genome[i] * item.Weight
			totalValue += genome[i] * item.Value
			fmt.Printf("- Adding %s: weight = %v, value = %v, accumulated weight = %v, accumulated value = %v\n",
				item.Name, item.Weight, item.Value, totalWeight, totalValue)
		}
	}

	fmt.Printf("- Total weight = %v, Total value = %v\n", totalWeight, totalValue)
}
