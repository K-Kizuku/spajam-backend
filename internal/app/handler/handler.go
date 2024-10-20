package handler

type Root struct {
	UserHandler IUserHandler
	ChatHandler IChatHandler
}

func New(UserHandler IUserHandler, ChatHandler IChatHandler) *Root {
	return &Root{
		UserHandler: UserHandler,
		ChatHandler: ChatHandler,
	}
}
