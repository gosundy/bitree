package bitree

import "testing"

const (
	count = 1000000
)

func TestBitNode_Set(t *testing.T) {
	bitmap := NewBitTree()
	i := uint32(0)
	for i = 0; i < count; i++ {
		if i%2 == 0 {
			err := bitmap.Set(i)
			if err != nil {
				t.Fatal(err)
			}
		}
	}
	for i = 0; i < count; i++ {
		isExists, err := bitmap.Get(i)
		if err != nil {
			t.Fatal(err)
		}
		if i%2 == 0 {
			if isExists == false {
				t.Fatalf("expect exists, acutal not exists")
			}
		} else {
			if isExists == true {
				t.Fatalf("expect not exists, acutal exists")
			}
		}
	}

}
