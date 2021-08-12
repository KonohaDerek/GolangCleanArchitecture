package board_test

import (
	"fmt"
	"testing"

	"example.com/GolangCleanArchitecture/src/Application/Interface/Repository/board"
	"example.com/GolangCleanArchitecture/src/Application/usecase/board"
	exit_code "example.com/GolangCleanArchitecture/src/Domain/enums"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func Test_should_succeed_when_create_board_in_team(t *testing.T) {

	teamId, err := uuid.NewUUID()
	repository := &irepository.BoardRepository{}
	boardName := "kanban"
	fmt.Printf("Your UUID is: %s", teamId)
	assert.Nil(t, err)

	input := new(board.CreateBoardUseCaseInput)

	input.SetTeamId(teamId.String())
	input.SetName(boardName)
	output, err := board.CreateBoardUseCase(*input, repository)
	assert.Nil(t, err)
	if assert.NotNil(t, output) {
		assert.NotNil(t, output.GetId())
		assert.Equal(t, exit_code.SUCCESS, output.GetExitCode())
	}

	board := repository.FindBoardById(output.GetId())
	assert.Equal(t, output.GetId(), board.GetBoardId())
	assert.Equal(t, boardName, board.GetName())
	assert.Equal(t, teamId.String(), board.GetTeamId())

	// 測試
	// // assert equality
	// assert.Equal(t, 123, 123, "they should be equal")

	// // assert inequality
	// assert.NotEqual(t, 123, 456, "they should not be equal")

	// // assert for nil (good for errors)
	// assert.Nil(t, object)

	// // assert for not nil (good when you expect something)
	// if assert.NotNil(t, object) {

	//   // now we know that object isn't nil, we are safe to make
	//   // further assertions without causing any errors
	//   assert.Equal(t, "Something", object.Value)

	// }

}
