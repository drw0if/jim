package actions

import (
	"fmt"
	"jim/pkg/models"
	"jim/pkg/rainbow"
	"jim/pkg/utils"
	"strings"
	"testing"

	"github.com/go-playground/assert"
)

func TestRun(t *testing.T) {

	rainbow.Blank()

	command := models.Command{
		Name:  "to_run",
		Value: "echo 1",
	}

	command.Save()

	// correct test
	correctMockResponse := "jim is launching > [powershell -c echo 1 ]\r"
	correctArgs := []string{"to_run"}

	correctResponseData := utils.InterceptStdout(func() {

		if !Run.ArgumentsCheck(correctArgs) {
			fmt.Println("wrong format")
			return
		}

		Run.Value(correctArgs)

	})

	assert.Equal(t, strings.TrimSpace(correctResponseData), strings.TrimSpace(correctMockResponse))

	// wrong test
	wrongMockResponse := "wrong format"
	wrongArgs := []string{}

	// check if the args check works

	wrongResponseData := utils.InterceptStdout(func() {

		if !Run.ArgumentsCheck(wrongArgs) {
			fmt.Println("wrong format")
			return
		}

		Run.Value(wrongArgs)

	})

	assert.Equal(t, wrongResponseData, wrongMockResponse)
}
