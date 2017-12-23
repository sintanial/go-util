package intutil

import (
	"strconv"
	"math/rand"
	"strings"
)

func ParseInt(s string, def ...int) int {
	val, err := strconv.Atoi(s)
	if err != nil && len(def) > 0 {
		return def[0]
	}

	return val
}

func Default(i int, def int) int {
	if i == 0 {
		return def
	}

	return i
}

func Contains(search int, list []int) bool {
	for i := 0; i < len(list); i++ {
		if search == list[i] {
			return true
		}
	}

	return false
}

func Shuffle(list []int) {
	for i := range list {
		j := rand.Intn(i + 1)
		list[i], list[j] = list[j], list[i]
	}
}

func IndexOf(search int, list []int) int {
	for i := 0; i < len(list); i++ {
		if search == list[i] {
			return i
		}
	}

	return -1
}

func LeadingZero(i int) string {
	res := strconv.Itoa(i)
	if i < 10 {
		res = "0" + res
	}

	return res
}

func RandRange(min int, max int) int {
	if max-min <= 0 {
		return 0
	}

	return min + rand.Intn(max-min)
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}

	if x == 0 {
		return 0
	}
	return x
}

func Btoi(b bool) int {
	if b {
		return 1
	}

	return 0
}

type Range struct {
	Min int
	Max int
}

func ParseRange(s string, mindef, maxdef int) Range {
	var err error

	ss := strings.Split(s, "-")
	min, max := mindef, maxdef

	if len(ss) > 0 {
		min, err = strconv.Atoi(ss[0])
		if err != nil {
			min = mindef
		}
	}
	if len(ss) == 1 {
		max = min
	}
	if len(ss) > 1 {
		max, err = strconv.Atoi(ss[1])
		if err != nil {
			max = maxdef
		}
	}

	return Range{min, max}
}

func (self Range) Rand() int {
	return RandRange(self.Min, self.Max)
}

func (self Range) In(i int) bool {
	return i >= self.Min && i <= self.Max
}

func (self Range) RandSub(i int) int {
	if self.In(i) {
		return 0
	}

	diff := self.Rand() - i
	if diff > 0 {
		return diff
	}

	return 0
}
