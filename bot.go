package discordbot

type Bot struct {
	Name        string
	Description string
	Token       string
}

func NewBot() *Bot {
	return &Bot{}
}

func (b *Bot) Run() error {
	return nil
}

func (b *Bot) Stop() {

}
