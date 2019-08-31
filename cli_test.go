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

func (testCommand TestCommand) Usage() string                       { return "usage message" }
func (testCommand TestCommand) Run([]string, []string) (int, error) { return 0, nil }

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

func Test_Run_InvalidCommand(t *testing.T) {
	commandName := "command_name"
	c, err := cli.NewCli(commandName, "1.0")
	if err != nil {
		t.Errorf("NewCli with valid params returned unexpected error: %s", err)
	}

	os.Args = []string{commandName, "arg1", "arg2"}

	if exitStatus, err := c.Run(); err.Error() != "invalid command" {
		t.Errorf("Run returned unexpected status or error: %d, %s", exitStatus, err)
	}
	if len(c.Args) != 1 {
		t.Errorf("c contains unexpected number of arguments: %s", c.Args)
	}
}

func Test_Run_WithVersionFlag(t *testing.T) {
	commandName := "command_name"
	c, err := cli.NewCli(commandName, "1.0")
	if err != nil {
		t.Errorf("NewCli with valid params returned unexpected error: %s", err)
	}

	os.Args = []string{commandName, "--version", "subcommand", "arg"}

	exitStatus, err := c.Run()
	if exitStatus != 0 || err != nil {
		t.Errorf("Run returned unexpected status or error: %d, %s", exitStatus, err)
	}
	if len(c.Args) != 1 {
		t.Errorf("c contains unexpected number of arguments: %s", c.Args)
	}
}

func Test_Run_WithHelpFlag(t *testing.T) {
	commandName := "command_name"
	c, err := cli.NewCli(commandName, "4.5.5")
	if err != nil {
		t.Errorf("NewCli with valid params returned unexpected error: %s", err)
	}

	os.Args = []string{commandName, "--help", "subcommand", "arg"}

	exitStatus, err := c.Run()
	if exitStatus != 0 || err != nil {
		t.Errorf("Run returned unexpected status or error: %d, %s", exitStatus, err)
	}
	if len(c.Args) != 1 {
		t.Errorf("c contains unexpected number of arguments: %s", c.Args)
	}
}

func Test_Run_TestCommand(t *testing.T) {
	commandName := "command_name"
	c, err := cli.NewCli(commandName, "4.5.5")
	if err != nil {
		t.Errorf("NewCli with valid params returned unexpected error: %s", err)
	}

	subCommand := "test"
	c.AddCommand(subCommand, func() (cli.Command, error) { return TestCommand{}, nil })

	os.Args = []string{commandName, "--someflag", subCommand, "somearg"}

	exitStatus, err := c.Run()
	if exitStatus != 0 || err != nil {
		t.Errorf("Run returned unexpected status or error: %d, %s", exitStatus, err)
	}
	if len(c.Flags) != 1 {
		t.Errorf("c contains unexpected number of flags %s", c.Flags)
	}
	if len(c.Args) != 1 {
		t.Errorf("c contains unexpected number of arguments: %s", c.Args)
	}
}
