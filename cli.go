package cli

import (
	"errors"
	"os"
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
		Name:     name,
		Version:  version,
		Commands: make(map[string]CommandFactory),
	}, nil
}

func (c *Cli) AddCommand(name string, factory CommandFactory) {
	c.Commands[name] = factory
}

func (c *Cli) Run() (exitStatus int, err error) {
	c.Args = os.Args[1:]
	// TODO
	return 0, nil
}
