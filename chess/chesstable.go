package chess

import (
	"time"

	"github.com/matiwinnetou/boarder/core"
)

type ChessTable struct {
	*core.Table
	Board *ChessBoard
	Player1 *core.Player
	Player2 *core.Player
}

func NewChessTable(tableNo int32, turnInSeconds time.Duration) *ChessTable {
	turnDuration := time.Minute * turnInSeconds
	table := &core.Table{TableNo: tableNo, TurnDuration: turnDuration}

	return &ChessTable{table, nil, nil, nil}
}

func NewChessTableWith3MinsTurnTime(tableNo int32) *ChessTable {
	turnDuration := time.Minute * 3
	table := &core.Table{TableNo: tableNo, TurnDuration: turnDuration}

	return &ChessTable{table, nil, nil, nil}
}

func (ct *ChessTable) PlayerCount() int {
	if ct.Player1 != nil && ct.Player2 != nil {
		return 2
	}
	if ct.Player1 != nil || ct.Player2 != nil {
		return 1
	}

	return 0
}
