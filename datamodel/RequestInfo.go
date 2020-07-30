package datamodel

import (
	"fmt"
	"github.com/goinggo/mapstructure"
)

type RequestInfo struct {
	Code    int         `json:"code"`
	Type    string      `json:"type"`
	Data    interface{} `json:"data"`
	Remarks string      `json:"remarks"`
}

func MapToRequestInfo(_data interface{}) (*RequestInfo, bool) {
	__data, ok := _data.(map[string]interface{})
	var requestInfo RequestInfo
	if !ok {
		return nil, false
	}
	//将 map 转换为指定的结构体
	if err := mapstructure.Decode(__data, &requestInfo); err != nil {
		fmt.Println(err)
		return nil, false
	}
	return &requestInfo, true
}
