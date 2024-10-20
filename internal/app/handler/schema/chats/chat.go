package chatschema

type Chat struct {
	ChatID  string `json:"chat_id"`
	UserID1 string `json:"user_id1"`
	UserID2 string `json:"user_id2"`
	Content string `json:"content"`
}
