package chess

import (
	"github.com/matiwinnetou/boarder/core"
)

type ChessTable struct {
	*core.Table
	Player1 *core.Player
	Player2 *core.Player
}

func NewChessTable(tableNo int32) *ChessTable {
	table := &core.Table{TableNo: tableNo}

	return &ChessTable{table, nil, nil}
}

func (ct *ChessTable) PlayerCount() int {
	if (ct.Player1 != nil && ct.Player2 != nil) {
		return 2
	}
	if (ct.Player1 != nil || ct.Player2 != nil) {
		return 1
	}

	return 0
}
