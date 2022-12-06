/**
 * @Author: 10512203@qq.com
 * @Description:
 * @File: request
 * @Version: 1.0.0
 */

package easysocket

import "google.golang.org/protobuf/proto"

type IRequest interface {
	GetSession() ISession
	GetData() []byte
	GetDataJson() string
	GetMsgId() int32
	SetMsg(msgId int32, data []byte)
	SendMsg(msgId int32, message proto.Message) error
	SendMsgJson(msgId int32, jsonStr string) error
}

type Request struct {
	sess ISession
	msg  IMessage
}

// GetSession 获取请求连接信息
func (r *Request) GetSession() ISession {
	return r.sess
}

// GetData 获取请求消息的数据
func (r *Request) GetData() []byte {
	return r.msg.GetData()
}

// GetData 获取请求消息的数据
func (r *Request) GetDataJson() string {
	return r.msg.GetDataJson()
}

// GetMsgId 获取请求的消息的ID
func (r *Request) GetMsgId() int32 {
	return r.msg.GetMsgId()
}

func (r *Request) SetMsg(msgId int32, data []byte) {
	r.msg.SetMsgId(msgId)
	r.msg.SetData(data)
}

func (r *Request) SendMsg(msgId int32, message proto.Message) error {
	buffer, err := proto.Marshal(message)

	if err != nil {
		return err
	}

	return r.sess.SendBuffMsg(msgId, buffer)
}

func (r *Request) SendMsgJson(msgId int32, jsonStr string) error {
	//buffer, err := proto.Marshal(message)

	// if err != nil {
	// 	return err
	// }

	return r.sess.SendJsonMsg(msgId, jsonStr)
}
