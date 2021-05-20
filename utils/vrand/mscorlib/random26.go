package mscorlib

import (
	"math"
)

const (
	MBIG  int = 2147483647
	MSEED int = 161803398
	MZ    int = 0
)

type Random26 struct {
	seed      int
	seedArray [56]int

	inext  int
	inextp int
}

func NewRandom26(seed int) *Random26 {
	var rand_obj *Random26 = &Random26{
		seed:      seed,
		seedArray: [56]int{},
		inext:     0,
		inextp:    31,
	}
	var ii int = 0
	var mj int = MSEED - int(math.Abs(float64(seed)))
	var mk int = 1
	rand_obj.seedArray[55] = mj
	for i := 1; i < 56; i++ {
		ii = (21 * i) % 55
		(*rand_obj).seedArray[ii] = mk
		mk = mj - mk
		if mk < 0 {
			mk += MBIG
		}
		mj = (*rand_obj).seedArray[ii]
	}
	for k := 1; k < 5; k++ {
		for i := 1; i < 56; i++ {
			(*rand_obj).seedArray[i] -= (*rand_obj).seedArray[1+(i+30)%55]
			if (*rand_obj).seedArray[i] < 0 {
				(*rand_obj).seedArray[i] += MBIG
			}
		}
	}
	return rand_obj
}

func (r *Random26) Sample() float32 {
	if r.inext+1 >= 56 {
		r.inext = 1
	} else {
		r.inext += 1
	}
	if r.inextp+1 >= 56 {
		r.inextp = 1
	} else {
		r.inextp += 1
	}
	var retVal int = r.seedArray[r.inext] - r.seedArray[r.inextp]
	if retVal < 0 {
		retVal += MBIG
	}
	r.seedArray[r.inext] = retVal
	return float32(retVal) * 1.0 / float32(MBIG)
}
