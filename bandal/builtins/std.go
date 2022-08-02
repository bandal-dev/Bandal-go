package std

import (
	//_list "container/list"
	"fmt"
)

/*
---
* to make *

1. stuct
2. map
3. for
4. typedef
5. fn int, float, str, etc...
6. fn len
7. module regex
8. list replace, append[push], pop, index, etc...
---
*/

type List struct {
	_list []interface{}
}

func (l *List) Init() *List {
	l._list = []interface{}{}
	return l
}

func list() *List {
	return new(List).Init()
}

func (l *List) append(v interface{}) *List {
	l._list = append(l._list, v)
	return l
}

func (l *List) show() {
	fmt.Print("[")
	for i := range l._list {
		fmt.Print(l._list[i])
		if len(l._list)-1 != i {
			fmt.Print(", ")
		}
	}
	fmt.Println("]")
}

func (l *List) get(v int) interface{} {
	//exception
	return l._list[v]
}
func main() {
	l := list()
	l.append(1)
	l.append("hello")
	l.show()
	fmt.Println(l.get(1))
}
