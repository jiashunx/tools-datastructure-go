package utils

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func getNodeString(o *Node) string {
	return strings.Join([]string{
		fmt.Sprintf("%p", o.prev),
		fmt.Sprintf("%p", o),
		fmt.Sprintf("%p", o.next),
	}, ", ")
}

func getLinkedListString(list *LinkedList) string {
	strArr := []string{}
	node := list.first
	for {
		if node == nil {
			break
		}
		strArr = append(strArr, getNodeString(node))
		node = node.next
	}
	return strings.Join(strArr, "\n")
}

func TestLinkedList(t *testing.T) {
	ast := assert.New(t)
	list := NewLinkedList()
	s := "dd"
	s_ptr := &s
	_ = list.Add(s_ptr)
	ast.Equal(1, list.Size())
	fmt.Println("1 element\n" + getLinkedListString(list))

	m := make(map[string]string)
	m_ptr := &m
	_ = list.Add(m_ptr)
	ast.Equal(2, list.Size())
	fmt.Println("2 element\n" + getLinkedListString(list))

	i := 10
	i_ptr := &i
	_ = list.AddFirst(i_ptr)
	ast.Equal(3, list.Size())
	fmt.Println("3 element\n" + getLinkedListString(list))

	b := true
	b_ptr := &b
	_ = list.AddTo(2, b_ptr)
	ast.Equal(4, list.Size())
	fmt.Println("4 element\n" + getLinkedListString(list))

	// i_ptr, s_ptr, b_ptr, m_ptr
	first, _ := list.GetFirst()
	first_i_ptr := (first).(*int)
	ast.Equal(10, *first_i_ptr)

	second, _ := list.Get(1)
	second_s_ptr := (second).(*string)
	ast.Equal("dd", *second_s_ptr)

	third, _ := list.Get(2)
	third_b_ptr := (third).(*bool)
	ast.Equal(true, *third_b_ptr)

	fourth, _ := list.GetLast()
	fourth_m_ptr := (fourth).(*map[string]string)
	ast.Equal(m, *fourth_m_ptr)

	list_copy := list.Copy()
	ast.Equal(4, list_copy.Size())

	fmt.Println("list_copy\n" + getLinkedListString(list_copy))
	_, _ = list_copy.Remove(0)
	fmt.Println("remove 1 element\n" + getLinkedListString(list_copy))
	_, _ = list_copy.RemoveFirst()
	fmt.Println("remove 2 element\n" + getLinkedListString(list_copy))
	_, _ = list_copy.RemoveLast()
	fmt.Println("remove 3 element\n" + getLinkedListString(list_copy))
	ast.Equal(1, list_copy.Size())

	list.Clear()
	ast.True(list.IsEmpty())
	ast.False(list_copy.IsEmpty())
}
