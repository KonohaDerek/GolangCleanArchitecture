package repostory

import (
	entities "example.com/GolangCleanArchitecture/src/entities"
)

type BoardRepository struct {
	Board string
}

func FindBoardById(id string) entities.Board {
	return entities.Board{}
}
