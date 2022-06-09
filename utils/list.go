package utils

import (
	"errors"
	"fmt"
	"reflect"
	"sync"
)

type Node struct {
	prev    *Node
	element interface{}
	next    *Node
}

func NewNode(element interface{}) *Node {
	return &Node{
		element: element,
	}
}

func (o *Node) SetElement(element interface{}) {
	o.element = element
}

func (o *Node) GetElement() interface{} {
	return o.element
}

func (o *Node) SetPrev(prev *Node) {
	o.prev = prev
}

func (o *Node) GetPrev() *Node {
	return o.prev
}

func (o *Node) SetNext(next *Node) {
	o.next = next
}

func (o *Node) GetNext() *Node {
	return o.next
}

func (o *Node) Ptr() string {
	return reflect.ValueOf(o).String()
}

type LinkedList struct {
	first  *Node
	last   *Node
	size   int
	rwLock *sync.RWMutex
}

func NewLinkedList() *LinkedList {
	return &LinkedList{
		rwLock: &sync.RWMutex{},
	}
}

func (o *LinkedList) Add(element interface{}) error {
	return o.AddLast(element)
}

func (o *LinkedList) AddTo(index int, element interface{}) error {
	o.rwLock.Lock()
	defer o.rwLock.Unlock()
	if size := o.Size(); index < 0 || index >= size {
		return errors.New(fmt.Sprintf("index %d out of range.", index))
	}
	node := NewNode(element)
	i := 0
	fn := o.first
	for {
		if i == index-1 {
			break
		}
		i += 1
		fn = fn.next
	}
	next := fn.next
	fn.next = node
	node.next = next
	if next != nil {
		next.prev = node
	}
	o.size += 1
	return nil
}

func (o *LinkedList) AddFirst(element interface{}) error {
	node := NewNode(element)
	o.rwLock.Lock()
	defer o.rwLock.Unlock()
	if o.first == nil || o.last == nil {
		o.first = node
		o.last = node
	} else {
		first := o.first
		first.SetPrev(node)
		node.SetNext(first)
		o.first = node
	}
	o.size += 1
	return nil
}

func (o *LinkedList) AddLast(element interface{}) error {
	node := NewNode(element)
	o.rwLock.Lock()
	defer o.rwLock.Unlock()
	if o.first == nil || o.last == nil {
		o.first = node
		o.last = node
	} else {
		last := o.last
		last.SetNext(node)
		node.SetPrev(last)
		o.last = node
	}
	o.size += 1
	return nil
}

func (o *LinkedList) Clear() {
	o.rwLock.Lock()
	defer o.rwLock.Unlock()
	o.first = nil
	o.last = nil
	o.size = 0
}

func (o *LinkedList) Copy() *LinkedList {
	o.rwLock.RLock()
	defer o.rwLock.RUnlock()
	list := NewLinkedList()
	f := o.first
	for {
		if f == nil {
			break
		}
		_ = list.AddLast(f.element)
		f = f.next
	}
	return list
}

func (o *LinkedList) Get(index int) (interface{}, error) {
	o.rwLock.RLock()
	defer o.rwLock.RUnlock()
	if size := o.Size(); index < 0 || index >= size {
		return nil, errors.New(fmt.Sprintf("index %d out of range.", index))
	}
	i, n := 0, o.first
	for {
		if i == index {
			break
		}
		i += 1
		n = n.next
	}
	element := n.element
	return element, nil
}

func (o *LinkedList) GetFirst() (interface{}, error) {
	o.rwLock.RLock()
	defer o.rwLock.RUnlock()
	if o.first == nil {
		return nil, errors.New("list is empty")
	}
	element := o.first.element
	return element, nil
}

func (o *LinkedList) GetLast() (interface{}, error) {
	o.rwLock.RLock()
	defer o.rwLock.RUnlock()
	if o.last == nil {
		return nil, errors.New("list is empty")
	}
	element := o.last.element
	return element, nil
}

func (o *LinkedList) Pop() (interface{}, error) {
	return o.RemoveLast()
}

func (o *LinkedList) Push(element interface{}) error {
	return o.AddLast(element)
}

func (o *LinkedList) Remove(index int) (interface{}, error) {
	o.rwLock.Lock()
	defer o.rwLock.Unlock()
	if size := o.Size(); index < 0 || index >= size {
		return nil, errors.New(fmt.Sprintf("index %d out of range.", index))
	}
	i := 0
	fn := o.first
	for {
		if i == index {
			break
		}
		i += 1
		fn = fn.next
	}
	prev := fn.prev
	next := fn.next
	if prev == nil {
		o.first = next
	} else {
		prev.next = next
	}
	if next == nil {
		o.last = prev
	} else {
		next.prev = prev
	}
	o.size -= 1
	element := fn.element
	return element, nil
}

func (o *LinkedList) RemoveFirst() (interface{}, error) {
	o.rwLock.Lock()
	defer o.rwLock.Unlock()
	if o.first == nil {
		return nil, errors.New("list is empty")
	}
	f := o.first
	n := f.next
	if n != nil {
		n.prev = nil
	} else {
		o.last = nil
	}
	o.first = n
	o.size -= 1
	return f.element, nil
}

func (o *LinkedList) RemoveLast() (interface{}, error) {
	o.rwLock.Lock()
	defer o.rwLock.Unlock()
	if o.last == nil {
		return nil, errors.New("list is empty")
	}
	l := o.last
	p := l.prev
	if p != nil {
		p.next = nil
	} else {
		o.first = nil
	}
	o.last = p
	o.size -= 1
	return l.element, nil
}

func (o *LinkedList) Size() int {
	return o.size
}

func (o *LinkedList) IsEmpty() bool {
	return o.Size() == 0
}
