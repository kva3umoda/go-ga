package main

import (
	"fmt"
	"testing"

	"ga-book/examples/helper"
)

func Test_loadCities(t *testing.T) {
	cities := loadCities(helper.GetCurDir()+"/bayg29.tsp")
	fmt.Println(cities)
	genome := loadOptimalPath(helper.GetCurDir()+"/bayg29.opt.tour")
	fmt.Println(genome)
}
