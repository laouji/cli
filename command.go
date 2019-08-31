package cli

type Command interface {
	Usage() string
	Run(flags, args []string) (int, error)
}

type CommandFactory func() (Command, error)
