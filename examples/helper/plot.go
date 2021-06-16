package helper

import (
	"github.com/kva3umoda/go-ga/stat"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

func PlotFitness(fname string, stat *stat.Fitness) {
	p := plot.New()

	p.Title.Text = "Plotutil example"
	p.X.Label.Text = "fitness"
	p.Y.Label.Text = "generation"
	// Draw a grid behind the data
	p.Add(plotter.NewGrid())

	err := plotutil.AddLinePoints(p,
		"MAX", points(stat.Max()),
		"AVG", points(stat.Mean()),
		"MIN", points(stat.Min()),
	)
	if err != nil {
		panic(err)
	}

	// Save the plot to a PNG file.
	if err := p.Save(10*vg.Inch, 10*vg.Inch, fname); err != nil {
		panic(err)
	}
}

// randomPoints returns some random x, y points.
func points(v []float64) plotter.XYs {
	pts := make(plotter.XYs, len(v))
	for i := range pts {
		pts[i].X = float64(i)
		pts[i].Y = v[i]
	}
	return pts
}

