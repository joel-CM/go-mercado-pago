package order

type PayerOrder struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	NickName string `json:"nick_name"`
}
