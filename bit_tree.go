package bitree

import (
	"sync"

	"github.com/gosundy/bitmap"
)

type BitNode struct {
	mm     *sync.Map
	bitmap *bitmap.BitMap
}

func NewBitTree() *BitNode {
	return &BitNode{mm: &sync.Map{}}
}
func (node *BitNode) Set(data uint32) error {
	root := node
	//find second tier position
	idx0 := data >> 24
	idx0Node := &BitNode{mm: &sync.Map{}}
	_idx0Node, _ := root.mm.LoadOrStore(idx0, idx0Node)
	// second tier
	idx1 := data << 8 >> 24
	idx1Node := &BitNode{bitmap: bitmap.NewBitMap(256 * 256)}
	idx0Node = _idx0Node.(*BitNode)
	_idx1Node, _ := idx0Node.mm.LoadOrStore(idx1, idx1Node)
	idx1Node = _idx1Node.(*BitNode)
	// third and fourth tier
	idx2 := data << 16 >> 16
	err := idx1Node.bitmap.Set(int64(idx2))
	if err != nil {
		return err
	}
	return nil
}
func (node *BitNode) Get(data uint32) (bool, error) {
	root := node
	//find second tier position
	idx0 := data >> 24
	_idx0Node, isExists := root.mm.Load(idx0)
	if !isExists {
		return false, nil
	}
	idx0Node := _idx0Node.(*BitNode)
	// second tier
	idx1 := data << 8 >> 24
	_idx1Node, isExists := idx0Node.mm.Load(idx1)
	if !isExists {
		return false, nil
	}
	idx1Node := _idx1Node.(*BitNode)
	// third and fourth tier
	idx2 := data << 16 >> 16
	if idx1Node.bitmap == nil {
		return false, nil
	}
	isExists, err := idx1Node.bitmap.Get(int64(idx2))
	if err != nil {
		return false, err
	}
	if !isExists {
		return false, nil
	}
	return true, nil

}
