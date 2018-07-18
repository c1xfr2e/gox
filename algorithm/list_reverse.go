package algorithm

import "fmt"

type List struct {
	value int
	next  *List
}

func (L *List) reverse() *List {
	cur := L
	var prev *List = nil
	for cur != nil {
		next := cur.next
		cur.next = prev
		prev, cur = cur, next
	}
	return prev
}

func TestListReverse() {
	L := &List{
		1, &List{2, &List{3, &List{4, nil}}},
	}

	K := L.reverse()
	fmt.Println(K.value, K.next.value, K.next.next.value, K.next.next.next)
}
