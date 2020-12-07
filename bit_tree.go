package bitree

import (
	"sync"

	"github.com/gosundy/bitmap"
)

type BitNode struct {
	mm     map[byte]*BitNode
	bitmap *bitmap.BitMap
	mu     sync.RWMutex
}

func NewBitTree() *BitNode {
	return &BitNode{mm: make(map[byte]*BitNode)}
}
func (node *BitNode) Set(data uint32) error {
	root := node
	//find second tier position
	idx0 := byte(data >> 24)
	var idx0Node *BitNode
	if len(node.mm) == 256 {
		idx0Node = root.mm[idx0]
	} else {
		root.mu.RLock()
		idx0Node = root.mm[idx0]
		root.mu.RUnlock()
		if idx0Node == nil {
			root.mu.Lock()
			idx0Node = root.mm[idx0]
			if idx0Node == nil {
				idx0Node = &BitNode{mm: make(map[byte]*BitNode)}
				root.mm[idx0] = idx0Node
			}
			root.mu.Unlock()
		}
	}

	// second tier
	idx1 := byte(data << 8 >> 24)
	var idx1Node *BitNode
	if len(idx0Node.mm) == 256 {
		idx1Node = idx0Node.mm[idx1]
	} else {
		idx0Node.mu.RLock()
		idx1Node = idx0Node.mm[idx1]
		idx0Node.mu.RUnlock()
		if idx1Node == nil {
			idx0Node.mu.Lock()
			idx1Node = idx0Node.mm[idx1]
			if idx1Node == nil {
				idx1Node = &BitNode{bitmap: bitmap.NewBitMap(256 * 256)}
				idx0Node.mm[idx1] = idx1Node
			}
			idx0Node.mu.Unlock()
		}
	}

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
	idx0 := byte(data >> 24)
	root.mu.RLock()
	idx0Node, isExists := root.mm[idx0]
	root.mu.RUnlock()
	if !isExists {
		return false, nil
	}
	idx0Node.mu.RLock()
	// second tier
	idx1 := byte(data << 8 >> 24)
	idx1Node, isExists := idx0Node.mm[idx1]
	if !isExists {
		return false, nil
	}
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
func (node *BitNode) Reset(data uint32) (bool, error) {
	root := node
	//find second tier position
	idx0 := byte(data >> 24)
	root.mu.RLock()
	idx0Node, isExists := root.mm[idx0]
	root.mu.RUnlock()
	if !isExists {
		return false, nil
	}
	idx0Node.mu.RLock()
	// second tier
	idx1 := byte(data << 8 >> 24)
	idx1Node, isExists := idx0Node.mm[idx1]
	if !isExists {
		return false, nil
	}
	// third and fourth tier
	idx2 := data << 16 >> 16
	if idx1Node.bitmap == nil {
		return false, nil
	}
	idx1Node.mu.Lock()
	_, err := idx1Node.bitmap.ResetN(int64(idx2))
	if err != nil {
		return false, err
	}
	if !isExists {
		return false, nil
	}
	return true, nil

}
