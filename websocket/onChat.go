package websocket

import "kGoChat/datamodel"

func (c *WebsocketController) onChat(_data interface{}) {
	resultInfo := datamodel.ResultInfo{}
	resultInfo.Type = "Chat"
	if c.user.User == "" || len(c.user.User) == 0 {
		resultInfo.Type = "ChatBack"
		resultInfo.Code = -3
		resultInfo.Message = "未登录"
		_ = c.Conn.Emit("Chat", resultInfo)
		return
	}
	_requestInfo, ok := datamodel.MapToRequestInfo(_data)
	if !ok {
		resultInfo.Code = -1
		resultInfo.Message = "错误的参数"
		_ = c.Conn.Emit("Chat", resultInfo)
		return
	}
	requestInfo := *_requestInfo
	msg := requestInfo.Data.(map[string]interface{})
	user := msg["user"].(string)

	var flag = false
	for i := range websocketControllers {
		if user == websocketControllers[i].user.User {
			resultInfo.Code = 0
			msg["user"] = c.user.User
			resultInfo.Data = msg
			_ = websocketControllers[i].Conn.Emit("Chat", resultInfo)
			flag = true
			break
		}
	}
	resultInfo.Type = "ChatBack"
	if flag {
		resultInfo.Code = 0
		resultInfo.Message = "发送消息成功"
	} else {
		resultInfo.Code = -1
		resultInfo.Message = "发送消息失败,未找到用户"
	}
	_ = c.Conn.Emit("Chat", resultInfo)
}
