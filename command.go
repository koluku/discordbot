package discordbot

type Commands struct {
	Commands []Command
}

type Usage interface {
	Help()
}

type Command struct {
	Command func()
	Usage   string
}

func (c *Command) Run() error {
	return nil
}

func (c *Command) Help() {

}

type CommandGroup struct {
	Commands []Command
	Usage    string
}

func (cg *CommandGroup) Help() {

}
