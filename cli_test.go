package cli_test

import (
	"os"
	"testing"

	"github.com/laouji/cli"
)

func Test_NewCli(t *testing.T) {
	_, err := cli.NewCli("command_name", "")
	if err != cli.ErrInvalidVersionString {
		t.Errorf("NewCli with blank version string returned unexpected error: %s", err)
	}

	_, err = cli.NewCli("command_name", "3.0")
	if err != nil {
		t.Errorf("NewCli with valid params returned unexpected error: %s", err)
	}
}

type TestCommand struct{}

func (testCommand TestCommand) Usage() string             { return "usage message" }
func (testCommand TestCommand) Run([]string) (int, error) { return 0, nil }

func Test_AddCommand(t *testing.T) {
	c, err := cli.NewCli("command_name", "someversion")
	if err != nil {
		t.Errorf("NewCli with valid params returned unexpected error: %s", err)
	}

	key := "commandName"
	c.AddCommand(key, func() (cli.Command, error) { return TestCommand{}, nil })
	if factory, ok := c.Commands[key]; !ok {
		t.Errorf("command '%s' not found", key)
	} else {
		factory()
	}
}

func Test_Run(t *testing.T) {
	commandName := "command_name"
	c, err := cli.NewCli(commandName, "1.0")
	if err != nil {
		t.Errorf("NewCli with valid params returned unexpected error: %s", err)
	}

	os.Args = []string{commandName, "arg1", "arg2"}

	exitStatus, err := c.Run()
	if exitStatus != 0 || err != nil {
		t.Errorf("Run returned unexpected status or error: %d, %s", exitStatus, err)
	}
	if len(c.Args) != 2 {
		t.Errorf("c contains unexpected number of arguments: %s", c.Args)
	}
}
