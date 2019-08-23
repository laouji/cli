package cli_test

import (
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
