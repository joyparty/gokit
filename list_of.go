package gokit

import "container/list"

// ListOf 泛型链表
type ListOf[T any] struct {
	list *list.List
}

// NewListOf 创建一个新的泛型链表
func NewListOf[T any]() *ListOf[T] {
	return &ListOf[T]{list: list.New()}
}

// Back 返回链表的最后一个元素
func (l *ListOf[T]) Back() *ElementOf[T] {
	return newElementOf[T](l.list.Back())
}

// Front 返回链表的第一个元素
func (l *ListOf[T]) Front() *ElementOf[T] {
	return newElementOf[T](l.list.Front())
}

// Init 初始化链表
func (l *ListOf[T]) Init() {
	l.list.Init()
}

// InsertAfter 在mark之后插入一个元素
func (l *ListOf[T]) InsertAfter(v T, mark *ElementOf[T]) *ElementOf[T] {
	return newElementOf[T](l.list.InsertAfter(v, mark.element))
}

// InsertBefore 在mark之前插入一个元素
func (l *ListOf[T]) InsertBefore(v T, mark *ElementOf[T]) *ElementOf[T] {
	return newElementOf[T](l.list.InsertBefore(v, mark.element))
}

// Len 返回链表的长度
func (l *ListOf[T]) Len() int {
	return l.list.Len()
}

// MoveAfter 将元素e移动到mark之后
func (l *ListOf[T]) MoveAfter(e, mark *ElementOf[T]) {
	l.list.MoveAfter(e.element, mark.element)
}

// MoveBefore 将元素e移动到mark之前
func (l *ListOf[T]) MoveBefore(e, mark *ElementOf[T]) {
	l.list.MoveBefore(e.element, mark.element)
}

// MoveToBack 将元素e移动到链表的最后一个元素
func (l *ListOf[T]) MoveToBack(e *ElementOf[T]) {
	l.list.MoveToBack(e.element)
}

// MoveToFront 将元素e移动到链表的第一个元素
func (l *ListOf[T]) MoveToFront(e *ElementOf[T]) {
	l.list.MoveToFront(e.element)
}

// PushBack 将元素v插入到链表的最后一个元素
func (l *ListOf[T]) PushBack(v T) *ElementOf[T] {
	return newElementOf[T](l.list.PushBack(v))
}

// PushBackList 将链表other的元素插入到链表的最后一个元素
func (l *ListOf[T]) PushBackList(other *ListOf[T]) {
	l.list.PushBackList(other.list)
}

// PushFront 将元素v插入到链表的第一个元素
func (l *ListOf[T]) PushFront(v T) *ElementOf[T] {
	return newElementOf[T](l.list.PushFront(v))
}

// PushFrontList 将链表other的元素插入到链表的第一个元素
func (l *ListOf[T]) PushFrontList(other *ListOf[T]) {
	l.list.PushFrontList(other.list)
}

// Remove 删除元素e
func (l *ListOf[T]) Remove(e *ElementOf[T]) T {
	return l.list.Remove(e.element).(T)
}

// ElementOf 泛型链表元素
type ElementOf[T any] struct {
	element *list.Element
}

func newElementOf[T any](element *list.Element) *ElementOf[T] {
	if element == nil {
		return nil
	}
	return &ElementOf[T]{element: element}
}

// Next 下一个元素
func (e *ElementOf[T]) Next() *ElementOf[T] {
	return &ElementOf[T]{element: e.element.Next()}
}

// Prev 上一个元素
func (e *ElementOf[T]) Prev() *ElementOf[T] {
	return &ElementOf[T]{element: e.element.Prev()}
}

// Value 返回元素的值
func (e *ElementOf[T]) Value() T {
	return e.element.Value.(T)
}
