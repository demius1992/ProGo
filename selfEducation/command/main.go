package main

import "fmt"

type Gender string

const (
	male   Gender = "male"
	female Gender = "female"
)

type Users struct {
	gender   Gender
	Commands []Command
}

func (u *Users) executeCommands() {
	for _, command := range u.Commands {
		command.execute()
	}
}

func main() {
	g := NewGym()

	tasks := []Command{
		g.UseBarbell(20),
		g.UseHorizontalBar(15),
		g.UseTreadmill(30),
		g.UseBarbell(20),
		g.UseHorizontalBar(15),
	}

	users := []*Users{
		&Users{gender: male},
		&Users{gender: female},
	}

	for i, task := range tasks {
		user := users[i%len(users)]
		user.Commands = append(user.Commands, task)
	}

	for _, c := range users {
		fmt.Println(c.gender, ":")
		c.executeCommands()
	}
}
