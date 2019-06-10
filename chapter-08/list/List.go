package list

import "fmt"

type List struct {
	Value int
	Next *List
}

type Node = List

func (this *List) Last() *List {
	for this.Next != nil {
		this = this.Next
	}

	return this
}

func (this *List) Print() {
	fmt.Print("[")
	if this != nil {
		fmt.Print(this.Value)
		this = this.Next
	}

	for ; this != nil; this = this.Next {
		fmt.Print(", ", this.Value)
	}
	fmt.Print("]")
}

func (this *List) CyclicNode() *Node {
	if this == nil || this.Next == nil || this.Next == nil {
		return nil
	}

	// Slow pointer moves one step and fast pointer two a time.
	slow, fast := this.Next, this.Next.Next

	// Keep moving till the last node is hit or two pointers point to the
	// same node.
	for slow != nil && fast != nil && slow != fast {
		slow = slow.Next

		if fast.Next == nil {
			break
		}

		fast = fast.Next.Next
	}

	if slow != fast {
		return nil
	}

	// Start traversing from the head and meet points together.
	// The confluence of the two pointers is the beginning of the cycle.
	for this != slow {
		this = this.Next
		slow = slow.Next
	}

	return slow
}

func InitList(values ...int) *List {
	var head List

	current := &head
	for _, v := range values {
		current.Next = new(List)
		current.Next.Value = v
		current = current.Next
	}

	return head.Next
}

func OverlappingNode(list1, list2 *List) *Node {
	length1, length2 := 0, 0

	// Find length of two the lists.
	for head := list1; head != nil; head = head.Next {
		length1++
	}

	for head := list2; head != nil; head = head.Next {
		length2++
	}

	if length1 > length2 {
		list1, list2 = list2, list1
		length1, length2 = length2, length1
	}

	// Variable list2 is the longer list.

	for i := 1; i <= length2-length1; i++ {
		list2 = list2.Next
	}

	// Both list1 and list2 from here on are of the same length.
	// Traverse both of them together at the same pace to see where
	// they overlap, if at all.
	for list1 != nil && list1 != list2 {
		list1, list2 = list1.Next, list2.Next
	}

	// If the variable does not point to nil, then the loop broke when
	// the overlap was found.
	return list1
}
