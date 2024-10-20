package chatschema

type CreateRequest struct {
	UserID1 string `json:"user_id1"`
	UserID2 string `json:"user_id2"`
	Content string `json:"content"`
}
