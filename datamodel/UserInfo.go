package datamodel

type UserInfo struct {
	UUID    string      `json:"uuid"`
	User    string      `json:"user"`
	Passwd  string      `json:"passwd"`
	Name    string      `json:"name"`
	Remarks interface{} `json:"remarks"`

	WebSocketID string
}
