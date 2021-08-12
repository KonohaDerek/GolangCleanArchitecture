package board

type Board struct {
	Id     string
	Name   string
	TeamId string
}

func (board *Board) GetBoardId() string {
	return board.Id
}

func (board *Board) GetName() string {
	return board.Name
}

func (board *Board) GetTeamId() string {
	return board.TeamId
}
