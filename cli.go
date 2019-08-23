package cli

import (
	"errors"
)

var ErrInvalidVersionString = errors.New("version string cannot be blank")

type Cli struct {
	Name     string
	Version  string
	Args     []string
	Commands map[string]CommandFactory
}

func NewCli(name, version string) (cli *Cli, err error) {
	if version == "" {
		return nil, ErrInvalidVersionString
	}

	return &Cli{
		Name:    name,
		Version: version,
	}, nil
}

func (c *Cli) Run() (exitStatus int, err error) {
	// TODO
	return 0, nil
}
