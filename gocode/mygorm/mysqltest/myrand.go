package mysqltest

import (
	"math/rand"
	"time"
)

type MyRand struct {
}

func (s *MyRand) RandInt(min, max int64) int64 {
	rand.Seed(time.Now().UnixNano())

	return min + rand.New(rand.NewSource(time.Now().UnixNano())).Int63n(max-min)
}

func (s *MyRand) GetCn(size int) string {

	a := make([]rune, size)
	for i := range a {
		a[i] = rune(s.RandInt(19968, 40869))
	}
	return string(a)
}
