package lists

type doublyLinkedList struct {
	length     int
	head, tail *node
}

func (d *doublyLinkedList) prepend(item any) {
	node := &node{data: item}
	if d.length == 0 {
		d.tail = node
	} else {
		d.head.prev = node
		node.next = d.head
	}

	d.head = node
	d.length++
}

func (d *doublyLinkedList) append(item any) {
	node := &node{data: item}
	if d.length == 0 {
		d.head = node
	} else {
		d.tail.next = node
		node.prev = d.tail
	}

	d.tail = node
	d.length++
}

func (d *doublyLinkedList) get(idx int) any {
	if idx >= d.length {
		// throw some error
		return nil
	} else if idx == 0 {
		return d.head.data
	} else if idx == d.length-1 {
		return d.tail.data
	}

	curr := d.head
	for curr != nil && idx > 0 {
		curr = curr.next
		idx--
	}

	return curr.data

}

func (d *doublyLinkedList) insertAt(item any, idx int) {
	if idx >= d.length {
		// throw some error
		return
	} else if idx == 0 {
		d.prepend(item)
	} else if idx == d.length-1 {
		d.append(item)
	}

	curr := d.head
	for curr != nil && idx > 0 {
		curr = curr.next
		idx--
	}

	node := &node{data: item, prev: curr.prev, next: curr}
	node.prev.next = node
	curr.prev = node
	d.length++
}

func (d *doublyLinkedList) remove(idx int) any {
	if idx >= d.length {
		// throw some error
		return nil
	}

	d.length--
	if d.length == 0 {
		result := d.head.data
		d.head = nil
		d.tail = nil
		return result
	}

	if idx == 0 {
		result := d.head.data
		d.head = d.head.next
		return result
	} else if idx == d.length-1 {
		result := d.tail.data
		d.tail = d.tail.prev
		return result
	}

	curr := d.head

	for curr != nil && idx > 0 {
		curr = curr.next
		idx--
	}

	result := curr.data
	curr.prev.next = curr.next
	curr.next.prev = curr.prev
	curr.next = nil
	curr.prev = nil

	return result
}
