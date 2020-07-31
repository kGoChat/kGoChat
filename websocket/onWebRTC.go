package websocket

import (
	"kGoChat/datamodel"
)

func (c *WebsocketController) onWebRTC(_data interface{}) {
	resultInfo := datamodel.ResultInfo{}
	resultInfo.Type = "WebRTC"
	_requestInfo, ok := datamodel.MapToRequestInfo(_data)
	if !ok {
		resultInfo.Code = -1
		resultInfo.Message = "错误的参数"
		_ = c.Conn.Emit("WebRTC", resultInfo)
		return
	}
	requestInfo := *_requestInfo
	resultInfo.Type = requestInfo.Type
	resultInfo.FnId = requestInfo.FnId
	if c.user.User == "" || len(c.user.User) == 0 {
		resultInfo.Type = "WebRTCBack"
		resultInfo.Code = -3
		resultInfo.Message = "未登录"
		_ = c.Conn.Emit("WebRTC", resultInfo)
		return
	}

	msg := requestInfo.Data.(map[string]interface{})
	user := msg["user"].(string)

	var flag = false
	for i := range websocketControllers {
		if user == websocketControllers[i].user.User {
			resultInfo.Code = 0
			msg["user"] = c.user.User
			resultInfo.Data = msg
			resultInfo.FnId = 0
			_ = websocketControllers[i].Conn.Emit("WebRTC", resultInfo)
			flag = true
			break
		}
	}

	resultInfo.Data = nil
	resultInfo.FnId = _requestInfo.FnId
	resultInfo.Type = "WebRTCBack"
	if flag {
		resultInfo.Code = 0
		resultInfo.Message = "发送消息成功"
	} else {
		resultInfo.Code = -1
		resultInfo.Message = "发送消息失败,未找到用户"
	}
	_ = c.Conn.Emit("WebRTC", resultInfo)

}
