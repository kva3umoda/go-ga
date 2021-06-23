package main

import (
	"bufio"
	"log"
	"math"
	"os"
	"strconv"
	"strings"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

type City struct {
	Name string
	X    float64
	Y    float64
}

type MapCities struct {
	cities []City
}

func NewMapCities(fname string) *MapCities {
	cities := loadCities(fname)
	return &MapCities{cities: cities}
}

func (m *MapCities) Distance(i, j int) float64 {
	return math.Sqrt(math.Pow(m.cities[i].X-m.cities[j].X, 2.0) + math.Pow(m.cities[i].Y-m.cities[j].Y, 2.0))
}

func (m *MapCities) Plot(fname string, genome []float64) error {
	p := plot.New()

	p.Title.Text = "Plotutil example"
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"
	// Draw a grid behind the data
	//bs, err := plotter.NewPolygon(bubbleData, vg.Points(1), vg.Points(20))
	p.Add(plotter.NewGrid())
	for _, city := range m.cities {
		point := plotter.Values{city.X - 10.0, city.Y - 10.0, city.X + 10.0, city.Y + 10.0}

		err := plotutil.AddBoxPlots(p, 2.0, city.Name, point)
		if err != nil {
			return err
		}
	}

	// Save the plot to a PNG file.
	if err := p.Save(10*vg.Inch, 10*vg.Inch, fname); err != nil {
		panic(err)
	}
	return nil
}

func (m *MapCities) TotalDistance(genome []float64) float64 {
	distance := 0.0
	for i := range genome {
		if i == 0 {
			distance += m.Distance(int(genome[0]), int(genome[len(genome)-1]))
			continue
		}
		distance += m.Distance(int(genome[i-1]), int(genome[i]))
	}
	return distance
}

func loadCities(fname string) []City {
	cities := make([]City, 0)

	file, err := os.Open(fname)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "DISPLAY_DATA_SECTION" || line == "NODE_COORD_SECTION" {
			break
		}
	}

	for scanner.Scan() {
		if scanner.Text() == "EOF" {
			break
		}
		line := scanner.Text()
		strs := strings.Split(line, " ")
		var (
			city City
			num  int
		)
	loop:
		for _, str := range strs {
			if strings.Trim(str, "") == "" {
				continue
			}
			num++
			switch num {
			case 1:
				city.Name = str
			case 2:
				city.X, err = strconv.ParseFloat(str, 64)
				if err != nil {
					log.Fatal(err)
				}
			case 3:
				city.Y, err = strconv.ParseFloat(str, 64)
				if err != nil {
					log.Fatal(err)
				}
				cities = append(cities, city)
				break loop
			}
		}
	}

	return cities
}

func loadOptimalPath(fname string) []float64 {
	genome := make([]float64, 0)

	file, err := os.Open(fname)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "TOUR_SECTION" {
			break
		}
	}

	for scanner.Scan() {
		if scanner.Text() == "EOF" {
			break
		}
		line := scanner.Text()
		line = strings.Trim(line, "")

		f, err := strconv.ParseFloat(line, 64)
		if err != nil {
			log.Fatal(err)
		}
		if f < 0 {
			continue
		}
		genome = append(genome, f-1)
	}
	return genome
}
