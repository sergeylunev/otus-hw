package hw04lrucache

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
}

type ListItem struct {
	Value interface{}
	Next  *ListItem
	Prev  *ListItem
}

type list struct {
	first  *ListItem
	last   *ListItem
	length int
}

func (l *list) Len() int {
	return l.length
}

func (l *list) Front() *ListItem {
	return l.first
}

func (l *list) PushFront(v interface{}) *ListItem {
	li := ListItem{Value: v}

	if l.Len() == 0 {
		l.first = &li
		l.last = &li
	} else {
		li.Next = l.first
		l.first.Prev = &li

		l.first = &li
	}

	l.length++

	return l.first
}

func (l *list) Back() *ListItem {
	return l.last
}

func (l *list) PushBack(v interface{}) *ListItem {
	li := ListItem{Value: v}

	if l.Len() == 0 {
		l.last = &li
		l.first = &li
	} else {
		li.Prev = l.last
		l.last.Next = &li

		l.last = &li
	}

	l.length++

	return l.last
}

func (l *list) MoveToFront(i *ListItem) {
	l.PushFront(i.Value)
	l.Remove(i)

	i = nil
}

func (l *list) Remove(i *ListItem) {
	if l.Len() == 1 {
		l.first = nil
		l.last = nil
		l.length = 0

		return
	}

	if i.Prev != nil {
		i.Prev.Next = i.Next
	}

	if i.Next != nil {
		i.Next.Prev = i.Prev
	}

	l.length--
}

func NewList() List {
	return new(list)
}
