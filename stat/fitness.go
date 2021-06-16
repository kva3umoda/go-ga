package stat

import "sync"

type Fitness struct {
	max  []float64
	mean struct {
		mean  []float64
		sum   []float64
		count []float64
	}
	min []float64

	lock sync.RWMutex
}

func NewFitness(defaultSize int) *Fitness {
	f := &Fitness{
		max: make([]float64, 0, defaultSize),
		min: make([]float64, 0, defaultSize),
	}
	f.mean.sum = make([]float64, 0, defaultSize)
	f.mean.count = make([]float64, 0, defaultSize)
	f.mean.mean = make([]float64, 0, defaultSize)

	return f
}

func (f *Fitness) Max() []float64 {
	return f.max
}

func (f *Fitness) Mean() []float64 {
	return f.mean.mean
}

func (f *Fitness) Min() []float64 {
	return f.min
}

func (f *Fitness) Add(generation int, value float64) {
	f.lock.Lock()
	defer f.lock.Unlock()

	f.calcMax(generation, value)
	f.calcMin(generation, value)
	f.calcMean(generation, value)
}

func (f *Fitness) LastMax() float64 {
	f.lock.RLock()
	defer f.lock.RUnlock()

	if len(f.max) == 0 {
		return 0.0
	}
	return f.max[len(f.max)-1]
}

func (f *Fitness) calcMax(generation int, value float64) {
	if cap(f.max) < (generation + 1) {
		tmp := make([]float64, len(f.max), generation+1)
		copy(tmp, f.max)
		f.max = tmp
	}

	f.max = f.max[:generation+1]
	f.max[generation] = max(f.max[generation], value)
}

func (f *Fitness) LastMin() float64 {
	f.lock.RLock()
	defer f.lock.RUnlock()

	if len(f.min) == 0 {
		return 0.0
	}
	return f.min[len(f.min)-1]
}

func (f *Fitness) calcMin(generation int, value float64) {
	if cap(f.min) < (generation + 1) {
		tmp := make([]float64, len(f.min), generation+1)
		copy(tmp, f.min)
		f.min = tmp
	}

	if len(f.min) < (generation + 1) {
		f.min = f.min[:generation+1]
		f.min[generation] = value
		return
	}

	f.min[generation] = min(f.min[generation], value)
}

func (f *Fitness) LastMean() float64 {
	f.lock.RLock()
	defer f.lock.RUnlock()

	if len(f.mean.mean) == 0 {
		return 0.0
	}
	return f.mean.mean[len(f.mean.mean)-1]
}

func (f *Fitness) calcMean(generation int, value float64) {
	if cap(f.mean.mean) < (generation + 1) {
		tmp := make([]float64, len(f.mean.mean), generation+1)
		copy(tmp, f.mean.mean)
		f.mean.mean = tmp
	}

	if cap(f.mean.count) < (generation + 1) {
		tmp := make([]float64, len(f.mean.count), generation+1)
		copy(tmp, f.mean.count)
		f.mean.count = tmp
	}

	if cap(f.mean.sum) < (generation + 1) {
		tmp := make([]float64, len(f.mean.sum), generation+1)
		copy(tmp, f.mean.sum)
		f.mean.sum = tmp
	}

	f.mean.sum = f.mean.sum[:generation+1]
	f.mean.count = f.mean.count[:generation+1]
	f.mean.mean = f.mean.mean[:generation+1]

	f.mean.sum[generation] += value
	f.mean.count[generation] += 1.0
	f.mean.mean[generation] = f.mean.sum[generation] / f.mean.count[generation]
}
