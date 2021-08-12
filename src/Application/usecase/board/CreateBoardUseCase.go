package board

import (
	"example.com/GolangCleanArchitecture/src/Application/Interface/Repository/board"
	"example.com/GolangCleanArchitecture/src/Domain/entities"
	exit_code "example.com/GolangCleanArchitecture/src/Domain/enums"
	"github.com/google/uuid"
)

type CreateBoardUseCaseInput struct {
	TeamId string
	Name   string
}

type CreateBoardUseCaseOutput struct {
	id       string
	exitCode int
}

func (input *CreateBoardUseCaseInput) SetTeamId(id string) {
	input.TeamId = id
}

func (input *CreateBoardUseCaseInput) SetName(name string) {
	input.Name = name
}

func (output *CreateBoardUseCaseOutput) GetId() string {
	return output.id
}

func (output *CreateBoardUseCaseOutput) GetExitCode() int {
	return output.exitCode
}

func CreateBoardUseCase(input CreateBoardUseCaseInput, repository irepository.IBoardRepository) (CreateBoardUseCaseOutput, error) {
	output := CreateBoardUseCaseOutput{}

	newId, err := uuid.NewUUID()
	if err == nil {
		output.exitCode = exit_code.SUCCESS
		output.id = newId.String()
	} else {
		output.exitCode = exit_code.Faild
	}

	newboard := board.Board{
		Id:     newId.String(),
		TeamId: input.TeamId,
		Name:   input.Name,
	}
	repository.Add(newboard)
	return output, err
}
