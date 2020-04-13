package discordbot

type Messenger interface {
	Send()
}

type Message struct {
	Message string
}
