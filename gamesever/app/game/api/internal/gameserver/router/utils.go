package router

import (
	"fmt"

	"github.com/aceld/zinx/ziface"
)

func getLogMsg(req ziface.IRequest) string {
	msg := fmt.Sprintf("msgid=%d,ConnID=%d,ip=%s,err=", req.GetMsgID(), req.GetConnection().GetConnID(), req.GetConnection().RemoteAddr())
	return msg
}
