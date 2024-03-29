package notification

type Notification struct {
	ID          int    `json:"id"`
	LiveMode    bool   `json:"live_mode"`
	Type        string `json:"type"`
	DateCreated string `json:"date_created"`
	UserId      string    `json:"user_id"`
	ApiVersion  string `json:"api_version"`
	Action      string `json:"action"`
	Data        struct {
		ID string `json:"id"`
	}
}
