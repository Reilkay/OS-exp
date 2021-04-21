package main

import (
	"container/heap"
	"fmt"
)

type Scheduler struct {
	// 五状态优先队列
	New, Ready, Running, Blocked, Exit ProcessQueue
	// 最大内存
	TotalMemory int
	// 当前内存
	UsingMemory int
	// ID库
	PIDs PIDLib
}

var _ Manager = (*Scheduler)(nil)

func NewScheduler(Memory int) Scheduler {
	newScheduler := new(Scheduler)
	newScheduler.TotalMemory = Memory
	newScheduler.PIDs = newPIDLib()
	return *newScheduler
}

// 重写输出
func (s Scheduler) String() string {
	return fmt.Sprintf(`当前状态：
TotalMemory: %vMB
UsingMemory: %vMB
New: %v
Ready: %v
Blocked: %v
Running: %v
Exit: %v`, s.TotalMemory, s.UsingMemory, s.New, s.Ready, s.Blocked, s.Running, s.Exit)
}

func (s Scheduler) PrintStatus() {
	fmt.Println(s)
	fmt.Println("回车以继续")
	fmt.Scanln()
}

// TODO: 实现Manager接口
// 创建进程
func (s *Scheduler) Create() {
	var (
		memory   int
		priority int
		name     string
	)
	fmt.Print("请输入进程内存占用(int)：")
	fmt.Scanln(&memory)
	fmt.Print("请输入进程名(string)：")
	fmt.Scanln(&name)
	fmt.Print("请输入进程优先级(int)：")
	fmt.Scanln(&priority)
	heap.Push(&s.New, NewProcess(s.PIDs.IDGenerator(), name, memory, priority))
	fmt.Println("添加进程到“New”中")
	s.PrintStatus()
	s.admit()
}

func (s *Scheduler) admit() {
	if len(s.New) >= 1 {
		processAdmit := heap.Pop(&s.New).(Process)
		// 判断内存是否足够
		if s.UsingMemory+processAdmit.Memory <= s.TotalMemory {
			heap.Push(&s.Ready, processAdmit)
			s.UsingMemory += processAdmit.Memory
			fmt.Println("将“New”中的新进程放入“Ready”")
			s.PrintStatus()
		} else {
			heap.Push(&s.New, processAdmit)
		}
	}
	s.dispatch()
}

// 调度
func (s *Scheduler) dispatch() {
	// 若有进程就绪且无程序运行，则调度进程到 “Running”
	if len(s.Running) == 0 && len(s.Ready) >= 1 {
		heap.Push(&s.Running, heap.Pop(&s.Ready).(Process))
		fmt.Println("将“Ready”中优先级最高的进程放入“Running”")
		s.PrintStatus()
	}
}

// 时间片超时
func (s *Scheduler) Timeout() {
	// 如果当前有运行任务
	if len(s.Running) == 1 {
		heap.Push(&s.Ready, heap.Pop(&s.Running).(Process))
		fmt.Println("将“Running”进程放回“Ready”")
		s.PrintStatus()
		s.admit()
	} else {
		// getError("当前无运行进程")
	}
}

func (s *Scheduler) EventWait() {
	// 如果当前有运行任务
	if len(s.Running) == 1 {
		heap.Push(&s.Blocked, heap.Pop(&s.Running).(Process))
		fmt.Println("将“Running”进程放到“Blocked”")
		s.PrintStatus()
		// 尝试调度
		s.admit()
	} else {
		// getError("当前无运行进程")
	}
}

func (s *Scheduler) EventOccurs() {
	// 如果当前有阻塞任务
	if len(s.Blocked) >= 1 {
		heap.Push(&s.Ready, heap.Pop(&s.Blocked))
		fmt.Println("将“Blocked”中最高优先级的进程放入“Ready”中")
		s.PrintStatus()
		// 尝试调度
		s.admit()
	} else {
		// getError("阻塞进程列表为空")
	}
}

func (s *Scheduler) Release() {
	// 如果当前有运行任务
	if len(s.Running) == 1 {
		processRelease := heap.Pop(&s.Running).(Process)
		heap.Push(&s.Exit, processRelease)
		s.UsingMemory -= processRelease.Memory
		s.PIDs.IDDestory(processRelease.ID)
		fmt.Println("将“Running”进程放到“Exit”中")
		s.PrintStatus()
		// 尝试调度
		s.admit()
	} else {
		// getError("当前无运行进程")
	}
}
