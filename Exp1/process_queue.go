package main

import (
	"container/heap"
	"fmt"
)

type Process struct {
	// 进程ID
	ID int
	// 进程名
	Name string
	// 占用内存
	Memory int
	// 优先级
	Priority int
}

func NewProcess(ID int, Name string, Memory int, Priority int) Process {
	return Process{
		ID: ID, Name: Name, Memory: Memory, Priority: Priority,
	}
}

func (p Process) String() string {
	return fmt.Sprintf("(ID: %d Name: %s Memory: %d Priority:%d)", p.ID, p.Name, p.Memory, p.Priority)
}

type ProcessQueue []Process

var _ heap.Interface = (*ProcessQueue)(nil)

func (pq ProcessQueue) Len() int {
	return len(pq)
}
func (pq ProcessQueue) Less(i int, j int) bool {
	return pq[i].Priority < pq[j].Priority
}
func (pq ProcessQueue) Swap(i int, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *ProcessQueue) Push(x interface{}) {
	*pq = append(*pq, x.(Process))
}
func (pq *ProcessQueue) Pop() interface{} {
	t := (*pq)[len(*pq)-1]
	*pq = (*pq)[:len(*pq)-1]
	return t
}
