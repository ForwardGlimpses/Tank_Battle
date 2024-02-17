package tank

import "fmt"

type Tank struct {
	dx int
	dy int
	Hp int
}

func New() *Tank {
	return &Tank{}
}

func (t *Tank) Move(direction int) {
	if direction == 0 {
		t.dx -= 1
	} else if direction == 1 {
		t.dx += 1
	} else if direction == 2 {
		t.dy -= 1
	} else {
		t.dy += 1
	}
}
func (t *Tank) Update() {
	for i := 1; i <= 600; i++ {
		for j := 1; j <= 600; j++ {
			if t.dx == i && t.dy == j {
				fmt.Print("T")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}
