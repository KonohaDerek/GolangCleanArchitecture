package irepository

import (
	entities "example.com/GolangCleanArchitecture/src/Domain/entities"
)

type IBoardRepository interface {
	FindBoardById(id string) entities.Board

	Add(entities.Board)
}

var boards map[string]entities.Board

type BoardRepository struct {
}

func (b *BoardRepository) FindBoardById(id string) entities.Board {
	return boards[id]
}

func (b *BoardRepository) Add(board entities.Board) {
	if boards == nil {
		boards = make(map[string]entities.Board)
	}
	boards[board.GetBoardId()] = board
}
