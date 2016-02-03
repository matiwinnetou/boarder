package main

import (
	"fmt"

	"github.com/matiwinnetou/boarder/chess"
	"github.com/matiwinnetou/boarder/core"
)

func main() {
	t1 := chess.NewChessTable3Minutes(1)
	t1.Player1 = &core.Player{Name: "Jan"}
	t1.Player2 = &core.Player{Name: "Stefan"}

	fmt.Printf("%d\n", t1.TableNo)
	fmt.Printf("%d\n", t1.PlayerCount())
	fmt.Printf("%s\n", t1.TurnDuration)
}
