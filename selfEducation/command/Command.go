package main

import "fmt"

type Command interface {
	execute()
}

type UseTreadmillCommand struct {
	n   int
	gym *Gym
}

func (u *UseTreadmillCommand) execute() {
	fmt.Println("is training on  treadmill", u.n, " minutes")
}

type UseHorizontalBarCommand struct {
	n   int
	gym *Gym
}

func (u *UseHorizontalBarCommand) execute() {
	fmt.Println("is training on  horizontal bar", u.n, " minutes")
}

type UseBarbellCommand struct {
	n   int
	gym *Gym
}

func (u *UseBarbellCommand) execute() {
	fmt.Println("is training on  barbell", u.n, "minutes")
}
