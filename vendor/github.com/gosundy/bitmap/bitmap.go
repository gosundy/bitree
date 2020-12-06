package bitmap

import (
	"errors"
)

type BitMap struct {
	bits     []byte
	cap      int64
	oneCount int64
}

var OverflowErr = errors.New("value is more than bitmap cap")

func NewBitMap(cap int64) *BitMap {
	bm := &BitMap{}
	bm.bits = make([]byte, cap/8+1)
	bm.cap = cap
	return bm
}

// set bits x position as 1
func (bm *BitMap) Set(x int64) error {
	_, err := bm.SetN(x)
	return err
}

// set bits x position as 1, if already sed, return false.
func (bm *BitMap) SetN(x int64) (bool, error) {
	if x > bm.cap {
		return false, OverflowErr
	}
	div := x / 8
	mod := byte(x % 8)
	flag := byte(1 << mod)
	v := bm.bits[div]
	v &= flag
	if v > 0 {
		return false, nil
	}
	bm.bits[div] |= flag
	bm.oneCount++
	return true, nil
}

// set bits x position as 0, if had not sed before, return false.
func (bm *BitMap) ResetN(x int64) (bool, error) {
	if x > bm.cap {
		return false, OverflowErr
	}
	div := x / 8
	mod := byte(x % 8)
	flag := byte(1 << mod)
	v := bm.bits[div]
	v &= flag
	if v == 0 {
		return false, nil
	}
	bm.bits[div] &= ^flag
	bm.oneCount--
	return true, nil
}

//get x position result, if data is one, return true, else return false
func (bm *BitMap) Get(x int64) (bool, error) {
	if x > bm.cap {
		return false, OverflowErr
	}
	div := x / 8
	mod := byte(x % 8)
	flag := byte(1 << mod)
	v := bm.bits[div]
	v &= flag
	if v > 0 {
		return true, nil
	} else {
		return false, nil
	}

}
func (bm *BitMap) Len() int64 {
	return bm.oneCount
}
