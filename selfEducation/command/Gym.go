package main

type Gym struct {
	TreadmillCount     int
	HorizontalBarCount int
	BarbellCount       int
}

func NewGym() *Gym {
	return &Gym{}
}

func (g *Gym) UseTreadmill(n int) Command {
	return &UseTreadmillCommand{
		n:   n,
		gym: g,
	}
}

func (g *Gym) UseHorizontalBar(n int) Command {
	return &UseHorizontalBarCommand{
		n:   n,
		gym: g,
	}
}

func (g *Gym) UseBarbell(n int) Command {
	return &UseBarbellCommand{
		n:   n,
		gym: g,
	}
}
