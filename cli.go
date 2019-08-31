package cli

import (
	"errors"
	"os"
	"strings"
)

var (
	ErrInvalidVersionString = errors.New("version string cannot be blank")
	ErrInvalidCommand       = errors.New("invalid command")
)

type Cli struct {
	Name     string
	Version  string
	Args     []string
	Flags    []string
	Commands map[string]CommandFactory

	hasVersion bool
	hasHelp    bool
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
	commandName := c.parseArgs()

	if c.hasVersion == true {
		os.Stderr.Write([]byte(c.Version + "\n"))
		return 0, nil
	}

	if c.hasHelp == true {
		c.printUsage()
		return 0, nil
	}

	commandFactory, ok := c.Commands[commandName]
	if !ok {
		c.printUsage()
		return 1, ErrInvalidCommand
	}

	command, err := commandFactory()
	if err != nil {
		return 1, err
	}
	command.Run(c.Flags, c.Args)

	return 0, nil
}

func (c *Cli) parseArgs() (command string) {
	for _, arg := range os.Args[1:] {
		if arg == "-v" || arg == "-version" || arg == "--version" {
			c.hasVersion = true
			continue
		}
		if arg == "-h" || arg == "-help" || arg == "--help" {
			c.hasHelp = true
			continue
		}
		if strings.HasPrefix(arg, "-") {
			c.Flags = append(c.Flags, arg)
			continue
		}

		// first seen argument is assumed to be a subcommand
		if command == "" {
			command = arg
			continue
		}
		c.Args = append(c.Args, arg)
	}
	return command
}

func (c *Cli) printUsage() {
	// TODO generate usage message from commands
	os.Stderr.Write([]byte("placeholder for usage message\n"))
}
