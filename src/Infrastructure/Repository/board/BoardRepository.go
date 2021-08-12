package repostory

import (
	entities "example.com/GolangCleanArchitecture/src/Domain/entities"
)

type BoardRepository struct  {
}

func (b *BoardRepository) FindBoardById(id string) entities.Board {
	return entities.Board{}
}
