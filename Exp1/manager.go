package main

type Manager interface {
	admit()
	dispatch()

	Create()
	Release()
	Timeout()
	EventOccurs()
	EventWait()
}
