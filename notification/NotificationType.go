package notification

type Notification struct {
	ID          string `json:"id"`
	LiveMode    bool   `json:"live_mode"`
	Type        string `json:"type"`
	DateCreated string `json:"date_created"`
	UserId      int    `json:"user_id"`
	ApiVersion  string `json:"api_version"`
	Action      string `json:"action"`
	Data        struct {
		ID string `json:"id"`
	}
}