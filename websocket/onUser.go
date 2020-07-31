package websocket

import "kGoChat/datamodel"

func (c *WebsocketController) onUser(_data interface{}) {
	resultInfo := datamodel.ResultInfo{}

	_requestInfo, ok := datamodel.MapToRequestInfo(_data)
	if !ok {
		resultInfo.Code = -1
		resultInfo.Message = "错误的参数"
		_ = c.Conn.Emit("User", resultInfo)
		return
	}
	requestInfo := *_requestInfo
	resultInfo.FnId = requestInfo.FnId
	resultInfo.Type = requestInfo.Type
	user := requestInfo.Data.(map[string]interface{})

	switch requestInfo.Type {
	case "Login":
		{
			name := user["user"].(string)
			var flag = false
			for i := range websocketControllers {
				if c != websocketControllers[i] && name == websocketControllers[i].user.User {
					flag = true
					break
				}
			}
			if flag {
				resultInfo.Code = -2
				resultInfo.Message = "用户名重复"
			} else {
				c.user.User = name
				c.user.Name = name
				resultInfo.Message = "登录成功"
				resultInfo.Data = c.user
			}
		}
		break
	default:
		resultInfo.Code = -1
		resultInfo.Data = nil
		resultInfo.Message = "错误的 [Type] 参数"
		break
	}
	_ = c.Conn.Emit("User", resultInfo)

}
