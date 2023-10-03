package main

import "fmt"

type Employee struct {
	ID   int
	Name string
	Sal  int
	Next *Employee
}

type EmpHashTable struct {
	Table []*Employee
	Size  int
}

func NewTable(size int) *EmpHashTable {
	return &EmpHashTable{
		Table: make([]*Employee, size),
		Size:  size,
	}
}

// 插入
func (ht *EmpHashTable) Insert(emp *Employee) {
	Index := int(emp.ID) % ht.Size
	if ht.Table[Index] == nil {
		ht.Table[Index] = emp
	} else {
		current := ht.Table[Index]
		for current.Next != nil {
			current = current.Next
		}
		current.Next = emp
	}
}

// 查找
func (ht *EmpHashTable) search(id int) (emp *Employee) {
	Index := id % ht.Size
	if ht.Table[Index].ID == id {
		return ht.Table[Index]
	} else {
		current := ht.Table[Index]
		for current != nil {
			if current.ID == id {
				return current
			}
			current = current.Next
		}
	}
	return nil
}

// 删除
func (ht *EmpHashTable) del(id int) bool {
	Index := id % ht.Size
	current := ht.Table[Index]
	var prev *Employee

	for current != nil {
		if current.ID == id {
			if prev == nil {
				//删除的是第一个元素
				ht.Table[Index] = current.Next
			} else {
				//中间或者末尾
				prev.Next = current.Next
			}
			return true
		}
		prev = current
		current = current.Next
	}
	return false
}

// 修改
func (ht *EmpHashTable) change(id int, newEmp *Employee) bool {
	Index := id % ht.Size
	current := ht.Table[Index]
	for current != nil {
		if current.ID == id {
			current.Name = newEmp.Name
			current.Sal = newEmp.Sal
			return true
		}
		current = current.Next
	}
	return false
}

// 遍历
func (ht *EmpHashTable) DisPlayAll() {
	for _, emp := range ht.Table {
		current := emp
		for current != nil {
			fmt.Printf("ID:%d,Name:%s,sal:%d\n", current.ID, current.Name, current.Sal)
			current = current.Next
		}
	}
}

func main() {
	empTable := NewTable(5)
	empTable.Insert(&Employee{ID: 101, Name: "x14n", Sal: 3000})
	empTable.Insert(&Employee{ID: 102, Name: "x14n", Sal: 3000})
	empTable.Insert(&Employee{ID: 103, Name: "x14n", Sal: 3000})
	empTable.Insert(&Employee{ID: 104, Name: "x14n", Sal: 3000})
	empTable.Insert(&Employee{ID: 105, Name: "x14n", Sal: 3000})
	empTable.Insert(&Employee{ID: 106, Name: "x14n", Sal: 3000})
	empTable.Insert(&Employee{ID: 107, Name: "x14n", Sal: 3000})
	empTable.Insert(&Employee{ID: 108, Name: "x14n", Sal: 3000})
	empTable.Insert(&Employee{ID: 109, Name: "x14n", Sal: 3000})

	emp := empTable.search(101)
	if emp == nil {
		fmt.Println("no found")
	} else {
		fmt.Println("find")
		fmt.Printf("ID:%d,Name:%s,sal:%d\n", emp.ID, emp.Name, emp.Sal)
	}

	empTable.DisPlayAll()

	fmt.Println("//删除")
	empTable.del(102)
	fmt.Println("//修改")
	empTable.change(101, &Employee{Name: "ahahah", Sal: 10000000})
	fmt.Println("////////")
	empTable.search(101)
	fmt.Println("////////")
	empTable.DisPlayAll()
}
