package logic

import (
	"gitee.com/godLei6/things/shared/errors"
	"gitee.com/godLei6/things/src/dmsvr/dm"
	"github.com/spf13/cast"
	"strings"
)
type LoginDevice struct {
	ClientID 	string	//clientID
	ProductID 	int64	//产品id
	StrProductID string //产品id的字符串版
	DeviceName 	string	//设备名称
	SdkAppID   	int64	//appid 直接填 12010126
	ConnID		string	//随机6字节字符串 帮助查bug
	Expiry 		int64	//过期时间 unix时间戳
}

func GetLoginDevice(userName string) (*LoginDevice,error){
	keys :=strings.Split(userName,";")
	if len(keys) != 4 || len(keys[0]) < 11{
		return nil,errors.Parameter.AddDetail("userName not right")
	}
	lg,err := GetClientIDInfo(keys[0])
	if err != nil {
		return nil, err
	}
	lg.SdkAppID = cast.ToInt64(keys[1])
	lg.ConnID = keys[2]
	lg.Expiry = cast.ToInt64(keys[3])
	return lg,nil
}

func GetClientIDInfo(ClientID string)(*LoginDevice,error){
	ProductID := dm.GetInt64ProductID(ClientID[0:11])
	if ProductID < 0 {
		return nil,errors.Parameter.AddDetail("product id not right")
	}
	DeviceName := ClientID[11:]
	lg:= &LoginDevice{
		ClientID	: ClientID,
		ProductID	: ProductID,
		DeviceName	: DeviceName,
		StrProductID: ClientID[0:11],
	}
	return lg,nil
}