package cli

type Command interface {
	Usage() string
	Run(args []string) (int, error)
}

type CommandFactory func() (Command, error)
