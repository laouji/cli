package cli

type Command interface {
	Usage() string
	Run(args []string) (Command, error)
}

type CommandFactory func() (Command, error)
