package collection

type LinkedNode[T any] struct {
    next *LinkedNode[T]
    value T
}

type LinkedList[T any] struct {
    head *LinkedNode[T]
    tail *LinkedNode[T]
}

func (l *LinkedList[T]) Size() uint {
    var size uint = 0
    node := l.head
    for node != nil {
        node = node.next
        size++
    }

    return size
}

func (l *LinkedList[T]) Push(element T) {
    if l.tail == nil {
        l.head = &LinkedNode[T]{value: element}
        l.tail = l.head
    } else {
        next := &LinkedNode[T]{value: element}
        l.tail.next = next
        l.tail = next
    }
}

func (l *LinkedList[T]) Insert(element T, index uint) bool {
    if index == 0 {
        node := &LinkedNode[T]{next: l.head, value: element}
        l.head = node
        if l.tail == nil {
            l.tail = node
        }

        return true
    }

    node := l.head.next 
    for i := uint(0); node != nil; i++ {
        if i == index-1 {
            newNode := &LinkedNode[T]{next: node.next, value: element}
            if l.tail == node {
                l.tail = newNode
            }
            node.next = newNode
            return true
        }
    }
    return false
}

func (l *LinkedList[T]) Get(index uint) *T {
    node := l.head
    for i := uint(0); node != nil; i++ {
        if i == index {
            return &node.value
        }
        node = node.next
    }

    return nil
}

func (l *LinkedList[T]) Pop() *T {
    if l.head == nil {
        return nil
    }

    target := l.head
    if l.tail == l.head {
        l.tail = nil
    }
    l.head = l.head.next
    return &target.value
}

func (l *LinkedList[T]) Remove(index uint) *T {
    if index == 0 {
        return l.Pop()
    }

    node := l.head
    for i := uint(0); node != nil; i++ {
        if i == index-1 {
            target := node.next
            node.next = target.next
            if l.tail == target {
                l.tail = node
            }
            return &target.value
        }
    }
    return nil
}

func (l *LinkedList[T]) RemoveLast() *T {
    if l.tail == nil {
        return nil
    }

    node := l.head
    for node.next != l.tail {
        node = node.next
    }

    target := node.next
    node.next = nil
    l.tail = node

    return &target.value
}

func EmptyList[T any]() *LinkedList[T] {
    return &LinkedList[T]{}
}

func ListFromArray[T any](array []T) *LinkedList[T] {
    list := EmptyList[T]()
    previous := list.head
    for _, element := range array {
        node := &LinkedNode[T]{ value: element }
        if list.head == nil {
            list.head = node
        }
        if list.tail == nil {
            list.tail = node
        }
        if previous != nil {
            previous.next = node
        }
        if list.tail == previous {
            list.tail = node
        }

        previous = node
    }

    return list
}

func ArrayFromList[T any](list *LinkedList[T]) []T {
    array := make([]T, 0)
    node := list.head
    for node != nil {
        array = append(array, node.value)
        node = node.next
    }
    return array
}
