package guildedgo

type CommandsBuilder struct {
	Commands []Command
}

type Command struct {
	CommandName string
	Action      func(client *Client, v *ChatMessageCreated)
}

type CommandService interface {
	AddCommand(command *Command)
	AddCommands(commands *CommandsBuilder)
}

type commandService struct {
	client *Client
}

var _ CommandService = &commandService{}

func (service *commandService) AddCommand(command *Command) {
	service.client.Command(command.CommandName, command.Action)
}

func (service *commandService) AddCommands(builder *CommandsBuilder) {
	for _, command := range builder.Commands {
		service.client.Command(command.CommandName, command.Action)
	}
}
