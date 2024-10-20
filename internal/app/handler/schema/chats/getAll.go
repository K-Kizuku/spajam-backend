package chatschema

type GetAllByUserIDRequest struct {
}

type GetAllByUserIDResponse struct {
	Chats []Chat `json:"chats"`
}
