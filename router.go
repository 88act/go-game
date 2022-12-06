/**
 * @Author: 10512203@qq.com
 * @Description:
 * @File: router
 * @Version: 1.0.0
 * @Date: 2022/4/7 11:20
 */

package easysocket

import "google.golang.org/protobuf/proto"

type IRouter interface {
	PreHandle(request IRequest, message proto.Message)
	Handle(request IRequest, message proto.Message)
	PostHandle(request IRequest, message proto.Message)
}

type BaseRouter struct{}

func (r *BaseRouter) PreHandle(request IRequest, message proto.Message) {}

func (r *BaseRouter) Handle(request IRequest, message proto.Message) {}

func (r *BaseRouter) PostHandle(request IRequest, message proto.Message) {}
