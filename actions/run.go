package actions

import (
	"fmt"
	"jim/models"
	"jim/utils"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

var Run = &Action{
	Value: func(args []string) {

		if len(args) != 1 && len(args) != 2 {
			utils.Alertf("wrong format!!!")
			return
		}

		command := models.Command{}

		if err := FindCommandByName(args[0], &command); err != nil {
			return
		}

		if len(args) == 2 {
			run(command, strings.Join(args[1:], " "))
			return
		}

		run(command, "")
	},
	Description:     "run a command (not required)",
	HelpDescription: "wp",
	ArgumentsLen:    utils.CUSTOM_ARGUMENTS_LEN,
}

func run(command models.Command, args string) {

	models.DB().Save(&command)

	var c *exec.Cmd

	if runtime.GOOS == "windows" {

		c = exec.Command("powershell", "-c", command.Value, args)
	} else {

		shell, err := os.LookupEnv("SHELL")

		if !err {
			utils.Alertf("no shell found!!!")
			return
		}

		c = exec.Command(shell, "-c", command.Value, args)
	}

	c.Stdin = os.Stdin
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr

	if err := c.Run(); err != nil {
		fmt.Println(err.Error())
	}
}
