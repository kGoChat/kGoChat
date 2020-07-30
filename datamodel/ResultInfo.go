package datamodel

type ResultInfo struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Type    string      `json:"type"`
	Data    interface{} `json:"data"`
	Remarks string      `json:"remarks"`
}
