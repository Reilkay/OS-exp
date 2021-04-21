package main

type PIDLib struct {
	IDs   map[int]bool
	nowID int
}

func newPIDLib() PIDLib {
	newLib := new(PIDLib)
	newLib.nowID = 0
	newLib.IDs = make(map[int]bool)
	for i := 0; i < 100; i++ {
		newLib.IDs[i] = false
	}
	return *newLib
}

func (p PIDLib) IDGenerator() int {
	p.nowID++
	for p.IDs[p.nowID] {
		p.nowID++
		p.nowID %= 100
	}
	p.IDs[p.nowID] = true
	return p.nowID
}

func (p PIDLib) IDDestory(ID int) {
	p.IDs[ID] = true
}
