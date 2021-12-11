package list

import "fmt"

type element struct {
	prev *element
	next *element
	value interface{}
}

type LinkedList struct {
	curElem *element
	curIndex int
	len int
}

func NewLinkedList() (LinkedList) {
	return LinkedList{
		curElem: nil,
		curIndex: -1,
		len: 0,
	}
}

func (l *LinkedList) movePrev() bool {
	if l.curElem.prev == nil {
		return false
	}

	l.curElem = l.curElem.prev
	l.curIndex--;

	return true
}

func (l *LinkedList) moveNext() bool {
	if l.curElem.next == nil {
		return false
	}

	l.curElem = l.curElem.next
	l.curIndex++;

	return true
}

func (l *LinkedList) Add(value interface{})  {
	if l.curElem == nil {		
		l.curElem = &element{
			prev: nil,
			next: nil,
			value: value,
		}
		l.curIndex = 0
	}else {
		for l.curIndex < (l.len-1) {
			l.moveNext()
		}

		el := element{
			prev: l.curElem,
			next: nil,
			value: value,
		}

		l.curElem.next = &el
		l.curElem = &el
		l.curIndex++
	}

	l.len++
}

func (l *LinkedList) Get(item int) interface{} {
	if (item < 0 || item > l.len-1) {
		return nil
	}
	for item < l.curIndex {
		l.movePrev()
	}
	for item > l.curIndex{
		l.moveNext()
	}
	return l.curElem.value
}

func (l *LinkedList) Len() int {
	return l.len
}

func (l LinkedList) String() (s string) {
	s = "["
	for i:=0; i<l.len-1; i++ {
		s += fmt.Sprint(l.Get(i), ",")
	}
	s += fmt.Sprint(l.Get(l.len-1), "]")
	return s
}