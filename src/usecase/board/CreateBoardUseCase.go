package board

import (
	exit_code "example.com/GolangCleanArchitecture/src/enums"
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

func CreateBoardUseCase(input CreateBoardUseCaseInput) (CreateBoardUseCaseOutput, error) {
	output := CreateBoardUseCaseOutput{}
	newId, err := uuid.NewUUID()
	if err == nil {
		output.exitCode = exit_code.SUCCESS
		output.id = newId.String()
	} else {
		output.exitCode = exit_code.Faild
	}

	return output, err
}
