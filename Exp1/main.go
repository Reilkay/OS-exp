package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
)

func main() {
	var memory int
	fmt.Println("请输入总内存(int): ")
	fmt.Scan(&memory)
	manager := NewScheduler(memory)
	scanner := bufio.NewScanner(os.Stdin)
	for {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
		fmt.Println(manager)
		fmt.Println("请输入采用的操作：")
		if scanner.Scan() {
			switch scanner.Text() {
			case "Create":
				manager.Create()
			case "EventWait":
				manager.EventWait()
			case "EventOccur":
				manager.EventOccurs()
			case "Timeout":
				manager.Timeout()
			case "Release":
				manager.Release()
			case "EXIT":
				return
			}
		}
	}
}
