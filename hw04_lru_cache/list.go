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
	return l.last
}

func (l *list) MoveToFront(i *ListItem) {}

func (l *list) Remove(i *ListItem) {}

func NewList() List {
	return new(list)
}
