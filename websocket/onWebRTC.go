package websocket

import (
	"fmt"
	"kGoChat/datamodel"
)

func (c *WebsocketController) onWebRTC(_data interface{}) {
	resultInfo := datamodel.ResultInfo{}
	resultInfo.Type = "WebRTC"
	requestInfo, ok := datamodel.MapToRequestInfo(_data)
	fmt.Printf("%+v %v \n", requestInfo, ok)

	resultInfo.Data = *requestInfo
	_ = c.Conn.Emit("WebRTC", resultInfo)

}
