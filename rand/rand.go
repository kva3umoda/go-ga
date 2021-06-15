package rand

import (
	"math/rand"
	"sync"
	"time"
)

var globalRand = NewRandom()

type Random struct {
	rnd  *rand.Rand
	lock sync.Locker
}

func NewRandom() *Random {
	return &Random{
		rnd: rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

func (r *Random) Float() float64 {
	r.lock.Lock()
	v := r.rnd.Float64()
	r.lock.Unlock()

	return v
}

func (r *Random) Int(n int) int {
	r.lock.Lock()
	v := r.rnd.Intn(n)
	r.lock.Unlock()

	return v
}

func (r *Random) Shuffle(size int, swapFunc func(i, j int)) {
	r.lock.Lock()
	r.rnd.Shuffle(size, swapFunc)
	r.lock.Unlock()
}

func Float() float64 {
	return globalRand.Float()
}

func Int(n int) int {
	return globalRand.Int(n)
}

func Shuffle(size int, swapFunc func(i, j int)) {
	globalRand.Shuffle(size, swapFunc)
}
